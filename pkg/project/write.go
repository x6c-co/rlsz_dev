package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/levidurfee/rlsz_dev/pkg/github"
)

var tmpl *template.Template

func init() {
	t, err := template.ParseFiles(
		"web/templates/base.html",
		"web/templates/repo.html",
	)
	if err != nil {
		panic(err)
	}
	tmpl = t
}

func (p Project) Write() error {
	err := p.writeIndex()
	if err != nil {
		return err
	}
	err = p.writeMajors()
	if err != nil {
		return err
	}

	return nil
}

func (p Project) writeIndex() error {
	p.Index = true
	writeJson(p, filepath.Join("build", p.Path), p.Name+"index.json")

	return write(p, filepath.Join("build", p.Path), "index.html")
}

func (p Project) writeMajors() error {
	releases := p.Releases

	var tmpProject Project

	// Loop through each major release
	for _, major := range p.Majors {
		var temp []*github.Release
		// Loop through each release
		for _, release := range releases {
			// Check if release matches major
			ver := strings.Split(release.TagName, ".")
			if ver[0] != major.MajorVersion {
				continue
			}
			// Add it to a temp releases slice
			temp = append(temp, release)
		}
		tmpProject = p
		// Set p.Releases to temp releases slice
		tmpProject.Releases = temp
		// Update name / title
		tmpProject.Name = p.Name + " " + major.MajorVersion
		// Write file
		write(tmpProject, filepath.Join("build", p.Path), major.MajorVersion+".html")
		writeJson(tmpProject, filepath.Join("build", p.Path), major.MajorVersion+".json")
	}

	return nil
}

func write(p Project, path, filename string) error {
	var output bytes.Buffer
	err := tmpl.ExecuteTemplate(&output, "base", p)
	if err != nil {
		return err
	}

	os.MkdirAll(path, os.ModePerm)
	file, err := os.Create(path + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(output.Bytes())

	return nil
}

func writeJson(p Project, path, filename string) error {

	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+"/"+filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
