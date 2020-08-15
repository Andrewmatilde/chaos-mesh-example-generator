
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true

// TestChaos is the Schema for the Testchaos API
type TestChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func (h TestChaos) DeepCopyObject() runtime.Object {
	panic("implement me")
}

// +kubebuilder:object:root=true

// TestChaosList contains a list of TestChaos
type TestChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestChaos `json:"items"`
}

func (h TestChaosList) DeepCopyObject() runtime.Object {
	panic("implement me")
}

func init() {
	SchemeBuilder.Register(&TestChaos{}, &TestChaosList{})
}
