package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 敛众而击
// 对敌军群体(1-2人)造成兵刃伤害（伤害率164%），并有45%概率治疗自身（治疗率88%，受武力影响）
type GatherTheCrowdAndStrikeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GatherTheCrowdAndStrikeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.35
	return g
}

func (g GatherTheCrowdAndStrikeTactic) Prepare() {

}

func (g GatherTheCrowdAndStrikeTactic) Id() consts.TacticId {
	return consts.GatherTheCrowdAndStrike
}

func (g GatherTheCrowdAndStrikeTactic) Name() string {
	return "敛众而击"
}

func (g GatherTheCrowdAndStrikeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (g GatherTheCrowdAndStrikeTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GatherTheCrowdAndStrikeTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GatherTheCrowdAndStrikeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (g GatherTheCrowdAndStrikeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GatherTheCrowdAndStrikeTactic) Execute() {
	currentGeneral := g.tacticsParams.CurrentGeneral
	ctx := g.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)
	//对敌军群体(1-2人)造成兵刃伤害（伤害率164%）
	//找到敌军
	enemyGenerals := util.GetEnemyGeneralsOneOrTwoArr(g.tacticsParams)
	for _, general := range enemyGenerals {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.64)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: g.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: general,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticName:    g.Name(),
		})
	}
	//并有45%概率治疗自身（治疗率88%，受武力影响）
	if util.GenerateRate(0.45) {
		resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.88)
		util.ResumeSoldierNum(ctx, currentGeneral, resumeNum)
	}
}

func (g GatherTheCrowdAndStrikeTactic) IsTriggerPrepare() bool {
	return false
}
