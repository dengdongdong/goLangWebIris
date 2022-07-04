package model

import "time"

type User struct {
	//todo
	AdminId    int64     `xorm:"pk autoincr 'id'"  json:"id"` // 主键⾃增
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"timestamp"    json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`      //管理员密码
	CityName   string    `xorm:"varchar(12)" json:"city_name"` //管理员所在城市名称
	CityId     int64     `xorm:"index" json:"city_id"`         //设置索引
}
