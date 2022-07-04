package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"goWebIris/testWebIris/main/model"
)

type Adminservice interface {
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine) *adminService {
	return &adminService{
		mysqlEngine: db,
	}
}

//使用Mysql的实现业务逻辑；如果需要切换数据库的话，此处可以实现接口并实现相应的业务逻辑。
type adminService struct {
	//实现Mysql
	mysqlEngine *xorm.Engine
}

func (as *adminService) GetAdminCount() (int64, error) {
	var admin model.Admin
	count, err := as.mysqlEngine.Count(&admin)
	fmt.Println(count)
	if err != nil {
		panic(err.Error())
	}
	return count, err
}
func (as *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin
	//开启事物
	session := as.mysqlEngine.NewSession()
	session.Begin()
	_, err := as.mysqlEngine.Where("admin_name = ? and pwd =?", username, password).Get(&admin)
	if err != nil {
		session.Rollback()
		session.Close()
	}
	session.Commit()
	session.Close()
	return admin, admin.AdminId != 0
}
