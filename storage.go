package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)
const dbFile = "tasks.db"
const bucketName = "Tasks"

// openn db connection
func openDB() *bbolt.DB {
	db, err := bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
// add task
func addTask(title string) {
	db := openDB()
	defer db.Close()
	db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			b, _ = tx.CreateBucket([]byte(bucketName))
		}
		id, _ := b.NextSequence()
		task := Task{ID: int(id), Title: title, Done: false}
		data, _ := json.Marshal(task)
		b.Put(itob(int(id)), data)
		fmt.Println("Task added:", title)
		return nil
	})
}
// list tasks
func listTasks() {
	db := openDB()
	defer db.Close()

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			fmt.Println("No tasks found.")
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			var task Task
			json.Unmarshal(v, &task)
			fmt.Printf("%d. %s (Done: %v)\n", task.ID, task.Title, task.Done)
			return nil
		})
		return nil
	})
}

// remove task
func removeTask(id int) {
	db := openDB()
	defer db.Close()

	db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b != nil {
			b.Delete(itob(id))
			fmt.Println("Task removed:", id)
		}
		return nil
	})
}
// convert int to byte slice
func itob(v int) []byte {
	return []byte(fmt.Sprintf("%d", v))
}
