package structures

import (
	"hash/fnv"
)

type HashTable []*List

func NewHashTable() HashTable {
	return make(HashTable, 2^32)
}

func (t HashTable) Insert(key string, val interface{}) {
	h := hash([]byte(key))
	bucket := h % len(t)

	if t[bucket] == nil {
		l := List{}
		l.Insert(key)
		t[bucket] = &l
	} else {
		t[bucket].Insert(val)
	}
}

func (t HashTable) Delete(key string) {
	h := hash([]byte(key))
	bucket := h % len(t)

	l := t[bucket]
	if l != nil {
		l.Delete(key)
	}
}

func (t HashTable) Get(key string) interface{} {
	h := hash([]byte(key))
	bucket := h % len(t)

	l := t[bucket]
	if l != nil {
		res := l.Search(key)
		if res == nil {
			return nil
		}
		return res.Item
	}
	return nil
}

func hash(data []byte) int {
	h := fnv.New32a()
	h.Write(data)
	hashed := h.Sum32()

	return int(hashed)
}
