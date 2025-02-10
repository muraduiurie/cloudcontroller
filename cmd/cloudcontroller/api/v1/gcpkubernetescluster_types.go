// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCPKubernetesClusterList contains a list of GCPKubernetesCluster
// +kubebuilder:object:root=true
type GCPKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GCPKubernetesCluster `json:"items"`
}

// GCPKubernetesCluster is the Schema for the gcpkubernetesclusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gcpkubernetesclusters,shortName=gkc,singular=gcpkubernetescluster
type GCPKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPKubernetesClusterSpec   `json:"spec"`
	Status GCPKubernetesClusterStatus `json:"status"`
}

// GCPKubernetesClusterSpec defines the desired state of GCPKubernetesCluster
type GCPKubernetesClusterSpec struct {
	// Name of the GCP Kubernetes cluster
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// InitialNodeCount defines the number of nodes to create in this cluster. You must
	// ensure that your Compute Engine resource quota
	// +kubebuilder:validation:Required
	InitialNodeCount int32 `json:"initialNodeCount"`
	// Zone in which the GCP Kubernetes cluster resides
	// +kubebuilder:validation:Required
	Zone string `json:"zone"`
	// Autopilot enables the Autopilot mode for the cluster. By default it is disabled.
	// +kubebuilder:validation:Optional
	Autopilot bool `json:"autopilot,omitempty"`
	// ClusterIpv4Cidr defines the IP address range of the container pods in this cluster,
	// in CIDR (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically chosen
	// or specify a `/14` block in `10.0.0.0/8`.
	// +kubebuilder:validation:Optional
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`
	// Description of this cluster.
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// InitialClusterVersion defines the initial Kubernetes version for this cluster.
	// Valid versions are those found in validMasterVersions returned by
	// getServerConfig. The version can be upgraded over time; such upgrades are
	// reflected in currentMasterVersion and currentNodeVersion. Users may specify
	// either explicit versions offered by Kubernetes Engine or version aliases,
	// which have the following behavior: - "latest": picks the highest valid
	// Kubernetes version - "1.X": picks the highest valid patch+gke.N patch in the
	// 1.X version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y
	// version - "1.X.Y-gke.N": picks an explicit Kubernetes version - "","-":
	// picks the default Kubernetes version
	// +kubebuilder:validation:Optional
	InitialClusterVersion string `json:"initialClusterVersion,omitempty"`
	// Network of the Google Compute Engine network
	// (https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to
	// which the cluster is connected. If left unspecified, the `default` network
	// will be used.
	// +kubebuilder:validation:Optional
	Network string `json:"network,omitempty"`
	// NodePools associated with this cluster. This field should
	// not be set if "node_config" or "initial_node_count" are specified.
	// +kubebuilder:validation:Optional
	NodePools []*NodePool `json:"nodePools,omitempty"`
	// Subnetwork of the Google Compute Engine subnetwork
	// (https://cloud.google.com/compute/docs/subnetworks) to which the cluster is
	// connected.
	// +kubebuilder:validation:Optional
	Subnetwork string `json:"subnetwork,omitempty"`
}

// NodePool defines the node pool configuration
type NodePool struct {
	// Name of the node pool.
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`
	// Version of Kubernetes running on this NodePool's nodes. If
	// unspecified, it defaults as described here
	// (https://cloud.google.com/kubernetes-engine/versioning#specifying_node_version).
	// +kubebuilder:validation:Optional
	Version string `json:"version,omitempty"`
	// Config defines the node configuration of the pool.
	// +kubebuilder:validation:Optional
	Config *NodeConfig `json:"config,omitempty"`
	// InitialNodeCount defines the initial node count for the pool. You must ensure that
	// your Compute Engine resource quota (https://cloud.google.com/compute/quotas)
	// is sufficient for this number of instances. You must also have available
	// firewall and routes quota.
	// +kubebuilder:validation:Required
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`
}

// NodeConfig defines the node configuration
type NodeConfig struct {
	// DiskSizeGb defines the size of the disk attached to each node, specified in GB. The
	// smallest allowed disk size is 10GB. If unspecified, the default disk size is
	// 100GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`
	// DiskType is the type of the disk attached to each node (e.g. 'pd-standard',
	// 'pd-ssd' or 'pd-balanced') If unspecified, the default disk type is
	// 'pd-standard'
	// +kubebuilder:validation:Optional
	DiskType string `json:"diskType,omitempty"`
	// ImageType to use for this node. Note that for a given image
	// type, the latest version of it will be used. Please see
	// https://cloud.google.com/kubernetes-engine/docs/concepts/node-images for
	// available image types.
	// +kubebuilder:validation:Optional
	ImageType string `json:"imageType,omitempty"`
	// Labels is the map of Kubernetes labels (key/value pairs) to be applied to each
	// node. These will added in addition to any default label(s) that Kubernetes
	// may apply to the node. In case of conflict in label keys, the applied set
	// may differ depending on the Kubernetes version -- it's best to assume the
	// behavior is undefined and conflicts should be avoided. For more information,
	// including usage and the valid values, see:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// MachineType is the name of a Google Compute Engine machine type
	// (https://cloud.google.com/compute/docs/machine-types) If unspecified, the
	// default machine type is `e2-medium`.
	// +kubebuilder:validation:Optional
	MachineType string `json:"machineType,omitempty"`
}

// GCPKubernetesClusterStatus defines the observed state of GCPKubernetesCluster
type GCPKubernetesClusterStatus struct {
	// Phase is the current state of the GCP Kubernetes cluster
	// +kubebuilder:validation:Optional
	Phase string `json:"phase"`
}
