package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte, difficulty int) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	if difficulty >= 1 {
		pow := NewProofOfWork(block, difficulty)
		hash := pow.Run()
		block.Hash = hash
	}
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string, difficulty int) *Block {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, difficulty)
	return newBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, 0)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func checkFirstBlock(database *sql.DB) bool {
	rows, _ := database.Query("SELECT data FROM block WHERE id = 1")
	var data string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data)
		return true
	}
	return false
}

func createDatabase(database *sql.DB) *sql.Stmt {
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS block (id INTEGER PRIMARY KEY, timestamp INTEGER, data BLOB, hash TEXT    , prevHash TEXT)")
	statement.Exec()
	if !checkFirstBlock(database) {
		block := NewGenesisBlock()
		createBlock(block, database, statement)
	}
	return statement
}

func createBlock(block *Block, database *sql.DB, statement *sql.Stmt) *sql.Stmt {
	statement, _ = database.Prepare("INSERT INTO block (timestamp, data, hash, prevHash) VALUES (?, ?, ?, ?)")
	statement.Exec(block.Timestamp, block.Data, block.Hash, block.PrevBlockHash)
	return statement
}

func printList(database *sql.DB) {
	rows, _ := database.Query("SELECT id, data, hash, prevHash FROM block")
	var id int
	var data string
	var hash []byte
	var prevHash []byte
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &data, &hash, &prevHash)
		fmt.Println(strconv.Itoa(id))
		fmt.Println(data)
		fmt.Printf("hash: %x\n", hash)
	}
}

func main() {
	database, _ := sql.Open("sqlite3", "./blockchain.db")
	statement := createDatabase(database)
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listBlockCmd := flag.NewFlagSet("list", flag.ExitOnError)
	bc := NewBlockchain()
	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
		block := bc.AddBlock(os.Args[2], 0)
		statement = createBlock(block, database, statement)
	case "list":
		err := listBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
		printList(database)
	case "mine":
		difficulty, _ := strconv.ParseInt(os.Args[2], 10, 32)
		block := bc.AddBlock("mine", int(difficulty))
		statement = createBlock(block, database, statement)
	}
}
