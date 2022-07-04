package model

import "time"

type Admin struct {

	//如果field名称为Id，⽽且类型为int64，并没有定义tag，则会被xorm视为主键，并且拥有⾃增属性
	AdminId    int64     `xorm:"pk autoincr 'id'"  json:"id"` // 主键⾃增
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"timestamp"    json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`      //管理员密码
	CityName   string    `xorm:"varchar(12)" json:"city_name"` //管理员所在城市名称
	CityId     int64     `xorm:"index" json:"city_id"`
	//City       *City     `xorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
	//city City `xorm:"-"` //不映射该字段
}
type City struct {
	cityName      string
	CityLongitude float32
	CityLatitude  float32
}
