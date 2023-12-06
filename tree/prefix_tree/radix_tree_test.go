package tree

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

/* chatgpt

基数树（Radix Tree，也称为 Patricia 树）都是用于存储和搜索字符串集合的数据结构，它们有着相似的作用，但在某些情况下选择使用基数树可能更为合适。以下是一些使用基数树的主要场景和优点：

空间效率： 基数树通常比前缀树具有更好的空间效率。在前缀树中，可能会出现大量的分支节点，而基数树通过路径压缩的方式，将部分只有一个子节点的分支进行合并，从而减少了存储空间的使用。
前缀重叠： 当存储的字符串具有较长的公共前缀时，基数树能够更好地利用路径压缩，减少重复存储前缀的开销，提高了存储效率。
高效性能： 在某些情况下，基数树的查找、插入和删除操作可能比前缀树更加高效，尤其是对于大规模的字符串集合。
IPv4 和 IPv6 路由表： 基数树常被用于存储 IP 地址和路由信息，因为 IP 地址往往具有较长的前缀重叠，基数树能够有效地压缩存储这些信息。
*/

type radixNode struct {
	key      string
	value    any
	children map[string]*radixNode // 如果用数组就不需要再存储一份 key 到 map 中
}

type RadixTree struct {
	root *radixNode
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: &radixNode{
			children: make(map[string]*radixNode),
		},
	}
}

/*
	Insert 方法用于插入键值对到基数树中

从根节点开始，逐级检查当前节点的键与插入的键的前缀重叠情况：

  - 如果当前节点的键与插入的键完全相同，则更新该节点的值为新值。
  - 如果当前节点的键与插入的键部分重叠，则将当前节点的键截断，创建一个新的中间节点，并将之前的节点作为其子节点。
  - 如果当前节点的键与插入的键没有重叠，则在当前节点下创建一个新的叶子节点。
*/
func (t *RadixTree) Insert(key string, value any) {
	cur := t.root
LOOP:
	for len(cur.children) != 0 {
		for prefix, child := range cur.children {
			cprefix := commonPrefix(prefix, key)
			if cprefix != "" { // apart match at leat
				if prefix == key { // match totally
					// update value
					child.value = value
					return
				} else { // match apartly
					// check grandchild
					if len(cprefix) == len(prefix) {
						cur = child
						key = key[len(cprefix):]
						continue LOOP
					}
					// split
					newNode := &radixNode{
						key:      prefix[len(cprefix):],
						value:    child.value,
						children: child.children, // 把孩子也复制过来
					}
					child.key = cprefix // become parent
					child.value = value
					child.children = make(map[string]*radixNode) // 解除对原来孩子的引用
					child.children[newNode.key] = newNode

					delete(cur.children, prefix) // change the key in parent
					cur.children[cprefix] = child

					key = key[len(cprefix):]
					if len(key) > 0 {
						newNode2 := &radixNode{
							key:      key,
							value:    value,
							children: make(map[string]*radixNode),
						}
						child.children[key] = newNode2
					}
					return
				}
			}
		}
		// 子孩子遍历完了还没找到公共前缀
		break
	}

	// new node
	newNode := &radixNode{
		key:      key,
		value:    value,
		children: make(map[string]*radixNode),
	}
	cur.children[key] = newNode
}

/*
	Search 方法用于根据给定的键搜索对应的值

该方法遵循以下步骤：

  - 从根节点开始，逐级检查当前节点的键与搜索的键的前缀重叠情况。
  - 如果找到与搜索键完全匹配的节点，则返回该节点的值。
  - 如果当前节点的键与搜索的键部分重叠，则继续向下搜索。
*/
func (t *RadixTree) Search(key string) (any, bool) {
	cur := t.root
LOOP:
	for {
		for childkey, child := range cur.children {
			cp := commonPrefix(childkey, key)
			if len(cp) > 0 {
				if childkey == key {
					return child.value, true
				}
				key = key[len(cp):]
				cur = child
				continue LOOP
			}
		}
		return nil, false
	}
}

func longestCommonPrefix(str1, str2 string) string {
	i := 0
	for i < len(str1) && i < len(str2) && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}

func commonPrefix(s1, s2 string) string {
	i := 0
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}

func printTree(node *radixNode, level int) {
	if node == nil {
		return
	}
	fmt.Printf("%s- %s\n", strings.Repeat("  ", level), node.key)
	for _, child := range node.children {
		printTree(child, level+1)
	}
}

func TestRadixTree(t *testing.T) {
	table := []struct {
		key   string
		value any
		in    bool
		want  any
	}{
		{"apple", 1, true, 1},
		{"app", 2, true, 2},
		{"append", 3, true, 3},
		{"orange", 4, true, 4},
		{"orange", 40, true, 40},
		{"banana", 5, true, 5},
		{"b", 6, true, 6},
		{"bad", 7, true, 7},
		{"ap", 9, true, 8},
		{"a", nil, false, nil},
		{"grap", nil, false, nil},
	}
	tree := NewRadixTree()

	for _, v := range table {
		if v.in {
			tree.Insert(v.key, v.value)
		}
	}
	for _, v := range table {
		t.Run(v.key, func(t *testing.T) {
			value, exists := tree.Search(v.key)
			if exists != v.in && value != v.value {
				t.Logf("%v, %v | %v, %v\n", exists, value, v.in, v.want)
				t.Fail()
			}
		})
	}

	printTree(tree.root, 0)
}

func Benchmark(b *testing.B) {
	table := []struct {
		key   string
		value any
		in    bool
		want  any
	}{
		{"/giftwall/light/count", 1, true, 1},
		{"/giftwall/gift/top", 2, true, 2},
		{"/giftwall/gift/sponsor", 3, true, 3},
		{"/giftwall/light/status", 4, true, 4},
		{"/giftwall/gift/send", 40, true, 40},
		{"/giftwall/gift/remove/notify", 5, true, 5},
		{"/teenpatti/round/show", 6, true, 6},
		{"/teenpatti/round/result", 7, true, 7},
		{"/teenpatti/round/pot", 9, true, 8},
		{"/live/stream/destroy/notify", nil, false, nil},
		{"grap", nil, false, nil},
	}
	tree := NewRadixTree()
	hash := make(map[string]any)

	for _, v := range table {
		if v.in {
			tree.Insert(v.key, v.value)
			hash[v.key] = v.value
		}
	}

	printTree(tree.root, 0)

	b.Run("radix tree", func(b *testing.B) {
		b.ReportAllocs()
		r := rand.Intn(len(table))
		for i := 0; i < b.N; i++ {
			tree.Search(table[r].key)
		}
	})

	b.Run("hash", func(b *testing.B) {
		b.ReportAllocs()
		r := rand.Intn(len(table))
		for i := 0; i < b.N; i++ {
			_ = hash[table[r].key]
		}
	})
}
