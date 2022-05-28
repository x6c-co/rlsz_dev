package config

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Stamp string
	Hash  string
)

func init() {
	Stamp = strconv.FormatInt(time.Now().UnixMicro(), 16)

	f, err := os.Open("./build/assets/css/app.css")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	Hash = fmt.Sprintf("%x", h.Sum(nil))
}
