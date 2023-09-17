package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		err := errors.New("test error")

		defer func() {
			r := recover()

			assert.NotEqual(t, nil, r, "Expected program to panic, but program did not panic")
		}()

		HandleErr(err)
	})
}

func TestHandleNoErr(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {
		defer func() {
			r := recover()

			assert.Equal(t, nil, r, "Expected program to not panic, but program panicked")
		}()

		HandleErr(nil)
	})
}
