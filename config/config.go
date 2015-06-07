package config

var Config = struct {
	ConnectionString string
}{
  ConnectionString: "root:root@tcp(127.0.0.1:3306)/hitch?parseTime=true",
}
