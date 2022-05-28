package version

import (
	"regexp"
	"strconv"
	"strings"
)

var reg *regexp.Regexp

func init() {
	r, err := regexp.Compile(`[a-zA-Z\.-]`)
	if err != nil {
		panic(err)
	}
	reg = r
}

func StripVersion(i string) int {
	version := strings.Split(i, ".")
	i = reg.ReplaceAllString(version[0], "")

	v, _ := strconv.Atoi(i)

	return v
}
