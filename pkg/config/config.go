package config

import (
	"strconv"
	"time"
)

var Stamp string

func init() {
	Stamp = strconv.FormatInt(time.Now().UnixMicro(), 16)
}
