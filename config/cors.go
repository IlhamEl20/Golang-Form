package config

func AllowHeaders() string {

	return "Origin, Content-Type, Accept, Authorization, X-Requested-With, Cookie"

}

func AllowMethods() string {

	return "GET,POST,PUT,DELETE"

}
