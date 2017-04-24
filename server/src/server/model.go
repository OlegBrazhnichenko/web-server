package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const PATH_TO_DB  = "./tasks.db"
var world = []byte("world")

func openConnection(DBName string) *bolt.DB{
	db, err := bolt.Open(PATH_TO_DB, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getFromDB(key []byte ){
	db := openConnection(PATH_TO_DB)
	defer db.Close()

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ShowDB(){
	db := openConnection(PATH_TO_DB)
	defer db.Close()

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func setInDB(key, value []byte){
	db := openConnection(PATH_TO_DB)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		err := bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func updateInDB(key, value []byte){
	db := openConnection(PATH_TO_DB)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFromDB(key []byte){
	db := openConnection(PATH_TO_DB)
	defer db.Close()
}

func Register(){
	db := openConnection(PATH_TO_DB)
	defer db.Close()
}

func Login(){
	db := openConnection(PATH_TO_DB)
	defer db.Close()
}

func initDB(){
	db := openConnection(PATH_TO_DB)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(world)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
