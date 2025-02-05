// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCPNetworkList contains a list of GCPNetwork
// +kubebuilder:object:root=true
type GCPNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Items is the list of GCPNetworks
	Items []GCPNetwork `json:"items"`
}

// GCPNetwork is the Schema for the gcpnetworks API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gcpnetworks,shortName=gn,singular=gcpnetwork
type GCPNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GCPNetwork
	Spec GCPNetworkSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GCPNetwork
	Status GCPNetworkStatus `json:"status"`
}

// GCPNetworkSpec defines the desired state of GCPNetwork
type GCPNetworkSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GCP network
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	// AutoCreateSubnetworks
	AutoCreateSubnetworks bool `json:"autoCreateSubnetworks"`
}

// GCPNetworkStatus defines the observed state of GCPNetwork
type GCPNetworkStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GCP cluster
	Phase string `json:"phase"`
}
