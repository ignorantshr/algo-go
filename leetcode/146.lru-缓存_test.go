/*
* @lc app=leetcode.cn id=146 lang=golang
*
* [146] LRU 缓存
*
* https://leetcode.cn/problems/lru-cache/description/
*
  - algorithms
  - Medium (53.76%)
  - Likes:    2984
  - Dislikes: 0
  - Total Accepted:    551.4K
  - Total Submissions: 1M
  - Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
    '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'

*
* 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
*
* 实现 LRUCache 类：
*
*
*
*
* LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
* int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
* void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
* key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
*
*
* 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*
*
*
*
*
* 示例：
*
*
* 输入
* ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
* [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
* 输出
* [null, null, null, 1, null, -1, null, -1, 3, 4]
*
* 解释
* LRUCache lRUCache = new LRUCache(2);
* lRUCache.put(1, 1); // 缓存是 {1=1}
* lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
* lRUCache.get(1);    // 返回 1
* lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
* lRUCache.get(2);    // 返回 -1 (未找到)
* lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
* lRUCache.get(1);    // 返回 -1 (未找到)
* lRUCache.get(3);    // 返回 3
* lRUCache.get(4);    // 返回 4
*
*
*
*
* 提示：
*
*
* 1 <= capacity <= 3000
* 0 <= key <= 10000
* 0 <= value <= 10^5
* 最多调用 2 * 10^5 次 get 和 put
*
*
*/
package leetcode

import (
	"container/list"
	"testing"
)

// @lc code=start

type ilru interface {
	Get(key int) int
	Put(key int, value int)
}

type LRUCache struct {
	base ilru
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{base: NewLRUCache2(capacity)}
	return LRUCache{base: NewLRUCache3(capacity)}
	// return LRUCache{base: NewLRUCache1(capacity)}
}

func (this *LRUCache) Get(key int) int {
	return this.base.Get(key)
}

func (this *LRUCache) Put(key int, value int) {
	this.base.Put(key, value)
}

/* v2 使用库提供的链表 */
type LRUCache3 struct {
	cap int

	ring *ring146
	eles map[int]*ele146
}

type ele146 = list.Element
type ring146 = list.List

type eleVal146 struct {
	key int
	val int
}

func NewLRUCache3(capacity int) *LRUCache3 {
	return &LRUCache3{cap: capacity, ring: list.New(), eles: make(map[int]*ele146, capacity)}
}

func (this *LRUCache3) Get(key int) int {
	e, has := this.eles[key]
	if !has {
		return -1
	}

	this.moveToFront(e)
	return e.Value.(*eleVal146).val
}

func (this *LRUCache3) Put(key int, value int) {
	if e, has := this.eles[key]; has {
		e.Value.(*eleVal146).val = value
		this.moveToFront(e)
		return
	}

	if len(this.eles) == this.cap {
		this.removeFromTail()
	}
	e := &ele146{Value: &eleVal146{key: key, val: value}}
	this.insertToFront(e)
}

func (this *LRUCache3) moveToFront(e *ele146) {
	if e == this.ring.Front() {
		return
	}

	this.ring.MoveToFront(e)
}

func (this *LRUCache3) insertToFront(e *ele146) {
	e = this.ring.PushFront(e.Value)
	this.eles[e.Value.(*eleVal146).key] = e
}

func (this *LRUCache3) remove(e *ele146) {
	if e == nil {
		return
	}

	this.ring.Remove(e)
	delete(this.eles, e.Value.(*eleVal146).key)
}

func (this *LRUCache3) removeFromTail() {
	this.remove(this.ring.Back())
}

// 节点定义
type node146 struct {
	key  int
	val  int
	pre  *node146
	next *node146
}

/* v2 使用自定义封装链表 */
type LRUCache2 struct {
	cap int

	ring *doublelist
	eles map[int]*node146
}

func NewLRUCache2(capacity int) *LRUCache2 {
	return &LRUCache2{cap: capacity, ring: NewDoubleList(capacity), eles: make(map[int]*node146, capacity)}
}

