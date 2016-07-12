package trie

import ("testing")

func TestBasics(t *testing.T) {
    root := newTrieNode("","")
    trieInsert(root, "pool", "cool")

    _, err := trieLookup(root, "pool")
    if err != nil {
        t.Fatalf("Could not find key")
    }
}

func TestBasics2(t *testing.T) {
    root := newTrieNode("","")
    trieInsert(root, "pool", "cool")
    trieInsert(root, "ok", "cool")

    _, err := trieLookup(root, "ok")
    if err != nil {
        t.Fatalf("Could not find key")
    }
}