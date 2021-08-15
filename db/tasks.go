package db

import (
	"encoding/binary"
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

// CreateTask returns a key and an error.
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// b for bucket
		b := tx.Bucket(taskBucket)
		// only two reasons that an error occur and are dev reliant, we arent catching it.
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetAllTasks
func GetAllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		// we can continually call next, similar to linked list.
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// ===================================================================================>
// vvvvv Converters over here vvvvv

// Converting integer to Byte
// using Big Endian
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
