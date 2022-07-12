/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GlsstorageSpec defines the desired state of Glsstorage
type SC struct {
	stype          string `json:"stype,omitempty"`
	gluster_host   string `json:"gluster-host,omitempty"`
	gluster_volume string `json:"gluster-volume,omitempty"`
	nvmetype       string `json:"nvme-type,omitempty"`
	nvme_host      string `json:"nvme-host,omitempty"`
	nvme_fs        string `json:"nvme-fs,omitempty"`
}

type PVC struct {
	accessMode string `json:"accessmode,omitempty"`
	limit      string `json:"limit,omitempty"`
}
type GlsstorageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Glsstorage. Edit glsstorage_types.go to remove/update
	Name string `json:"name,omitempty"`
	Sc   SC     `json:"sc,omitempty"`
	Pvc  PVC    `json:"pvc,omitempty"`
}

// GlsstorageStatus defines the observed state of Glsstorage
type GlsstorageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Glsstorage is the Schema for the glsstorages API
type Glsstorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlsstorageSpec   `json:"spec,omitempty"`
	Status GlsstorageStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GlsstorageList contains a list of Glsstorage
type GlsstorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Glsstorage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Glsstorage{}, &GlsstorageList{})
}
