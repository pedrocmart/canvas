package http_test

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewHandler(t *testing.T) {
	checkIs := is.New(t)
	checkIs.Equal(true, true)
}
