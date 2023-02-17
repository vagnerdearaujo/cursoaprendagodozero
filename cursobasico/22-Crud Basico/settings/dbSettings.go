package settings

type DBSettings struct {
	ServerPort           string
	DBUser               string
	DBPasswd             string
	DBDriver             string
	DBName               string
	DBServer             string
	ConnectionParameters string
}

func MySQLSettings() DBSettings {
	return DBSettings{"5932",
		"golang_devbook",
		"devbook_golang",
		"mysql",
		"devbook",
		"172.18.0.2:3306",
		"?charset=utf8&parseTime=True&loc=Local"}
}
