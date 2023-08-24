// idl/api.thrift
namespace go api

include "idl/common.thrift"
include "idl/enum.thrift"

//============= 模拟对战 BEGIN ==============

struct BattleExecuteRequest {
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
	1: double ForceBase
	2: double ForceRate
	3: double IntelligenceBase
	4: double IntelligenceRate
	5: double CharmBase
	6: double CharmRate
	7: double CommandBase
	8: double CommandRate
	9: double PoliticsBase
	10: double PoliticsRate
	11: double SpeedBase
	12: double SpeedRate
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
	1: double ForceAddition
	2: double IntelligenceAddition
	3: double CommandAddition
	4: double SpeedAddition
}

// 建筑科技阵营加成
struct BuildingTechGroupAddition {
	1: double GroupWeiGuoRate
	2: double GroupShuGuoRate
	3: double GroupWuGuoRate
	4: double GroupQunXiongRate
}

struct BattleExecuteResponse {
  1: common.Meta meta
    //我军武将对战信息
  2: list<BattleGeneralStatistics> FightingGeneralStatistics
    //敌军武将对战信息
  3: list<BattleGeneralStatistics> EnemyGeneralStatistics
}

// 对战武将统计
struct BattleGeneralStatistics {
	//武将基础信息
	1:BattleGeneral BattleGeneral
	//战法统计
	2:list<BattleGeneralTacticStatistics> TacticStatistics
	//普攻统计
	3:BattleGeneralTacticStatistics AttackStatistics
}

// 对战战法统计
struct BattleGeneralTacticStatistics {
	//释放次数
	1: i64 TriggerCnt
	//伤害量
	2: i64 DamageNum
	//恢复量
	3: i64 ResumeNum
}

//============= 模拟对战 END ==============

//============= 查询战法列表 BEGIN ==============
struct TacticQueryRequest{
   1: i64  Id      // 主键ID
   2: string Name    // 战法名称
   3: enum.TacticQuality  Quality // 战法品质
   4: enum.TacticsSource  Source  // 战法来源
   5: enum.TacticsType  Type    // 战法类型
   6: list<enum.TacticsSource>  Sources    // 战法类型列表

   100: i64 PageNo,
   101: i64 PageSize
}

struct TacticQueryResponse{
    1: common.Meta Meta
    //战法信息列表
    2: list<Tactics> TacticList
}
//============= 查询战法列表 END ==============

//============= 查询武将列表 BEGIN ==============
struct GeneralQueryRequest{
     1: optional i64  Id               // 武将ID
     2: optional string Name             // 姓名
     3: optional enum.Gender   Gender           // 性别
     4: optional enum.ControlLevel  Control          // 统御
     5: optional enum.Group   Group            // 阵营
     6: optional enum.GeneralQuality   Quality          // 品质
     7: optional list<enum.GeneralTag>  Tags             // 标签
     8: optional enum.Enable  IsSupportDynamics// 是否支持动态
     9: optional enum.Enable  IsSupportCollect // 是否支持典藏

     100: i64 PageNo,
     101: i64 PageSize
}

struct GeneralQueryResponse{
    1: common.Meta Meta
    //武将信息列表
    2: list<BattleGeneral> GeneralList
}
//============= 查询武将列表 END ==============

//============= 查询兵书列表 BEGIN ==============
struct GeneralWarBookQueryRequest{
    //武将ID
    1: i64 GeneralId
    //兵书类型
    2: enum.WarbookType WarbookType
}

struct GeneralWarBookQueryResponse{
    1: common.Meta Meta
    //兵书信息列表<map<兵书类型,map<层级,兵书list>>>
    2: map<i64,map<i64,list<WarBook>>> WarBookMapList
}

struct WarBook {
    1: i64 Id
    2: string Name
    3: i64 Type
    4: i64 Level
}
//============= 查询兵书列表 END ==============

//============= 查询特技列表 BEGIN ==============
struct SpecialTechQueryRequest{
    1: string Name
    2: i64 Id
    3: enum.EquipType Type

    100: i64 PageNo,
    101: i64 PageSize
}

struct SpecialTechQueryResponse{
    1: common.Meta meta
    //特技信息列表
    2: list<SpecialTech> SpecialTechList
}

struct SpecialTech {
    1: i64 Id
    2: string Name
    3: enum.EquipType Type
}
//============= 查询特技列表 END ==============

//============= 查询推荐阵容列表 BEGIN ==============
struct RecTeamQueryRequest{
    1: string Name

    100: i64 PageNo,
    101: i64 PageSize
}

struct RecTeamQueryResponse{
    1: common.Meta meta
    //武将信息列表
    2: list<RecTeamGeneral> RecTeamGeneralList
}

struct RecTeamGeneral {
    1: list<i64> GeneralIds
    2: list<i64> TacticIds
    3: list<i64> WarbookIds
    4: string Name
    5: i64 Id
}
//============= 查询推荐阵容列表 END ==============

service ApiService {
    //模拟对战
    BattleExecuteResponse BattleExecute(1: BattleExecuteRequest request) (api.post="/v1/battle/execute");
    //查询战法列表
    TacticQueryResponse TacticQuery(1:TacticQueryRequest request)(api.get="/v1/tactic/query");
    //查询武将列表
    GeneralQueryResponse GeneralQuery(1:GeneralQueryRequest request)(api.get="/v1/general/query");
    //查询兵书列表
    GeneralWarBookQueryResponse GeneralWarBookQuery(1:GeneralWarBookQueryRequest request)(api.get="/v1/general_warbook/query");
    //查询特技列表
    SpecialTechQueryResponse SpecialTechQuery(1:SpecialTechQueryRequest request)(api.get="/v1/special_tech/query");
    //推荐阵容列表
    RecTeamQueryResponse RecTeamQuery(1:RecTeamQueryRequest request)(api.get="/v1/rec_team/query");
}