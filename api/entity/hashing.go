package entity

import (
	"crypto/sha256"
	"doc-api/env"
	"encoding/hex"
)

type (
	Hashing interface {
		Hash(val string) string
	}

	hashing struct {
		salt    string
		stretch int
	}
)

func NewHashing(salt env.HASH_SALT, stretch env.HASH_STRETCH) Hashing {
	return &hashing{
		salt:    string(salt),
		stretch: int(stretch),
	}
}

func (h *hashing) Hash(val string) string {
	hashed := val
	for i := 0; i < h.stretch; i++ {
		hashed = hash_sha256(hashed + h.salt)
	}
	return hashed
}

func hash_sha256(val string) string {
	b := []byte(val)
	sha256 := sha256.Sum256(b)
	hashed := hex.EncodeToString(sha256[:])
	return hashed
}
