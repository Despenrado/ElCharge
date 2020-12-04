package models

import (
	"testing"
)

func TestEncryptString(t *testing.T) {
	str := "test"
	encEsting, err := EncryptString(str)
	if err != nil {
		t.Error(err)
	}
	if encEsting == str {
		t.Error("are equal")
	}
}
