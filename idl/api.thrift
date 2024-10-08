// idl/api.thrift
namespace go api

include "idl/common.thrift"
include "idl/enum.thrift"
include "idl/jym.thrift"

//============= 模拟对战 BEGIN ==============

struct BattleDoRequest {
	/** 队伍信息 **/
	// 出战队伍信息
	1:BattleTeam FightingTeam
	// 对战队伍信息
	2:BattleTeam EnemyTeam
}

// 对战队伍信息
struct BattleTeam {
	/** 队伍基础信息 **/
	//队伍类型
	1:enum.TeamType TeamType
	//队伍兵种
	2:enum.ArmType ArmType
	//队伍武将信息
	3:list<BattleGeneral> BattleGenerals
	//兵战科技-属性加成
	4:BuildingTechAttrAddition BuildingTechAttrAddition
	//协力科技-阵营加成
	5:BuildingTechGroupAddition BuildingTechGroupAddition
	//队伍总兵力
	6:i64 SoliderNum
	//队伍剩余兵力
	7:i64 RemainNum
	//队伍名称
	8:string Name
}

// 对战武将信息
struct BattleGeneral {
	//基础信息
	1:MetadataGeneral BaseInfo
	//佩戴战法
	2:list<Tactics> EquipTactics
	//佩戴兵书
	3: list<WarBook> WarBooks
	//佩戴装备特技
	4: list<SpecialTech> SpecialTechs
	//武将对战加成
	5:BattleGeneralAddition Addition
	//是否主将
	6: bool IsMaster
	//携带兵力
	7: i64 SoldierNum
	//兵种适性
    8: enum.ArmsAbility ArmsAbility
    //剩余兵力
    9: i64 RemainNum
}

// 武将对战加成
struct BattleGeneralAddition {
	//1. 武将加成
	//1.a. 加点加成
	1:AbilityAttr AbilityAttr
	//1.b. 等级加成
	2:i64 GeneralLevel
	//1.c. 红度加成
	3:enum.GeneralStarLevel GeneralStarLevel
	//2. 装备加成
	//TODO
	//3. 特技加成
	//TODO
}

// 能力属性
struct AbilityAttr {
	1: string ForceBase
	2: string IntelligenceBase
	3: string CommandBase
	4: string SpeedBase
}

// 武将资料
struct MetadataGeneral {
	//mongodb 持久化字段
	1:i64 Id
	2:string Name
	3:enum.Gender Gender
	4:enum.Group Group
	5:list<enum.GeneralTag> GeneralTag
	6:string AvatarUrl
	7:AbilityAttr AbilityAttr
	8:ArmsAttr ArmsAttr

	//业务字段
	9:enum.GeneralBattleType GeneralBattleType
	//自带战法
	10:Tactics SelfTactic
	11:enum.GeneralQuality GeneralQuality,
	//动态
	12:bool IsSupportDynamics
	//典藏
	13:bool IsSupportCollect
}

// 战法资料
struct Tactics {
	1:i64                Id
	2:string             Name
	3:enum.TacticsSource TacticsSource
	4:enum.TacticsType   Type
	5:enum.TacticQuality Quality
	6:double TriggerRate
	7:string Desc
}

// 兵种属性
struct ArmsAttr{
	1:enum.ArmsAbility Cavalry
	2:enum.ArmsAbility Mauler
	3:enum.ArmsAbility Archers
	4:enum.ArmsAbility Spearman
	5:enum.ArmsAbility Apparatus
}

// 建筑科技属性加成
struct BuildingTechAttrAddition {
	1: string ForceAddition
	2: string IntelligenceAddition
	3: string CommandAddition
	4: string SpeedAddition
}

// 建筑科技阵营加成
struct BuildingTechGroupAddition {
	1: string GroupWeiGuoRate
	2: string GroupShuGuoRate
	3: string GroupWuGuoRate
	4: string GroupQunXiongRate
}

