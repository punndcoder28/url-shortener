package helpers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleErr(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {
		defer func() {
			r := recover()

			assert.Equal(t, nil, r, "Program did not panic")
		}()

		HandleErr(nil)
	})

	t.Run("WithError", func(t *testing.T) {
		err := errors.New("test error")

		defer func() {
			r := recover()

			assert.NotEqual(t, nil, r, "Program recovered from panic successfully")
		}()

		HandleErr(err)
	})
}
