package orderedmap

import (
	"iter"
	"slices"
)

type OrderedMap[K comparable, V any] struct {
	m   map[K]V
	ord []K
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		m: make(map[K]V),
	}
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.m)
}

func (om *OrderedMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, k := range om.ord {
			if !yield(k) {
				return
			}
		}
	}
}
func (om *OrderedMap[K, V]) KVPairs() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range om.ord {
			if !yield(k, om.m[k]) {
				return
			}
		}
	}
}

// Get - fetch a value from the ordered map
func (om *OrderedMap[K, V]) Get(k K) (V, bool) {
	v, ok := om.m[k]
	return v, ok
}

// Set - unconditionally set the value of k to v.  Returns true
// if the key was newly added, false if it was an update
func (om *OrderedMap[K, V]) Set(k K, v V) bool {
	added := false
	if _, exists := om.m[k]; !exists {
		om.ord = append(om.ord, k)
		added = true
	}
	om.m[k] = v
	return added
}

// Insert - inserts the key if it doesn't exist, returns false otherwise
func (om *OrderedMap[K, V]) Insert(k K, v V) bool {
	if _, exists := om.m[k]; exists {
		return false
	}
	return om.Set(k, v)
}

// Update - updates the key if it exists, returns false otherwise
func (om *OrderedMap[K, V]) Update(k K, v V) bool {
	if _, exists := om.m[k]; !exists {
		return false
	}
	om.m[k] = v
	return true
}

// Delete - delete the key if it exists, returns false otherwise
func (om *OrderedMap[K, V]) Delete(k K) bool {
	if _, exists := om.m[k]; !exists {
		return false
	}
	delete(om.m, k)
	om.ord = slices.DeleteFunc(om.ord, func(ordK K) bool {
		return ordK == k
	})
	return true
}
