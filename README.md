# fyleaf


### 难点
    每个包下面的 agent.go 是负责读取数据,写入数据等

### 计划实现的功能
  *  log --> glog，按大小分割，每个文件100MB   完成  2018.5.1
  *  去掉 chanrpc
  *  去掉 cluster
  *  去掉 recordile
  * 增加 etcd
  * 增加 tcp 接口,用来获取程序运行状态(telnet 分析工具go tool pprof 20180508_22_40_42.gprof)
    增加 telnet 修改配置文件(方便调试,优先级低),或者 telnet只有看服务器状态权限
  *  增加 版本发现(consul+自主接口)  正在写  2018.5.6
  *  增加 goroutine 使用接口,内存使用, cpu使用,运行时间,服务名称版本号   完成  2018.5.1 (cpu暂时没有找到方法收集)
    增加 计划任务(这个可以通过 telnet 查看)
  *  增加 models(mongodb,mysql,redis)  redis等后期在实现
  *  增加 配置文件支持的格式 (ini json)       完成  2018.5.5
    增加 network (copy leaf)
    增加 peer （ws/tcp）
  *  所有模块都增加test,和README
   增加 gate //负责启动wss tcp，路由


### 后期计划实现的功能
    goroutine池
    通过tcp接口查看程序运行状态


### 计划使用teleport 的设计方式
https://raw.githubusercontent.com/henrylee2cn/teleport/master/doc/teleport_framework.png

# 修改日志
### 2018.5.1
    * 添加 glog
    
### 2018.5.4 
    * 添加 glog 日志定期清除功能和框架错误码

### 2018.5.5
    * json解析使用 github.com/json-iterator/go 性能是原生的6倍（官方说法）
    ```
        jsoniter.Marshal(&data)
        jsoniter.Unmarshal(input, &data)
    ```
    * 增加conf(json,ini)两总配置文件解析
    ```
        a := conf.GetInstance()
    	a.ParseConf("../cmd/server.json","json")  //json,ini支持两种格式
    ```
    
### 2018.5.6 
    * 在每个子包下面创建 init() 并将 errstr 在这初始化
    (使用错误码 提升程序性能,具体提升多少,还需要测试一下)
    * 添加 etcd   配置文件获取
    * 添加 consul 版本发现
    * 搞fyleaf入口文件
    * copy leaf module
    * 改动一些conf代码 可以在获取配置文件之后 初始化glog
    * 开始写 console 和 gate
    
    
### 2018.5.7 
    * json.go 增加默认值,加载 ini 和 json配置,尽量让程序行为一直
    * 改动 tcp_server,使用atomic.LoadInt32 代替len(server.conns)
    * 注释 wss_server
    
    
### 2018.5.8 
    增加 tcp 接口,用来获取程序运行状态

### 使用到库
     github.com/xuri/glc              //glog日志清除,修改  (改进后融入)
     github.com/golang/glog           //日志库,修改        (改进后融入)
     github.com/json-iterator/go      //json解析
     github.com/Unknwon/goconfig      //ini配置文件
     github.com/gorilla/websocket
     github.com/go-sql-driver/mysql
     gopkg.in/mgo.v2
     github.com/go-redis/redis         //后期加入redis支持

### 测试(备注一下)
    go test -run='Test_IniConf'
    go test -file mysql_test.go
    go test


### *.gprof 分析
    go tool pprof 20180508_22_40_42.gprof
    具体使用 google pprof

### 参与贡献
    1.Fork 本项目
    2.新建 Feat_xxx 分支
    3.提交代码
    4.新建 Pull Request