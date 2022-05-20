package project

import (
	"sort"
	"strings"

	"github.com/levidurfee/rlsz_dev/pkg/config"
	"github.com/levidurfee/rlsz_dev/pkg/github"
)

// Project represents a project a project on GitHub.
type Project struct {
	Name  string `json:"name"`
	Path  string `json:"-"`
	Stamp string `json:"-"`

	Repository *github.Repository `json:"repository"`
	Releases   []*github.Release  `json:"releases"`
	Majors     []string           `json:"majors"`

	Index bool
}

// Build calls the other methods on the Project struct that do API calls and
// populate the fields.
func (p *Project) Build() error {
	err := p.getRepository()
	if err != nil {
		return err
	}

	err = p.getReleases()
	if err != nil {
		return err
	}

	p.getMajors()

	p.Stamp = config.Stamp

	return nil
}

// getRepository calls the github API to get information about the repository.
func (p *Project) getRepository() error {
	repo, err := github.GetRepository(p.Path)
	if err != nil {
		return err
	}

	p.Repository = repo

	return nil
}

// getReleases calls the github API and grabs all the releases.
func (p *Project) getReleases() error {
	releases, err := p.Repository.GetReleases()
	if err != nil {
		return err
	}

	for _, release := range releases {
		release.ParseMarkdown()
		release.Process()
	}

	p.Releases = releases

	return nil
}

func (p *Project) getMajors() {
	var majors []string

	for _, release := range p.Releases {
		version := strings.Split(release.TagName, ".")

		if !contains(majors, version[0]) {
			majors = append(majors, version[0])
		}
	}

	sort.Slice(majors, func(i, j int) bool {
		return majors[i] > majors[j]
	})

	p.Majors = majors
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
