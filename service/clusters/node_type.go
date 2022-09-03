package clusters

import (
	"fmt"
	"strings"
)

// NodeTypeRequest is a wrapper for local filtering of node types
type NodeTypeRequest struct {
	MinMemoryGB           int32  `json:"min_memory_gb,omitempty"`
	GBPerCore             int32  `json:"gb_per_core,omitempty"`
	MinCores              int32  `json:"min_cores,omitempty"`
	MinGPUs               int32  `json:"min_gpus,omitempty"`
	LocalDisk             bool   `json:"local_disk,omitempty"`
	Category              string `json:"category,omitempty"`
	PhotonWorkerCapable   bool   `json:"photon_worker_capable,omitempty"`
	PhotonDriverCapable   bool   `json:"photon_driver_capable,omitempty"`
	Graviton              bool   `json:"graviton,omitempty"`
	IsIOCacheEnabled      bool   `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool   `json:"support_port_forwarding,omitempty"`
	VCPU                  bool   `json:"vcpu,omitempty"`
}

// sort NodeTypes within this struct
func (ntl *ListNodeTypesResponse) sort() {
	sortByChain(ntl.NodeTypes, func(i int) sortCmp {
		var localDisks, localDiskSizeGB int32
		// if ntl.NodeTypes[i].InstanceTypeId != nil {
		// 	localDisks = ntl.NodeTypes[i].NodeInstanceType.LocalDisks
		// 	localDiskSizeGB = ntl.NodeTypes[i].NodeInstanceType.LocalDiskSizeGB
		// }
		return sortChain{
			boolAsc(ntl.NodeTypes[i].IsDeprecated),
			intAsc(localDisks),
			intAsc(localDiskSizeGB),
			intAsc(ntl.NodeTypes[i].MemoryMb),
			intAsc(ntl.NodeTypes[i].NumCores),
			//intAsc(ntl.NodeTypes[i].NumGPUs),
			strAsc(ntl.NodeTypes[i].InstanceTypeId),
		}
	})
}

func (ntl *ListNodeTypesResponse) Smallest(r NodeTypeRequest) (string, error) {
	// error is explicitly ingored here, because Azure returns
	// apparently too big of a JSON for Go to parse
	if len(ntl.NodeTypes) == 0 {
		if r.VCPU {
			return "vcpu-worker", nil
		}
		return "", fmt.Errorf("cannot determine smallest node type with empty response")
	}
	ntl.sort()
	for _, nt := range ntl.NodeTypes {
		gbs := int32(nt.MemoryMb / 1024)
		if r.VCPU && !strings.HasPrefix(nt.NodeTypeId, "vcpu") {
			continue
		}
		if !r.VCPU && strings.HasPrefix(nt.NodeTypeId, "vcpu") {
			continue
		}
		if r.MinMemoryGB > 0 && gbs < r.MinMemoryGB {
			continue
		}
		if r.GBPerCore > 0 && (gbs/int32(nt.NumCores)) < r.GBPerCore {
			continue
		}
		if r.MinCores > 0 && int32(nt.NumCores) < r.MinCores {
			continue
		}
		if r.MinGPUs > 0 && int32(nt.NumGpus) < r.MinGPUs {
			continue
		}
		if r.LocalDisk && nt.NodeInstanceType != nil &&
			(nt.NodeInstanceType.LocalDisks < 1 &&
				nt.NodeInstanceType.LocalNvmeDisks < 1) {
			continue
		}
		if r.Category != "" && !strings.EqualFold(nt.Category, r.Category) {
			continue
		}
		if r.IsIOCacheEnabled && nt.IsIoCacheEnabled != r.IsIOCacheEnabled {
			continue
		}
		if r.SupportPortForwarding && nt.SupportPortForwarding != r.SupportPortForwarding {
			continue
		}
		if r.PhotonDriverCapable && nt.PhotonDriverCapable != r.PhotonDriverCapable {
			continue
		}
		if r.PhotonWorkerCapable && nt.PhotonWorkerCapable != r.PhotonWorkerCapable {
			continue
		}
		if r.Graviton && nt.IsGraviton != r.Graviton {
			continue
		}
		return nt.NodeTypeId, nil
	}
	return "", fmt.Errorf("cannot determine smallest node type")
}
