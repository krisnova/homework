package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/kris-nova/job2021/cs/bst"
)

// hash.Table{}
// map[string]interface{} <-- We are implementing

// "Charlie" -> abcdef12345 -> A bucket -> Charlie

// [ Table ]
//    |    a0001      / - Record1
//    | - [ Bucket ] {  - Record2
//    |               \ - Record3

//    |    b0001      / - Record4
//    | - [ Bucket ] {  - Record5
//    |               \ - Record6

//    |    c0001      / - Record7
//    | - [ Bucket ] {  - Record8
//    |               \ - Record9
//    |

const DefaultBucketCount = 1024

type Table struct {
	BucketCount int
	Name        string
	Buckets     []*Bucket
	Length      int // 0 indexed
}

type Bucket struct {
	ID      int
	Records *bst.BTree
	Length  int
}

type Record struct {
	Key   string
	Value interface{}
	//Hash  string // TODO Could return md5 string for hash
}

func NewTable(name string, bucketCount int) *Table {
	if bucketCount < 1 {
		bucketCount = DefaultBucketCount
	}
	table := &Table{
		Name:        name,
		BucketCount: bucketCount,
	}
	// Initialize the Table with BucketCount buckets
	for i := bucketCount; i > 0; i-- {
		table.Buckets = append(table.Buckets, &Bucket{
			ID:      i - 1, // Pass in 512->511,1->0
			Records: bst.NewBTree(fmt.Sprintf("%d", i-1)),
		})
	}
	return table
}

func (t *Table) Get(key string) *Record {
	n := t.Hash(key)
	bucket := t.Buckets[n]
	node := bucket.Records.Search(key)
	if node == nil {
		return nil
	}
	return node.Data.(*Record)
}

// Set is used to add a key/value pair to the hash table
//
// Currently uses a (slow) linear search to search for keys
//    Linear Search
//      Worst Case O(n) <-- Slow as fuck on large sets
//      Best  Case O(1)
//func (t *Table) Set(record *Record)
func (t *Table) Set(key string, value interface{}) {
	// Calculate n
	n := t.Hash(key)
	// Create a new record
	record := &Record{
		Key:   key,
		Value: value,
	}
	t.Buckets[n].Records.Insert(&bst.Node{
		Key:  key,
		Data: record,
	})
}

// Hash is used to calculate an int value for a given string where the integer
// result is a factor of the table's dynamic bucket count
// TODO currently the Hash() method does not balance and there is a high standard deviation
func (t *Table) Hash(key string) int {
	// Calculate MD5 sum and form an int result
	h := md5.New()
	h.Write([]byte(key))
	md5 := hex.EncodeToString(h.Sum(nil))
	// Now that we have a idempotent hash, we can calculate an integer value
	bigInt := big.NewInt(0)
	bigInt.SetString(md5, 16) // Probably get away with base 2 or base 16
	i64 := bigInt.Int64()
	n := int(i64)
	// n is the int result of our arithmetic
	if n < 0 {
		n = n * -1 // Take the absolute value of n
	}
	// modulus for our bucket count
	//
	// We use modulus here so that our int value (n)
	// is always divisible by the bucketCount such that
	// bucketCount + 1 = bucketCount[0]
	//
	// n = 7
	// bucketCount = 3
	//     [0] 1,4,7 <--- n=7 gets bucket 0
	//     [1] 2,5
	//     [2] 3,6
	//
	// This allows for upwards maximum of values of n
	// while retaining a finite amount of buckets that
	// algorithmically can be calculated.
	x := n % (t.BucketCount)
	return x // Result should alwayws be a factor of t.BucketCount
}

func (r *Record) String() string {
	if r == nil || r.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", r.Value)
}
