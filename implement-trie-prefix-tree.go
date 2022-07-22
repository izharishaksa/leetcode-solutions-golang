package leetcode_solutions_golang

//https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	words map[string]bool
	index map[string]bool
}

func NewTrie() Trie {
	return Trie{
		words: make(map[string]bool, 0),
		index: make(map[string]bool, 0),
	}
}

func (this *Trie) Insert(word string) {
	this.words[word] = true
	prefix := ""
	for _, c := range word {
		prefix += string(c)
		this.index[prefix] = true
	}
}

func (this *Trie) Search(word string) bool {
	_, isExist := this.words[word]
	if isExist {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	_, isExist := this.index[prefix]
	if isExist {
		return true
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
