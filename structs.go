package main

import (
	"time"
)

type DBDetails struct {
	val int
	ver float32
}

type Block struct {
	BlockNO       int
	PrevBlockHash string
	txns          []Txn
	TimeStamp     time.Time
	CommitStatus  bool
}

type Txn struct {
	Id   string
	Data Data
}

type Data struct {
	TxnID   int
	Value   int     `json:"val"`
	Version float32 `json:"ver"`
	Valid   bool
	Hash    string
}
