package controller

// Context はcontrollerでリクエストを処理するための定義です。
type Context interface {
	Param(key string) string
	JSON(conde int, obj interface{})
}
