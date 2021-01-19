package hash

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// InterfaceHash will calculate an MD5 sum of a given
// go interface (or pointer to an interface) and
// can be used to calculate a unique value of a
// graph despite the order of the vertices from
// the root.
//
// TODO we could use reflection and do something
// like reflect.ValueOf(data).Type().Kind() == reflect.Struct
func InterfaceHash(data interface{}) (string, error) {
	jBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("unable to marshal: %v", err)
	}
	return fmt.Sprintf("%d", md5.Sum(jBytes)), nil
}
