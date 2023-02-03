package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

//req
type BattleLogicContextRequest struct {
	/** 对战阶段 **/
	BattlePhase consts.BattlePhase
	/** 回合信息 **/
	BattleRound consts.BattleRound
	/** 武将信息 **/
	// 出战武将信息
	FightingGenerals []*vo.BattleGeneral
	// 对战武将信息
	EnemyGenerals []*vo.BattleGeneral
}

//resp
type BattleLogicContextResponse struct {
}

//resp

// 对战上下文环境
type BattleLogicContext struct {
	//上下文
	Ctx context.Context
	// 入参
	ReqParam *BattleLogicContextRequest
	// 执行方法
	Funcs []func()
	// 执行错误
	RunErr error
}

func NewBattleLogicContext(ctx context.Context, req *BattleLogicContextRequest) *BattleLogicContext {
	runCtx := &BattleLogicContext{
		ReqParam: req,
	}
	//注入方法执行顺序
	runCtx.Funcs = []func(){
		//初始化元数据
		runCtx.initMetadata,
		//对战准备阶段处理
		runCtx.processBattlePreparePhase,
		//对战对阵阶段处理
		runCtx.processBattleFightingPhase,
	}

	return runCtx
}

func (runCtx *BattleLogicContext) Run() error {
	for _, f := range runCtx.Funcs {
		f()
		if runCtx.RunErr != nil {
			hlog.CtxErrorf(runCtx.Ctx, "BattleLogicContext Func=%v Run err:%v", f, runCtx.RunErr)
			return runCtx.RunErr
		}
	}
	return nil
}

func (runCtx *BattleLogicContext) initMetadata() {

}

//属性加点处理
func (runCtx *BattleLogicContext) handleAbilityAttrAddition(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase + general.Addition.AbilityAttr.ForceBase
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase + general.Addition.AbilityAttr.IntelligenceBase
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase + general.Addition.AbilityAttr.CommandBase
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase + general.Addition.AbilityAttr.SpeedBase
}

//武将等级处理
func (runCtx *BattleLogicContext) handleGeneralLevel(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase +
			general.BaseInfo.AbilityAttr.ForceRate*float64(general.Addition.GeneralLevel)
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase +
			general.BaseInfo.AbilityAttr.IntelligenceRate*float64(general.Addition.GeneralLevel)
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase +
			general.BaseInfo.AbilityAttr.CommandRate*float64(general.Addition.GeneralLevel)
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase +
			general.BaseInfo.AbilityAttr.SpeedRate*float64(general.Addition.GeneralLevel)
}

//兵种适性处理
func (runCtx *BattleLogicContext) handleArmAbility(general *vo.BattleGeneral) {
	switch general.ArmType {
	//骑兵
	case consts.ArmType_Cavalry:
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Cavalry, general.BaseInfo.AbilityAttr)
	//盾兵
	case consts.ArmType_Mauler:
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Mauler, general.BaseInfo.AbilityAttr)
	//弓兵
	case consts.ArmType_Archers:
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Archers, general.BaseInfo.AbilityAttr)
	//枪兵
	case consts.ArmType_Spearman:
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Spearman, general.BaseInfo.AbilityAttr)
	//器械
	case consts.ArmType_Apparatus:
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Apparatus, general.BaseInfo.AbilityAttr)
	}
}

func (runCtx *BattleLogicContext) processBattlePreparePhase() {
	//出战武将
	for _, general := range runCtx.ReqParam.FightingGenerals {
		//士气加成
		//武将加点加成
		runCtx.handleAbilityAttrAddition(general)
		//武将等级加成
		runCtx.handleGeneralLevel(general)
		//红度加成
		//缘分加成
		//兵种适性加成
		runCtx.handleArmAbility(general)
		//装备加成
		//特技加成
	}

	//对战武将
	for _, general := range runCtx.ReqParam.EnemyGenerals {
		//士气加成
		//武将加点加成
		runCtx.handleAbilityAttrAddition(general)
		//武将等级加成
		runCtx.handleGeneralLevel(general)
		//红度加成
		//缘分加成
		//兵种适性加成
		runCtx.handleArmAbility(general)
		//装备加成
		//特技加成
	}

	hlog.CtxInfof(runCtx.Ctx, "fighting general => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.FightingGenerals))
	hlog.CtxInfof(runCtx.Ctx, "enemy general => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.EnemyGenerals))
}

func (runCtx *BattleLogicContext) processBattleFightingPhase() {

}
