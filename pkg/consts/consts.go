package consts

const (

	//Logger format for echo logger middleware
	LoggerFormat = "${time_rfc3339}, method = ${method}, uri = ${uri}," +
		" status = ${status}, remote_ip = ${remote_ip}\n"

	HostAddress        = "0.0.0.0:8080"
	ConnStr            = "user=postgres password=7396 dbname=db_tp sslmode=disable"
	//ConnStr            = "host=db user=postgres password=7396 dbname=sunrise_db sslmode=disable"//"host=my_postgres user=postgres password=7396 dbname=sunrise_db sslmode=disable"
	NumberOfPinsOnPage = 10
)
