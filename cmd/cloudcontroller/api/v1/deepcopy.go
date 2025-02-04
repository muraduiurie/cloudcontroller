package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ---------------------------------------------------
// GKECluster
// ---------------------------------------------------
func (in *GKECluster) DeepCopyInto(out *GKECluster) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GKEClusterSpec{
		Name:             in.Spec.Name,
		InitialNodeCount: in.Spec.InitialNodeCount,
	}
}

func (in *GKECluster) DeepCopyObject() runtime.Object {
	out := GKECluster{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GKEClusterList) DeepCopyObject() runtime.Object {
	out := GKEClusterList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GKECluster, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

// ---------------------------------------------------
// GKEInstance
// ---------------------------------------------------
func (in *GKEInstance) DeepCopyInto(out *GKEInstance) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GKEInstanceSpec{
		Name: in.Spec.Name,
	}
}

func (in *GKEInstance) DeepCopyObject() runtime.Object {
	out := GKEInstance{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GKEInstanceList) DeepCopyObject() runtime.Object {
	out := GKEInstanceList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GKEInstance, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

// ---------------------------------------------------
// GKENetwork
// ---------------------------------------------------
func (in *GKENetwork) DeepCopyInto(out *GKENetwork) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GKENetworkSpec{
		Name:                  in.Spec.Name,
		AutoCreateSubnetworks: in.Spec.AutoCreateSubnetworks,
	}
}

func (in *GKENetwork) DeepCopyObject() runtime.Object {
	out := GKENetwork{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GKENetworkList) DeepCopyObject() runtime.Object {
	out := GKENetworkList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GKENetwork, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
