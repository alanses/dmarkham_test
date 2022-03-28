package services

import (
	"fmt"
	"testing"
)

func TestGetRandomName(t *testing.T) {
	randomNameService := NewRandomNameService()

	result, err := randomNameService.GetRandomName()

	if err != nil {
		t.Error(err, "error into get random name")
	}

	fmt.Println(result)
}
