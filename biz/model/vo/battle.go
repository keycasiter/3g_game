package vo

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
)

//对战队伍信息
type BattleTeam struct {
	ArmType        consts.ArmType
	BattleGenerals []*BattleGeneral
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