struct BattleDoResponse {
  1: common.Meta meta
  //对战数据统计
  2: BattleResultStatistics BattleResultStatistics
  //对战过程数据 map<战斗阶段,map<战斗回合,战报内容list>>
  3: map<i64,map<i64,list<string>>> BattleProcessStatistics
}

//对战数据统计
struct BattleResultStatistics  {
	//我军统计
	TeamBattleStatistics FightingTeam
	//敌军统计
	TeamBattleStatistics EnemyTeam
}

//队伍对战数据统计
struct TeamBattleStatistics  {
	//***队伍原始数据***
	//队伍信息
	BattleTeam BattleTeam

	//***对战数据***
	//对战结果
	i64 BattleResult
	//对战统计信息
	list<GeneralBattleStatistics> GeneralBattleStatisticsList
}

//武将对战数据统计
struct GeneralBattleStatistics {
	//战法统计
	list<TacticStatistics> TacticStatisticsList
	//普攻统计
	TacticStatistics GeneralAttackStatistics
}

struct TacticStatistics {
	//战法ID
	i64 TacticId
	//战法名称
	string TacticName
	//战法品质
	i64 TacticQuality
	//发动次数
	i64 TriggerTimes
	//杀敌
	i64 KillSoliderNum
	//救援
	i64 ResumeSoliderNum
}

//============= 模拟对战 END ==============

//============= 模拟对战列表 BEGIN ==============
struct BattleListRequest{
   1: string  Uid      // 用户ID

   100: i64 PageNo,
   101: i64 PageSize
}

struct BattleListResponse{
    1: common.Meta Meta
    //对战记录列表
    2: list<BattleRecordInfo> TacticList
}

struct BattleRecordInfo {
    //对战结果
    1: enum.BattleResult BattleResult
	// 出战队伍信息
	2:BattleTeam FightingTeam
	// 对战队伍信息
	3:BattleTeam EnemyTeam
}

//============= 模拟对战列表 END ==============

//============= 查询战法列表 BEGIN ==============
struct TacticListRequest{
   1: i64  Id      // 战法ID
   2: string Name    // 战法名称
   3: enum.TacticQuality  Quality // 战法品质
   4: enum.TacticsSource  Source  // 战法来源
   5: enum.TacticsType  Type    // 战法类型
   6: list<enum.TacticsSource>  Sources    // 战法类型列表

   100: i64 PageNo,
   101: i64 PageSize
}

struct TacticListResponse{
    1: common.Meta Meta
    //战法信息列表
    2: list<Tactics> TacticList
}
//============= 查询战法列表 END ==============

//============= 查询武将列表 BEGIN ==============
struct GeneralListRequest{
     1: optional i64  Id               // 武将ID
     2: optional string Name             // 姓名
     3: optional enum.Gender   Gender           // 性别
     4: optional enum.ControlLevel  Control          // 统御
     5: optional enum.Group   Group            // 阵营
     6: optional enum.GeneralQuality   Quality          // 品质
     7: optional list<enum.GeneralTag>  Tags             // 标签
     8: optional enum.Enable  IsSupportDynamics// 是否支持动态
     9: optional enum.Enable  IsSupportCollect // 是否支持典藏
     10: list<i64> Ids //武将ID列表

     100: i64 PageNo,
     101: i64 PageSize
}

struct GeneralListResponse{
    1: common.Meta Meta
    //武将信息列表
    2: list<BattleGeneral> GeneralList
}
//============= 查询武将列表 END ==============

//============= 查询兵书列表 BEGIN ==============
struct GeneralWarBookListRequest{
    //武将ID
    1: i64 GeneralId
    //兵书类型
    2: enum.WarbookType WarbookType
}

struct GeneralWarBookListResponse{
    1: common.Meta Meta
    //兵书信息列表<map<兵书类型,map<层级,兵书list>>>
    2: map<i64,map<i64,list<WarBook>>> WarBookMapList
}

struct WarBook {
    1: i64 Id
    2: string Name
    3: i64 Type
    4: i64 Level
    5: string Desc
}
//============= 查询兵书列表 END ==============

