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

// 陷阵突袭
// 战斗中，使自己的普通攻击目标有68%几率锁定为敌军主将，自身突击战法的发动几率提高15%并使自己成功发动突击战法后，对目标额外发动一次兵刃攻击（95%）
// 自身为主将时，获得6%倒戈
// 被动，100%
type RaidInFormationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RaidInFormationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RaidInFormationTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	// 战斗中，使自己的普通攻击目标有68%几率锁定为敌军主将，自身突击战法的发动几率提高15%并使自己成功发动突击战法后，对目标额外发动一次兵刃攻击（95%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		// 战斗中，使自己的普通攻击目标有68%几率锁定为敌军主将
		if util.GenerateRate(0.68) {
			enemyMasterGeneral := util.GetEnemyMasterGeneral(r.tacticsParams)
			r.tacticsParams.CurrentSufferGeneral = enemyMasterGeneral
		}
		//自身突击战法的发动几率提高15%
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsAssaultTriggerImprove, &vo.EffectHolderParams{
			TriggerRate:    0.15,
			FromTactic:     r.Id(),
			ProduceGeneral: currentGeneral,
		})
		//并使自己成功发动突击战法后，对目标额外发动一次兵刃攻击（95%）
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AssaultTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerAssaultResp := &vo.TacticsTriggerResult{}
			triggerAssaultGeneral := params.CurrentGeneral
			sufferGeneral := r.tacticsParams.CurrentSufferGeneral

			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     r.tacticsParams,
				AttackGeneral:     triggerAssaultGeneral,
				SufferGeneral:     sufferGeneral,
				DamageType:        consts.DamageType_Weapon,
				DamageImproveRate: 0.95,
				TacticId:          r.Id(),
				TacticName:        r.Name(),
			})

			return triggerAssaultResp
		})

		return triggerResp
	})
	// 自身为主将时，获得6%倒戈
	if currentGeneral.IsMaster {
		util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
			EffectRate:     0.06,
			FromTactic:     r.Id(),
			ProduceGeneral: currentGeneral,
		})
	}
}

func (r RaidInFormationTactic) Id() consts.TacticId {
	return consts.RaidInFormation
}

func (r RaidInFormationTactic) Name() string {
	return "陷阵突袭"
}

func (r RaidInFormationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (r RaidInFormationTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RaidInFormationTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RaidInFormationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (r RaidInFormationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RaidInFormationTactic) Execute() {
}

func (r RaidInFormationTactic) IsTriggerPrepare() bool {
	return false
}
