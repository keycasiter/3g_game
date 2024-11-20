package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 众望所归
// 为我军群体2人恢复一定兵力(治疗率72%，受统率影响)，并使我军武力最高单体发动1次兵刃攻击（伤害率86%），
// 智力最高单体对敌军单体发动1次谋略攻击(伤害率86%，受智力影响)，若我军武力最高者与智力最高者为同一人，则两次攻击伤害率均降至72%
type BeThePopularChoiceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a BeThePopularChoiceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.4
	return a
}

func (a BeThePopularChoiceTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//为我军群体2人恢复一定兵力(治疗率72%，受统率影响)
	pairGeneralArr := util.GetPairGeneralsTwoArr(currentGeneral, a.tacticsParams)
	for _, general := range pairGeneralArr {
		resume := cast.ToInt64(float64(general.SoldierNum) * (0.72 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100))
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  a.tacticsParams,
			ProduceGeneral: currentGeneral,
			SufferGeneral:  general,
			ResumeNum:      resume,
			TacticId:       a.Id(),
		})
	}

	//若我军武力最高者与智力最高者为同一人，则两次攻击伤害率均降至72%
	higestForceGeneral := util.GetPairGeneralWhoIsHighestForce(currentGeneral, a.tacticsParams)
	higestIntelligenceGeneral := util.GetPairGeneralWhoIsHighestIntelligence(currentGeneral, a.tacticsParams)
	dmgBaseRate := 0.86
	if higestForceGeneral.BaseInfo.UniqueId == higestIntelligenceGeneral.BaseInfo.UniqueId {
		dmgBaseRate = 0.72
	}
	//使我军武力最高单体发动1次兵刃攻击（伤害率86%）
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, a.tacticsParams)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     a.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: dmgBaseRate,
		TacticId:          a.Id(),
		TacticName:        a.Name(),
	})

	//智力最高单体对敌军单体发动1次谋略攻击(伤害率86%，受智力影响)
	dmgBaseRate += currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
	enemyGeneral = util.GetEnemyOneGeneralByGeneral(currentGeneral, a.tacticsParams)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     a.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Strategy,
		DamageImproveRate: dmgBaseRate,
		TacticId:          a.Id(),
		TacticName:        a.Name(),
	})
}

func (a BeThePopularChoiceTactic) Id() consts.TacticId {
	return consts.BeThePopularChoice
}

func (a BeThePopularChoiceTactic) Name() string {
	return "众望所归"
}

func (a BeThePopularChoiceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a BeThePopularChoiceTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a BeThePopularChoiceTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a BeThePopularChoiceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a BeThePopularChoiceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a BeThePopularChoiceTactic) Execute() {
}

func (a BeThePopularChoiceTactic) IsTriggerPrepare() bool {
	return false
}
