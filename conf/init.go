package conf



var errstr struct {
	F_NotType string
	F_OnceLoad string
	F_INIParse string
	F_JsonParse string

	E_JsonParse string
	E_OnceLoad string
	E_Reload string
	E_SaveFaild string
	E_DelKeyFaild string
	E_CreateFaild string
	E_WriteFileFaild string
	E_CopyConfFaild string

	I_Lock string
	I_Unlock string
	I_ModifyKey string
	I_DelKeySuccess string
	I_DelKeyFaild string
	I_ReloadConf string
	I_CopyConfSuccess string
}

func init()  {
	errstr.F_NotType = "不支持类型格式"
	errstr.F_OnceLoad = "第一次载入配置文件失败"
	errstr.F_INIParse = "ini解析失败"
	errstr.F_JsonParse = "json解析失败"

	errstr.E_JsonParse = "json解析失败"
	errstr.E_OnceLoad = "第一次载入配置文件失败"
	errstr.E_Reload = "重新载入配置文件失败"
	errstr.E_SaveFaild = "保存配置到本地失败"
	errstr.E_DelKeyFaild = "key不存在 or 删除key失败"
	errstr.E_CreateFaild = "临时文件创建错误"
	errstr.E_WriteFileFaild = "配置文件写入临时文件错误"
	errstr.E_CopyConfFaild = "copy 配置文件错误"

	errstr.I_Lock = "开始-读取配置文件加锁"
	errstr.I_Unlock = "完成-读取配置文件解锁"
	errstr.I_ModifyKey = "修改配置值"
	errstr.I_DelKeySuccess = "删除key成功"
	errstr.I_DelKeyFaild = "key不存在 or 删除key失败"
	errstr.I_ReloadConf = "reload 配置文件"
	errstr.I_CopyConfSuccess = "copy配置文件成功"

}