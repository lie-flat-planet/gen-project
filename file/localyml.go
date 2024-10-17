package file

type LocalYML struct {
}

func (*LocalYML) Name() string {
	return "local.yml"
}

func (*LocalYML) Content() string {
	return `Mysql_DbName: dbName # you own's mysql db name
Mysql_Host: 127.0.0.1:3306 # your own's mysql host
Mysql_Password: xxx # your own's mysql password
Redis_Host: 127.0.0.1:6379 # your own's redis password
Redis_Password: xxx # your own redis password
`
}
