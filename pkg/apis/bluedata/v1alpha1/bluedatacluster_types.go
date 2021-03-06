


package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the BlueDataCluster resource schema definition
// as a go struct

type BlueDataLabel struct {
   Name string        `json:"name,omitempty"`
   Description string `json:"description,omitemtpy"`
}

type BlueDataNodegroupRoleConfig struct {
    RoleId    string `json:"role_id,omitempty"`
    NodeCount int    `json:"node_count,omitempty"`
    Flavor    string `json:"flavor,omitempty"`
}

type BlueDataNodegroup struct {
    CatalogEntryDistroId  string                        `json:"catalog_entry_distro_id,omitempty"`
    RoleConfigs           []BlueDataNodegroupRoleConfig `json:"role_configs,omitempty"`
}

// BlueDataClusterSpec defines the desired state of BlueDataCluster
type BlueDataClusterSpec struct {
    // INSERT YOUR CODE HERE - define desired state schema
    Label           BlueDataLabel     `json:"label,omitempty"`
    //Debug           bool              `json:"debug,omitempty"`
    Nodegroup       BlueDataNodegroup `json:"nodegroup,omitempty"`
    TwoPhaseDelete  bool              `json:"two_phase_delete,omitempty"`
}

// BlueDataClusterStatus defines the observed state of BlueDataCluster
type BlueDataClusterStatus struct {
    // INSERT YOUR CODE HERE - define observed state schema
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlueDataCluster
// +k8s:openapi-gen=true
// +resource:path=bluedataclusters
type BlueDataCluster struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   BlueDataClusterSpec  `json:"spec,omitempty"`
    Status BlueDataClusterStatus `json:"status,omitempty"`
}
