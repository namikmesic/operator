/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SmartContractSpec defines the desired state of SmartContract
type SmartContractSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SmartContract. Edit smartcontract_types.go to remove/update
	Bytecode  string   `json:"bytecode"`
	Arguments []string `json:"arguments"`
	GasLimit  uint64   `json:"gasLimit"`
	GasPrice  uint64   `json:"gasPrice"`
}

// SmartContractStatus defines the observed state of SmartContract
type SmartContractStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Address string `json:"address,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SmartContract is the Schema for the smartcontracts API
type SmartContract struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SmartContractSpec   `json:"spec,omitempty"`
	Status SmartContractStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SmartContractList contains a list of SmartContract
type SmartContractList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SmartContract `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SmartContract{}, &SmartContractList{})
}
