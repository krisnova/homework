package main

import (
	"fmt"

	"github.com/kris-nova/job2021/cs/junkdrawer"
	"github.com/kubicorn/kubicorn/pkg/namer"

	"github.com/kris-nova/job2021/cs/hash"
)

func main() {
	table := hash.NewTable("Nóva", 128)
	// Add 128*128 single kb records to the table
	generateKeyValue(table, 128*12, 1024)
	fmt.Println("Table length: ", table.Length)
	for _, bucket := range table.Buckets {
		fmt.Printf("Bucket (%d) length (%d)\n", bucket.ID, bucket.Length)
	}
}

var keys []string

func generateKeyValue(table *hash.Table, itemCount int, itemSize int) error {
	for i := itemCount; i > 0; i-- {
		key := namer.RandomName()
		value := junkdrawer.RandomString(itemSize)
		//fmt.Printf("%s : %d\n", key, itemSize)
		// Insert into table
		table.Set(key, value)
		keys = append(keys, key)
	}
	return nil
}
