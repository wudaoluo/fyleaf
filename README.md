# fyleaf


### 计划实现的功能
    log --> glog，按大小分割，每个文件100MB
    去掉 chanrpc
    去掉 cluster
    去掉 recordile
    增加 etcd
    增加 telnet 修改配置文件(方便调试,优先级低),或者 telnet只有看服务器状态权限
    增加 版本发现(consul+自主接口)
    增加 goroutine 使用接口,内存使用, cpu使用,运行时间,服务名称
    增加 计划任务(这个可以通过 telnet 查看)
    增加 models(mongodb,mysql,redis)
    增加 配置文件支持的格式 (ini json)