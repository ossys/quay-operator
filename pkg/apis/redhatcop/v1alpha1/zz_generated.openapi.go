// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystem":       schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystem(ref),
		"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemSpec":   schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystemSpec(ref),
		"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemStatus": schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystemStatus(ref),
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystem(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayEcosystem is the Schema for the quayecosystems API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
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
							Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemSpec", "github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystemSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayEcosystemSpec defines the desired state of QuayEcosystem",
				Properties: map[string]spec.Schema{
					"quay": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Quay"),
						},
					},
					"redis": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Redis"),
						},
					},
					"clair": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Clair"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Clair", "github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Quay", "github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.Redis"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayEcosystemStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayEcosystemStatus defines the observed state of QuayEcosystem",
				Properties: map[string]spec.Schema{
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hostname": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemCondition"),
									},
								},
							},
						},
					},
					"setupComplete": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/quay-operator/pkg/apis/redhatcop/v1alpha1.QuayEcosystemCondition"},
	}
}
