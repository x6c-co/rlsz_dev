package github

import (
	"encoding/json"
	"fmt"
	"path"
	"time"
)

type Release struct {
	ID          uint   `json:"id"`
	TagName     string `json:"tag_name"`
	Body        string `json:"body"`
	Draft       bool   `json:"draft"`
	PreRelease  bool   `json:"prerelease"`
	CreatedAt   string `json:"created_at"`
	PublishedAt string `json:"published_at"`
	HTMLURL     string `json:"html_url"`
	NodeID      string `json:"node_id"`

	HTML      string
	CardClass string

	// Specifies the commitish value that determines where the Git tag is
	// created from. Can be any branch or commit SHA. Unused if the Git tag
	// already exists. Default: the repository’s default branch (usually master).
	TargetCommitish string `json:"target_commitish"`
}

func (r *Repository) GetReleases() ([]*Release, error) {
	releases := []*Release{}

	page := 1
	for {
		rls := []*Release{}
		body, err := get("https://" + path.Join(hostname, reposPath, r.Path, "releases") + "?per_page=100&page=" + fmt.Sprintf("%d", page))
		if err != nil {
			return releases, err
		}

		err = json.Unmarshal(body, &rls)
		if err != nil {
			return releases, err
		}

		releases = append(releases, rls...)

		if len(rls) < 100 {
			break
		}

		page++
	}

	return releases, nil
}

func (r *Release) Process() {
	format := "2006-01-02"
	t, _ := time.Parse(time.RFC3339, r.CreatedAt)
	r.CreatedAt = t.Format(format)

	t, _ = time.Parse(time.RFC3339, r.PublishedAt)
	r.PublishedAt = t.Format(format)
}

func (r *Release) ParseMarkdown() {
	r.HTML = parseMarkdown([]byte(r.Body))
}
