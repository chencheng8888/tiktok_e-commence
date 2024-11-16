package model

//user服务的相关操作
//只是用来返回其操作的name

// User 实现了casbin.CheckActer 接口
type User struct{}

func (u User) String() string {
	return "user"
}

// act

func (u User) Create() string {
	return "CREATE"
}
func (u User) Login() string {
	return "LOGIN"
}
func (u User) Logout() string {
	return "LOGOUT"
}
func (u User) Delete() string {
	return "DELETE"
}
func (u User) Update() string {
	return "UPDATE"
}
func (u User) GetInfo() string {
	return "GET"
}

// CheckActExist 检查对于user,act是否合法
func (u User) CheckActExist(act string) bool {
	switch act {
	case u.Create(), u.Login(), u.Logout(), u.Delete(), u.Update(), u.GetInfo():
	default:
		return false
	}
	return true
}
