package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {

	trie := initTrie()

	t.Run("Test init function", func(t *testing.T) {
		assert.NotEqual(t, trie, (*TrieNode)(nil))
	})

	t.Run("Test no string search", func(t *testing.T) {
		res := trie.search("hi")
		assert.Equal(t, res, false)
	})

	t.Run("Test insert string", func(t *testing.T) {
		res := trie.insert("hi")
		assert.Equal(t, res, true)
	})

	t.Run("Test insert another string", func(t *testing.T) {
		res := trie.insert("hill")
		assert.Equal(t, res, true)
	})

	t.Run("Test valid string search", func(t *testing.T) {
		res := trie.search("hill")
		assert.Equal(t, res, true)
	})

	t.Run("Test invalid string search", func(t *testing.T) {
		trie.insert("hill")
		res := trie.search("hil")
		assert.Equal(t, res, false)
	})
}
