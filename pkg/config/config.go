package config

var XwDB = map[string]interface{}{
	"ENGINE": "mysql",
	"HOST":   "data.fxt.cn",
	//"HOST":"localhost",
	"PORT": 3306,
	"NAME": "sub_readonly", //数据库名称
	"USER": "fxt_read",
	//"USER":"root",
	"PASSWORD": "jNecQlKGCwPWrOt8",
}