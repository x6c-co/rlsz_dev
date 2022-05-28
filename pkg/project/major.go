package project

import (
	"sort"
	"strings"

	"github.com/levidurfee/rlsz_dev/pkg/github"
	"github.com/levidurfee/rlsz_dev/pkg/version"
)

func (p *Project) getMajors() {
	var majors []*github.Release
	var mjrs []string

	for _, release := range p.Releases {
		version := strings.Split(release.TagName, ".")

		if !contains(mjrs, version[0]) {
			release.MajorVersion = version[0]
			mjrs = append(mjrs, version[0])
			majors = append(majors, release)
		}
	}

	sort.Slice(majors, func(i, j int) bool {
		vi := version.StripVersion(majors[i].TagName)
		vj := version.StripVersion(majors[j].TagName)

		return vi > vj
	})

	p.Majors = majors
}
