package project

type Group struct {
	Language string
	Title    string
	Order    int

	Projects []*Project
}

// GroupPHP represents PHP Frameworks
var GroupPHP = Group{
	Language: "PHP",
	Title:    "PHP Frameworks",
	Order:    1,
}

// GroupJavascript represents JS Frameworks
var GroupJavascript = Group{
	Language: "JS",
	Title:    "Frontend Frameworks",
	Order:    2,
}
