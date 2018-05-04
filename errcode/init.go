package errcode



func init() {
	//初始化错误字符串
	config := GetErrInstance()
	config.Load()
}
