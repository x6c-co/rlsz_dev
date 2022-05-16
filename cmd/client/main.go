package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/levidurfee/rlsz_dev/pkg/app"
	"github.com/levidurfee/rlsz_dev/pkg/config"
	"github.com/levidurfee/rlsz_dev/pkg/project"
)

func main() {
	project.GroupPHP.Projects = []*project.Project{
		{
			Name: "Laravel",
			Path: "laravel/laravel",
		},
		{
			Name: "Laravel Framework",
			Path: "laravel/framework",
		},
		{
			Name: "Slim",
			Path: "slimphp/Slim",
		},
		{
			Name: "Symfony PHP framework",
			Path: "symfony/symfony",
		},
	}

	project.GroupJavascript.Projects = []*project.Project{
		{
			Name: "Angular",
			Path: "angular/angular",
		},
		{
			Name: "React",
			Path: "facebook/react",
		},
		{
			Name: "Vue",
			Path: "vuejs/core",
		},
	}

	groups := []project.Group{
		project.GroupPHP,
		project.GroupJavascript,
	}

	app := app.App{
		Groups: groups,
		Stamp:  config.Stamp,
	}

	for _, group := range groups {
		for _, project := range group.Projects {
			err := project.Build()
			if err != nil {
				log.Println(err)
			}

			// Create files for each individual project
			err = project.Write()
			if err != nil {
				log.Println(err)
			}
		}
	}

	// Create index.html
	t, err := template.ParseFiles("web/templates/home.html", "web/templates/base.html")
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	err = t.ExecuteTemplate(&output, "base", app)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./build/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(output.Bytes())
}
