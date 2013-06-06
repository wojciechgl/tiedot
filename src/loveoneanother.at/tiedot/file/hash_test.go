package file

import (
	"os"
	"testing"
)

func TestPutGet(t *testing.T) {
	tmp := "/tmp/tiedot_hash_test"
	os.Remove(tmp)
	defer os.Remove(tmp)
	ht, err := OpenHash(tmp, 2, 2)
	if err != nil {
		t.Errorf("Failed to open: %v", err)
	}
	for i := uint64(0); i < 30; i++ {
		ht.Put(i, i)
	}
	for i := uint64(0); i < 30; i++ {
		keys, vals := ht.Get(i, 0, func(a, b uint64) bool {
			return true
		})
		if !(len(keys) == 1 && keys[0] == i && len(vals) == 1 && vals[0] == i) {
			t.Errorf("Get failed on key %d, got %v and %v", i, keys, vals)
		}
	}
	ht.File.Close()
}

func TestPutGet2(t *testing.T) {
	tmp := "/tmp/tiedot_hash_test"
	os.Remove(tmp)
	defer os.Remove(tmp)
	ht, err := OpenHash(tmp, 2, 2)
	if err != nil {
		t.Errorf("Failed to open: %v", err)
	}
	ht.Put(1, 1)
	ht.Put(1, 2)
	ht.Put(1, 3)
	ht.Put(2, 1)
	ht.Put(2, 2)
	ht.Put(2, 3)
	keys, vals := ht.Get(1, 0, func(a, b uint64) bool {
		return true
	})
	if !(len(keys) == 3 && len(vals) == 3) {
		t.Errorf("Get failed, got %v, %v", keys, vals)
	}
	keys, vals = ht.Get(2, 2, func(a, b uint64) bool {
		return true
	})
	if !(len(keys) == 2 && len(vals) == 2) {
		t.Errorf("Get failed, got %v, %v", keys, vals)
	}
	keys, vals = ht.Get(1, 0, func(a, b uint64) bool {
		return b >= 2
	})
	if !(len(keys) == 2 && len(vals) == 2) {
		t.Errorf("Get failed, got %v, %v", keys, vals)
	}
	ht.File.Close()
}

func TestPutRemove(t *testing.T) {
	tmp := "/tmp/tiedot_hash_test"
	os.Remove(tmp)
	defer os.Remove(tmp)
	ht, err := OpenHash(tmp, 2, 2)
	if err != nil {
		t.Errorf("Failed to open: %v", err)
	}
	ht.Put(1, 1)
	ht.Put(1, 2)
	ht.Put(1, 3)
	ht.Put(2, 1)
	ht.Put(2, 2)
	ht.Put(2, 3)
	ht.Remove(1, 1, func(a, b uint64) bool {
		return true
	})
	ht.Remove(2, 2, func(a, b uint64) bool {
		return b >= 2
	})
	keys, vals := ht.Get(1, 0, func(a, b uint64) bool {
		return true
	})
	if !(len(keys) == 2 && len(vals) == 2) {
		t.Errorf("Did not delete, still have %v, %v", keys, vals)
	}
	keys, vals = ht.Get(2, 0, func(a, b uint64) bool {
		return true
	})
	if !(len(keys) == 1 && len(vals) == 1) {
		t.Errorf("Did not delete, still have %v, %v", keys, vals)
	}
	ht.File.Close()
}

func TestGetAll(t *testing.T) {
	tmp := "/tmp/tiedot_hash_test"
	os.Remove(tmp)
	defer os.Remove(tmp)
	ht, err := OpenHash(tmp, 2, 2)
	if err != nil {
		t.Errorf("Failed to open: %v", err)
	}
	ht.Put(1, 1)
	ht.Put(1, 2)
	ht.Put(1, 3)
	ht.Put(2, 1)
	ht.Put(2, 2)
	ht.Put(2, 3)
	keys, vals := ht.GetAll()
	if !(len(keys) == 6 && len(vals) == 6) {
		t.Errorf("Did not have all, got only %v, %v", keys, vals)
	}
	ht.File.Close()
}
