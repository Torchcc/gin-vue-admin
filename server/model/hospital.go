// 自动生成模板Hospital
package model

// 如果含有time.Time 请自行import time包
type Hospital struct {
	ID         uint    `json:"id" form:"id" gorm:"column:id;comment:'id'"`
	Name       string  `json:"name" form:"name" gorm:"column:name;comment:'医院名称';type:varchar(256)"`
	Address    string  `json:"address" form:"address" gorm:"column:address;comment:'医院地址，街道门牌';type:varchar(20)"`
	AvatarUrl  string  `json:"avatar_url" form:"avatar_url" gorm:"column:avatar_url;comment:'医院头像';type:varchar(128)"`
	Level      int     `json:"level" form:"level" gorm:"column:level;comment:'0.无级别;1-公立三甲 2-公立医院 3-民营医院 4-专业机构';type:tinyint(4)"`
	ProvinceId int     `json:"province_id" form:"province_id" gorm:"column:province_id;comment:'省行政id';type:int(10)"`
	CityId     int     `json:"city_id" form:"city_id" gorm:"column:city_id;comment:'市行政id';type:int(10)"`
	CountyId   int     `json:"county_id" form:"county_id" gorm:"column:county_id;comment:'区行政id';type:int(10)"`
	TownId     int     `json:"town_id" form:"town_id" gorm:"column:town_id;comment:'镇行政id';type:int(10)"`
	Longitude  float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:'经度';type:decimal(10,7)"`
	Latitude   float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:'纬度';type:decimal(10,7)"`
	CreateTime int     `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间';type:int(10)"`
	UpdateTime int     `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间';type:int(10)"`
	IsDeleted  int     `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:'是否删除';type:tinyint(4)"`
}

type ListHospitalOutputEle struct {
	ID           uint    `json:"id" form:"id" gorm:"column:id;comment:'id'"`
	Name         string  `json:"name" form:"name" gorm:"column:name;comment:'医院名称';type:varchar(256)"`
	Address      string  `json:"address" form:"address" gorm:"column:address;comment:'医院地址，街道门牌';type:varchar(20)"`
	AvatarUrl    string  `json:"avatar_url" form:"avatar_url" gorm:"column:avatar_url;comment:'医院头像';type:varchar(128)"`
	Level        int     `json:"level" form:"level" gorm:"column:level;comment:'0.无级别;1-公立三甲 2-公立医院 3-民营医院 4-专业机构';type:tinyint(4)"`
	ProvinceId   int     `json:"province_id" form:"province_id" gorm:"column:province_id;comment:'省行政id';type:int(10)"`
	CityId       int     `json:"city_id" form:"city_id" gorm:"column:city_id;comment:'市行政id';type:int(10)"`
	CountyId     int     `json:"county_id" form:"county_id" gorm:"column:county_id;comment:'区行政id';type:int(10)"`
	TownId       int     `json:"town_id" form:"town_id" gorm:"column:town_id;comment:'镇行政id';type:int(10)"`
	ProvinceName string  `json:"province_name"`
	CityName     string  `json:"city_name"`
	CountyName   string  `json:"county_name"`
	TownName     string  `json:"town_name"`
	LevelName    string  `json:"level_name"`
	Longitude    float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:'经度';type:decimal(10,7)"`
	Latitude     float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:'纬度';type:decimal(10,7)"`
	CreateTime   int     `json:"create_time" form:"create_time" gorm:"column:create_time;comment:'创建时间';type:int(10)"`
	UpdateTime   int     `json:"update_time" form:"update_time" gorm:"column:update_time;comment:'更新时间';type:int(10)"`
}

func (Hospital) TableName() string {
	return "mkh_hospital"
}

func (ListHospitalOutputEle) TableName() string {
	return "mkh_hospital"
}
