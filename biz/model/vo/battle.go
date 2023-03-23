package vo

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
)

// 对战队伍信息
type BattleTeam struct {
	/** 队伍基础信息 **/
	//队伍类型
	TeamType consts.TeamType
	//队伍兵种
	ArmType consts.ArmType
	//队伍武将信息
	BattleGenerals []*BattleGeneral
	//兵战科技-属性加成
	BuildingTechAttrAddition BuildingTechAttrAddition
	//协力科技-阵营加成
	BuildingTechGroupAddition BuildingTechGroupAddition

	/** 队伍对战中信息 **/

}

// 对战武将信息
type BattleGeneral struct {
	//基础信息
	BaseInfo *po.MetadataGeneral
	//佩戴战法
	EquipTactics []*po.Tactics
	//武将对战加成
	Addition *BattleGeneralAddition
}

// 武将对战加成
type BattleGeneralAddition struct {
	//1. 武将加成
	//1.a. 加点加成
	AbilityAttr po.AbilityAttr
	//1.b. 等级加成
	GeneralLevel consts.GeneralLevel
	//1.c. 红度加成
	GeneralStarLevel consts.GeneralStarLevel
	//1.d. 缘分加成
	Predestination consts.Predestination
	//2. 装备加成
	//TODO
	//3. 特技加成
	//TODO
}

// 建筑科技属性加成
type BuildingTechAttrAddition struct {
	ForceAddition        float64
	IntelligenceAddition float64
	CommandAddition      float64
	SpeedAddition        float64
}

// 建筑科技阵营加成
type BuildingTechGroupAddition struct {
	GroupWeiGuoRate   float64
	GroupShuGuoRate   float64
	GroupWuGuoRate    float64
	GroupQunXiongRate float64
}

// 对战队伍效果
type BattlingTeamEffect struct {
}

// 对战武将效果
type BattlingGeneralEffect struct {
}
