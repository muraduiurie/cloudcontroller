package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCPKubernetesClusterList contains a list of GCPKubernetesCluster.
// +kubebuilder:object:root=true
type GCPKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GCPKubernetesCluster `json:"items"`
}

// GCPKubernetesCluster is the Schema for the gcpkubernetesclusters API.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gcpkubernetesclusters,shortName=gkc,singular=gcpkubernetescluster
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.phase"
type GCPKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPKubernetesClusterSpec   `json:"spec"`
	Status GCPKubernetesClusterStatus `json:"status,omitempty"`
}

type GCPKubernetesClusterSpec struct {
	// ClusterName of the GCP Kubernetes cluster.
	// +kubebuilder:validation:Required
	ClusterName string `json:"clusterName"`
	// InitialNodeCount defines the number of nodes to create in this cluster.
	// +kubebuilder:validation:Required
	InitialNodeCount int64 `json:"initialNodeCount"`
	// Zone in which the GCP Kubernetes cluster resides.
	// +kubebuilder:validation:Required
	Zone string `json:"zone"`
	// Autopilot enables the Autopilot mode for the cluster.
	// +kubebuilder:validation:Optional
	Autopilot bool `json:"autopilot,omitempty"`
	// ClusterIpv4Cidr defines the IP address range of the container pods in this cluster.
	// +kubebuilder:validation:Optional
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`
	// Description of this cluster.
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// InitialClusterVersion defines the initial Kubernetes version for this cluster.
	// +kubebuilder:validation:Optional
	InitialClusterVersion string `json:"initialClusterVersion,omitempty"`
	// Network of the Google Compute Engine network which the cluster is connected.
	// +kubebuilder:validation:Optional
	Network string `json:"network,omitempty"`
	// NodePools associated with this cluster.
	// +kubebuilder:validation:Optional
	NodePools []*NodePool `json:"nodePools,omitempty"`
	// Subnetwork of the Google Compute Engine subnetwork connected.
	// +kubebuilder:validation:Optional
	Subnetwork string `json:"subnetwork,omitempty"`
}

type NodePool struct {
	// NodeName of the node pool.
	// +kubebuilder:validation:Required
	NodeName string `json:"nodeName,omitempty"`
	// Version of Kubernetes running on this NodePool's nodes.
	// +kubebuilder:validation:Optional
	Version string `json:"version,omitempty"`
	// Config defines the node configuration of the pool.
	// +kubebuilder:validation:Optional
	Config *NodeConfig `json:"configmap,omitempty"`
	// InitialNodeCount defines the initial node count for the pool.
	// +kubebuilder:validation:Required
	InitialNodeCount int64 `json:"nodeCount,omitempty"`
}

type NodeConfig struct {
	// DiskSizeGb defines the size of the disk attached to each node, specified in GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`
	// DiskType is the type of the disk attached to each node.
	// +kubebuilder:validation:Optional
	DiskType string `json:"diskType,omitempty"`
	// ImageType to use for this node.
	// +kubebuilder:validation:Optional
	ImageType string `json:"imageType,omitempty"`
	// Labels is the map of Kubernetes labels (key/value pairs) to be applied to each node.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// MachineType is the name of a Google Compute Engine machine type.
	// +kubebuilder:validation:Optional
	MachineType string `json:"machineType,omitempty"`
}

type ClusterStatus string

const (
	ClusterStatusUnspecified  ClusterStatus = "STATUS_UNSPECIFIED"
	ClusterStatusProvisioning ClusterStatus = "PROVISIONING"
	ClusterStatusRunning      ClusterStatus = "RUNNING"
	ClusterStatusReconciling  ClusterStatus = "RECONCILING"
	ClusterStatusStopping     ClusterStatus = "STOPPING"
	ClusterStatusError        ClusterStatus = "ERROR"
	ClusterStatusDegraded     ClusterStatus = "DEGRADED"
)

type GCPKubernetesClusterStatus struct {
	// Phase is the current state of the GCP Kubernetes cluster
	// +kubebuilder:validation:Optional
	Phase ClusterStatus `json:"phase,omitempty"`
}
