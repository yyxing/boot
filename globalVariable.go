package boot

const (
	MysqlDriverFormatter = "%s:%s@%s"
	Int64Min             = ^int(^uint64(0) >> 1)
	Int64Max             = int(^uint64(0) >> 1)
	ServerPortKey        = "server.port"
)
