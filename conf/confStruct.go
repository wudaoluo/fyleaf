package conf



//不太适合修改的值 或者 基本不会改动的值
const LenStackBuf = 4096
const LogV int32 = 3
const ConsolePrompt string = "fyLeaf# "
const ProfilePath string = ""

type Server struct {
	Version     string
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int32
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
