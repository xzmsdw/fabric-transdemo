package model

// Selling 销售要约
// 需要确定ObjectOfSale是否属于Seller
// 买家初始为空
// Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的销售
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //销售对象(正在出售的房地产RealEstateID)
	Seller        string  `json:"seller"`        //发起销售人、卖家(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与销售人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

// SellingStatusConstant 销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// Donating 捐赠要约
// 需要确定ObjectOfDonating是否属于Donor
// 需要指定受赠人Grantee，并等待受赠人同意接收
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //捐赠对象(正在捐赠的房地产RealEstateID)
	Donor            string `json:"donor"`            //捐赠人(捐赠人AccountId)
	Grantee          string `json:"grantee"`          //受赠人(受赠人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //捐赠状态
}

// DonatingStatusConstant 捐赠状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "捐赠中", //捐赠人发起捐赠合约，等待受赠人确认受赠
		"cancelled":     "已取消", //捐赠人在受赠人确认受赠之前取消捐赠或受赠人取消接收受赠
		"done":          "完成",  //受赠人确认接收，交易完成
	}
}

// Crops 作物信息
type Crops struct {
	//作物ID
	CropsId string `json:"crops_id"`
	//作物名称
	CropsName string `json:"crops_name"`
	//所在地
	Address string `json:"address"`
	//生长开始日期
	RegisterTime string `json:"register_time"`
	//年度
	Year string `json:"year"`
	//农户名字
	FarmerName string `json:"farmer_name"`
	//农户ID
	FarmerID string `json:"farmer_id"`
	//联系电话
	FarmerTel string `json:"farmer_tel"`
	//肥料名称
	FertilizerName string `json:"fertilizer_name"`
	//种植方式
	PlatMode string `json:"plant_mode"`
	//是否套袋种植
	BaggingStatus string `json:"bagging_status"`
	//育苗周期
	GrowSeedlingsCycle string `json:"grow_seedlings_cycle"`
	//灌溉周期
	IrrigationCycle string `json:"irrigation_cycle"`
	//施肥周期
	ApplyFertilizerCycle string `json:"apply_fertilizer_cycle"`
	//除草周期
	WeedCycle string `json:"weed_cycle"`
	//备注
	Remarks string `json:"remarks"`
}

// CropsGrowInfo 作物生长信息
type CropsGrowInfo struct {
	//生长情况唯一ID
	CropsGrowId string `json:"crops_grow_id"`
	//作物ID
	CropsBakId string `json:"crops_bak_id"`
	//记录时间
	RecordTime string `json:"record_time"`
	//作物生长图片URL
	CropsGrowPhotoUrl string `json:"crops_grow_photo_url"`
	//温度
	Temperature string `json:"temperature"`
	//生长情况
	GrowStatus string `json:"grow_status"`
	//水分
	WaterContent string `json:"water_content"`
	//光照情况
	IlluminationStatus string `json:"illumination_status"`
	//备注
	Remarks string `json:"remarks"`
}

// Transport 物流信息
type Transport struct {
	//物流ID(此链中唯一)
	TransportId string `json:"transport_id"`
	//司机ID
	DriverId string `json:"driver_id"`
	//司机名字
	DriverName string `json:"driver_name"`
	//司机电话
	DriverTel string `json:"driver_tel"`
	//所属部门
	DriverDept string `json:"driver_dept"`
	//货物ID
	CropsId string `json:"crops_id"`
	//物流信息上链时间
	TransportToChainTime string `json:"transport_to_chain_time"`
	//物流路过地址
	TransportToAddress string `json:"transport_to_address"`
	//备注（始发地，途中，目的地）
	Remarks string `json:"remarks"`
}

// Machining 加工信息
type Machining struct {
	//加工ID
	MachiningId string `json:"machining_id"`
	//原料厂商负责人
	Leader string `json:"leader"`
	//货物ID
	CropsId string `json:"crops_id"`
	//厂商负责人tel
	LeaderTel string `json:"leader_tel"`
	//厂商名称
	FactoryName string `json:"factory_name"`
	//检测结果
	TestingResult string `json:"testing_result"`
	//入库时间
	InFactoryTime string `json:"in_factory_time"`
	//出库时间
	OutFactoryTime string `json:"out_factory_time"`
	//质检过程图片
	TestingPhotoUrl string `json:"testing_photo_url"`
	//备注
	Remarks string `json:"remarks"`
}

// Sell 交易信息
type Sell struct {
	//交易ID
	SellID string `json:"sell_id"`
	//作物ID
	CropsID string `json:"crops_id"`
	//销售商ID
	SellerID string `json:"seller_id"`
	//买家ID
	BuyerID string `json:"buyer_id"`
	//交易价格
	Price string `json:"price"`
}

// Commodity 商品信息
type Commodity struct {
	//作物ID
	CropsID string `json:"crops_id"`
	//销售商ID
	SellerID string `json:"seller_id"`
	//价格
	Price string `json:"price"`
	//销售状态
	State string `json:"state"`
}
type Capital struct {
	UserID  string `json:"user_id"`
	Balance string `json:"balance"`
}
type User struct {
	UserID string `json:"user_id"`
	Passwd string `json:"passwd"`
}

const (
	FarmerKey     = "farmer_key"
	FarmerGrowKey = "farmer_grow_key"
	DriverKey     = "driver_key"
	MaterialKey   = "material_key"
)
