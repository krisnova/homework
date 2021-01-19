package hash

import (
	"crypto/md5"
	"fmt"
	"math/big"
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
}

type Bucket struct {
	ID      int
	Records []*Record
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
			ID: i - 1, // Pass in 512->511,1->0
			// Records: <>
		})
	}
	return table
}

func (t *Table) Get(key string) *Record {
	n := t.Hash(key)
	bucket := t.Buckets[n]
	// Linear search
	// TODO add a more effecient search later :)
	for _, record := range bucket.Records {
		// TODO Implement a more effecient way of comparing record keys
		if record.Key == key {
			return record
		}
	}
	return nil
}

//func (t *Table) Set(record *Record)
func (t *Table) Set(key string, value interface{}) {
	// Calculate n
	n := t.Hash(key)
	// Create a new record
	record := &Record{
		Key:   key,
		Value: value,
	}
	for i := len(t.Buckets[n].Records); i > 0; i-- {
		// TODO Implement a more effecient way of comparing record keys
		if t.Buckets[n].Records[i-1].String() == key {
			// Record key exists, so we update
			t.Buckets[n].Records[i-1] = record
			return
		}
	}

	// Add new record to table
	t.Buckets[n].Records = append(t.Buckets[n].Records, record)
}

// Hash is used to calculate an int value for a given string where the integer
// result is a factor of the table's dynamic bucket count
func (t *Table) Hash(key string) int {
	// Calculate MD5 sum and form an int result
	bytes := md5.Sum([]byte(key))
	md5 := fmt.Sprintf("%d", bytes)
	// Now that we have a idempotent hash, we can calculate an integer value
	bigInt := big.NewInt(0)
	bigInt.SetString(md5, 8) // Probably get away with base 2 or base 16
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
	x := n%t.BucketCount + 1
	return x // Result should alwayws be a factor of t.BucketCount
}

func (r *Record) String() string {
	if r == nil || r.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", r.Value)
}
