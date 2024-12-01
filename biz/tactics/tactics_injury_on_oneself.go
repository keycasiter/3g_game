package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 苦肉计
// 对自己造成兵刃伤害（伤害率40%），使敌军单体进入灼烧（伤害率122%，受智力影响）及混乱状态
// 并使自己受到兵刃伤害降低30%，持续2回合
// 主动，40%
type InjuryOnOneselfTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InjuryOnOneselfTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.4
	return i
}

func (i InjuryOnOneselfTactic) Prepare() {

}

func (i InjuryOnOneselfTactic) Id() consts.TacticId {
	return consts.InjuryOnOneself
}

func (i InjuryOnOneselfTactic) Name() string {
	return "苦肉计"
}

func (i InjuryOnOneselfTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i InjuryOnOneselfTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i InjuryOnOneselfTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i InjuryOnOneselfTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InjuryOnOneselfTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InjuryOnOneselfTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)
	// 对自己造成兵刃伤害（伤害率40%），
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     i.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     currentGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 0.4,
		TacticId:          i.Id(),
		TacticName:        i.Name(),
	})
	//使敌军单体进入灼烧（伤害率122%，受智力影响）及混乱状态
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, i.tacticsParams)
	//灼烧
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     i.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Firing,
				TacticId:   i.Id(),
			}) {
				//伤害
				fireDmgRate := revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.22
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     i.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     revokeGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: fireDmgRate,
					TacticId:          i.Id(),
					TacticName:        i.Name(),
					EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
				})
			}

			return revokeResp
		})
	}
	//混乱
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     i.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Chaos,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
	// 并使自己受到兵刃伤害降低30%，持续2回合
	if util.BuffEffectWrapSet(ctx, enemyGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     i.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
}

func (i InjuryOnOneselfTactic) IsTriggerPrepare() bool {
	return false
}

func (a InjuryOnOneselfTactic) SetTriggerPrepare(triggerPrepare bool) {
}