func (this *LRUCache2) Get(key int) int {
	n, has := this.eles[key]
	if !has {
		return -1
	}

	this.moveToFront(n)
	return n.val
}

func (this *LRUCache2) Put(key int, value int) {
	if n, has := this.eles[key]; has {
		n.val = value
		this.moveToFront(n)
		return
	}

	if len(this.eles) == this.cap {
		this.removeFromTail()
	}
	n := &node146{key: key, val: value}
	this.insertToFront(n)
}

func (this *LRUCache2) moveToFront(n *node146) {
	if n == this.ring.Front() {
		return
	}

	this.ring.Remove(n)
	this.ring.PushFront(n)
}

func (this *LRUCache2) insertToFront(n *node146) {
	this.ring.PushFront(n)
	this.eles[n.key] = n
}

func (this *LRUCache2) remove(n *node146) {
	if n == nil {
		return
	}

	this.ring.Remove(n)
	delete(this.eles, n.key)
}

func (this *LRUCache2) removeFromTail() {
	this.remove(this.ring.Tail())
}

// 双向链表
type doublelist struct {
	root *node146 // 哨兵节点

	size int
}

func NewDoubleList(size int) *doublelist {
	r := &node146{}
	r.next = r
	r.pre = r
	return &doublelist{r, 0}
}

func (l *doublelist) Front() *node146 {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (l *doublelist) Tail() *node146 {
	if l.size == 0 {
		return nil
	}
	return l.root.pre
}

func (l *doublelist) PushFront(n *node146) {
	if n == nil {
		return
	}

	n.next = l.root.next
	n.pre = l.root
	l.root.next.pre = n
	l.root.next = n
	l.size++
}

func (l *doublelist) Remove(n *node146) *node146 {
	if n == nil {
		return nil
	}

	n.pre.next = n.next
	n.next.pre = n.pre
	n.pre = nil
	n.next = nil
	l.size--
	return n
}

func (l *doublelist) RemoveTail() *node146 {
	return l.Remove(l.Tail())
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

/*
	v1 版本

未将底层数据结构与操作分离
未封装完整的底层数据结构的操作方法
*/
type LRUCache1 struct {
	cap int

	ring *node146
	eles map[int]*node146
}

func NewLRUCache1(capacity int) *LRUCache1 {
	return &LRUCache1{cap: capacity, eles: make(map[int]*node146, capacity)}
}

func (this *LRUCache1) Get(key int) int {
	n, has := this.eles[key]
	if !has {
		return -1
	}

	this.putOnHead(n)
	return n.val
}

func (this *LRUCache1) Put(key int, value int) {
	n, has := this.eles[key]
	if !has {
		// 不在列表中
		if len(this.eles) == this.cap {
			this.removeNode(this.ring.pre, true)
		}
		n = &node146{key: key, val: value}
		this.eles[key] = n
	} else {
		// 在列表中
		n.val = value
	}
	this.putOnHead(n)
}

func (this *LRUCache1) putOnHead(n *node146) {
	if n == this.ring {
		return
	}

	if this.ring == nil {
		n.next = n
		n.pre = n
		this.ring = n
	} else {
		n = this.removeNode(n, false)
		n.next = this.ring
		n.pre = this.ring.pre
		this.ring.pre.next = n
		this.ring.pre = n
		this.ring = n
	}
}

func (this *LRUCache1) removeNode(n *node146, real bool) *node146 {
	if n == nil || (n.next == nil && n.pre == nil) {
		return n
	}

	if real {
		delete(this.eles, n.key)
	}
	n.pre.next = n.next
	n.next.pre = n.pre
	n.pre = nil
	n.next = nil
	return n
}

func TestLRUCache(t *testing.T) {
	capacity := 2
	lRUCache := Constructor146(capacity)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	t.Log(lRUCache.Get(1))
	lRUCache.Put(3, 3)
	t.Log(lRUCache.Get(2))
	lRUCache.Put(4, 4)
	t.Log(lRUCache.Get(1))
	t.Log(lRUCache.Get(3))
	t.Log(lRUCache.Get(4))
}
