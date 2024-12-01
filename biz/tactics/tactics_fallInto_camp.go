package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 陷阵营
// 将盾兵进阶为无往不利的陷阵营：
// 我军全体武力/统率提高22点，战斗前3回合获得急救状态，受到伤害时有30%概率获得治疗（治疗率60%，受智力影响）
// 若高顺统领，则治疗率将额外受统率影响
type FallIntoCampTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FallIntoCampTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FallIntoCampTactic) Prepare() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	//我军全体武力/统率提高22点
	//找到我军全体
	pairGenerals := util.GetPairGeneralArr(currentGeneral, f.tacticsParams)
	for _, general := range pairGenerals {
		//武力
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    22,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		})
		//统率
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectValue:    22,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		})
		//战斗前3回合获得急救状态，受到伤害时有30%概率获得治疗（治疗率60%，受智力影响）
		//若高顺统领，则治疗率将额外受统率影响
		effectRate := 0.6
		if consts.GaoShun == consts.General_Id(currentGeneral.BaseInfo.Id) {
			effectRate += currentGeneral.BaseInfo.AbilityAttr.CommandBase / 100 / 100
		}
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_EmergencyTreatment, &vo.EffectHolderParams{
			TriggerRate:    0.3,
			EffectRate:     effectRate,
			EffectRound:    3,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_EmergencyTreatment,
					TacticId:   f.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (f FallIntoCampTactic) Id() consts.TacticId {
	return consts.FallIntoCamp
}

func (f FallIntoCampTactic) Name() string {
	return "陷阵营"
}

func (f FallIntoCampTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FallIntoCampTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FallIntoCampTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FallIntoCampTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (f FallIntoCampTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Mauler,
	}
}

func (f FallIntoCampTactic) Execute() {
}

func (f FallIntoCampTactic) IsTriggerPrepare() bool {
	return false
}

func (a FallIntoCampTactic) SetTriggerPrepare(triggerPrepare bool) {
}
