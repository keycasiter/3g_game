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

// 暴敛四方
// 对敌军单体发动2次兵刃伤害（伤害率102%），如果目标处于震慑状态，对其造成禁疗效果，持续2回合，2次攻击目标独立判定
// 主动 45%
type OverwhelmingAllDirectionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OverwhelmingAllDirectionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	o.tacticsParams = tacticsParams
	o.triggerRate = 0.45
	return o
}

func (o OverwhelmingAllDirectionsTactic) Prepare() {
}

func (o OverwhelmingAllDirectionsTactic) Id() consts.TacticId {
	return consts.OverwhelmingAllDirections
}

func (o OverwhelmingAllDirectionsTactic) Name() string {
	return "暴敛四方"
}

func (o OverwhelmingAllDirectionsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (o OverwhelmingAllDirectionsTactic) GetTriggerRate() float64 {
	return o.triggerRate
}

func (o OverwhelmingAllDirectionsTactic) SetTriggerRate(rate float64) {
	o.triggerRate = rate
}

func (o OverwhelmingAllDirectionsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (o OverwhelmingAllDirectionsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (o OverwhelmingAllDirectionsTactic) Execute() {
	ctx := o.tacticsParams.Ctx
	currentGeneral := o.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		o.Name(),
	)
	// 对敌军单体发动2次兵刃伤害（伤害率102%），如果目标处于震慑状态，对其造成禁疗效果，持续2回合，2次攻击目标独立判定
	for i := 0; i < 2; i++ {
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, o.tacticsParams)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     o.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Weapon,
			DamageImproveRate: 1.02,
			TacticId:          o.Id(),
			TacticName:        o.Name(),
		})
		if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Awe) {
			util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     o.Id(),
				ProduceGeneral: currentGeneral,
			})
		}
	}
}

func (o OverwhelmingAllDirectionsTactic) IsTriggerPrepare() bool {
	return false
}

func (a OverwhelmingAllDirectionsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
