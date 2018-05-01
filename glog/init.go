package glog



//var MaxSize uint64 = 1024*1024*100  //100MB
//var logDir = flag.String("log_dir", "", "If non-empty, write log files in this directory")

//func init() {
//	flag.BoolVar(&logging.toStderr, "logtostderr", false, "log to standard error instead of files")
//	flag.BoolVar(&logging.alsoToStderr, "alsologtostderr", false, "log to standard error as well as files")
//	flag.Var(&logging.verbosity, "v", "log level for V logs")
//	flag.Var(&logging.stderrThreshold, "stderrthreshold", "logs at or above this threshold go to stderr")
//	flag.Var(&logging.vmodule, "vmodule", "comma-separated list of pattern=N settings for file-filtered logging")
//	flag.Var(&logging.traceLocation, "log_backtrace_at", "when logging hits line file:N, emit a stack trace")
//	//var logDir = flag.String("log_dir", "", "If non-empty, write log files in this directory")
//	// Default stderrThreshold is ERROR.
//	logging.stderrThreshold = errorLog
//
//	logging.setVState(0, nil, false)
//	go logging.flushDaemon()
//}



var logDir string
var MaxSize uint64 = 1024*1024*100  //100MB

//日志级别 INFO,WARNING,ERROR,FATAL
//日志在console中的输出级别
//level uint32
func Init(level,log_dir string,v Level) {
	logging.toStderr = false
	logging.alsoToStderr = false
	logDir = log_dir

	switch level {
	case "INFO":
		logging.stderrThreshold = infoLog
	case "WARNING":
		logging.stderrThreshold = warningLog
	case "ERROR":
		logging.stderrThreshold = errorLog
	case "FATAL":
		logging.stderrThreshold = fatalLog
	default:
		logging.stderrThreshold = fatalLog
	}

	logging.setVState(v, nil, false)
	go logging.flushDaemon()
}


