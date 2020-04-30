package cf

type Verify struct {
	Keys string `ini:"keys"`
}

type HTTPServerConfig struct {
	Addr 	string	`ini:"addr"`
	AllowIp string	`ini:"allow_ip"`
}

type MysqlConfig struct {
	DBHost      string `ini:"db_host"`
	DBName      string `ini:"db_name"`
	User        string `ini:"user"`
	Pwd	        string `ini:"pwd"`
	Collation   string `ini:"collation"`
	MaxIdleConn string `ini:"max_idle_conn"`
	MaxOpenConn string `ini:"max_open_conn"`
}

type RedisConfig struct {
	Addr     string `ini:"addr"`
	Index    int    `ini:"index"`
	Pwd	     string `ini:"pwd"`
	PoolSize int    `ini:"pool_size"`
}

type CacheRedisConfig struct {
	Addr     string `ini:"addr"`
	Index    int    `ini:"index"`
	Pwd	     string `ini:"pwd"`
	PoolSize int    `ini:"pool_size"`
}

type TaskRedisConfig struct {
	Addr     string `ini:"addr"`
	Index    int    `ini:"index"`
	Pwd	     string `ini:"pwd"`
	PoolSize int    `ini:"pool_size"`
}

type OSSConfig struct {
	AccessKeyID     string `ini:"access_id"`
	AccessKeySecret string `ini:"access_secret"`
	RoleArn         string `ini:"role_arn"`
	SessionName     string `ini:"session_name"`
}

type Trace struct {
	TraceId     string
	SpanId      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

type TraceContext struct {
	Trace
	CSpanId string
}