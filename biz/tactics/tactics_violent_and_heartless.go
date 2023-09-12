package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 暴戾无仁
// 普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率196%），并使其进入混乱状态（攻击和战法无差别选择目标）,持续1回合
// 发动概率35%
type ViolentAndHeartlessTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (v ViolentAndHeartlessTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	v.tacticsParams = tacticsParams
	v.triggerRate = 0.35
	return v
}

func (v ViolentAndHeartlessTactic) Prepare() {
}

func (v ViolentAndHeartlessTactic) Id() consts.TacticId {
	return consts.ViolentAndHeartless
}

func (v ViolentAndHeartlessTactic) Name() string {
	return "暴戾无仁"
}

func (v ViolentAndHeartlessTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (v ViolentAndHeartlessTactic) GetTriggerRate() float64 {
	return v.triggerRate
}

func (v ViolentAndHeartlessTactic) SetTriggerRate(rate float64) {
	v.triggerRate = rate
}

func (v ViolentAndHeartlessTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (v ViolentAndHeartlessTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (v ViolentAndHeartlessTactic) Execute() {
	ctx := v.tacticsParams.Ctx
	currentGeneral := v.tacticsParams.CurrentGeneral
	currentRound := v.tacticsParams.CurrentRound

	//普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率196%），并使其进入混乱状态（攻击和战法无差别选择目标）
	//持续1回合
	//伤害
	enemyGenerals := util.GetEnemyGeneralArr(v.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.96)
	for _, enemyGeneral := range enemyGenerals {
		//伤害
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: v.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			Damage:        dmg,
			DamageType:    consts.DamageType_Weapon,
			TacticId:      v.Id(),
			TacticName:    v.Name(),
		})
		//混乱状态
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
			EffectRate: 1.0,
			FromTactic: v.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral
				revokeRound := params.CurrentRound

				if currentRound+1 == revokeRound {
					util.DebuffEffectWrapRemove(ctx, revokeGeneral, consts.DebuffEffectType_Chaos, v.Id())
				}

				return revokeResp
			})
		}
	}
}

func (v ViolentAndHeartlessTactic) IsTriggerPrepare() bool {
	return false
}
