package config

var Config = make(map[string]string)

func init() {
	Config["ConnectionString"] = "root:root@tcp(127.0.0.1:3306)/hitch?parseTime=true"
}
