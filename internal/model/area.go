package model

type Area struct {
	ID         int64   `gorm:"id"`          // 自增编号
	Level      int32   `gorm:"level"`       // 区域级别
	ParentCode int64   `gorm:"parent_code"` // 上级区域编码
	AreaCode   int64   `gorm:"area_code"`   // 区域编码
	ZipCode    int32   `gorm:"zip_code"`    // 邮政编码
	CityCode   string  `gorm:"city_code"`   // 区号
	Name       string  `gorm:"name"`        // 名称
	ShortName  string  `gorm:"short_name"`  // 简称
	MergerName string  `gorm:"merger_name"` // 组合名
	Pinyin     string  `gorm:"pinyin"`      // 拼音
	Lng        float64 `gorm:"lng"`         // 经度
	Lat        float64 `gorm:"lat"`         // 纬度
}

func (m *Area) TableName() string {
	return "area"
}
