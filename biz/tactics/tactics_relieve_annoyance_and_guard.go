package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 解烦卫
// 将枪兵进阶为战无不胜的解烦卫，战斗中，我军全体速度提升36点，
// 且我军速度最快的武将普通攻击后，有30%概率对敌军单体造成一次伤害（伤害率36%，受武力和智力较高一项影响，额外附带该武将速度属性40%伤害），
// 若未生效则我军单体回复一定兵力（治疗率72%，受武力和智力较高一项影响），若韩当统领，则基础伤害率提升至72%
type RelieveAnnoyanceAndGuardTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f RelieveAnnoyanceAndGuardTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f RelieveAnnoyanceAndGuardTactic) Prepare() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	// 将枪兵进阶为战无不胜的解烦卫，战斗中，我军全体速度提升36点，
	// 且我军速度最快的武将普通攻击后，有30%概率对敌军单体造成一次伤害（伤害率36%，受武力和智力较高一项影响，额外附带该武将速度属性40%伤害），
	// 若未生效则我军单体回复一定兵力（治疗率72%，受武力和智力较高一项影响），若韩当统领，则基础伤害率提升至72%

	// 将枪兵进阶为战无不胜的解烦卫，战斗中，我军全体速度提升36点，
	pairGenerals := util.GetPairGeneralArr(currentGeneral, f.tacticsParams)
	for _, general := range pairGenerals {
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
			EffectValue: 36,
			FromTactic:  f.Id(),
		})
	}
	// 且我军速度最快的武将普通攻击后，有30%概率对敌军单体造成一次伤害（伤害率36%，受武力和智力较高一项影响，额外附带该武将速度属性40%伤害），
	highestSpeedGeneral := util.GetPairGeneralWhoIsHighestSpeed(currentGeneral, f.tacticsParams)
	util.TacticsTriggerWrapRegister(highestSpeedGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		triggerRate := 0.3
		if util.GenerateRate(triggerRate) {
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, f.tacticsParams)

			_, val := util.GetGeneralHighestBetweenForceOrIntelligence(triggerGeneral)
			effectRate := val / 100 / 100
			dmgRate := 0.36 + effectRate

			appendDamageValue := triggerGeneral.BaseInfo.AbilityAttr.SpeedBase * 0.4

			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     f.tacticsParams,
				AttackGeneral:     triggerGeneral,
				SufferGeneral:     enemyGeneral,
				DamageType:        consts.DamageType_Weapon,
				DamageImproveRate: dmgRate,
				TacticId:          f.Id(),
				TacticName:        f.Name(),
				AppendDamageValue: appendDamageValue,
			})
		}

		return triggerResp
	})
}

func (f RelieveAnnoyanceAndGuardTactic) Id() consts.TacticId {
	return consts.RelieveAnnoyanceAndGuard
}

func (f RelieveAnnoyanceAndGuardTactic) Name() string {
	return "解烦卫"
}

func (f RelieveAnnoyanceAndGuardTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (f RelieveAnnoyanceAndGuardTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f RelieveAnnoyanceAndGuardTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f RelieveAnnoyanceAndGuardTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (f RelieveAnnoyanceAndGuardTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (f RelieveAnnoyanceAndGuardTactic) Execute() {
}

func (f RelieveAnnoyanceAndGuardTactic) IsTriggerPrepare() bool {
	return false
}
