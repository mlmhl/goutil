package io

import (
	"testing"
)


const (
	name = "test"
)
var (
	err error
)

func TestCreate(t *testing.T) {
	t.Log("Test: Create file...")

	_, err = Create(name)
	if err != nil {
		t.Fatalf("Create file error: %v", err)
	}
	if !IsExist(name) {
		t.Fatalf("Create file failed")
	}

	t.Log("Passed...")
}

func TestRemove(t *testing.T) {
	t.Log("Test: Remove file...")

	err = Remove(name)
	if err != nil {
		t.Fatalf("Remove file error: %v", err)
	}
	if IsExist(name) {
		t.Fatalf("Remove file failed")
	}

	t.Log("Passed...")
}