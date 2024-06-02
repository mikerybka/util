package util

type GithubRepository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	FullName string `json:"full_name"`
}
