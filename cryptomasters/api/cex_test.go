package api

import (
	"testing"

	"dw/go/crypto/api/currencyapi"
)

func TestGetRate(t *testing.T) {
	_, err := currencyapi.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}
