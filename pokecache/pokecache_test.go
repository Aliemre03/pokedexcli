package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	c := NewCache()
	if c == nil {
		t.Errorf("Cache is nil")
	}
}

func TestAddToCache(t *testing.T) {
	c := NewCache()
	c.Add("test", []byte("test"))
	if c.cache["test"].val == nil {
		t.Errorf("Cache entry is nil")
	}
}

func TestGetFromCache(t *testing.T) {
	c := NewCache()
	c.Add("test", []byte("test"))
	val, ok := c.Get("test")
	if !ok {
		t.Errorf("Cache entry not found")
	}
	if string(val) != "test" {
		t.Errorf("Cache entry is not equal to test")
	}
}

func TestGetFromCacheNotFound(t *testing.T) {
	c := NewCache()
	_, ok := c.Get("test")
	if ok {
		t.Errorf("Cache entry found")
	}
}
