package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 战法名称：士别三日
// 战法描述：战斗前3回合，无法进行普通攻击但获得30%概率规避效果
// 第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
type ThreeDaysOfSeparationTactic struct {
	tacticsParams model.TacticsParams
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t ThreeDaysOfSeparationTactic) Prepare() {
	//战斗前3回合，无法进行普通攻击
	currentGeneral := t.tacticsParams.CurrentGeneral
	currentGeneral.DeBuffEffectTriggerMap[consts.DebuffEffectType_CanNotGeneralAttack][consts.Battle_Round_First] = 1.0
	currentGeneral.DeBuffEffectTriggerMap[consts.DebuffEffectType_CanNotGeneralAttack][consts.Battle_Round_Second] = 1.0
	currentGeneral.DeBuffEffectTriggerMap[consts.DebuffEffectType_CanNotGeneralAttack][consts.Battle_Round_Third] = 1.0
	//战斗前3回合，获得30%概率规避效果
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade] += 0.3
	currentGeneral.BuffEffectTriggerMap[consts.BuffEffectType_Evade_Disappear][consts.Battle_Round_Fourth] -= 0.3
}

func (t ThreeDaysOfSeparationTactic) Execute() {
	//第4回合
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		currentGeneral := t.tacticsParams.CurrentGeneral
		//提高自己68点智力
		currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_IncrIntelligence] += 68

		//找到敌人全体
		enemyGenerals := util.GetEnemyGeneralArr(t.tacticsParams)
		//对敌军全体造成谋略伤害(伤害率180%，受智力影响)
		for _, general := range enemyGenerals {
			// TODO 受智力影响
			damage := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.8
			general.SoldierNum = util.TacticsDamage(general.SoldierNum, cast.ToInt64(damage))
		}
	}
}

func (t ThreeDaysOfSeparationTactic) Trigger() {
	return
}

func (t ThreeDaysOfSeparationTactic) Name() string {
	return "士别三日"
}

func (t ThreeDaysOfSeparationTactic) Id() int64 {
	return ThreeDaysOfSeparation
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
