package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

var UseDB *leveldb.DB

func CreateDB() {
	db, err := leveldb.OpenFile("./database", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	UseDB = db
}

func PutValues() {
	for i := 1; i <= 1000; i++ {
		id := "SIM" + strconv.Itoa(i)
		tempDBDetails := DBDetails{
			val: i,
			ver: 1.0,
		}
		tempDBDetails1, err := json.Marshal(tempDBDetails)
		DBkey := []byte(id)
		if err != nil {
			fmt.Print(err)
		}
		UseDB.Put(DBkey, tempDBDetails1, nil)
	}
}

func InitializeDatabase() {
	CreateDB()
	PutValues()
}
