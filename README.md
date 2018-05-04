# fyleaf


### 计划实现的功能
    log --> glog，按大小分割，每个文件100MB   完成  2018.5.1
    去掉 chanrpc
    去掉 cluster
    去掉 recordile
    增加 etcd
    增加 telnet 修改配置文件(方便调试,优先级低),或者 telnet只有看服务器状态权限
    增加 版本发现(consul+自主接口)
    增加 goroutine 使用接口,内存使用, cpu使用,运行时间,服务名称版本号   完成  2018.5.1 (cpu暂时没有找到方法收集)
    增加 计划任务(这个可以通过 telnet 查看)
    增加 models(mongodb,mysql,redis)  redis等后期在实现
    增加 配置文件支持的格式 (ini json)
    增加 network (copy leaf)
    增加 peer （ws/tcp
    所有模块都增加test,和README


### 计划使用teleport 的设计方式
https://raw.githubusercontent.com/henrylee2cn/teleport/master/doc/teleport_framework.png

# 修改日志
### 2018.5.1
    添加 glog
    
### 2018.5.4 
    添加 glog 日志定期清除功能和框架错误码

### 融合进框架的库
     github.com/xuri/glc     glog日志清除,修改
     github.com/golang/glog  日志库,修改