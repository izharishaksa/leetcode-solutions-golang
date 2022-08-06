//https://leetcode.com/problems/word-search-ii/

package word_search_ii

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
var boardArr [][]byte
var result []string
var isAdded map[string]bool

func findWords(board [][]byte, words []string) []string {
	boardArr = board
	isAdded = make(map[string]bool, 0)
	trie = NewTriImpl()
	for _, word := range words {
		trie.Insert(word)
	}
	result = make([]string, 0)
	for j := 0; j < len(board); j++ {
		for k := 0; k < len(board[0]); k++ {
			isVisited := make([][]bool, len(board))
			for j := 0; j < len(board); j++ {
				isVisited[j] = make([]bool, len(board[0]))
			}
			searchWord(j, k, "", isVisited)
		}
	}
	return result
}

func searchWord(row, col int, prefix string, isVisited [][]bool) {
	if row < 0 || row >= len(boardArr) || col < 0 || col >= len(boardArr[0]) {
		return
	}
	if isVisited[row][col] {
		return
	}
	prefix += string(boardArr[row][col])
	if !trie.StartsWith(prefix) {
		return
	}
	if trie.Search(prefix) && !isAdded[prefix] {
		result = append(result, prefix)
		isAdded[prefix] = true
	}
	isVisited[row][col] = true
	searchWord(row+1, col, prefix, isVisited)
	searchWord(row-1, col, prefix, isVisited)
	searchWord(row, col+1, prefix, isVisited)
	searchWord(row, col-1, prefix, isVisited)
	isVisited[row][col] = false
}
