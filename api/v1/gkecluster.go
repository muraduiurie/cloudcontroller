package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GKEClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []GKECluster `json:"items"`
}

type GKECluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GKEClusterSpec `json:"spec"`
}

type GKEClusterSpec struct {
	Name             string `json:"name"`
	InitialNodeCount int32  `json:"initialNodeCount"`
}
