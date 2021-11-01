package unique_identifier

import (
	"testing"
)

func TestUiID(t *testing.T) {
	id, err := UiID()
	t.Run("generate unique identifier", func(t *testing.T) {
		if err != nil {
			t.Error("error: in generate unique identifier")
		}
	})

	t.Run("unique identifier generate incorrect", func(t *testing.T) {
		if len(id) == 17 {
			t.Error("error: unique identifier generate incorrect")
		}
	})
}
