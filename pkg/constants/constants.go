package constants
const (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	UserTableName = "user"

	statusSuccess = 0
	statusFail = 1

	statusMsgSuccess = "success"
	statusMsgFail = "fail"
)