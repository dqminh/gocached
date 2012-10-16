package gocached

import (
	"testing"
)

func TestSet(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("test", []byte("value"))
	item, _ := cache.Get("test")
	if string(item.value) != "value" {
		t.Errorf("failed to return correct value. Expected %s, got %s", "test",
			string(item.value))
	}
}

func TestAdd(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("test", []byte("value"))
	_ = cache.Add("test", []byte("another value"))
	item, _ := cache.Get("test")
	if string(item.value) != "value" {
		t.Errorf("failed to return correct value. Expected %s, got %s", "value",
			string(item.value))
	}
}

func TestReplace(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("test", []byte("value"))
	st, _ := cache.Replace("test", []byte("another value"))
	item, _ := cache.Get("test")
	if string(item.value) != "another value" {
		t.Errorf("failed to replace %s with %s", "value", "another value")
	}
	st, _ = cache.Replace("invalid data", []byte("another value"))
	if st != false {
		t.Errorf("should not replace non-existent data")
	}
}

func TestAppend(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("test", []byte("value"))
	st, _ := cache.Append("test", []byte("value"))
	if st != true {
		t.Errorf("failed to append anything")
	}
	item, _ := cache.Get("test")
	if string(item.value) != "valuevalue" {
		t.Errorf("failed to append. Expected %s, got %s", "valuevalue",
			string(item.value))
	}
	st, _ = cache.Append("anothertest", []byte("value"))
	if st != false {
		t.Errorf("should not append to non-existent key-value pair")
	}
}

func TestPrepend(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("test", []byte("value"))
	st, _ := cache.Prepend("test", []byte("ok"))
	if st != true {
		t.Errorf("failed to prepend anything")
	}
	item, _ := cache.Get("test")
	if string(item.value) != "okvalue" {
		t.Errorf("failed to append. Expected %s, got %s", "okvalue",
			string(item.value))
	}
	st, _ = cache.Append("anothertest", []byte("value"))
	if st != false {
		t.Errorf("should not prepend to non-existent key-value pair")
	}
}

func TestGetNotExistentKey(t *testing.T) {
	cache := NewCache()
	item, _ := cache.Get("test")
	if item != nil {
		t.Errorf("expected nil, got %s", item.value)
	}
}
