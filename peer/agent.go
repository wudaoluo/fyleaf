package peer



type Agent interface {
	Run()
	OnClose()
}
