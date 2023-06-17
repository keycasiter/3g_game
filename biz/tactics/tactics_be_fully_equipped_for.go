package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 整装待发
// 战斗中，偶数回合，恢复我军群体（2人）兵力（治疗率88%，受智力影响）
type BeFullyEquippedForTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BeFullyEquippedForTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BeFullyEquippedForTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_BeFullyEquippedFor_Prepare, &vo.EffectHolderParams{
		FromTactic:     b.Id(),
		ProduceGeneral: currentGeneral,
	})

	//战斗中，偶数回合，恢复我军群体（2人）兵力（治疗率88%，受智力影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound%2 == 0 {
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				b.Name(),
				consts.BuffEffectType_BeFullyEquippedFor_Prepare,
			)
			pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, b.tacticsParams)
			for _, general := range pairGenerals {
				resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.88)
				util.ResumeSoldierNum(ctx, general, resumeNum)
			}
		}

		return triggerResp
	})
}

func (b BeFullyEquippedForTactic) Id() consts.TacticId {
	return consts.BeFullyEquippedFor
}

func (b BeFullyEquippedForTactic) Name() string {
	return "整装待发"
}

func (b BeFullyEquippedForTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BeFullyEquippedForTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeFullyEquippedForTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeFullyEquippedForTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (b BeFullyEquippedForTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeFullyEquippedForTactic) Execute() {

}

func (b BeFullyEquippedForTactic) IsTriggerPrepare() bool {
	return false
}
