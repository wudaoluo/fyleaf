package utils

import (
	"time"
	"fyleaf/glog"
	"io/ioutil"
	"os"
	"strings"
	"path"
)

type gclear struct {
	path     string
	suffix   string
	interval time.Duration   //多长时间运行
	reserve  time.Duration   //保留多久的文件
	c        chan struct{}
}

//默认保留7天 ,3小时检查一次
func NewGlogClear(path,suffix string,t ...time.Duration) *gclear {
	reserve := time.Hour*24*7
	interval := time.Hour*3
	switch len(t) {
	case 2:
		reserve = t[0]
		interval = t[1]
	case 1:
		reserve = t[0]
	case 0:
	default:
		break
	}

	c:= &gclear{
		path:path,
		suffix:suffix,
		interval:interval,
		reserve:reserve,
		c:make(chan struct{}),
	}

	go c.cleaner()
	return c
}



func (c *gclear) cleaner() {
	timer := time.NewTimer(c.interval)
	for {
		select {
		case <-c.c:
			glog.Info("停止时间轮训 timer")
			timer.Stop()
			return
		case <-timer.C:
			c.clean()
		}

		timer.Reset(c.interval)
	}

}


func (c *gclear) Close() {
	if c.c != nil {
		glog.Warning("关闭c.c chan")
		close(c.c)
		c.c = nil
	}else {
		glog.Warning("c.c 已经关闭了,多次关闭,出现 bug了")
	}
}


func (c *gclear) clean() {
	exists, err := c.exists(c.path)
	if err != nil {
		glog.Error(err)
		return
	}
	if !exists {
		return
	}

	files, err := ioutil.ReadDir(c.path)
	if err != nil {
		glog.Error(err)
		return
	}
	c.check(files)
}

// exists returns whether the given file or directory exists or not
func (c *gclear) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//至少要保留一个日志文件和连接文件
func (c *gclear) check(files []os.FileInfo) {
	filesLen := len(files)
	if filesLen <= 2 {
		glog.Info("日志文件少于2个不删除了")
		return
	}
	for i, f := range files {
		if i == filesLen -2 {
			return
		}
		prefix := strings.HasSuffix(f.Name(), c.suffix)
		if prefix  {
			c.drop(f)
		}
	}
}

// drop check the log file creation time and delete the file if the conditions
// are met.
func (c *gclear) drop(f os.FileInfo) {
	if time.Since(f.ModTime()) > c.reserve {
		pathfile:=path.Join(c.path,f.Name())
		glog.Info("删除是日志",pathfile)
		err := os.Remove(pathfile)
		if err != nil {
			glog.Error("删除有问题了",err)
		}

	}
}