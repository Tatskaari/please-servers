package trie

import (
	"testing"

	"google.golang.org/grpc"
	"github.com/stretchr/testify/assert"
)

func TestEmptyTrie(t *testing.T) {
	trie := New(callback)
	assert.Error(t, trie.Check())
}

func TestOneCompleteLevel(t *testing.T) {
	trie := New(callback)
	assert.NoError(t, trie.AddRange("0", "f", "127.0.0.1:443"))
	assert.NoError(t, trie.Check())
}

func TestOneCompleteLevelTwoParts(t *testing.T) {
	trie := New(callback)
	assert.NoError(t, trie.AddAll(map[string]string{
		"0-6": "127.0.0.1:443",
		"7-f": "127.0.0.1:443",
	}))
	assert.NoError(t, trie.Check())
	assert.EqualValues(t, trie.Get("0123"), trie.Get("6789"))
	assert.NotEqual(t, trie.Get("2345"), trie.Get("7890"))
}

func TestTwoCompleteLevels(t *testing.T) {
	trie := New(callback)
	assert.NoError(t, trie.AddAll(map[string]string{
		"00-0a": "127.0.0.1:443",
		"0b-0f": "127.0.0.1:443",
		"10-ff": "127.0.0.1:443",
	}))
	assert.NoError(t, trie.Check())
	assert.NotNil(t, trie.Get("abcd"))
	assert.Equal(t, trie.Get("abcd"), trie.Get("bcde"))
	assert.NotEqual(t, trie.Get("0abc"), trie.Get("0bcd"))
}

func TestOffset(t *testing.T) {
	trie := New(callback)
	assert.NoError(t, trie.AddAll(map[string]string{
		"00-3f": "127.0.0.1:443",
		"40-7f": "127.0.0.1:443",
		"80-af": "127.0.0.1:443",
		"b0-ff": "127.0.0.1:443",
	}))
	assert.NoError(t, trie.Check())
	assert.NotNil(t, trie.Get("0000"))
	// An offset of 1 isn't enough to move to the next replica.
	assert.Equal(t, trie.Get("0000"), trie.GetOffset("0000", 1))
	// But 4 is
	assert.NotEqual(t, trie.Get("0000"), trie.GetOffset("0000", 4))
	assert.Equal(t, trie.Get("4000"), trie.GetOffset("0000", 4))
}

func callback(address string) (*grpc.ClientConn, error) {
	return nil, nil
}
