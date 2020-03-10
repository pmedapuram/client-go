/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// This represents a particular channel of distribution for a given package. E.g., Debian's jessie-backports dpkg mirror.
type PackageDistribution struct {
	// Required. The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
	CpeUri string `json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were built.
	Architecture *PackageArchitecture `json:"architecture,omitempty"`
	// The latest available version of this package in this distribution channel.
	LatestVersion *PackageVersion `json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description string `json:"description,omitempty"`
}
