package model

//item服务的相关操作
//只是用来返回其操作的name

// Item 实现了casbin.CheckActer 接口
type Item struct{}

func (i Item) String() string {
	return "item"
}

// act

func (i Item) Create() string {
	return "CREATE"
}
func (i Item) UpdateInfo() string {
	return "UPDATE"
}
func (i Item) Delete() string {
	return "DELETE"
}
func (i Item) GetInfo() string {
	return "GET"
}

// CheckActExist 检查对于item,act是否合法
func (i Item) CheckActExist(act string) bool {
	switch act {
	case i.Create(), i.UpdateInfo(), i.Delete(), i.GetInfo():
	default:
		return false
	}
	return true
}
