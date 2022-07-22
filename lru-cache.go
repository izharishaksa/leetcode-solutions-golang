package leetcode_solutions_golang

import "container/list"

type LRUCache struct {
	q        *list.List
	used     map[int]*list.Element
	data     map[int]int
	capacity int
}

//https://leetcode.com/problems/lru-cache/
func NewLruCache(capacity int) LRUCache {
	return LRUCache{
		used:     make(map[int]*list.Element, capacity+1),
		data:     make(map[int]int, capacity),
		capacity: capacity,
		q:        list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	val, isExist := lru.data[key]
	if isExist {
		lru.q.MoveToBack(lru.used[key])

		return val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	el, isExist := lru.used[key]
	if isExist {
		lru.q.Remove(el)
	}

	lru.data[key] = value
	newEl := lru.q.PushBack(key)
	lru.used[key] = newEl

	if len(lru.data) > lru.capacity {
		el := lru.q.Front()
		lruKey := el.Value.(int)
		lru.q.Remove(el)
		delete(lru.data, lruKey)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
