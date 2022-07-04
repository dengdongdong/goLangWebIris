package dataSource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"goWebIris/testWebIris/main/model"
	"xorm.io/core"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:cape@tcp(127.0.0.1:3306)/golangiristest?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//设置数据库名称映射规则，
	engine.SetMapper(core.SnakeMapper{}) //驼峰式名称【结构体大写字母，在数据库中映射下划线隔开】
	//engine.SetMapper(core.SameMapper{})  //数据库字段跟结构体相同
	// GonicMapper implements IMapper. It will consider initialisms when mapping names.
	// E.g. id -> ID, user -> User and to table names: UserID -> user_id, MyUID -> my_uid
	//engine.SetMapper(core.GonicMapper{}) //

	//数据库迁移【根据结构体生成，表结构】
	engine.Sync2(
		new(model.Admin),
	)
	//defer engine.Close()
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxOpenConns(10)
	return engine
}
