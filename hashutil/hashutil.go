package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 возвращает hex‑представление SHA‑256 для строки s.
func HashSHA256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
