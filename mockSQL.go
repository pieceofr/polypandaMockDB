package main

import (
	"database/sql"
	"fmt"
	"log"
)

/*SQLConnect SQL Database connect*/
func SQLConnect(name, passwd, address, dbname string) (*sql.DB, error) {
	url := name + ":" + passwd + "@tcp(" + address + ":3306)/" + dbname
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*SQLDisconnect SQL Database disconnect*/
func SQLDisconnect(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

/*SQLClearTable clear table content*/
func SQLClearTable(db *sql.DB, table string) error {
	sqlStatement := fmt.Sprintf("TRUNCATE TABLE %s", table)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

/*InsertMultiplePandas insert multiple records of panda into database*/
func InsertMultiplePandas(db *sql.DB, pandas []Panda) {
	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
	}
	defer tx.Rollback()
	insertStr := "INSERT INTO panda (pandaIndex, genes, birthtime, cooldown, rank," +
		" motherID, fatherID, generation, owner, ownername, photourl, snapurl)" +
		" VALUES (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? )"
	stm, err := tx.Prepare(insertStr)
	if err != nil {
		log.Panicln(err)
	}
	defer stm.Close()
	for _, p := range pandas {
		_, err = stm.Exec(p.PandaIndex, encodeNumberToHexString(p.Genes), p.Birthtime, p.Cooldown, p.Rank,
			p.MotherID, p.FatherID, p.Generation, ethAddrToHexString(p.Owner), p.Ownername, p.Photourl, p.Snapurl)
		if err != nil {
			log.Panicln(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Panicln(err)
	}
}
