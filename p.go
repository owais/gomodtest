package gomodtest

import (
	"time"

	"github.com/jinzhu/now"
)

func Get() string {
	return now.Monday().Format(time.UnixDate)
}
