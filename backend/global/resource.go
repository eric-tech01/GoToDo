package global

// 全局资源信息，从LICENSE文件中解密出来的
var HARD_RESOURCE HardResource = HardResource{License: License{}}

type HardResource struct {
	License License `json:"license"`

	//主控信息（第三方的）
	BoardModel   string `json:"boardModel"`   //主控型号：
	BoardVersion string `json:"boardVersion"` //主控版本 V6.0
	BoardVendor  string `json:"boardVendor"`  //主控供应商:

}

type License struct {
	Key      string `json:"key"`
	Expired  int64  `json:"expired"`
	CreateAt int64  `json:"createAt"`
}
