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

	// +kubebuilder:validation:Required
	// Items is the list of GCPKubernetesClusters
	Items []GCPKubernetesCluster `json:"items"`
}

// GCPKubernetesCluster is the Schema for the gcpkubernetesclusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gcpkubernetesclusters,shortName=gkc,singular=gcpkubernetescluster
type GCPKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GCPKubernetesCluster
	Spec GCPKubernetesClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GCPKubernetesCluster
	Status GCPKubernetesClusterStatus `json:"status"`
}

// GCPKubernetesClusterSpec defines the desired state of GCPKubernetesCluster
type GCPKubernetesClusterSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GCP Kubernetes cluster
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	// InitialNodeCount is the number of nodes to create in the GCP Kubernetes cluster
	InitialNodeCount int32 `json:"initialNodeCount"`
}

// GCPKubernetesClusterStatus defines the observed state of GCPKubernetesCluster
type GCPKubernetesClusterStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GCP Kubernetes cluster
	Phase string `json:"phase"`
}
