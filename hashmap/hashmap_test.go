package hashmap

import (
	"testing"
)

func TestPutAndGetInt(t *testing.T) {
	t.Parallel()
	hmap := NewHashMap[int]()
	hmap.Put("key1", 1)
	hmap.Put("key2", 2)
	if v, ok := hmap.Get("key1"); ok {
		if v != 1 {
			t.Errorf("expected 1, but got %d", v)
		}
	} else {
		t.Error("not contains key1")
	}

	if v, ok := hmap.Get("key2"); ok {
		if v != 2 {
			t.Errorf("expected 2, but got %d", v)
		}
	} else {
		t.Error("not contains key2")
	}
}

func TestPutAndGetString(t *testing.T) {
	t.Parallel()
	hmap := NewHashMap[string]()
	hmap.Put("key1", "v1")
	hmap.Put("key2", "v2")
	if v, ok := hmap.Get("key1"); ok {
		if v != "v1" {
			t.Errorf("expected v1, but got %s", v)
		}
	} else {
		t.Error("not contains key1")
	}

	if v, ok := hmap.Get("key2"); ok {
		if v != "v2" {
			t.Errorf("expected v2, but got %s", v)
		}
	} else {
		t.Error("not contains key2")
	}
}

func TestDuplicationAndLength(t *testing.T) {
	t.Parallel()
	hmap := NewHashMap[string]()
	hmap.Put("key1", "v1")
	hmap.Put("key2", "v2")
	if hmap.Length() != 2 {
		t.Errorf("unexpected size %d, should be 2", hmap.Length())
	}
	hmap.Put("key2", "v2")
	if hmap.Length() != 2 {
		t.Errorf("unexpected size %d, should be 2", hmap.Length())
	}
}

func TestKeys(t *testing.T) {
	t.Parallel()

	truthMap := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	hmap := NewHashMap[int]()
	for k, v := range truthMap {
		hmap.Put(k, v)
	}

	if len(truthMap) != len(hmap.Keys()) {
		t.Error("Map/keys mismatch")
	}

	for _, key := range hmap.Keys() {
		if _, ok := truthMap[key]; !ok {
			t.Error("Key", key, "not in truthMap")
		}

		delete(truthMap, key)
	}

	if len(truthMap) != 0 {
		t.Error("key mismatch")
	}
}