//============= 查询特技列表 BEGIN ==============
struct SpecialTechListRequest{
    1: string Name
    2: i64 Id
    3: enum.EquipType Type
    4: enum.EquipLevel Level

    100: i64 PageNo,
    101: i64 PageSize
}

struct SpecialTechListResponse{
    1: common.Meta meta
    //特技信息列表
    2: list<SpecialTech> SpecialTechList
}

struct SpecialTech {
    1: i64 Id
    2: string Name
    3: enum.EquipType Type
    4: enum.EquipLevel Level
}
//============= 查询特技列表 END ==============

//============= 查询推荐阵容列表 BEGIN ==============
struct RecTeamListRequest{
    1: string Name
    2: enum.Group Group
    3: enum.ArmType ArmType

    100: i64 PageNo,
    101: i64 PageSize
}

struct RecTeamListResponse{
    1: common.Meta meta
    //武将信息列表
    2: list<RecTeamGeneral> RecTeamGeneralList
}

struct RecTeamGeneral {
    1: list<BattleGeneral> GeneralList
    2: string Name
    3: i64 Id
    4: i64 Group
    5: i64 ArmType
}
//============= 查询推荐阵容列表 END ==============

//============= 查询推荐战法列表 BEGIN ==============
struct RecTacticListRequest{
    1: i64 GeneralId

    100: i64 PageNo,
    101: i64 PageSize
}

struct RecTacticListResponse{
    1: common.Meta meta
    //战法信息列表
    2: list<Tactics> RecTacticList
}

//============= 查询推荐战法列表 END ==============

//============= 查询推荐兵书列表 BEGIN ==============
struct RecWarBookListRequest{
    1: i64 GeneralId
}

struct RecWarBookListResponse{
    1: common.Meta meta
    //兵书信息列表<map<兵书类型,map<层级,兵书list>>>
    2: map<i64,map<i64,list<WarBook>>> WarBookMapList
}
//============= 查询推荐兵书列表 END ==============

//============= 查询推荐特技列表 BEGIN ==============
struct RecSpecialTechListRequest{
    1: i64 GeneralId
}

struct RecSpecialTechListResponse{
    1: common.Meta meta
    //特技信息列表
    2: list<SpecialTech> SpecialTechList
}
//============= 查询推荐特技列表 END ==============

//============= 用户登录 BEGIN ==============
struct UserLoginRequest{
    1: string Code
    2: string NickName
    3: string AvatarUrl
}

struct UserLoginResponse{
    1: common.Meta meta
    2: string NickName
    3: string AvatarUrl
    4: string WxOpenId
    5: i64 Level
    6: string Uid
}

//============= 用户登录 END ==============

//============= 用户信息查询 BEGIN ==============
struct UserInfoDetailRequest{
    1: string Code
}

struct UserInfoDetailResponse{
    1: common.Meta meta
    2: UserInfo UserInfo //用户信息
    3: BattleStatisticsInfo BattleStatisticsInfo //模拟对战信息
    4: LotteryStatisticsInfo LotteryStatisticsInfo //模拟抽卡信息
}

struct UserInfo {
    1: string Uid
    2: string NickName
    3: string AvatarUrl
    4: string WxOpenId
    5: i64 Level
}

struct BattleStatisticsInfo{
    1: list<GeneralRecord> HighFreqGeneralList //高频使用武将列表
    2: list<TacticsRecord> HighFreqTacticsList //高频使用战法列表
    3: list<TeamRecord> HighFreqTeamList //高频使用队伍列表

    50:double WinRate //胜率
}

struct GeneralRecord{
    1: MetadataGeneral General //武将
    2: i64 Times //次数
}

struct TacticsRecord{
    1: Tactics Tactics //战法
    2: i64 Times //场次
}

struct TeamRecord{
    1: BattleTeam BattleTeam //阵容
    2: i64 Times //场次
}

