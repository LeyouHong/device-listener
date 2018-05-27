package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"fmt"
)

func NewDataBase() *leveldb.DB {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}

	return db
}

func Put(db *leveldb.DB, key []byte, val []byte) {
	db.Put(key, val, nil)
}

func Get(db *leveldb.DB, key []byte) []byte {
	data, _ := db.Get(key, nil)
	return data
}

func Search(db *leveldb.DB, key string, start string, end string,res map[string]string) {
	start = key + start
	end = key + end
	fmt.Println(start)
	fmt.Println(end)
	iter := db.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(end)}, nil)
	for iter.Next() {
		res[string(iter.Key())] = string(iter.Value())
	}
	iter.Release()
}

func CloseDataBase(db *leveldb.DB) {
	defer db.Close()
}