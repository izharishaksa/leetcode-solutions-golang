//https://leetcode.com/problems/word-break-ii/

package word_break_ii

import (
	"strings"
)

type Trie struct {
	words map[string]bool
	index map[string]bool
}

func NewTriImpl() Trie {
	return Trie{
		words: make(map[string]bool, 0),
		index: make(map[string]bool, 0),
	}
}

func (t *Trie) Insert(word string) {
	t.words[word] = true
	prefix := ""
	for _, c := range word {
		prefix += string(c)
		t.index[prefix] = true
	}
}

func (t *Trie) Search(word string) bool {
	_, isExist := t.words[word]
	if isExist {
		return true
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	_, isExist := t.index[prefix]
	if isExist {
		return true
	}
	return false
}

var trie Trie
var result []string

func wordBreak(s string, wordDict []string) []string {
	trie = NewTriImpl()
	for _, word := range wordDict {
		trie.Insert(word)
	}
	result = make([]string, 0)
	searchWord("", string(s[0]), s, 0)
	return result
}

func searchWord(curResult, prefix, word string, pos int) {
	if !trie.StartsWith(prefix) {
		return
	}
	if pos == len(word)-1 {
		if trie.Search(prefix) {
			result = append(result, strings.Trim(curResult+" "+prefix, " "))
		}
		return
	}
	if trie.Search(prefix) {
		searchWord(curResult+" "+prefix, string(word[pos+1]), word, pos+1)
	}
	searchWord(curResult, prefix+string(word[pos+1]), word, pos+1)
}