struct LotteryStatisticsInfo{
    1: list<GeneralRecord> HighFreqGeneralList //高频抽中武将列表
    2: list<string> HighFreqCardPoolList //高频抽取卡池列表

    50:double Lev5HitRate //五星率
}

//============= 用户信息查询 END ==============

//============= 武将抽卡 BEGIN ==============
struct GeneralLotteryDoRequest{
    //抽取卡池枚举
    1: i64 GeneralLotteryPool
    //抽取次数
    2: i64 RollTimes
    //用户uid
    3: string Uid
}

struct GeneralLotteryDoResponse{
    1: common.Meta meta
    //抽取的武将信息
    2:list<GeneralLotteryDoInfo> GeneralLotteryInfoList
    //保底统计
    3:i64 ProtectedMustHitNum
    //五星武将出现数
    4:i64 Hit5LevGeneralNum
    //五星武将出现率
    5:double Hit5LevGeneralRate
    //连续不出橙次数累计
    6:i64 NotHitLev5Times
}

struct GeneralLotteryDoInfo {
	//武将信息
	1:MetadataGeneral GeneralInfo
	//抽中次数
	2:i64 HitNum
	//本次抽中占比
	3:double HitRate
	//游戏设置概率
	4:double LotteryRate
}

//============= 武将抽卡 END ==============

//============= 武将抽卡记录 BEGIN ==============
struct GeneralLotteryListRequest{
    //用户uid
    1: string Uid
}

struct GeneralLotteryListResponse{
    1: common.Meta meta
    //高频抽取的武将信息
    2:list<GeneralLotteryDoInfo> HighFreqGeneralLotteryInfoList
}

//============= 武将抽卡记录 END ==============

//============= 武将卡池查询 BEGIN ==============
struct GeneralLotteryInfoListRequest{
    //卡池枚举
    1: i64 GeneralLotteryPool
}

struct GeneralLotteryInfoListResponse{
    1: common.Meta meta
    //卡池武将信息
    2:list<GeneralLotterInfoQueryInfo> GeneralLotteryPoolInfoList
}

struct GeneralLotterInfoQueryInfo {
	//武将信息
	1:list<MetadataGeneral> GeneralInfoList
    //卡池枚举
    2: i64 GeneralLotteryPool
    //卡池名称
    3: string GeneralLotteryPoolName
}

//============= 武将卡池查询 END ==============

//============= 武将概率查询 BEGIN ==============
struct GeneralLotteryRateListRequest{
    1: i64 GeneralId
}

struct GeneralLotteryRateListResponse{
    1: common.Meta meta
    //武将概率信息
    2:list<GeneralLotteryRateListInfo> GeneralLotteryRateInfoList
}

struct GeneralLotteryRateListInfo {
	//武将概率
	1:double LotteryRate,
    //卡池枚举
    2: i64 GeneralLotteryPool
    //卡池名称
    3: string GeneralLotteryPoolName
}

//============= 武将概率查询 END ==============

//============= 卡池查询 BEGIN ==============
struct GeneralLotteryPoolListRequest{
}

struct GeneralLotteryPoolListResponse{
    1: common.Meta meta
    //卡池武将信息
    2:list<GeneralLotterPoolQueryInfo> GeneralLotteryPoolInfoList
}

struct GeneralLotterPoolQueryInfo {
    //卡池枚举
    1: i64 GeneralLotteryPool
    //卡池名称
    2: string GeneralLotteryPoolName
}

//============= 武将抽卡 END ==============

//============= 用户卡池抽奖次数重置 BEGIN ==============
struct GeneralLotteryUserDataResetRequest{
    //用户id
    1: string Uid
    //卡池枚举
    2: i64 GeneralLotteryPool
}

struct GeneralLotteryUserDataResetResponse{
    1: common.Meta meta
}

//============= 用户卡池抽奖次数重置 END ==============

//============= 用户卡池抽奖信息查询 BEGIN ==============
struct GeneralLotteryUserDataQueryRequest{
    //用户id
    1: string Uid
    //卡池枚举
    2: i64 GeneralLotteryPool
}

