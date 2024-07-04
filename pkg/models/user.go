package models

type (
	User struct {
		ID                int    `json:"id"`
		Login             string `json:"login"`
		FullName          string `json:"full_name"`
		Email             string `json:"email"`
		AvatarURL         string `json:"avatar_url"`
		Language          string `json:"language"`
		IsAdmin           bool   `json:"is_admin"`
		LastLogin         string `json:"last_login"`
		Created           string `json:"created"`
		Restricted        bool   `json:"restricted"`
		Active            bool   `json:"active"`
		ProhibitLogin     bool   `json:"prohibit_login"`
		Location          string `json:"location"`
		Website           string `json:"website"`
		Description       string `json:"description"`
		Visibility        string `json:"visibility"`
		FollowersCount    int    `json:"followers_count"`
		FollowingCount    int    `json:"following_count"`
		StarredReposCount int    `json:"starred_repos_count"`
		Permission        string `json:"permission"`
		Username          string `json:"username"`
	}
	CreateUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUser struct {
		Username string `json:"login_name"`
		Active   bool   `json:"active"`
	}
)
