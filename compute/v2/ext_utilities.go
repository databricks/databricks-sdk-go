// TODO : Add the missing methods and implement the methods
// This file has not been completely shifted from the SDK-Beta
// as we still don't have the wait for state methods in the SDK-mod
package compute

func (c *ClusterDetails) IsRunningOrResizing() bool {
	return c.State == StateRunning || c.State == StateResizing
}
