package unique_identifier

import (
	"crypto/rand"
	"fmt"
)

func UiID() (string, error) {
	list := make([]byte, 16)
	_, err := rand.Read(list)
	if err != nil {
		return "error in generate uuid", err
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", list[0:4], list[4:6], list[6:8], list[8:10], list[10:]), nil
}
