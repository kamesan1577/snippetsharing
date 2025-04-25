package library

import (
	"crypto/md5"
	"fmt"
	"time"
)

func GetHash() string {
	seed := time.Now().String()
	hash := fmt.Sprintf("%x", md5.Sum([]byte(seed)))
	return hash
}
