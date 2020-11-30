package models

import (
	"database/sql"
	"log"
	"math/rand"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/chichat?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannnot genrate UUID", err)
	}

	u[8] = (u[8] | 0x40)& 0x7F

	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:]
	return
}

func Encrypt(plantext string)(cryptext string){
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return

}
