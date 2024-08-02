package sqlite

import (
	"database/sql/driver"
	"encoding/json"
)

type StringArray []string

func (t *StringArray) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}
func (t StringArray) Value() (driver.Value, error) {
	return json.Marshal(t)
}
