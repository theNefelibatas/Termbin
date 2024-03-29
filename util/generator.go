package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenDigest(content []byte) (string, error) {
	digest := sha256.Sum256(content)
	digestStr := hex.EncodeToString(digest[:])
	return digestStr, nil
}

func GenShortID(length int) (string, error) {
	var stringBuilder = strings.Builder{}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(alphabet))
		stringBuilder.WriteByte(alphabet[index])
	}
	return stringBuilder.String(), nil
}

func GenUUID() (string, error) {
	u, err := uuid.NewRandom()
	return u.String(), err
}
