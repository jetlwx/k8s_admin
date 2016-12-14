package common

//IsLogin check the user or login
func IsLogin(loginname interface{}) bool {
	if loginname == nil {
		return false
	}
	if loginname.(string) == "" {
		return false
	}
	return true
}
