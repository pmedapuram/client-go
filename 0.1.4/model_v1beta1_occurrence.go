/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

import (
	"time"
)

// An instance of an analysis type that has been found on a resource.
type V1beta1Occurrence struct {
	// Output only. The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
	Name string `json:"name,omitempty"`
	// Required. Immutable. The resource for which the occurrence applies.
	Resource *V1beta1Resource `json:"resource,omitempty"`
	// Required. Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
	NoteName string `json:"note_name,omitempty"`
	// Output only. This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
	Kind *V1beta1NoteKind `json:"kind,omitempty"`
	// A description of actions that can be taken to remedy the note.
	Remediation string `json:"remediation,omitempty"`
	// Output only. The time this occurrence was created.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Output only. The time this occurrence was last updated.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Describes a security vulnerability.
	Vulnerability *V1beta1vulnerabilityDetails `json:"vulnerability,omitempty"`
	// Describes a verifiable build.
	Build *V1beta1buildDetails `json:"build,omitempty"`
	// Describes how this resource derives from the basis in the associated note.
	DerivedImage *V1beta1imageDetails `json:"derived_image,omitempty"`
	// Describes the installation of a package on the linked resource.
	Installation *V1beta1packageDetails `json:"installation,omitempty"`
	// Describes the deployment of an artifact on a runtime.
	Deployment *V1beta1deploymentDetails `json:"deployment,omitempty"`
	// Describes when a resource was discovered.
	Discovered *V1beta1discoveryDetails `json:"discovered,omitempty"`
	// Describes an attestation of an artifact.
	Attestation *V1beta1attestationDetails `json:"attestation,omitempty"`
}
