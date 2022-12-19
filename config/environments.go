package config

import "os"

type Environments struct {
	AppPort string
}

func ENV() Environments {
	return Environments{
		AppPort: os.Getenv("APP_PORT"),
	}
}
