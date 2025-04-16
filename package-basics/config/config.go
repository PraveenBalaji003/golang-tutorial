package config

const (
	USER_NAME string = "postgres"
	password  string = "ZAQ1@wsx"
)

func GetUserName() string {
	return USER_NAME
}

func getPassword() string {
	return password
}
