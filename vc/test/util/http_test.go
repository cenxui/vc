package util

import (
	"testing"
)

func TestGet(t *testing.T) {
	defer func() {
		if p := recover();p != nil {
			t.Log("recover")
		} else {
			t.Log("none")
		}
	}()
	panic("error")
	t.Log("complete")
}

