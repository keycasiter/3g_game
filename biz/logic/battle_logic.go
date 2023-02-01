package logic

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

//req
type BattleLogicContextRequest struct {
}

//resp

// 对战上下文环境
type BattleLogicContext struct {
	/** 对战阶段 **/
	BattlePhase consts.BattlePhase
	/** 回合信息 **/
	BattleRound consts.BattleRound
	/** 武将信息 **/
	// 出战武将信息
	FightingGenerals []*vo.BattleGeneral
	// 对战武将信息
	EnemyGenerals []*vo.BattleGeneral

	// 执行方法
	Funcs []func()
}

func NewBattleLogicContext(battlePhase consts.BattlePhase,
	battleRound consts.BattleRound,
	fightingGenerals []*vo.BattleGeneral,
	enemyGenerals []*vo.BattleGeneral) *BattleLogicContext {
	runCtx := &BattleLogicContext{}
	//注入方法执行顺序
	runCtx.Funcs = []func(){
		//初始化元数据
		runCtx.initMetadata,
		//对战准备阶段处理
		runCtx.processBattlePreparePhase,
		//对战对阵阶段处理
		runCtx.processBattleFightingPhase,
	}

	return
}

func (runCtx *BattleLogicContext) Run() {

}

func (runCtx *BattleLogicContext) initMetadata() {
	//出战武将
	for i, general := range runCtx.FightingGenerals {
		//1. 武将加成
		//a. 加点加成
		runCtx.handleAbilityAttrAddition(general)
		//b. 等级加成
		//c. 红度加成
		//d. 缘分加成
		//e. 兵种适性加成
		//2. 装备加成
		//3. 特技加成
	}
	//对战武将
}

func (runCtx *BattleLogicContext) handleAbilityAttrAddition(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase + general.Addition.AbilityAttrAddition.ForceBase
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase + general.Addition.AbilityAttrAddition.IntelligenceBase
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase + general.Addition.AbilityAttrAddition.CommandBase
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase + general.Addition.AbilityAttrAddition.SpeedBase
}

func (runCtx *BattleLogicContext) handleGeneralLevel(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase + int(general.BaseInfo.AbilityAttr.ForceRate)*general.Addition.GeneralLevel
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase + general.Addition.AbilityAttrAddition.IntelligenceBase
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase + general.Addition.AbilityAttrAddition.CommandBase
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase + general.Addition.AbilityAttrAddition.SpeedBase
}

func (runCtx *BattleLogicContext) processBattlePreparePhase() {

}

func (runCtx *BattleLogicContext) processBattleFightingPhase() {

}
