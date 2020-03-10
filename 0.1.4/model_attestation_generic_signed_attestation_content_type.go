/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// AttestationGenericSignedAttestationContentType : Type of the attestation plaintext that was signed.   - CONTENT_TYPE_UNSPECIFIED: `ContentType` is not set.  - SIMPLE_SIGNING_JSON: Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted in `plaintext` is a JSON blob conforming to the linked schema.
type AttestationGenericSignedAttestationContentType string

// List of attestationGenericSignedAttestationContentType
const (
	CONTENT_TYPE_UNSPECIFIED_AttestationGenericSignedAttestationContentType AttestationGenericSignedAttestationContentType = "CONTENT_TYPE_UNSPECIFIED"
	SIMPLE_SIGNING_JSON_AttestationGenericSignedAttestationContentType AttestationGenericSignedAttestationContentType = "SIMPLE_SIGNING_JSON"
)
