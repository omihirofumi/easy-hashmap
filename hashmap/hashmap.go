// Package hashmap provides a just easy hashmap.
package hashmap

import (
	"bytes"
	"fmt"
)

// MAP_SIZE is the size of the hash table.
const MAP_SIZE = 30

// Node is a HashMap node.
// The Type of the value can be a string or a numeric.
type Node[T interface{ string | int }] struct {
	key   string
	value T
	next  *Node[T]
}

// HashMap is a hashMap。
type HashMap[T string | int] struct {
	nodes []*Node[T]
}

// hash calculates the hash value based on the "Jenkins hash function".
func hash(key string) uint32 {
	var hash uint32 = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash
}

// NewHashMap creates a HashMap with values of the specified type.
func NewHashMap[T string | int]() *HashMap[T] {
	return &HashMap[T]{nodes: make([]*Node[T], MAP_SIZE)}
}

// getIndex returns the index of the key.
func getIndex(key string) int {
	return int(hash(key)) % MAP_SIZE
}

// Put puts a key and value to HashMap.
func (h *HashMap[T]) Put(key string, value T) {
	index := getIndex(key)

	if h.nodes[index] != nil {
		node := h.nodes[index]
		for node != nil {
			if node.key == key {
				node.value = value
				return
			}
			node = node.next
		}
		node.next = &Node[T]{key: key, value: value}
	} else {
		h.nodes[index] = &Node[T]{key: key, value: value}
	}
}

// Get retrieves the value of the key from HashMap.
func (h *HashMap[T]) Get(key string) (value T, ok bool) {
	index := getIndex(key)

	if h.nodes[index] != nil {
		node := h.nodes[index]
		for node != nil {
			if node.key == key {
				value = node.value
				ok = true
				return
			}
			node = node.next
		}
	}
	return
}

// String returns the contents of HashMap as a string.
func (h *HashMap[T]) String() string {
	var b bytes.Buffer
	for _, node := range h.nodes {
		for node != nil {
			b.WriteString(fmt.Sprintf("%s=%v; ", node.key, node.value))
			node = node.next
		}
	}
	return b.String()
}

// Length returns the size of HashMap.
func (h *HashMap[T]) Length() (length int) {
	for _, node := range h.nodes {
		for node != nil {
			length++
			node = node.next
		}
	}
	return
}

// Keys returns the list of keys.
func (h *HashMap[T]) Keys() []string {
	length := h.Length()
	keyList := make([]string, 0, length)

	for _, node := range h.nodes {
		for node != nil {
			keyList = append(keyList, node.key)
			node = node.next
		}
	}
	return keyList
}
