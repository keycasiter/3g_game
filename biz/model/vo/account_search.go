package vo

// ################################## GetSgzGameZoneItemList ######################################
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

//https://m.jiaoyimao.com/api2/sgzItemList2022/getSgzGameZoneItemList?gameId=1009207&fcid=10103&osId=3&sort=composite_sorting_score
//&extConditions={"stage_sum":"10","stage":"2","hero":"10014,10016","equip_skill":"28025,28041"}
//&stdCatId=1844450&jymCatId=1844455&filterLowQuality=false
//&keyword=2000区段+1000区段&enforcePlat=3&cid=1844455&page=1&platformId=3
type AccountSearchReq struct {
	//检索条件 {"hero":"10014,10005"}
	DefiniteHeros []string
	//特技要求 {"equip_skill":"28025,28041"}
	DefiniteSkill []string
	//红度要求
	DefiniteStage string
	//总红度要求 {"stage_sum":"10"}
	DefiniteTotalStage string
	//关键字  keyword=2000区段
	Keyword string
	//价格区间 ["5000","12000"]
	PriceRange string
	//查询商品列表页数
	PageSize int
	//指定武将是否必须觉醒
	IsDefiniteHeroMustAwake bool
	//指定武将是否必须开三兵书
	IsDefiniteHeroMustTalent3 bool
	//特技要求
	MustSpecialTech []string
	//指定战法
	MustTactic []string
	//可跨服、公示
	CrossServerAndPublic bool
	//五星武将数量
	FiveStarHeroNum string
	//S战法数量
	SskillNum string
}

type AccountCheckReq struct {
	//检索条件 {"stage":"5","hero":"10014"}
	DefiniteHeros []string
	//红度要求
	DefiniteStage string
	//价格区间 ["5000","12000"]
	PriceRange string
	//指定武将是否必须觉醒
	IsDefiniteHeroMustAwake bool
	//指定武将是否必须开三兵书
	IsDefiniteHeroMustTalent3 bool
	//特技要求
	MustSpecialTech []string
	//指定战法
	MustTactic []string
	//检测的商品链接
	CheckGoodsUrl string
}

type ExtConditions struct {
	//进阶次数
	Stage string `json:"stage"`
	//武将ID
	Hero string `json:"hero"`
	//总进阶次数
	StageSum string `json:"stage_sum"`
	//特技
	EquipSkill string `json:"equip_skill"`
	//五星武将数量
	FiveStarHeroNum string `json:"five_star_hero"`
	//S战法数量
	SskillNum string `json:"s_skill"`
	//可跨服，非公示
	CrossServerAndPublic string `json:"fastCondition"`
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
	Price float64 `json:"price"`
	//商品Url链接
	DetailUrl string `json:"detailUrl"`
}

type GetSgzGameZoneItemListRespResult struct {
	TotalCnt    int64                                       `json:"totalCnt"`
	HasNextPage bool                                        `json:"hasNextPage"`
	GoodsList   []GetSgzGameZoneItemListRespResultGoodsInfo `json:"goodsList"`
}

//################################## AccountItemInfo ######################################

type AccountItemInfo struct {
	ApiData ApiData `json:"apiData"`
}

type ApiData struct {
	//基础信息
	ItemBaseInfo ItemBaseInfo `json:"itemBaseInfo"`
	//商品质量
	ItemQuality ItemQuality `json:"itemQuality"`
	//商品卖点
	SellPointTags []string `json:"sellPointTags"`
	//商品卖点2
	SecondSellPointTags []string `json:"secondSellPointTags"`
	//综合属性
	MultiplePropertyInfo MultiplePropertyInfo `json:"multiplePropertyInfo"`
	//灵犀角色信息
	ItemLingxiRoleDetail ItemLingxiRoleDetail `json:"itemLingxiRoleDetail"`
}

type ItemLingxiRoleDetail struct {
	S3RoleCustomizeInfo S3RoleCustomizeInfo `json:"s3RoleCustomizeInfo"`
}

type S3RoleCustomizeInfo struct {
	//武将信息
	Heros []Heros `json:"heros"`
	//非事件战法
	Skills []Skills `json:"skills"`
	//事件战法
	EventSkills []EventSkills `json:"eventSkills"`
	//库藏
	Storage Storage `json:"storage"`
	//资产
	Currencies []Currencies `json:"currencies"`
}

type Currencies struct {
	CurrencyName string `json:"currencyName"`
	Amount       int64  `json:"amount"`
}

type Skills struct {
	SkillId int64  `json:"skillId"`
	Name    string `json:"name"`
}

type EventSkills struct {
	SkillId int64  `json:"skillId"`
	Name    string `json:"name"`
}

type Storage struct {
	//装备等
	Equipments []Equipments `json:"equipments"`
	//材料
	Materials []Materials `json:"materials"`
}

type Materials struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Star   int64  `json:"star"`
	Amount int64  `json:"amount"`
}

type Equipments struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	Star          int64    `json:"star"`
	AttrDesc      string   `json:"attrDesc"`
	SkillDesc     string   `json:"skillDesc"`
	SkillDescList []string `json:"skillDescList"`
}

type Heros struct {
	//武将ID
	Id      int64  `json:"id"`
	HeroId  int64  `json:"heroId"`
	Name    string `json:"name"`
	Star    int64  `json:"star"`
	IsAwake bool   `json:"isAwake"`
	//阵营
	Camp int64 `json:"camp"`
	//进阶
	Stage int64 `json:"stage"`
	//典藏/SP
	Prefix          string `json:"prefix"`
	SeasonTag       string `json:"seasonTag"`
	IsUnlockTalent  bool   `json:"isUnlockTalent"`
	IsUnlockTalent3 bool   `json:"isUnlockTalent3"`
}

type MultiplePropertyInfo struct {
	//资产等信息
	DigestInfo DigestInfo `json:"digestInfo"`
	//角色、出生服、赛季服等信息
	GamePropertyInfo string `json:"gamePropertyInfo"`
}

type DigestInfo struct {
	DigestGroupNodeList []DigestGroupNodeList `json:"digestGroupNodeList"`
}

type DigestGroupNodeList struct {
	GroupName   string        `json:"groupName"`
	GroupValues []GroupValues `json:"GroupValues"`
	DisplayType string        `json:"displayType"`
}

type GroupValues struct {
	GroupKey string   `json:"groupKey"`
	Values   []Values `json:"values"`
}

type Values struct {
	Value string `json:"value"`
}

type ItemQuality struct {
	FavoriteNum int64 `json:"favoriteNum"`
}

type ItemBaseInfo struct {
	ItemId     string  `json:"itemId"`
	Title      string  `json:"title"`
	SellPrice  float64 `json:"sellPrice"`
	ServerName string  `json:"serverName"`
}
