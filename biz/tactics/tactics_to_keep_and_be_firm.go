package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//守而必固
//战斗开始时，嘲讽敌军主将（强迫目标普通攻击自己），并使自己的统率提升40点，持续4回合
type ToKeepAndBeFirmTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToKeepAndBeFirmTactic) IsTriggerPrepare() bool {
	return false
}

func (t ToKeepAndBeFirmTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ToKeepAndBeFirmTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//找到敌军主将嘲讽
	enemyMaster := util.GetEnemyMasterGeneral(t.tacticsParams)
	enemyMaster.TauntByGeneral = currentGeneral

	if util.DebuffEffectWrapSet(ctx, enemyMaster, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
		EffectRate: 1.0,
		FromTactic: t.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(enemyMaster, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound
			triggerGeneral := params.CurrentGeneral
			//第五回合消失
			if triggerRound == consts.Battle_Round_Fifth {
				util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_Taunt, t.Id())

				//统率效果消失
				currentGeneral.BaseInfo.AbilityAttr.CommandBase -= 40
				hlog.CtxInfof(ctx, "[%s]的统率降低了40",
					currentGeneral.BaseInfo.Name,
				)
			}

			return triggerResp
		})
	}

	//统率提高
	currentGeneral.BaseInfo.AbilityAttr.CommandBase += 40
	hlog.CtxInfof(ctx, "[%s]的统率提高了40",
		currentGeneral.BaseInfo.Name,
	)
}

func (t ToKeepAndBeFirmTactic) Id() consts.TacticId {
	return consts.ToKeepAndBeFirm
}

func (t ToKeepAndBeFirmTactic) Name() string {
	return "守而必固"
}

func (t ToKeepAndBeFirmTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ToKeepAndBeFirmTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToKeepAndBeFirmTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToKeepAndBeFirmTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t ToKeepAndBeFirmTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToKeepAndBeFirmTactic) Execute() {
	return
}
