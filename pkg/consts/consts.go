package consts

const (
	LoggerFormat = "${time_rfc3339}, method = ${method}, uri = ${uri}," +
		" status = ${status}, remote_ip = ${remote_ip}\n"

	HostAddress        = "0.0.0.0:5000"
	ConnStr            = "user=forum password=forum dbname=forum sslmode=disable"
)
