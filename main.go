package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	blockchain := CreateBlockchain(5)
	fmt.Println((blockchain.chain))
	blockchain.addBlock("Zach", "Patrick" , 100)
	fmt.Println((blockchain.chain[len(blockchain.chain) -1]))
	blockchain.addBlock("Patrick", "Zach", 100)
	fmt.Println((blockchain.chain[len(blockchain.chain) -1]))
	fmt.Println(blockchain.isValid())
}

type Block struct {
	data map[string]interface{}
	hash string
	prevHash string
	timestamp time.Time
	pow int
}

type Blockchain struct {
	genBlock Block
	chain []Block
	difficulty int
}

func (block Block) calcHash() string {
	data, err := json.Marshal(block.data)
	if err != nil {
		return "No data in block"
	}
	
	byteData := []byte(block.prevHash + string(data) + block.timestamp.String() + strconv.Itoa(block.pow))
	blockHash := sha256.Sum256(byteData)
	return fmt.Sprintf("%x", blockHash)

}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.hash, strings.Repeat("0", difficulty)) {
		block.pow++
		block.hash = block.calcHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genBlock := Block{
		hash: "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genBlock: genBlock,
		chain: []Block{genBlock},
		difficulty: difficulty,
	}
}

func (blockchain *Blockchain) addBlock(from string, to string, amount float64){
	blockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}
	
	lastBlock := blockchain.chain[len(blockchain.chain)-1]
	newBlock := Block {
		data: blockData,
		prevHash: lastBlock.hash,
		timestamp: time.Now(),
	}

	newBlock.mine(blockchain.difficulty)

	blockchain.chain = append(blockchain.chain, newBlock)

}

func (blockchain Blockchain) isValid() bool {
	for i:= range blockchain.chain[1:] {
		prevBlock := blockchain.chain[i]
		currBlock := blockchain.chain[i+1]
		if currBlock.hash != currBlock.calcHash() || currBlock.prevHash != prevBlock.hash {
			return false
		}
	}

	return true
}