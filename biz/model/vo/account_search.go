package vo

type GetSgzGameZoneItemListReq struct {
	GameId     int64 `json:"gameId"`
	Fcid       int64 `json:"fcid"`
	OsId       int64 `json:"osid"`
	Cid        int64 `json:"cid"`
	PlatformId int64 `json:"platformId"`
	//排序
	Sort          string `json:"sort"`
	ExtConditions string `json:"extConditions"`
	StdCatId      int64  `json:"stdCatId"`
	JymCatId      int64  `json:"jymCatId"`
	//
	FilterLowQuality string `json:"filterLowQuality"`
	//关键字
	Keyword string `json:"keyword"`
	//价格区间 ["5000","12000"]
	PriceRange  string `json:"priceRange"`
	enforcePlat int64  `json:"enforcePlat"`
	//翻页 从1开始
	Page int64 `json:"page"`
}

type GetSgzGameZoneItemListResp struct {
	success bool                             `json:"success"`
	result  GetSgzGameZoneItemListRespResult `json:"result"`
}

type GetSgzGameZoneItemListRespResultGoodsInfo struct {
	//商品ID
	goodsId int64 `json:"goodsId"`
	//商品标题
	realTitle string `json:"realTitle"`
	//当前所在服
	serverName string `json:"serverName"`
	//赛季服
	seasonServerName string `json:"seasonServerName"`
	//价格
	price int64 `json:"price"`
	//商品Url链接
	detailUrl string `json:"detailUrl"`
}

type GetSgzGameZoneItemListRespResult struct {
	totalCnt    int64                                       `json:"totalCnt"`
	hasNextPage bool                                        `json:"hasNextPage"`
	goodsList   []GetSgzGameZoneItemListRespResultGoodsInfo `json:"goodsList"`
}
