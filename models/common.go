package models

// ResponseData 返回的数据统一格式
type ResponseData[T interface{}] struct {
	Code    int64  `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}
