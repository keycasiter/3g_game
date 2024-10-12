// idl/jym.thrift
namespace go jym

include "idl/common.thrift"
include "idl/enum.thrift"

//============= 账户搜索 BEGIN ==============
struct AccountSearchRequest {
	//阵容
	1: string LineUpList
	//检索条件 {"hero":"10014,10005"}
	2:string DefiniteHeroList
	//特技要求 {"equip_skill":"28025,28041"}
	3:string DefiniteSkillList
	//红度要求
    4:string DefiniteStage
	//总红度要求 {"stage_sum":"10"}
	5:string DefiniteTotalStage
	//关键字  keyword=2000区段
	6:string Keyword
	//价格区间 ["5000","12000"]
	7:string PriceRange
	//翻页数量
	8:i64 PageNum
	//查询商品数
	9:i64 GoodsNum
	//指定武将是否必须觉醒
	10: bool IsDefiniteHeroMustAwake
	//指定武将是否必须开三兵书
	11: bool IsDefiniteHeroMustTalent3
	//特技要求
	12:string MustSpecialTechList
	//指定战法
	13:string MustTacticList
	//可跨服、公示
	14: bool CrossServerAndPublic
	//五星武将数量
	15: string FiveStarHeroNum
	//S战法数量
	16: string SskillNum
}

struct AccountSearchResponse{
    1: common.Meta meta
	2: list<ApiData> ApiDatas
}

struct ApiData {
	//基础信息
	1:ItemBaseInfo ItemBaseInfo
	//商品质量
	2:ItemQuality ItemQuality
	//商品卖点
	3:list<string> SellPointTags
	//商品卖点2
	4:list<string> SecondSellPointTags
	//综合属性
	5:MultiplePropertyInfo MultiplePropertyInfo
	//灵犀角色信息
	6:ItemLingxiRoleDetail ItemLingxiRoleDetail
}

struct ItemLingxiRoleDetail {
	S3RoleCustomizeInfo S3RoleCustomizeInfo
}

struct S3RoleCustomizeInfo {
	//武将信息
	1:list<Heros> Heros
	//非事件战法
	2:list<Skills> Skills
	//事件战法
	3:list<EventSkills> EventSkills
	//库藏
	4:Storage Storage
	//资产
	5:list<Currencies> Currencies
}

struct Currencies {
	1:string CurrencyName
	2:i64 Amount
}

struct Skills {
	1:i64 SkillId
	2:string Name
}

struct EventSkills {
	1:i64 SkillId
	2:string Name
}

struct Storage {
	//装备等
	1:list<Equipments> Equipments
	//材料
	2:list<Materials> Materials
}

struct Materials {
	1:i64 Id
	2:string Name
	3:i64 Star
	4:i64 Amount
}

struct Equipments {
	1:i64 Id
	2:string Name
	3:i64 Star
	4:string AttrDesc
	5:string SkillDesc
	6:list<string> SkillDescList
}

struct Heros {
	//武将ID
	1:i64 Id
	2:i64 HeroId
	3:string Name
	4:i64 Star
	5:bool IsAwake
	//阵营
	6:i64 Camp
	//进阶
	7:i64 Stage
	//典藏/SP
	8:string Prefix
	9:string SeasonTag
	10:bool IsUnlockTalent
	11:bool IsUnlockTalent3
}

struct MultiplePropertyInfo {
	//资产等信息
	1:DigestInfo DigestInfo
	//角色、出生服、赛季服等信息
	2: string GamePropertyInfo
}

struct DigestInfo {
	1:list<DigestGroupNodeList> DigestGroupNodeList
}

struct DigestGroupNodeList {
	1: string GroupName
	2:list<GroupValues> GroupValues
	3: string DisplayType
}

struct GroupValues {
	1:string GroupKey
	2:list<Values> Values
}

struct Values {
	1: string Value
}

struct ItemQuality {
	1: i64 FavoriteNum
}

struct ItemBaseInfo {
	1:string       ItemId
	2:string       Title
	3:double      SellPrice
	4:string       ServerName
	5:string       StatusName
	6:string       OsTypeName
	7:CategoryInfo FirstCategoryInfo
	8:CategoryInfo SecondCategoryInfo
}

struct CategoryInfo {
	1: string CategoryName
}
//============= 账户搜索 END ==============