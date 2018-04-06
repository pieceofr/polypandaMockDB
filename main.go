package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// ethkey "github.com/ethereum/go-ethereum/accounts/keystore"
)

const initZeroGen uint = 10

func main() {
	// Generate Panada Data
	num := flag.Int("num", 1, "number of mock record to create")
	flag.Parse()
	log.Println("num:", *num)

	cf := initConfig()
	var pandas []Panda
	for i := 0; i < *num; i++ {
		panda := CreatePanda(uint32(i))
		panda.PandaIndex = uint32(i)
		if i >= int(initZeroGen) {
			if panda.Generation != 0 {
				rand.Seed(time.Now().UTC().UnixNano())
				panda.MotherID = uint32(rand.Int31n(int32(i - 1)))
				panda.FatherID = uint32(rand.Int31n(int32(i - 1)))
				for panda.MotherID != panda.FatherID {
					panda.FatherID = uint32(rand.Int31n(int32(i - 1)))
				}
			}
		} else { // Initial Generation 0
			panda.Generation = 0
		}
		pandas = append(pandas, panda)
		time.Sleep(1 * time.Nanosecond) // for random seed to be varied
	}
	//	printAllPandas(pandas)
	//Open DB
	db, err := SQLConnect(cf.SQLUser, cf.SQLPwd, cf.SQLEndpoint, cf.SQLDB)
	if err != nil {
		log.Println("[mock] Connect to SQL Fail! err:", err)
		return
	}
	log.Println("SQL Connect")
	defer func() {
		SQLDisconnect(db)
		log.Println("SQL Disconnect")
	}()
	//Clear Panda Table
	SQLClearTable(db, cf.SQLPandaTable)
	//Transcation of multiple records
	InsertMultiplePandas(db, pandas)
}

func printAllPandas(pandas []Panda) {
	printStr := ""
	for idx, val := range pandas {
		if idx == 0 {
			printStr = fmt.Sprintf("%s[\n%v,", printStr, val.GetString())
		} else if idx == len(pandas)-1 {
			printStr = fmt.Sprintf("%s\n%v\n]", printStr, val.GetString())
		}
	}
	log.Println(printStr)
}
