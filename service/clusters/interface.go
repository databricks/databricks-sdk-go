// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
)


// &lt;needs content added&gt;
type ClustersService interface {
    // Public version of editClusterOwner, allowing admins to change cluster
    // owner
    ChangeOwner(ctx context.Context, changeClusterOwner ChangeClusterOwner) error
    // Creates a new Spark cluster. This method will acquire new instances from
    // the cloud provider if necessary. This method is asynchronous; the
    // returned ``cluster_id`` can be used to poll the cluster status. When this
    // method returns, the cluster will be in a ``PENDING`` state. The cluster
    // will be usable once it enters a ``RUNNING`` state. Note: Databricks may
    // not be able to acquire some of the requested nodes, due to cloud provider
    // limitations (account limits, spot price, ...) or transient network
    // issues. If Databricks acquires at least 85% of the requested on-demand
    // nodes, cluster creation will succeed. Otherwise the cluster will
    // terminate with an informative error message. An example request: ..
    // code:: { &#34;cluster_name&#34;: &#34;my-cluster&#34;, &#34;spark_version&#34;:
    // &#34;2.0.x-scala2.10&#34;, &#34;node_type_id&#34;: &#34;r3.xlarge&#34;, &#34;spark_conf&#34;: {
    // &#34;spark.speculation&#34;: true }, &#34;aws_attributes&#34;: { &#34;availability&#34;: &#34;SPOT&#34;,
    // &#34;zone_id&#34;: &#34;us-west-2a&#34; }, &#34;num_workers&#34;: 25 } See below as an example
    // for an autoscaling cluster. Note that this cluster will start with `2`
    // nodes, the minimum. .. code:: { &#34;cluster_name&#34;: &#34;autoscaling-cluster&#34;,
    // &#34;spark_version&#34;: &#34;2.0.x-scala2.10&#34;, &#34;node_type_id&#34;: &#34;r3.xlarge&#34;,
    // &#34;autoscale&#34; : { &#34;min_workers&#34;: 2, &#34;max_workers&#34;: 50 } }
    Create(ctx context.Context, createCluster CreateCluster) (*CreateClusterResponse, error)
    // Terminates a Spark cluster given its id. The cluster is removed
    // asynchronously. Once the termination has completed, the cluster will be
    // in a ``TERMINATED`` state. If the cluster is already in a ``TERMINATING``
    // or ``TERMINATED`` state, nothing will happen. An example request: ..
    // code:: { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34; }
    Delete(ctx context.Context, deleteCluster DeleteCluster) error
    // Edits the configuration of a cluster to match the provided attributes and
    // size. A cluster can be edited if it is in a ``RUNNING`` or ``TERMINATED``
    // state. If a cluster is edited while in a ``RUNNING`` state, it will be
    // restarted so that the new attributes can take effect. If a cluster is
    // edited while in a ``TERMINATED`` state, it will remain ``TERMINATED``.
    // The next time it is started using the ``clusters/start`` API, the new
    // attributes will take effect. An attempt to edit a cluster in any other
    // state will be rejected with an ``INVALID_STATE`` error code. Clusters
    // created by the Databricks Jobs service cannot be edited. An example
    // request: .. code:: { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34;, &#34;num_workers&#34;:
    // 10, &#34;spark_version&#34;: &#34;3.3.x-scala2.11&#34;, &#34;node_type_id&#34;: &#34;i3.2xlarge&#34; }
    Edit(ctx context.Context, editCluster EditCluster) error
    // Retrieves a list of events about the activity of a cluster. This API is
    // paginated. If there are more events to read, the response includes all
    // the parameters necessary to request the next page of events. An example
    // request: ``/clusters/events?cluster_id=1202-211320-brick1`` An example
    // response: { &#34;events&#34;: [ { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34;,
    // &#34;timestamp&#34;: 1509572145487, &#34;event_type&#34;: &#34;RESTARTING&#34;, &#34;event_details&#34;:
    // { &#34;username&#34;: &#34;admin&#34; } }, ... { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34;,
    // &#34;timestamp&#34;: 1509505807923, &#34;event_type&#34;: &#34;TERMINATING&#34;, &#34;event_details&#34;:
    // { &#34;termination_reason&#34;: { &#34;code&#34;: &#34;USER_REQUEST&#34;, &#34;parameters&#34;: [
    // &#34;username&#34;: &#34;admin&#34; ] } } ], &#34;next_page&#34;: { &#34;cluster_id&#34;:
    // &#34;1202-211320-brick1&#34;, &#34;end_time&#34;: 1509572145487, &#34;order&#34;: &#34;DESC&#34;,
    // &#34;offset&#34;: 50 }, &#34;total_count&#34;: 303 } Example request to retrieve the next
    // page of events
    // ``/clusters/events?cluster_id=1202-211320-brick1&amp;end_time=1509572145487&amp;order=DESC&amp;offset=50``
    Events(ctx context.Context, getEvents GetEvents) (*GetEventsResponse, error)
    // Retrieves the information for a cluster given its identifier. Clusters
    // can be described while they are running, or up to 60 days after they are
    // terminated. An example request:
    // ``/clusters/get?cluster_id=1202-211320-brick1``
    Get(ctx context.Context, getRequest GetRequest) (*ClusterInfo, error)
    // Returns information about all pinned clusters, currently active clusters,
    // up to 70 of the most recently terminated interactive clusters in the past
    // 7 days, and up to 30 of the most recently terminated job clusters in the
    // past 7 days. For example, if there is 1 pinned cluster, 4 active
    // clusters, 45 terminated interactive clusters in the past 7 days, and 50
    // terminated job clusters in the past 7 days, then this API returns the 1
    // pinned cluster, 4 active clusters, all 45 terminated interactive
    // clusters, and the 30 most recently terminated job clusters.
    List(ctx context.Context, listRequest ListRequest) (*ListClustersResponse, error)
    // Returns a list of supported Spark node types. These node types can be
    // used to launch a cluster.
    ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error)
    // Returns a list of availability zones where clusters can be created in
    // (ex: us-west-2a). These zones can be used to launch a cluster.
    ListZones(ctx context.Context) (*ListAvailableZonesResponse, error)
    // Permanently deletes a Spark cluster. This cluster is terminated and
    // resources are asynchronously removed. In addition, users will no longer
    // see permanently deleted clusters in the cluster list, and API users can
    // no longer perform any action on permanently deleted clusters. An example
    // request: .. code:: { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34; }
    PermanentDelete(ctx context.Context, permanentDeleteCluster PermanentDeleteCluster) error
    // Pinning a cluster ensures that the cluster will always be returned by the
    // ListClusters API. Pinning a cluster that is already pinned will have no
    // effect. This API can only be called by workspace admins. An example
    // request: ``/clusters/pin?cluster_id=1202-211320-brick1``
    Pin(ctx context.Context, pinCluster PinCluster) error
    // Resizes a cluster to have a desired number of workers. This will fail
    // unless the cluster is in a ``RUNNING`` state. An example request: ..
    // code:: { &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34;, &#34;num_workers&#34;: 30 }
    Resize(ctx context.Context, resizeCluster ResizeCluster) error
    // Restarts a Spark cluster given its id. If the cluster is not currently in
    // a ``RUNNING`` state, nothing will happen. An example request: .. code:: {
    // &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34; }
    Restart(ctx context.Context, restartCluster RestartCluster) error
    // Returns the list of available Spark versions. These versions can be used
    // to launch a cluster.
    SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error)
    // Starts a terminated Spark cluster given its id. This works similar to
    // `createCluster` except: - The previous cluster id and attributes are
    // preserved. - The cluster starts with the last specified cluster size. -
    // If the previous cluster was an autoscaling cluster, the current cluster
    // starts with the minimum number of nodes. - If the cluster is not
    // currently in a ``TERMINATED`` state, nothing will happen. - Clusters
    // launched to run a job cannot be started. An example request: .. code:: {
    // &#34;cluster_id&#34;: &#34;1202-211320-brick1&#34; }
    Start(ctx context.Context, startCluster StartCluster) error
    // Unpinning a cluster will allow the cluster to eventually be removed from
    // the ListClusters API. Unpinning a cluster that is not pinned will have no
    // effect. This API can only be called by workspace admins. An example
    // request: ``/clusters/unpin?cluster_id=1202-211320-brick1``
    Unpin(ctx context.Context, unpinCluster UnpinCluster) error
}
