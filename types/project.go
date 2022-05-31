package types

import "time"

type ProjectVisibility string

const (
	ProjectVisibility_Private  = ProjectVisibility("private")
	ProjectVisibility_Public   = ProjectVisibility("public")
	ProjectVisibility_Internal = ProjectVisibility("internal")
)

type ProjectList []*Project

type Project struct {
	ID                                        int               `json:"id,omitempty"`
	Description                               interface{}       `json:"description,omitempty"`
	DefaultBranch                             string            `json:"default_branch,omitempty"`
	Visibility                                ProjectVisibility `json:"visibility,omitempty"`
	SSHURLToRepo                              string            `json:"ssh_url_to_repo,omitempty"`
	HTTPURLToRepo                             string            `json:"http_url_to_repo,omitempty"`
	WebURL                                    string            `json:"web_url,omitempty"`
	ReadmeURL                                 string            `json:"readme_url,omitempty"`
	TagList                                   []string          `json:"tag_list,omitempty"`
	Topics                                    []string          `json:"topics,omitempty"`
	Owner                                     Owner             `json:"owner,omitempty"`
	Name                                      string            `json:"name,omitempty"`
	NameWithNamespace                         string            `json:"name_with_namespace,omitempty"`
	Path                                      string            `json:"path,omitempty"`
	PathWithNamespace                         string            `json:"path_with_namespace,omitempty"`
	IssuesEnabled                             bool              `json:"issues_enabled,omitempty"`
	OpenIssuesCount                           int               `json:"open_issues_count,omitempty"`
	MergeRequestsEnabled                      bool              `json:"merge_requests_enabled,omitempty"`
	JobsEnabled                               bool              `json:"jobs_enabled,omitempty"`
	WikiEnabled                               bool              `json:"wiki_enabled,omitempty"`
	SnippetsEnabled                           bool              `json:"snippets_enabled,omitempty"`
	CanCreateMergeRequestIn                   bool              `json:"can_create_merge_request_in,omitempty"`
	ResolveOutdatedDiffDiscussions            bool              `json:"resolve_outdated_diff_discussions,omitempty"`
	ContainerRegistryEnabled                  bool              `json:"container_registry_enabled,omitempty"`
	ContainerRegistryAccessLevel              string            `json:"container_registry_access_level,omitempty"`
	SecurityAndComplianceAccessLevel          string            `json:"security_and_compliance_access_level,omitempty"`
	CreatedAt                                 time.Time         `json:"created_at,omitempty"`
	LastActivityAt                            time.Time         `json:"last_activity_at,omitempty"`
	CreatorID                                 int               `json:"creator_id,omitempty"`
	Namespace                                 Namespace         `json:"namespace,omitempty"`
	ImportStatus                              string            `json:"import_status,omitempty"`
	Archived                                  bool              `json:"archived,omitempty"`
	AvatarURL                                 string            `json:"avatar_url,omitempty"`
	SharedRunnersEnabled                      bool              `json:"shared_runners_enabled,omitempty"`
	ForksCount                                int               `json:"forks_count,omitempty"`
	StarCount                                 int               `json:"star_count,omitempty"`
	RunnersToken                              string            `json:"runners_token,omitempty"`
	CiDefaultGitDepth                         int               `json:"ci_default_git_depth,omitempty"`
	CiForwardDeploymentEnabled                bool              `json:"ci_forward_deployment_enabled,omitempty"`
	PublicJobs                                bool              `json:"public_jobs,omitempty"`
	SharedWithGroups                          []interface{}     `json:"shared_with_groups,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool              `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	AllowMergeOnSkippedPipeline               bool              `json:"allow_merge_on_skipped_pipeline,omitempty"`
	RestrictUserDefinedVariables              bool              `json:"restrict_user_defined_variables,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool              `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	RemoveSourceBranchAfterMerge              bool              `json:"remove_source_branch_after_merge,omitempty"`
	RequestAccessEnabled                      bool              `json:"request_access_enabled,omitempty"`
	MergeMethod                               string            `json:"merge_method,omitempty"`
	SquashOption                              string            `json:"squash_option,omitempty"`
	AutocloseReferencedIssues                 bool              `json:"autoclose_referenced_issues,omitempty"`
	SuggestionCommitMessage                   interface{}       `json:"suggestion_commit_message,omitempty"`
	MergeCommitTemplate                       interface{}       `json:"merge_commit_template,omitempty"`
	SquashCommitTemplate                      interface{}       `json:"squash_commit_template,omitempty"`
	MarkedForDeletionAt                       string            `json:"marked_for_deletion_at,omitempty"`
	MarkedForDeletionOn                       string            `json:"marked_for_deletion_on,omitempty"`
	Statistics                                Statistics        `json:"statistics,omitempty"`
	ContainerRegistryImagePrefix              string            `json:"container_registry_image_prefix,omitempty"`
	Links                                     Links             `json:"_links,omitempty"`
}
type Owner struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
type Namespace struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	Kind     string `json:"kind,omitempty"`
	FullPath string `json:"full_path,omitempty"`
}
type Statistics struct {
	CommitCount           int `json:"commit_count,omitempty"`
	StorageSize           int `json:"storage_size,omitempty"`
	RepositorySize        int `json:"repository_size,omitempty"`
	WikiSize              int `json:"wiki_size,omitempty"`
	LfsObjectsSize        int `json:"lfs_objects_size,omitempty"`
	JobArtifactsSize      int `json:"job_artifacts_size,omitempty"`
	PipelineArtifactsSize int `json:"pipeline_artifacts_size,omitempty"`
	PackagesSize          int `json:"packages_size,omitempty"`
	SnippetsSize          int `json:"snippets_size,omitempty"`
	UploadsSize           int `json:"uploads_size,omitempty"`
}
type Links struct {
	Self          string `json:"self,omitempty"`
	Issues        string `json:"issues,omitempty"`
	MergeRequests string `json:"merge_requests,omitempty"`
	RepoBranches  string `json:"repo_branches,omitempty"`
	Labels        string `json:"labels,omitempty"`
	Events        string `json:"events,omitempty"`
	Members       string `json:"members,omitempty"`
	ClusterAgents string `json:"cluster_agents,omitempty"`
}
