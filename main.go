package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	PrevHash string
	ch       = make(chan Txn, 10)
	tmpch    = make(chan Txn)
	blockUse = Block{
		BlockNO:       0,
		PrevBlockHash: PrevHash,
	}
	now           time.Time
	UseDB         *leveldb.DB
	BlockCapacity int = 4
	router            = gin.Default()
	TxnNo             = 0
)

func main() {
	InitializeDatabase()
	InitializeBlock()
	router.POST("/post", Handler)
	router.GET("/blocks", GetAllBlocks)
	router.GET("/blocks/:blockNumber", GetBlock)
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
