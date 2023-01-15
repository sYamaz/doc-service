package env

import (
	"os"
	"strconv"
)

type (
	PORT           string
	HASH_SALT      string
	HASH_STRETCH   int
	JWT_SECRET_KEY string
)

func NewPort() PORT {
	v := os.Getenv("PORT")
	if v == "" {
		v = "3000"
	}
	return PORT(v)
}

func NewHashSalt() HASH_SALT {
	return HASH_SALT(os.Getenv("HASH_SALT"))
}

func NewHashStretch() HASH_STRETCH {
	v, err := strconv.Atoi(os.Getenv("HASH_STRETCH"))
	if err != nil {
		v = 1
	}
	return HASH_STRETCH(v)
}

func NewJwtSecretKey() JWT_SECRET_KEY {
	return JWT_SECRET_KEY(os.Getenv("JWT_SECRET_KEY"))
}
