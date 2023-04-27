package dto

import (
	"Backend/model"
	"database/sql/driver"
	"strings"
)

func str2bool(s string) bool {
	if s == "yes" {
		return true
	} else {
		return false
	}
}

func bool2str(b bool) string {
	if b == true {
		return "yse"
	} else {
		return "no"
	}
}

func (m *bools) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*m = nil
	for i := 0; i < len(ss); i++ {
		*m = append(*m, str2bool(ss[i]))
	}
	return nil
}

func (m bools) Value() (driver.Value, error) {
	var str model.Strs
	str = nil
	for i := 0; i < len(m); i++ {
		str = append(str, bool2str(m[i]))
	}
	return str, nil
}
