package gocached

import (
	"testing"
)

func TestSet(t *testing.T) {
	server := NewServer()
	_ = server.Set("test", "value")
	item, _ := server.Get("test")
	if string(item.value) != "value" {
		t.Errorf("failed to return correct value. Expected %s, got %s", "test",
			string(item.value))
	}
}

func TestAdd(t *testing.T) {
	server := NewServer()
	_ = server.Set("test", "value")
	_ = server.Add("test", "another value")
	item, _ := server.Get("test")
	if string(item.value) != "value" {
		t.Errorf("failed to return correct value. Expected %s, got %s", "value",
			string(item.value))
	}
}

func TestReplace(t *testing.T) {
	server := NewServer()
	_ = server.Set("test", "value")
	st, _ := server.Replace("test", "another value")
	item, _ := server.Get("test")
	if string(item.value) != "another value" {
		t.Errorf("failed to replace %s with %s", "value", "another value")
	}
	st, _ = server.Replace("invalid data", "another value")
	if st != false {
		t.Errorf("should not replace non-existent data")
	}
}

func TestAppend(t *testing.T) {
	server := NewServer()
	_ = server.Set("test", "value")
	st, _ := server.Append("test", "value")
	if st != true {
		t.Errorf("failed to append anything")
	}
	item, _ := server.Get("test")
	if string(item.value) != "valuevalue" {
		t.Errorf("failed to append. Expected %s, got %s", "valuevalue",
			string(item.value))
	}
	st, _ = server.Append("anothertest", "value")
	if st != false {
		t.Errorf("should not append to non-existent key-value pair")
	}
}

func TestPrepend(t *testing.T) {
	server := NewServer()
	_ = server.Set("test", "value")
	st, _ := server.Prepend("test", "ok")
	if st != true {
		t.Errorf("failed to prepend anything")
	}
	item, _ := server.Get("test")
	if string(item.value) != "okvalue" {
		t.Errorf("failed to append. Expected %s, got %s", "okvalue",
			string(item.value))
	}
	st, _ = server.Append("anothertest", "value")
	if st != false {
		t.Errorf("should not prepend to non-existent key-value pair")
	}
}

func TestGetNotExistentKey(t *testing.T) {
	server := NewServer()
	item, _ := server.Get("test")
	if item != nil {
		t.Errorf("expected nil, got %s", item.value)
	}
}
