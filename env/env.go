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

	DB_HOST   string
	DB_PORT   string
	DB_NAME   string
	DB_PASS   string
	DB_USER   string
	DB_SCHEMA string
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

func NewDBHost() DB_HOST {
	return DB_HOST(os.Getenv("DB_HOST"))
}

func NewDBPort() DB_PORT {
	return DB_PORT(os.Getenv("DB_PORT"))
}

func NewDBName() DB_NAME {
	return DB_NAME(os.Getenv("DB_NAME"))
}

func NewDBPass() DB_PASS {
	return DB_PASS(os.Getenv("DB_PASS"))
}

func NewDBUser() DB_USER {
	return DB_USER(os.Getenv("DB_USER"))
}

func NewDBSchema() DB_SCHEMA {
	return DB_SCHEMA(os.Getenv("DB_SCHEMA"))
}
