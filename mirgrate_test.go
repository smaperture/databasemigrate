package databasemigrate

import "testing"

func TestNoConnectionString(t *testing.T) {
	err := Go("", false)

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}

func TestInvalidConnectionString(t *testing.T) {
	err := Go("invalid", false)

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}
