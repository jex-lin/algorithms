package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://flaviocopes.com/golang-data-structure-hashtable/

type HashTable struct {
	items map[int]string
	lock  sync.RWMutex
}

func TestHashTable(t *testing.T) {
	t.Log("Test hash_table")
	list := []struct {
		items map[string]string
	}{
		{
			items: map[string]string{
				"aaa": "AAA",
				"bbb": "BBB",
				"ccc": "CCC",
				"ddd": "DDD",
				"eee": "EEE",
			},
		},
	}

	for _, item := range list {
		var h HashTable
		for k, v := range item.items {
			h.Put(k, v)
		}
		assert.Equal(t, 5, h.Size())
		assert.Equal(t, "DDD", h.Get("ddd"))
		h.Remove("ddd")
		assert.Equal(t, "", h.Get("ddd"))
	}
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (h *HashTable) Put(k string, v string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.items == nil {
		h.items = make(map[int]string)
	}
	i := hash(k)
	h.items[i] = v
}

func (h *HashTable) Remove(k string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(k)
	delete(h.items, i)
}

func (h *HashTable) Get(k string) string {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := hash(k)
	return h.items[i]
}

func (h *HashTable) Size() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.items)
}
