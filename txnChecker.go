package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func (newTxn Txn) DeriveHash(t DBDetails) {
	info, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err, t)
	}
	hash := sha256.Sum256((info))
	bs := fmt.Sprintf("%x", hash)
	newTxn.Data.Hash = bs
	tmpch <- newTxn
}

func (b Block) DeriveHash() {
	info, _ := json.Marshal(b)
	hash := sha256.Sum256([]byte(info))
	bs := fmt.Sprintf("%x", hash)
	PrevHash = bs
}

func Validator(inputTxn map[string]Data) {
	TxnNo++
	newTxn := Txn{}
	for id, value := range inputTxn {
		newTxn.Id = id
		newTxn.Data = value
		key := []byte(id)
		data, err := UseDB.Get(key, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		var tmp DBDetails
		json.Unmarshal(data, &tmp)
		if tmp.ver == newTxn.Data.Version {
			newTxn.Data.Valid = true
			newTxn.Data.Value = value.Value
			newTxn.Data.Version += 1.0
			tmp.ver += 1.0
			tmp.val = value.Value
		}
		newTxn.Data.TxnID = TxnNo
		go newTxn.DeriveHash(tmp)
		strData, err := json.Marshal(tmp)
		if err != nil {
			fmt.Println(err)
		}
		UseDB.Put(key, strData, nil)
	}
}
