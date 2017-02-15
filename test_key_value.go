package voodoo

import "testing"

func TestKeyValue(t *testing.T, kv func() (KeyValue, error)) {
	t.Run("kv/get-set", func(t *testing.T) {
		testKVGetSet(t, kv)
	})

	t.Run("kv/iter", func(t *testing.T) {
		testKVIter(t, kv)
	})
}

func testKVGetSet(t *testing.T, kv func() (KeyValue, error)) {
	k, err := kv()
	if err != nil {
		t.Error("cannot open kv store:", err)
	}

	err = k.Set([]byte("/vader/son"), []byte("luke"))
	if err != nil {
		t.Error("cannot set key", err)
	}
	err = k.Set([]byte("/vader/daughter"), []byte("leia"))
	if err != nil {
		t.Error("cannot set key", err)
	}

	val, err := k.Get([]byte("/vader/son"))
	if err != nil {
		t.Error("cannot get previously set value")
	}

	if string(val) != "luke" {
		t.Error("value from get != actual")
	}
}

func testKVIter(t *testing.T, kv func() (KeyValue, error)) {
	k, err := kv()
	if err != nil {
		t.Error("cannot open kv store:", err)
	}

	err = k.Set([]byte("/vader/son"), []byte("luke"))
	if err != nil {
		t.Error("cannot set key", err)
	}
	err = k.Set([]byte("/vader/daughter"), []byte("leia"))
	if err != nil {
		t.Error("cannot set key", err)
	}

	iter, err := k.Walk([]byte("/vader/"))
	if err != nil {
		t.Error("cannot get iterator", err)
	}

	iter.Next()
	if string(iter.Val()) != "leia" {
		t.Error("first value in iter not leia")
	}

	iter.Next()
	if string(iter.Val()) != "luke" {
		t.Error("second value in iter not luke")
	}
}
