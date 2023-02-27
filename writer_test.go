package iptables

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func BenchmarkWriter(b *testing.B) {
	writerEnv := os.Getenv("WRITER")
	var writer func(int) []byte

	if writerEnv == "OLD" {
		writer = OldWriter
	} else if writerEnv == "NEW" {
		writer = NewWriter
	} else {
		fmt.Printf("expected ENV variable WRITER with values [OLD|NEW]\n")
		return
	}

	for i := 1000; i <= 2000; i += 100 {
		b.Run(fmt.Sprintf("Rules-%d", i), func(b *testing.B) {
			writer(i + 1)
		})
	}
}

func TestOutput(t *testing.T) {
	for _, n := range []int{1, 5, 10} {
		assert.Equal(t, OldWriter(n), NewWriter(n))
	}
}
