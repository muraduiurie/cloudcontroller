package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

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
