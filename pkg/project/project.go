package project

import (
	"github.com/levidurfee/rlsz_dev/pkg/config"
	"github.com/levidurfee/rlsz_dev/pkg/github"
)

// Project represents a project a project on GitHub.
type Project struct {
	Name  string
	Path  string
	Stamp string

	Repository *github.Repository
	Releases   []*github.Release
	Majors     []*github.Release

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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
