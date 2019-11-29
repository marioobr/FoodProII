package utils

import "os"

func GetPort() string{
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}
