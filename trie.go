package trie

import ("fmt")

type trieNode struct {
    children map[rune]trieNode
    key string 
    val string
}

var emptyTrieNode = trieNode{}

func newTrieNode(key, val string) trieNode {
    return trieNode{
        children : make(map[rune]trieNode),
        key : key,
        val : val,
    }
}

func trieInsert(root trieNode, key, val string) error {
    if len(key) == 0 {
        return fmt.Errorf("Zero length key")
    }

    node := root
    i := 0
    keyr := []rune(key)
    for {
        if i == len(keyr)-1 {
            node.children[keyr[i]] = newTrieNode(key,val)
            return nil
        }

        if _,ok:= node.children[keyr[i]]; !ok {
            node.children[keyr[i]] = newTrieNode(key,val)
            node = node.children[keyr[i]]
            i++
            continue
        }

        node = node.children[keyr[i]]
        i++
    }
}

func trieLookup(root trieNode, key string) (trieNode, error) {
    if len(key) == 0 {
        return emptyTrieNode, fmt.Errorf("zero length key")
    }

    node := root
    i := 0
    keyr := []rune(key)
    for {
        if i == len(keyr) {
            return emptyTrieNode, fmt.Errorf("node not found")
        }

        if _,ok:= node.children[keyr[i]]; !ok {
            return emptyTrieNode, fmt.Errorf("node not found")
        }

        if n, ok:= node.children[keyr[i]]; i == len(keyr)-1 && ok {
            return n, nil
        }

        node = node.children[keyr[i]]
        i++
    }
}