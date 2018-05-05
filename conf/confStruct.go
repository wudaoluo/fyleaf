package conf



type Server struct {
	Version     string
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	Debug       bool
	Mysql       Mysql
}

//root:123456@tcp(10.211.55.4:3306)/game
type Mysql struct {
	DBname string
	DBaddr string
	DBport string
	DBuser string
	DBpasswd string
}
