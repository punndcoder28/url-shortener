package helpers

import (
	"errors"
	"testing"
)

func TestHandleErr(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {
		HandleErr(nil)
	})

	t.Run("WithError", func(t *testing.T) {
		err := errors.New("test error")

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("HandleErr did not panic as expected")
			}
		}()

		HandleErr(err)
	})
}
