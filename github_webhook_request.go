package util

type GithubWebhookRequest struct {
	Zen        string            `json:"zen"`
	HookID     int               `json:"hook_id"`
	Hook       *GithubWebhook    `json:"hook"`
	Repository *GithubRepository `json:"repository"`
}