struct GeneralLotteryUserDataQueryResponse{
    1: common.Meta meta
    2: GeneralLotteryUserDataQueryInfo GeneralLotteryDataQueryInfo
}

struct GeneralLotteryUserDataQueryInfo{
    //用户id
    1: string Uid
    //卡池枚举
    2: i64 GeneralLotteryPool
    //连续未中五星次数
    3: i64 NotHitLev5Times
}
//============= 用户卡池抽奖信息查询 END ==============

service ApiService {
    //**模拟对战**
    //模拟对战
    BattleDoResponse BattleDo(1: BattleDoRequest request) (api.post="/v1/battle/do");
    //模拟对战列表
    BattleListResponse BattleList(1: BattleListRequest request) (api.get="/v1/battle/list");

    //**信息查询**
    //查询战法列表
    TacticListResponse TacticList(1:TacticListRequest request)(api.get="/v1/tactic/list");
    //查询武将列表
    GeneralListResponse GeneralList(1:GeneralListRequest request)(api.get="/v1/general/list");
    //查询兵书列表
    GeneralWarBookListResponse GeneralWarBookList(1:GeneralWarBookListRequest request)(api.get="/v1/general_warbook/list");
    //查询特技列表
    SpecialTechListResponse SpecialTechList(1:SpecialTechListRequest request)(api.get="/v1/special_tech/list");

    //**推荐**
    //推荐阵容列表
    RecTeamListResponse RecTeamList(1:RecTeamListRequest request)(api.get="/v1/rec_team/list");
    //武将推荐战法列表
    RecTacticListResponse RecTacticList(1:RecTacticListRequest request)(api.get="/v1/rec_tactic/list");
    //武将推荐兵书列表
    RecWarBookListResponse RecWarBookList(1:RecWarBookListRequest request)(api.get="/v1/rec_warbook/list");
    //武将推荐特技列表
    RecSpecialTechListResponse RecSpecialTechList(1:RecSpecialTechListRequest request)(api.get="/v1/rec_special_tech/list");

    //**抽卡**
     //卡池查询
    GeneralLotteryPoolListResponse GeneralLotteryPoolList(1:GeneralLotteryPoolListRequest request)(api.get="/v1/lottery/general/pool_list");
    //卡池武将查询
    GeneralLotteryInfoListResponse GeneralLotteryInfoList(1:GeneralLotteryInfoListRequest request)(api.get="/v1/lottery/general/info_list");
    //武将概率查询
    GeneralLotteryRateListResponse GeneralLotteryRateList(1:GeneralLotteryRateListRequest request)(api.get="/v1/lottery/general/rate_list");
    //武将抽卡
    GeneralLotteryDoResponse GeneralLotteryDo(1:GeneralLotteryDoRequest request)(api.post="/v1/lottery/general/do");
    //武将抽卡列表
    GeneralLotteryListResponse GeneralLotteryList(1:GeneralLotteryListRequest request)(api.get="/v1/lottery/general/list");

    //用户武将抽卡数据重置
    GeneralLotteryUserDataResetResponse GeneralLotteryUserDataReset(1:GeneralLotteryUserDataResetRequest request)(api.post="/v1/lottery/general/user_data_reset");
    //用户武将抽卡数据查询
    GeneralLotteryUserDataQueryResponse GeneralLotteryUserDataQuery(1:GeneralLotteryUserDataQueryRequest request)(api.get="/v1/lottery/general/user_data_query");

    //**微信**
    //用户登录接口
    UserLoginResponse UserLogin(1:UserLoginRequest request)(api.post="/v1/user/login");
    //用户信息获取接口
    UserInfoDetailResponse UserInfoDetail(1:UserInfoDetailRequest request)(api.get="/v1/user/detail");

    //**jym**
    //账户商品搜索
     jym.AccountSearchResponse AccountSearch(1:jym.AccountSearchRequest request)(api.get="/jym/account/search");
}