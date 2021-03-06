/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// This represents how a particular software package may be installed on a system.
type PackageInstallation struct {
	// Output only. The name of the installed package.
	Name string `json:"name,omitempty"`
	// Required. All of the places within the filesystem versions of this package have been found.
	Location []V1beta1packageLocation `json:"location,omitempty"`
}
