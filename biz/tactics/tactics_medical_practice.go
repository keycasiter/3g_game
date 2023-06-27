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

// 青囊
// 战斗前4回合，使我军群体（2人）获得40统率（受智力影响）及急救效果，每次受到伤害时有50%几率回复一定兵力（治疗率88%，受智力影响）
// 指挥 100%
type MedicalPracticeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MedicalPracticeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 1.0
	return m
}

func (m MedicalPracticeTactic) Prepare() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	// 战斗前4回合，使我军群体（2人）获得40统率（受智力影响）及急救效果，每次受到伤害时有50%几率回复一定兵力（治疗率88%，受智力影响）
	pairGenerals := util.GetPairGeneralsTwoArrByGeneral(currentGeneral, m.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//提高统率
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectValue:    20,
			EffectRound:    4,
			FromTactic:     m.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrCommand,
					TacticId:   m.Id(),
				})

				return revokeResp
			})
		}
		//急救效果
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_EmergencyTreatment, &vo.EffectHolderParams{
			EffectRound:    4,
			FromTactic:     m.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_EmergencyTreatment,
					TacticId:   m.Id(),
				}) {
					if util.GenerateRate(0.5) {
						resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.88)
						util.ResumeSoldierNum(&util.ResumeParams{
							Ctx:            ctx,
							TacticsParams:  m.tacticsParams,
							ProduceGeneral: currentGeneral,
							SufferGeneral:  revokeGeneral,
							ResumeNum:      resumeNum,
						})
					}
				}

				return revokeResp
			})
		}
	}
}

func (m MedicalPracticeTactic) Id() consts.TacticId {
	return consts.MedicalPractice
}

func (m MedicalPracticeTactic) Name() string {
	return "青囊"
}

func (m MedicalPracticeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (m MedicalPracticeTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MedicalPracticeTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MedicalPracticeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (m MedicalPracticeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MedicalPracticeTactic) Execute() {

}

func (m MedicalPracticeTactic) IsTriggerPrepare() bool {
	return false
}
