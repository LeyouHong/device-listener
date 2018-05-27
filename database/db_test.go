package database

import (
	"testing"
	"bytes"
	"log"
)

func TestDB(t *testing.T) {
	db := NewDataBase()

	keyArray := [][]byte {
		[]byte("demo2018-05-21 14:49:07.942329535 -0400 EDT m=+0.017295054"),
		[]byte("demo2018-05-21 14:49:08.942329535 -0400 EDT m=+0.017295055"),
		[]byte("demo2018-05-21 14:49:09.942329535 -0400 EDT m=+0.017295056"),
		[]byte("demo2018-05-21 14:49:10.942329535 -0400 EDT m=+0.017295057"),
		[]byte("demo2018-05-21 14:49:11.942329535 -0400 EDT m=+0.017295058"),
	}

	valArray := [][]byte {
		[]byte("{\"id\":demo1526593921,\"name\":\"demo\",\"action\":\"register\",\"ts\":1526593921,\"addr\":\"172.17.0.2:8080\"}"),
		[]byte("{\"id\":demo1526593923,\"name\":\"demo\",\"action\":\"deregister\",\"ts\":1526593923,\"addr\":\"172.17.0.2:8080\"}"),
		[]byte("{\"id\":demo1526593925,\"name\":\"demo\",\"action\":\"register\",\"ts\":1526593925,\"addr\":\"172.17.0.2:8080\"}"),
		[]byte("{\"id\":demo1526593927,\"name\":\"demo\",\"action\":\"deregister\",\"ts\":1526593927,\"addr\":\"172.17.0.2:8080\"}"),
		[]byte("{\"id\":demo1526593929,\"name\":\"demo\",\"action\":\"register\",\"ts\":1526593929,\"addr\":\"172.17.0.2:8080\"}"),
	}

	for i := 0; i < len(keyArray); i++ {
		Put(db, keyArray[i], valArray[i])
	}

	for i := 0; i < len(keyArray); i++ {
		val := Get(db, keyArray[i])
		if !bytes.Equal(val, valArray[i]) {
			//log.Println(string(valArray[i]))
			//log.Println(string(val))
			log.Fatal("test database failed")
		}
	}

	m := make(map[string]string)
	Search(db, "demo", "2018-05-21 14:49:07", "2018-05-21 14:49:09", m)
	log.Println(m)

	CloseDataBase(db)
}