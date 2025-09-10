package fs

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashContents(contents []byte) string {
	sum := sha256.Sum256(contents)
	return hex.EncodeToString(sum[:])
}
