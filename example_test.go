package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Run("successs", func(t *testing.T) {
		input := "ebi"
		result := SayHello(input)
		assert.NotEqual(t, result, input, "testing gagal")
	})

	t.Run("faild", func(t *testing.T) {
		input := "ebi"
		result := SayHello(input)
		assert.NotEqual(t, result, input, "testing gagal")
	})
}
