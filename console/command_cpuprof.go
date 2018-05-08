package console

import (
	"os"
	"runtime/pprof"
)

// cpuprof
type CommandCPUProf struct{}

func (c *CommandCPUProf) name() string {
	return "cpuprof"
}

func (c *CommandCPUProf) help() string {
	return "CPU profiling for the current process"
}

func (c *CommandCPUProf) usage() string {
	return "cpuprof writes runtime profiling data in the format expected by \r\n" +
		"the pprof visualization tool\r\n\r\n" +
		"Usage: cpuprof start|stop\r\n" +
		"  start - enables CPU profiling\r\n" +
		"  stop  - stops the current CPU profile"
}

func (c *CommandCPUProf) run(args []string) string {
	if len(args) == 0 {
		return c.usage()
	}

	switch args[0] {
	case "start":
		fn := profileName() + ".cpuprof"
		f, err := os.Create(fn)
		if err != nil {
			return err.Error()
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			f.Close()
			return err.Error()
		}
		return fn
	case "stop":
		pprof.StopCPUProfile()
		return ""
	default:
		return c.usage()
	}
}

