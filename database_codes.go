package main

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

func CreateDB() {
	_, filename, _, _ := runtime.Caller(0)
	dbPath := path.Dir(filename) + "./database"
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	UseDB = db
}

func PutInitialValues() {
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

func Put(key []byte, value []byte) {
	UseDB.Put(key, value, nil)
}
func Get(key []byte) []byte {
	output, getErr := UseDB.Get(key, nil)
	if getErr != nil {
		fmt.Println(getErr.Error() + string(key))
	}
	return output
}

func Delete(key []byte) {
	UseDB.Delete(key, nil)
}

func InitializeDatabase() {
	CreateDB()
	PutInitialValues()
}
