package constants
const (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	PublicUrl = "http://192.168.47.153:8888/public/"
	UserTableName = "user"
	StatusSuccess = 0
	StatusFail = 1
	StatusMsgSuccess = "success"
	StatusMsgFail = "fail"
)