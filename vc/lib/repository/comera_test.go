package repository_test

import (
	"testing"
	"logitech.com/vc/lib/repository"
	"fmt"
	"encoding/json"
)

func TestScan(t *testing.T) {
	r := repository.Scan();
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}

