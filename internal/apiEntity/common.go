package apiEntity

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type CommonPageReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"20" v:"max:20#分页数量最大20条" dc:"分页数量，最大20"`
}

type CommonPageRes struct {
	Total int `json:"total" dc:"总数"`
}

type TimestampDate time.Time

func (a *TimestampDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tm, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*a = TimestampDate(tm)
	return nil
}

func (a *TimestampDate) MarshalJSON() ([]byte, error) {
	s := (time.Time)(*a)
	return json.Marshal(s.Unix())
}

func (a *TimestampDate) Timestamp() int64 {
	return (time.Time)(*a).Unix()
}

func (a *TimestampDate) Gtime() *gtime.Time {
	return gtime.NewFromTimeStamp(a.Timestamp())
}

type Uint64ToString uint64

type int64ToString int64

func (i *Uint64ToString) UnmarshalJSON(b []byte) error {
	// 去掉 JSON 字符串的引号
	str := string(b)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	// 将字符串转换为 uint64
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return errors.New("无法将字符串转换为 uint64")
	}
	// 将解析后的值赋给 *i
	*i = Uint64ToString(value)
	return nil
}

// CommonPaginationRes 通用分页响应结构
type CommonPaginationRes struct {
	Total int64 `json:"total" dc:"总数"`
	Page  int   `json:"page" dc:"当前页"`
	Size  int   `json:"size" dc:"每页数量"`
}
