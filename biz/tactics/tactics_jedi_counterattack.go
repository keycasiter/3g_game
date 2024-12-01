package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 绝地反击
// 战斗中，自己每次受到兵刃伤害后，武力提升6点，最大叠加10次；第5回合时，根据叠加次数对敌军全体造成兵刃伤害（伤害率120%，每次提高14%伤害率）
type JediCounterattackTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JediCounterattackTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JediCounterattackTactic) Prepare() {
	ctx := j.tacticsParams.Ctx
	currentGeneral := j.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		j.Name(),
	)
	//战斗中，自己每次受到兵刃伤害后，武力提升6点，最大叠加10次；第5回合时，根据叠加次数对敌军全体造成兵刃伤害（伤害率120%，每次提高14%伤害率）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}

		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    6,
			EffectTimes:    1,
			MaxEffectTimes: 10,
			FromTactic:     j.Id(),
			ProduceGeneral: currentGeneral,
		})
		if triggerRound == consts.Battle_Round_Fifth {
			times := int64(0)
			if effectParams, ok := util.BuffEffectOfTacticGet(triggerGeneral, consts.BuffEffectType_IncrForce, j.Id()); ok {
				for _, effectParam := range effectParams {
					times += effectParam.EffectTimes
				}
			}
			dmgRate := 1.2 + (0.14 * cast.ToFloat64(times))
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, j.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     j.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: dmgRate,
					TacticId:          j.Id(),
					TacticName:        j.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (j JediCounterattackTactic) Id() consts.TacticId {
	return consts.JediCounterattack
}

func (j JediCounterattackTactic) Name() string {
	return "绝地反击"
}

func (j JediCounterattackTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (j JediCounterattackTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JediCounterattackTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JediCounterattackTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (j JediCounterattackTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Spearman,
	}
}

func (j JediCounterattackTactic) Execute() {

}

func (j JediCounterattackTactic) IsTriggerPrepare() bool {
	return false
}

func (a JediCounterattackTactic) SetTriggerPrepare(triggerPrepare bool) {
}
