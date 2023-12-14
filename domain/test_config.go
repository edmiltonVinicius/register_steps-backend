package domain

func StartDependencies() {
	LoadEnv(true)
	ConnectDB()
	StartRedis()
}

func DownDependencies() {
	Conn.Close()
	RedisClient.Conn().Close()
}

func ClearTable(table string) {
	DB.Exec("DELETE FROM " + table)
}

func ClearRedis() {
	RedisClient.FlushAll(Environment.CTX)
}
