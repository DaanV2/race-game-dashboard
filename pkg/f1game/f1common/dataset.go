package f1common

import "golang.org/x/exp/constraints"

type DataSet[TKey constraints.Integer, T comparable] map[TKey]T

func (s DataSet[TKey, T]) Get(id TKey) (item T, ok bool) {
	item, ok = s[id]

	return
}

// FindByID tries to check if under a similar key, an item exists
// will return ok = true if found, will convert the key to the key type
func (s DataSet[TKey, T]) FindByID(lookup int) (key TKey, item T, ok bool) {
	key = TKey(lookup)
	item, ok = s[key]

	return
}

// FindByValue loops through the item to find the exact match
// will return ok = true if found, will convert the key to the key type
func (s DataSet[TKey, T]) FindByValue(match T) (key TKey, ok bool) {
	for k, v := range s {
		if v == match {
			return k, true
		}
	}

	return
}
