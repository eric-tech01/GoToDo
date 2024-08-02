package global

import (
	sjson "github.com/eric-tech01/simple-json"
)

type PaginationRsp struct {
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Total  int64       `json:"total"`
	Data   interface{} `json:"data"`
}

func FormatRsp(limit, offset int, total int64, data interface{}) PaginationRsp {
	return PaginationRsp{Limit: limit, Offset: offset, Total: total, Data: data}
}

type ResponseT *sjson.Json

// 返回结果包含 statusCode、code、data
func Success() ResponseT {
	return SuccessWithData(nil)
}
func SuccessWithData(data interface{}) ResponseT {
	j := sjson.New()
	j.Set("code", 200)
	j.Set("data", data)
	return j
}

func Failed(msg string) ResponseT {
	return FailedWithData(msg, nil)
}

func FailedWithData(msg string, data interface{}) ResponseT {
	j := sjson.New()
	j.Set("code", 500)
	j.Set("msg", msg)
	j.Set("data", data)
	return j
}
