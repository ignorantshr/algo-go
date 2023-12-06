package tree

import (
	"fmt"
	"testing"
)

// trie: 前缀树，https://en.wikipedia.org/wiki/Trie
// 下面是 chatgpt 的代码

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
}

func NewPrefixTree() *trieNode {
	return &trieNode{}
}

/*
	插入字符串

前缀树的插入操作包括以下步骤：

  - 从根节点开始遍历前缀树，查找字符在前缀树中的位置。
  - 如果字符在前缀树中不存在，则创建一个新的节点，并将该节点加入到前缀树中。
  - 继续遍历下一个字符，直到字符串的末尾位置。
  - 将最后一个节点标记为字符串的结尾。
*/
func (t *trieNode) Insert(word string) {
	cur := t
	for _, ch := range word {
		if cur.children == nil {
			cur.children = make(map[rune]*trieNode)
		}
		if _, ok := cur.children[ch]; !ok {
			cur.children[ch] = &trieNode{}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

/*
	查找字符串

前缀树的查找操作包括以下步骤：

  - 从根节点开始遍历前缀树，查找字符在前缀树中的位置。
  - 如果字符在前缀树中不存在，则字符串不存在于前缀树中。
  - 继续遍历下一个字符，直到字符串的末尾位置。
  - 如果最后一个节点被标记为字符串的结尾，则字符串存在于前缀树中。
*/
func (t *trieNode) Search(word string) bool {
	cur := t
	for _, ch := range word {
		if cur.children == nil {
			return false
		}
		if _, ok := cur.children[ch]; !ok {
			return false
		}
		cur = cur.children[ch]
	}
	return cur.isEnd
}

/*
	查找前缀匹配字符串

前缀树的前缀匹配操作包括以下步骤：

  - 从根节点开始遍历前缀树，查找前缀在前缀树中的位置。
  - 如果前缀在前缀树中不存在，则前缀没有匹配项。
  - 继续遍历下一个字符，直到无法继续匹配或者匹配到所有前缀相同的字符串。
  - 返回匹配到的字符串集合。
*/
func (t *trieNode) SearchPrefix(prefix string) []string {
	cur := t
	for _, ch := range prefix {
		if cur.children == nil {
			return nil
		}
		if _, ok := cur.children[ch]; !ok {
			return nil
		}
		cur = cur.children[ch]
	}
	return t.collectWords(cur, prefix)
}

func (t *trieNode) collectWords(node *trieNode, prefix string) []string {
	var words []string
	if node.isEnd {
		words = append(words, prefix)
	}

	for ch, n := range node.children {
		words = append(words, t.collectWords(n, prefix+string(ch))...)
	}
	return words
}

func Test(t *testing.T) {
	trie := NewPrefixTree()
	trie.Insert("apple")
	trie.Insert("app")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.Search("banana"))
	fmt.Println(trie.SearchPrefix("ap"))
	fmt.Println(trie.SearchPrefix("appl"))
}
