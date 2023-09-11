package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 落凤
// 对随机敌军单体造成兵刃攻击（伤害率250%），并计穷（无法发动主动战法）1回合
type FallingPhoenixTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FallingPhoenixTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FallingPhoenixTactic) Prepare() {
}

func (f FallingPhoenixTactic) Id() consts.TacticId {
	return consts.FallingPhoenix
}

func (f FallingPhoenixTactic) Name() string {
	return "落凤"
}

func (f FallingPhoenixTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FallingPhoenixTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FallingPhoenixTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FallingPhoenixTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FallingPhoenixTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FallingPhoenixTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	//对随机敌军单体造成兵刃攻击（伤害率250%），并计穷（无法发动主动战法）1回合
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, f.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.5)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: f.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    f.Name(),
		TacticId:      f.Id(),
	})
	//施加效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     f.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		//消失效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_NoStrategy,
				TacticId:   f.Id(),
			})

			return revokeResp
		})
	}
}

func (f FallingPhoenixTactic) IsTriggerPrepare() bool {
	return false
}
