package github

import (
	"encoding/json"
	"path"
)

const reposPath = "repos"

type Repository struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"full_name"`
	Description string `json:"description"`
	HTMLUrl     string `json:"html_url"`

	Targets []string
}

func GetRepository(repo string) (*Repository, error) {
	r := &Repository{}

	body, err := get("https://" + path.Join(hostname, reposPath, repo))
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
