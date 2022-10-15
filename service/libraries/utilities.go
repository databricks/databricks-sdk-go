package libraries

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/databricks/databricks-sdk-go/retries"
)

type Wait struct {
	ClusterID string
	Timeout   time.Duration
	IsRunning bool
	IsRefresh bool
}

func (library Library) String() string {
	if library.Whl != "" {
		return fmt.Sprintf("whl:%s", library.Whl)
	}
	if library.Jar != "" {
		return fmt.Sprintf("jar:%s", library.Jar)
	}
	if library.Pypi != nil && library.Pypi.Package != "" {
		return fmt.Sprintf("pypi:%s%s", library.Pypi.Repo, library.Pypi.Package)
	}
	if library.Maven != nil && library.Maven.Coordinates != "" {
		mvn := library.Maven
		return fmt.Sprintf("mvn:%s%s%s", mvn.Repo, mvn.Coordinates,
			strings.Join(mvn.Exclusions, ""))
	}
	if library.Egg != "" {
		return fmt.Sprintf("egg:%s", library.Egg)
	}
	if library.Cran != nil && library.Cran.Package != "" {
		return fmt.Sprintf("cran:%s%s", library.Cran.Repo, library.Cran.Package)
	}
	return "unknown"
}

func (cll *InstallLibraries) Sort() {
	sort.Slice(cll.Libraries, func(i, j int) bool {
		return cll.Libraries[i].String() < cll.Libraries[j].String()
	})
}

// ToLibraryList convert to envity for convenient comparison
func (cls ClusterLibraryStatuses) ToLibraryList() InstallLibraries {
	cll := InstallLibraries{ClusterId: cls.ClusterId}
	for _, lib := range cls.LibraryStatuses {
		cll.Libraries = append(cll.Libraries, *lib.Library)
	}
	cll.Sort()
	return cll
}

// IsRetryNeeded returns first bool if there needs to be retry.
// If there needs to be retry, error message will explain why.
// If retry does not need to happen and error is not nil - it failed.
func (cls ClusterLibraryStatuses) IsRetryNeeded(refresh bool) (bool, error) {
	pending := 0
	ready := 0
	errors := []string{}
	for _, lib := range cls.LibraryStatuses {
		if lib.IsLibraryForAllClusters {
			continue
		}
		switch lib.Status {
		// No action has yet been taken to install the library. This state should be very short lived.
		case "PENDING":
			pending++
		// Metadata necessary to install the library is being retrieved from the provided repository.
		case "RESOLVING":
			pending++
		// The library is actively being installed, either by adding resources to Spark
		// or executing system commands inside the Spark nodes.
		case "INSTALLING":
			pending++
		// The library has been successfully installed.
		case "INSTALLED":
			ready++
		// Installation on a Databricks Runtime 7.0 or above cluster was skipped due to Scala version incompatibility.
		case "SKIPPED":
			ready++
		// The library has been marked for removal. Libraries can be removed only when clusters are restarted.
		case "UNINSTALL_ON_RESTART":
			ready++
		//Some step in installation failed. More information can be found in the messages field.
		case "FAILED":
			if refresh {
				// we're reading library list on a running cluster and some of the libs failed to install
				continue
			}
			errors = append(errors, fmt.Sprintf("%s failed: %s", lib.Library, strings.Join(lib.Messages, ", ")))
			continue
		}
	}
	if pending > 0 {
		return true, fmt.Errorf("%d libraries are ready, but there are still %d pending", ready, pending)
	}
	if len(errors) > 0 {
		return false, fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	return false, nil
}

type UpdateLibraries struct {
	ClusterId string
	// The libraries to install.
	Install []Library
	// The libraries to install.
	Uninstall []Library

	Timeout time.Duration
}

func (a *LibrariesAPI) UpdateLibraries(ctx context.Context, update UpdateLibraries) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "update-libraries")
	if len(update.Uninstall) > 0 {
		err := a.Uninstall(ctx, UninstallLibraries{
			ClusterId: update.ClusterId,
			Libraries: update.Uninstall,
		})
		if err != nil {
			return fmt.Errorf("uninstall: %w", err)
		}
	}
	if len(update.Install) > 0 {
		err := a.Install(ctx, InstallLibraries{
			ClusterId: update.ClusterId,
			Libraries: update.Install,
		})
		if err != nil {
			return fmt.Errorf("install: %w", err)
		}
	}
	_, err := a.WaitForLibrariesInstalled(ctx, Wait{
		ClusterID: update.ClusterId,
		Timeout:   update.Timeout,
		IsRunning: true,
		IsRefresh: false,
	})
	return err
}

// clusterID string, timeout time.Duration, isActive bool, refresh bool
func (a *LibrariesAPI) WaitForLibrariesInstalled(ctx context.Context, wait Wait) (*ClusterLibraryStatuses, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "wait-for-libraries")
	result, err := retries.Poll(ctx, wait.Timeout, func() (*ClusterLibraryStatuses, *retries.Err) {
		status, err := a.ClusterStatusByClusterId(ctx, wait.ClusterID)
		if apierr.IsMissing(err) {
			// eventual consistency error
			return nil, retries.Continue(err)
		}
		if err != nil {
			return nil, retries.Halt(err)
		}
		if !wait.IsRunning {
			logger.Infof("Cluster %s is currently not running, so just returning list of %d libraries",
				wait.ClusterID, len(status.LibraryStatuses))
			return status, nil
		}
		retry, err := status.IsRetryNeeded(wait.IsRefresh)
		if retry {
			return status, retries.Continue(err)
		}
		if err != nil {
			return status, retries.Halt(err)
		}
		return status, nil
	})
	if err != nil {
		return nil, err
	}
	if wait.IsRunning {
		installed := []LibraryFullStatus{}
		cleanup := UninstallLibraries{
			ClusterId: wait.ClusterID,
			Libraries: []Library{},
		}
		// cleanup libraries that failed to install
		for _, v := range result.LibraryStatuses {
			if v.Status == "FAILED" {
				logger.Warnf("Removing failed library %s from %s",
					v.Library, wait.ClusterID)
				cleanup.Libraries = append(cleanup.Libraries, *v.Library)
				continue
			}
			installed = append(installed, v)
		}
		// and result contains only the libraries that were successfully installed
		result.LibraryStatuses = installed
		if len(cleanup.Libraries) > 0 {
			err = a.Uninstall(ctx, cleanup)
			if err != nil {
				return nil, fmt.Errorf("cannot cleanup libraries: %w", err)
			}
		}
	}
	return result, nil
}
