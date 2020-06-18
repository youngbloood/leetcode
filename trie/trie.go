package trie

type trieNode struct {
	isEnd bool
	next  map[rune]*trieNode
}

type Trie struct {
	nodes map[rune]*trieNode
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	return Trie{
		nodes: make(map[rune]*trieNode, 0),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	b := rune(word[0])
	if node, exist := this.nodes[b]; exist {
		this.insert(word[1:], node)
	} else {
		n := this.nodes
		for i, v := range word {
			n[v] = &trieNode{next: make(map[rune]*trieNode, 0)}
			if i == len(word)-1 {
				n[v].isEnd = true
			}
			n = n[v].next
		}
	}
}

func (this *Trie) insert(word string, node *trieNode) {
	if len(word) == 0 {
		node.isEnd = true
		return
	}
	b := rune(word[0])
	if val, exist := node.next[b]; exist {
		this.insert(word[1:], val)
	} else {
		n := node
		for i, v := range word {
			insertNode := &trieNode{next: make(map[rune]*trieNode, 0)}
			if i == len(word)-1 {
				insertNode.isEnd = true
			}
			n.next[v] = insertNode
			n = insertNode
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	return this.search(word[1:], this.nodes[rune(word[0])])
}

func (this *Trie) search(word string, node *trieNode) bool {
	if node == nil {
		return false
	}

	if len(word) == 0 && !node.isEnd {
		return false
	} else if len(word) == 0 && node.isEnd {
		return true
	}

	return this.search(word[1:], node.next[rune(word[0])])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	return this.startWith(prefix[1:], this.nodes[rune(prefix[0])])
}

func (this *Trie) startWith(prefix string, node *trieNode) bool {
	if node == nil {
		return false
	}
	if len(prefix) == 0 {
		return true
	}
	return this.startWith(prefix[1:], node.next[rune(prefix[0])])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
