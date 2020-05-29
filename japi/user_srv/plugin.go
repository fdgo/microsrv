package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "microservice/jzapi/lib/db"
	_ "microservice/jzapi/lib/redis"
	_ "microservice/jzapi/user_srv/initial"
)
