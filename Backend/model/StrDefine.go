package model

import (
	"database/sql/driver"
	"strings"
)

// 重写gorm Scanner/Valuer 接口

type Strs []string

func (m *Strs) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*m = ss
	return nil
}

func (m Strs) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}
