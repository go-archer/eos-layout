package response

type Area struct {
	ID         int64   `json:"id"`          // 行政代码
	ZipCode    int32   `json:"zip_code"`    // 邮政编码
	CityCode   string  `json:"city_code"`   // 区号
	Name       string  `json:"name"`        // 名称
	ShortName  string  `json:"short_name"`  // 简称
	MergerName string  `json:"merger_name"` // 组合名
	Pinyin     string  `json:"pinyin"`      // 拼音
	Lng        float64 `json:"lng"`         // 经度
	Lat        float64 `json:"lat"`         // 纬度
}

type AreaList struct {
	Items []*Area `json:"items"`
}
