package models

type (
	Repository struct {
		ID                        int             `json:"id"`
		Owner                     User            `json:"owner"`
		Name                      string          `json:"name"`
		FullName                  string          `json:"full_name"`
		Description               string          `json:"description"`
		Empty                     bool            `json:"empty"`
		Private                   bool            `json:"private"`
		Pinned                    bool            `json:"pinned"`
		IsStarred                 bool            `json:"is_starred"`
		Fork                      bool            `json:"fork"`
		Template                  bool            `json:"template"`
		Parent                    interface{}     `json:"parent"`
		Mirror                    bool            `json:"mirror"`
		Size                      int             `json:"size"`
		Language                  string          `json:"language"`
		LanguagesURL              string          `json:"languages_url"`
		HTMLURL                   string          `json:"html_url"`
		SSHURL                    string          `json:"ssh_url"`
		CloneURL                  string          `json:"clone_url"`
		OriginalURL               string          `json:"original_url"`
		Website                   string          `json:"website"`
		StarsCount                int             `json:"stars_count"`
		ForksCount                int             `json:"forks_count"`
		WatchersCount             int             `json:"watchers_count"`
		OpenIssuesCount           int             `json:"open_issues_count"`
		OpenPRCounter             int             `json:"open_pr_counter"`
		ReleaseCounter            int             `json:"release_counter"`
		DefaultBranch             string          `json:"default_branch"`
		Archived                  bool            `json:"archived"`
		CreatedAt                 string          `json:"created_at"`
		UpdatedAt                 string          `json:"updated_at"`
		Permissions               Permissions     `json:"permissions"`
		HasIssues                 bool            `json:"has_issues"`
		InternalTracker           InternalTracker `json:"internal_tracker"`
		HasWiki                   bool            `json:"has_wiki"`
		HasPullRequests           bool            `json:"has_pull_requests"`
		HasProjects               bool            `json:"has_projects"`
		IgnoreWhitespaceConflicts bool            `json:"ignore_whitespace_conflicts"`
		AllowMergeCommits         bool            `json:"allow_merge_commits"`
		AllowRebase               bool            `json:"allow_rebase"`
		AllowRebaseExplicit       bool            `json:"allow_rebase_explicit"`
		AllowSquashMerge          bool            `json:"allow_squash_merge"`
		DefaultMergeStyle         string          `json:"default_merge_style"`
		AvatarURL                 string          `json:"avatar_url"`
		Internal                  bool            `json:"internal"`
		MirrorInterval            string          `json:"mirror_interval"`
		MirrorUpdated             string          `json:"mirror_updated"`
		RepoTransfer              interface{}     `json:"repo_transfer"`
	}
	Permissions struct {
		Admin       bool `json:"admin"`
		Push        bool `json:"push"`
		Pull        bool `json:"pull"`
		IsMeInvited bool `json:"is_me_invited"`
	}
	InternalTracker struct {
		EnableTimeTracker                bool `json:"enable_time_tracker"`
		AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
		EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
	}
)
