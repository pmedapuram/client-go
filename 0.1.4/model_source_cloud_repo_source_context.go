/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A CloudRepoSourceContext denotes a particular revision in a Google Cloud Source Repo.
type SourceCloudRepoSourceContext struct {
	// The ID of the repo.
	RepoId *SourceRepoId `json:"repo_id,omitempty"`
	// A revision ID.
	RevisionId string `json:"revision_id,omitempty"`
	// An alias, which may be a branch or tag.
	AliasContext *SourceAliasContext `json:"alias_context,omitempty"`
}
