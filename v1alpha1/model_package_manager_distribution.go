/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type PackageManagerDistribution struct {
	// The cpe_uri in [cpe format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
	CpeUri string `json:"cpe_uri,omitempty"`
	Architecture *PackageManagerArchitecture `json:"architecture,omitempty"`
	// The latest available version of this package in this distribution channel.
	LatestVersion *VulnerabilityTypeVersion `json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description string `json:"description,omitempty"`
}
