package env

import "os"

type (
	PORT string
)

func NewPort() PORT {
	v := os.Getenv("PORT")
	if v == "" {
		v = "3000"
	}
	return PORT(v)
}
