package errcode



func init() {
	//初始化错误字符串
	errList := GetErrInstance()
	errList.Load()
}
