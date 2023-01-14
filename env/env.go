package env

import "os"

type (
	PORT string
)

func NewPort() PORT { return PORT(os.Getenv("PORT")) }
