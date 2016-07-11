function TrieNode(key, val){ 
    this.children = {};
    this.key = key; 
    this.val = val; 
}

function trieLookup(root,key){
    if (key.length === 0){
        throw new Error("Zero length key");
    }
    
    var node = root; 
    var i = 0; 

    while(true){ 
        if (!node.children[key[i]]){ 
            return; 
        } 
        if (i === key.length-1) { 
            return node.children[key[i]]; 
        } 
        node = node.children[key[i++]]; 
    }
}

function trieInsert(root, key, val) {
    if (key.length === 0){
        throw new Error("Zero length key");
    }

    var newNode = new TrieNode(key, val); 
    var node = root; 
    var i = 0; 

    while( true ){ 
        if (i === key.length-1) { 
            node.children[key[i]] = newNode; 
            return; 
        }

        if (node.children[key[i]]) { 
            node = node.children[key[i++]]; 
            continue; 
        }

        node.children[key[i]] = new TrieNode(key.substring(0, i)); 
        node = node.children[key[i++]]; 
    }
}