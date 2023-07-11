package config

import (
	"os"
	"time"
)

func InitTimeZone() {
	ict, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
