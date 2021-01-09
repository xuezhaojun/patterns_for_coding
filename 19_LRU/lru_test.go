package LRU

import "testing"

type LRUCache struct {
	m        map[int]*Node
	dhead    *Node
	dtail    *Node
	capacity int
	size     int
}

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	fakeHead := &Node{
		key: 0,
		val: 0,
	}
	fakeTail := &Node{
		key: 0,
		val: 0,
	}
	fakeHead.next = fakeTail
	fakeTail.pre = fakeHead
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node),
		dtail:    fakeTail,
		dhead:    fakeHead,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	curNode := this.m[key]
	this.removeNodeFromCache(curNode)
	this.addNodeToHead(curNode)

	return curNode.val
}

func (this *LRUCache) Put(key int, value int) {
	// get the node
	var node *Node
	if _, ok := this.m[key]; ok {
		node = this.m[key]
		node.val = value
		this.removeNodeFromCache(node)
		this.addNodeToHead(node)
	} else {
		node = &Node{
			key: key,
			val: value,
		}
		this.addNodeToHead(node)
	}
	this.deleteTail()
}

func (this *LRUCache) addNodeToHead(node *Node) {
	head := this.dhead
	node.next = head.next
	node.pre = head

	head.next.pre = node
	head.next = node

	this.m[node.key] = node
	this.size += 1
}

func (this *LRUCache) removeNodeFromCache(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = nil
	node.next = nil

	this.size -= 1
	delete(this.m, node.key)
}

func (this *LRUCache) deleteTail() {
	if this.capacity < this.size {
		key := this.dtail.pre.key

		this.dtail = this.dtail.pre
		this.dtail.next = nil
		this.size -= 1

		delete(this.m, key)
	}
}

func TestLRU(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)
	lRUCache.Get(3) // 返回 3
	lRUCache.Get(4) // 返回 4
}
