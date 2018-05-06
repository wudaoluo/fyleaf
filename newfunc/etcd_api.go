package newfunc

import (
	"github.com/coreos/etcd/client"
	"github.com/liu-junyong/go-logger/logger"
	"time"
	"context"
	"strings"
	"path"
	"os"
	"github.com/wudaoluo/goutils/sys"
	"fmt"
	"strconv"
)




var process_start_time string

//虚拟目录,每次启动先将etcd 数据放到这里,在删除 etcd_dir,再mv 这个目录 etcd_dir
var dir_timestamps = strconv.FormatInt(time.Now().Unix(),10)

func init() {
	process_start_time = sys_time()
}


//返回当前系统时间
func sys_time() string {
	time_format := "20060102-15:04:05"
	return time.Unix(time.Now().Unix(), 0).Format(time_format)
}

type worker interface {
	Watch(key string)
	GetRoot() bool
	SyncDir()
}

type work struct {
	conf_dir   string
	keysAPI    client.KeysAPI
	key_map    map[string]keyMD5
}


//key指纹
type keyMD5 struct {
	path string
	md5 string
}

func NewWorker(conf_dir string, endpoints []string) worker {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		logger.Error("初始化etcd失败",err)
	}

	w := &work{
		conf_dir:	conf_dir,
		keysAPI: 	client.NewKeysAPI(etcdClient),
		key_map:    make(map[string]keyMD5),
	}
	return w
}


type TreeNode struct {
	Name     string
	Key      string
	Dir      bool
	Path     string
	Value    string
	Children []*TreeNode
}


func (w *work)GetRoot() bool{
	api := w.keysAPI
	opts := &client.GetOptions{}
	opts.Recursive = true
	resp, err := api.Get(context.Background(),"/",opts)
	if err != nil {
		logger.Error("get 数据错误",err)
		return false
	}

	//w.GetAllNode(resp.Node)
	treeNodes := formatEtcdNodes(resp.Node)
	w.GetAllNode(treeNodes)


	return true
}


func formatEtcdNodes(node *client.Node) *TreeNode {
	treeNode := &TreeNode{}
	arr := strings.Split(node.Key, "/")
	count := len(arr)
	if count < 1 {
		count = 1
	}
	treeNode.Name = node.Key
	treeNode.Key = arr[count-1]
	treeNode.Path = path.Join(dir_timestamps,node.Key)  //临时目录
	treeNode.Dir = node.Dir
	treeNode.Value = node.Value

	for _, v := range node.Nodes {
		treeNode.Children = append(treeNode.Children, formatEtcdNodes(v))
	}
	//必须返回一个空数组供前端渲染
	if treeNode.Dir && len(treeNode.Children) == 0 {
		treeNode.Children = []*TreeNode{}
	}
	return treeNode
}


//程序刚开始运行时,同步 etcd 所有配置文件,存在本地
func (w *work) GetAllNode(treenode *TreeNode) {
	//创建目录下文件 文件名:key-时间戳,内容是 value
	if !treenode.Dir {
		filename := treenode.Path + "-" + process_start_time
		//保存所有 key到文件
		err := save_key(filename,treenode.Value)
		if err != nil {
			logger.Error(err)
		}else {
			logger.Debug("创建文件",filename,"成功")

			//创建文件指纹
			w.key_map[treenode.Name] = keyMD5{
				path:path.Join(w.conf_dir,treenode.Name) + "-" + process_start_time,
				md5:sys.Get_md5(treenode.Value),
			}
		}
	}

	if treenode.Dir && !sys.PathExists(treenode.Path){

		err := os.MkdirAll(treenode.Path,0755)
		if err != nil {
			logger.Error("创建目录失败",treenode.Path,err)
		} else {
			logger.Debug("创建目录",treenode.Path,"成功")
		}

	}

	for node := range treenode.Children {
		w.GetAllNode(treenode.Children[node])
	}
}


//写入文件
func save_key(filename,value string) error{
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件出现问题了",err)
	}

	_,err = f.WriteString(value)
	if err != nil {
		return fmt.Errorf("文件写入数据失败",err)
	}
	f.Sync()
	return nil
}


func (w *work) SyncDir() {
	//rm old;mv new old
	err := os.RemoveAll(w.conf_dir)
	if err != nil {
		logger.Error(err)
	}
	err = os.Rename(dir_timestamps,w.conf_dir)
	if err != nil {
		logger.Error(err)
	}

}

func (w *work)Watch(key string) {
	api := w.keysAPI
	watcher := api.Watcher(key, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			logger.Error("watch 错误",key,err)
			time.Sleep(time.Second*10)
			continue
		}

		ress := res
		go func() {
			//res.Action == "delete" res.Action == "expire"....
			fmt.Println("ress",w.key_map)
			if ress.Node.Dir {
				//删除目录
				//dir := path.Join(w.conf_dir,ress.Node.Key)
				//logger.Warn("删除目录",dir)
				//os.RemoveAll(dir)
			}else {
				//获取文件名称
				//if _, ok := w.key_map[ress.Node.Key]; ok {
				//	//删除文件
				//	old_filename := w.key_map[ress.Node.Key].path
				//	logger.Warn("删除文件",old_filename)
				//	os.Remove(old_filename)
				//	delete(w.key_map,ress.Node.Key)  //删除指纹
				//}
			}

			if ress.Action == "set" || ress.Action == "update" {
				if !ress.Node.Dir {
					//创建一个新的文件
					filename := path.Join(w.conf_dir,ress.Node.Key) + "-" + sys_time()
					logger.Debug("watchd到新的数据,创建新的文件",filename)
					save_key(filename,ress.Node.Value)
					//创建文件指纹
					w.key_map[ress.Node.Key] = keyMD5{
						path:path.Join(w.conf_dir,ress.Node.Key) + "-" + sys_time(),
						md5:sys.Get_md5(ress.Node.Value),
					}


				}else {
					//如果是目录且文件存在创建
					dirpath := path.Join(w.conf_dir,ress.Node.Key)
					logger.Debug("watchd到新的数据,创建新的目录",dirpath)
					if !sys.PathExists(dirpath) {
						err := os.MkdirAll(dirpath,0755)
						if err != nil {
							logger.Error("创建目录失败",dirpath,err)
						} else {
							logger.Debug("创建目录",dirpath,"成功")
						}
					}
				}
			}
		}()


	}
}


//指纹匹配测试
//func (w *work)Match(ctx context.Context,interval time.Duration) {
//	timer := time.NewTimer(time.Second * interval)
//	for {
//		// try to read from channel, block at most 5s.
//		// if timeout, print time event and go on loop.
//		// if read a message which is not the type we want(we want true, not false),
//		// retry to read.
//		//if !timer.Stop() {
//		//	select {
//		//	case <-timer.C: //try to drain from the channel
//		//	default:
//		//	}
//		//}
//		timer.Reset(time.Second * interval)
//		select {
//		case ctx.Done():
//			logger.Error("ctx.Done 退出指纹匹配进程")
//			break
//		case <-timer.C:
//			logger.Debug("指纹测试..")
//
//
//
//
//
//
//			continue
//		}
//	}
//}

