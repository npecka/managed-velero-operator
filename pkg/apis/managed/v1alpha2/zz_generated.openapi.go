//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.StorageBucket":       schema_pkg_apis_managed_v1alpha2_StorageBucket(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstall":       schema_pkg_apis_managed_v1alpha2_VeleroInstall(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallSpec":   schema_pkg_apis_managed_v1alpha2_VeleroInstallSpec(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallStatus": schema_pkg_apis_managed_v1alpha2_VeleroInstallStatus(ref),
	}
}

func schema_pkg_apis_managed_v1alpha2_StorageBucket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageBucket contains details of the storage bucket for backups",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the storage bucket created to store Velero backup details",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"provisioned": {
						SchemaProps: spec.SchemaProps{
							Description: "Provisioned is true once the bucket has been initially provisioned.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastSyncTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "LastSyncTimestamp is the time that the bucket policy was last synced.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"provisioned"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_managed_v1alpha2_VeleroInstall(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VeleroInstall is the Schema for the veleroinstalls API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallSpec", "github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.VeleroInstallStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_managed_v1alpha2_VeleroInstallSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VeleroInstallSpec defines the desired state of Velero",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_managed_v1alpha2_VeleroInstallStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VeleroInstallStatus defines the observed state of Velero",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageBucket": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageBucket contains details of the storage bucket for backups",
							Ref:         ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.StorageBucket"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha2.StorageBucket"},
	}
}
