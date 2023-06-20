package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 百计多谋
// 战斗中，每回合都有70%几率使随机敌军单体产生逃兵（伤害率140%，受智力影响，无视防御）
// 指挥 100%
type HundredStrategiesAndManyStrategiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HundredStrategiesAndManyStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 1.0
	return h
}

func (h HundredStrategiesAndManyStrategiesTactic) Prepare() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)
	//战斗中，每回合都有70%几率使随机敌军单体产生逃兵（伤害率140%，受智力影响，无视防御）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.7) {
			//敌军单体
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, h.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.4)
			util.TacticDamage(&util.TacticDamageParam{
				TacticsParams:  h.tacticsParams,
				AttackGeneral:  triggerGeneral,
				SufferGeneral:  enemyGeneral,
				DamageType:     consts.DamageType_Strategy,
				Damage:         dmg,
				TacticName:     h.Name(),
				IsIgnoreDefend: true,
			})
		}

		return triggerResp
	})
}

func (h HundredStrategiesAndManyStrategiesTactic) Id() consts.TacticId {
	return consts.HundredStrategiesAndManyStrategies
}

func (h HundredStrategiesAndManyStrategiesTactic) Name() string {
	return "百计多谋"
}

func (h HundredStrategiesAndManyStrategiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HundredStrategiesAndManyStrategiesTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HundredStrategiesAndManyStrategiesTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HundredStrategiesAndManyStrategiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (h HundredStrategiesAndManyStrategiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HundredStrategiesAndManyStrategiesTactic) Execute() {
}

func (h HundredStrategiesAndManyStrategiesTactic) IsTriggerPrepare() bool {
	return false
}
