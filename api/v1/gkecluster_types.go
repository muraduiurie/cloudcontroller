// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEClusterList contains a list of GKECluster
// +kubebuilder:object:root=true
type GKEClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Items is the list of GKEClusters
	Items []GKECluster `json:"items"`
}

// GKECluster is the Schema for the gkeclusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gkeclusters,shortName=gke,singular=gkecluster
type GKECluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GKECluster
	Spec GKEClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GKECluster
	Status GKEClusterStatus `json:"status"`
}

// GKEClusterSpec defines the desired state of GKECluster
type GKEClusterSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GKE cluster
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	// InitialNodeCount is the number of nodes to create in the GKE cluster
	InitialNodeCount int32 `json:"initialNodeCount"`
}

// GKEClusterStatus defines the observed state of GKECluster
type GKEClusterStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GKE cluster
	Phase string `json:"phase"`
}
