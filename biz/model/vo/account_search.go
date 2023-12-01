package vo

type GetSgzGameZoneItemListReq struct {
	GameId     int64 `json:"gameId"`
	Fcid       int64 `json:"fcid"`
	OsId       int64 `json:"osid"`
	Cid        int64 `json:"cid"`
	PlatformId int64 `json:"platformId"`
	//排序
	Sort string `json:"sort"`
	//检索条件 {"stage":"5","hero":"10014"}
	ExtConditions string `json:"extConditions"`
	StdCatId      int64  `json:"stdCatId"`
	JymCatId      int64  `json:"jymCatId"`
	//
	FilterLowQuality bool `json:"filterLowQuality"`
	//关键字
	Keyword string `json:"keyword"`
	//价格区间 ["5000","12000"]
	PriceRange  string `json:"priceRange"`
	enforcePlat int64  `json:"enforcePlat"`
	//翻页 从1开始
	Page int64 `json:"page"`
}

type ExtConditions struct {
	//进阶次数
	Stage string `json:"stage"`
	//武将ID
	Hero string `json:"hero"`
}

type GetSgzGameZoneItemListResp struct {
	Success bool                             `json:"success"`
	Result  GetSgzGameZoneItemListRespResult `json:"result"`
}

type GetSgzGameZoneItemListRespResultGoodsInfo struct {
	//商品ID
	GoodsId int64 `json:"goodsId"`
	//商品标题
	RealTitle string `json:"realTitle"`
	//当前所在服
	ServerName string `json:"serverName"`
	//赛季服
	SeasonServerName string `json:"seasonServerName"`
	//价格
	Price int64 `json:"price"`
	//商品Url链接
	DetailUrl string `json:"detailUrl"`
}

type GetSgzGameZoneItemListRespResult struct {
	TotalCnt    int64                                       `json:"totalCnt"`
	HasNextPage bool                                        `json:"hasNextPage"`
	GoodsList   []GetSgzGameZoneItemListRespResultGoodsInfo `json:"goodsList"`
}
