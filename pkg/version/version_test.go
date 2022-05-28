package version

import (
	"testing"
)

func TestStripVersion(t *testing.T) {
	versions := map[string]int{
		"v9":  9,
		"v1":  1,
		"v13": 13,
	}

	for k, v := range versions {
		version := StripVersion(k)

		if version != v {
			t.Error("Error")
		}
	}
}
