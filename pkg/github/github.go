package github

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	hostname = "api.github.com"
)

func get(url string) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(os.Getenv("GH_USERNAME"), os.Getenv("GH_PAT"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resBody, nil
}
