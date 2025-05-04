package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringSlice 自定义字符串切片类型
type StringSlice []string

// Value 写入数据库时转换为 JSON 字符串
func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(s)
}

// Scan 从数据库读取时解析 JSON 字符串到切片
func (s *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型不匹配")
	}
	return json.Unmarshal(bytes, s)
}
