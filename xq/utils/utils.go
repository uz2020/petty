package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"time"
)

func GenerateRandomBytesWithReader(size int, reader io.Reader) ([]byte, error) {
	if reader == nil {
		return nil, fmt.Errorf("provided reader is nil")
	}
	buf := make([]byte, size)
	if _, err := io.ReadFull(reader, buf); err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %v", err)
	}
	return buf, nil
}

func GenerateRandomBytes(size int) ([]byte, error) {
	return GenerateRandomBytesWithReader(size, rand.Reader)
}

func GenUserId() string {
	size := 8
	b, err := GenerateRandomBytes(size)
	if err != nil {
		log.Println("GenerateRandomBytesError", err.Error())
		t := time.Now().Format(time.RFC3339Nano)
		h := md5.New()
		h.Write([]byte(t))
		ts := hex.EncodeToString(h.Sum(nil))
		return ts[:16]
	}
	return hex.EncodeToString(b)
}
