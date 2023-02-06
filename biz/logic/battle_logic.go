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
	/** 队伍信息 **/
	// 出战队伍信息
	FightingTeam *vo.BattleTeam
	// 对战队伍信息
	EnemyTeam *vo.BattleTeam
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
func (runCtx *BattleLogicContext) handleArmAbility(teamArmType consts.ArmType, general *vo.BattleGeneral) {
	armType := consts.ArmType_Unknow
	switch teamArmType {
	//骑兵
	case consts.ArmType_Cavalry:
		armType = general.BaseInfo.ArmsAttr.Cavalry
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Cavalry, general.BaseInfo.AbilityAttr)
	//盾兵
	case consts.ArmType_Mauler:
		armType = general.BaseInfo.ArmsAttr.Mauler
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Mauler, general.BaseInfo.AbilityAttr)
	//弓兵
	case consts.ArmType_Archers:
		armType = general.BaseInfo.ArmsAttr.Archers
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Archers, general.BaseInfo.AbilityAttr)
	//枪兵
	case consts.ArmType_Spearman:
		armType = general.BaseInfo.ArmsAttr.Spearman
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Spearman, general.BaseInfo.AbilityAttr)
	//器械
	case consts.ArmType_Apparatus:
		armType = general.BaseInfo.ArmsAttr.Apparatus
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Apparatus, general.BaseInfo.AbilityAttr)
	}

	hlog.CtxInfof(runCtx.Ctx, "[%s]的【 %s 】兵种适性为【 %s 】，属性调整为原来的%s",
		general.BaseInfo.Name,
		util.TranslateArmType(teamArmType),
		util.TranslateArmsAbility(consts.ArmsAbility(armType)),
		util.TranslateArmsAbilityAddition(consts.ArmsAbility(armType)),
	)
}

func (runCtx *BattleLogicContext) processBattlePreparePhase() {
	//出战武将加成处理
	for _, general := range runCtx.ReqParam.FightingTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.FightingTeam.ArmType, general)
	}

	//对战武将加成处理
	for _, general := range runCtx.ReqParam.EnemyTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.EnemyTeam.ArmType, general)
	}

	hlog.CtxInfof(runCtx.Ctx, "fighting team => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.FightingTeam))
	hlog.CtxInfof(runCtx.Ctx, "enemy team => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.EnemyTeam))
}

func (runCtx *BattleLogicContext) handleGeneralAddition(armType consts.ArmType, general *vo.BattleGeneral) {
	//国土效果
	//TODO 部队兵力不足可携带兵力上限的65%，国土效果不生效
	//士气加成
	//TODO 士气满不影响任何东西，不满100，则降低伤害，其余不影响；
	//看了下战报分析了下，每减少0.1士气，降低伤害比例是0.07%，如果0士气则降低70%伤害，其余不影响
	//兵种适性加成
	runCtx.handleArmAbility(armType, general)
	//武将加点加成
	runCtx.handleAbilityAttrAddition(general)
	//武将等级加成
	runCtx.handleGeneralLevel(general)
	//红度加成
	//缘分加成
	//装备加成
	//特技加成
}

func (runCtx *BattleLogicContext) processBattleFightingPhase() {

}
