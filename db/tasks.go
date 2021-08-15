package db

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

// a string that represents the name of the bucket we are storing in.
// bolt expects byte slices rather than strings.
// const to byte conversion is fine, but takes extra steps.
var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// have the user pass in the db path.
// this is an exported fn.
func Init(dbPath string) error {
	// we want db to be package level, not scope level
	var err error

	// bolt options are a struct, we need to take their memory address.
	// I want to guess that it is because we are only creating a db one at a time, we dont take up the memory
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}

	// db.Update takes in a function: can return an err if there is one.
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
