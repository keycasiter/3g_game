package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：士别三日
// 战法描述：战斗前3回合，无法进行普通攻击但获得30%概率规避效果
// 第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
type ThreeDaysOfSeparationTactic struct {
	tacticsParams model.TacticsParams
}

func (t ThreeDaysOfSeparationTactic) Name() string {
	return "士别三日"
}

func (t ThreeDaysOfSeparationTactic) BuffEffect() map[int64]map[consts.BuffEffectType]map[consts.BattleRound]float64 {
	currentGeneralId := t.tacticsParams.CurrentGeneral.BaseInfo.Id
	return map[int64]map[consts.BuffEffectType]map[consts.BattleRound]float64{
		currentGeneralId: {
			//第4回合，提高自己68点智力
			consts.BuffEffectType_IncrIntelligence: {
				consts.Battle_Round_Fourth: 68,
			},
			//战斗前3回合,获得30%概率规避效果
			consts.BuffEffectType_Evade: {
				consts.Battle_Round_First:  0.3,
				consts.Battle_Round_Second: 0.3,
				consts.Battle_Round_Third:  0.3,
			},
		},
	}
	return nil
}

func (t ThreeDaysOfSeparationTactic) DebuffEffect() map[int64]map[consts.DebuffEffectType]map[consts.BattleRound]float64 {
	//战斗前3回合，无法进行普通攻击
	currentGeneralId := t.tacticsParams.CurrentGeneral.BaseInfo.Id
	return map[int64]map[consts.DebuffEffectType]map[consts.BattleRound]float64{
		currentGeneralId: {
			consts.DebuffEffectType_CanNotGeneralAttack: {
				consts.Battle_Round_First:  1.0,
				consts.Battle_Round_Second: 1.0,
				consts.Battle_Round_Third:  1.0,
			},
		},
	}
}

func (t ThreeDaysOfSeparationTactic) Damage() map[int64]map[consts.BattleRound]map[consts.DamageType]float64 {
	//第四回合，对敌军全体造成谋略伤害(伤害率180%，受智力影响)
	mm := make(map[int64]map[consts.BattleRound]map[consts.DamageType]float64, 0)
	//找到敌军全体武将进行伤害输出
	for _, general := range t.tacticsParams.EnemyGeneralMap {
		dmg := 1.8 * general.BaseInfo.AbilityAttr.IntelligenceBase
		mm[general.BaseInfo.Id][consts.Battle_Round_Fourth][consts.DamageType_Strategy] = dmg
	}
	return mm
}

func (t ThreeDaysOfSeparationTactic) Resume() map[int64]map[consts.BattleRound]float64 {
	return nil
}

func (t ThreeDaysOfSeparationTactic) GetCurrentRound() consts.BattleRound {
	return consts.Battle_Round_Unknow
}

func (t ThreeDaysOfSeparationTactic) LastTriggerRound() consts.BattleRound {
	return consts.Battle_Round_Unknow
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t ThreeDaysOfSeparationTactic) Id() int64 {
	return ThreeDaysOfSeparation
}

func (t ThreeDaysOfSeparationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThreeDaysOfSeparationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThreeDaysOfSeparationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThreeDaysOfSeparationTactic) TriggerRate() float64 {
	return 1.0
}
