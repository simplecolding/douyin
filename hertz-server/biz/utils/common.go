package utils

const (
	LocalURL = "192.168.59.153:8888"
	PlayURL = "http://"+LocalURL+"/public/"
	CoverTestURL = "http://"+LocalURL+"/public/cover/covertest.jpg"

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
)
