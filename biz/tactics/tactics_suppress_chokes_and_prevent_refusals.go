package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 战法名称：镇扼防拒
// 战法描述：每回合有50%概率（受智力影响）使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响），
// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态
type SuppressChokesAndPreventRefusalsTactic struct {
	tacticsParams model.TacticsParams
}

func (s SuppressChokesAndPreventRefusalsTactic) Init(tacticsParams model.TacticsParams) {
	s.tacticsParams = tacticsParams
}

func (s SuppressChokesAndPreventRefusalsTactic) Id() int64 {
	return SuppressChokesAndPreventRefusals
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SuppressChokesAndPreventRefusalsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SuppressChokesAndPreventRefusalsTactic) Handle() {
	//每回合有50%概率（受智力影响）使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响），
	// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态

	//获取自身智力
	resumeSoldiersNum := cast.ToInt64(s.tacticsParams.CurrentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.92)

	if util.GenerateRate(0.5) {
		for id, general := range s.tacticsParams.FightingGeneralMap {
			//优先其他副将
			if !general.IsMaster && id != cast.ToInt64(s.tacticsParams.CurrentGeneral.BaseInfo.Id) {
				//治疗
				general.SoldiersNum, resumeSoldiersNum = util.ResumeSoldiersNum(general.SoldiersNum, resumeSoldiersNum)
				hlog.CtxInfof(s.tacticsParams.Ctx, "恢复[%s]兵力[%d]", general.BaseInfo.Name, resumeSoldiersNum)
			}
		}
	}
}

func (s SuppressChokesAndPreventRefusalsTactic) TriggerRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageType() consts.DamageType {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageNum() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageRange() consts.GeneralNum {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsLockingMaster() bool {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsLockingVice() bool {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrDamageNum() int64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrDamageRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DecrDamageNum() int64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DecrDamageRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) ResumeMilitaryStrengthRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EnhancedStrategyDamageRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EnhancedWeaponDamageRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) SuperposeNum() int64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrForceNum() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrIntelligenceNum() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrCommandNum() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrSpeedNum() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EffectNextRounds() int64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) FrozenNextRounds() int64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DebuffEffect() consts.DebuffEffectType {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DebuffEffectRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) BuffEffect() consts.BuffEffectType {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) BuffEffectRate() float64 {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsGeneralAttack() bool {
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EffectNextRoundDamageRate() float64 {
	panic("implement me")
}
