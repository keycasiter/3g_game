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

// 骁健神行
// 对敌军单体造成缴械状态，持续1回合，且使自身获得必中状态，持续2回合，
// 如果目标已经被缴械则造成兵刃攻击（伤害率156%，受速度影响）
// 主动，45%
type VigorousAndWalkTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (v VigorousAndWalkTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	v.tacticsParams = tacticsParams
	v.triggerRate = 0.45
	return v
}

func (v VigorousAndWalkTactic) Prepare() {

}

func (v VigorousAndWalkTactic) Id() consts.TacticId {
	return consts.VigorousAndWalk
}

func (v VigorousAndWalkTactic) Name() string {
	return "骁健神行"
}

func (v VigorousAndWalkTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (v VigorousAndWalkTactic) GetTriggerRate() float64 {
	return v.triggerRate
}

func (v VigorousAndWalkTactic) SetTriggerRate(rate float64) {
	v.triggerRate = rate
}

func (v VigorousAndWalkTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (v VigorousAndWalkTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (v VigorousAndWalkTactic) Execute() {
	ctx := v.tacticsParams.Ctx
	currentGeneral := v.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		v.Name(),
	)
	// 如果目标已经被缴械则造成兵刃攻击（伤害率156%，受速度影响）
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, v.tacticsParams)
	if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_CancelWeapon) {
		dmgRate := currentGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100 + 1.56
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     v.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Weapon,
			DamageImproveRate: dmgRate,
			TacticId:          v.Id(),
			TacticName:        v.Name(),
		})
	} else {
		// 对敌军单体造成缴械状态，持续1回合，
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     v.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_CancelWeapon,
					TacticId:   v.Id(),
				})

				return revokeResp
			})
		}
	}
	//且使自身获得必中状态，持续2回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_MustHit, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     v.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_MustHit,
				TacticId:   v.Id(),
			})

			return revokeResp
		})
	}
}

func (v VigorousAndWalkTactic) IsTriggerPrepare() bool {
	return false
}

func (a VigorousAndWalkTactic) SetTriggerPrepare(triggerPrepare bool) {
}
