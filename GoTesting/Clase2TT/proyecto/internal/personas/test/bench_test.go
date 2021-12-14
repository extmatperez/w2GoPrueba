package internal

import (
	"crypto/sha1"
	"crypto/sha256"
	"testing"
)

func BenchmarkSum256(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}
func BenchmarkSum(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}
