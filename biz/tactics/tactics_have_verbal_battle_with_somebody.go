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

// 唇枪舌战
// 对敌军全体造成谋略攻击（伤害率60%，受智力影响）及嘲讽，持续1回合
// 主动 50%
type HaveVerbalBattleWithSomebodyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HaveVerbalBattleWithSomebodyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.5
	return h
}

func (h HaveVerbalBattleWithSomebodyTactic) Prepare() {

}

func (h HaveVerbalBattleWithSomebodyTactic) Id() consts.TacticId {
	return consts.HaveVerbalBattleWithSomebody
}

func (h HaveVerbalBattleWithSomebodyTactic) Name() string {
	return "唇枪舌战"
}

func (h HaveVerbalBattleWithSomebodyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HaveVerbalBattleWithSomebodyTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HaveVerbalBattleWithSomebodyTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HaveVerbalBattleWithSomebodyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HaveVerbalBattleWithSomebodyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HaveVerbalBattleWithSomebodyTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	// 对敌军全体造成谋略攻击（伤害率60%，受智力影响）及嘲讽，持续1回合
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, h.tacticsParams)
	for _, general := range enemyGenerals {
		//攻击
		dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.6
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     h.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     general,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          h.Id(),
			TacticName:        h.Name(),
		})
		//施加效果
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:   1,
			FromTactic:    h.Id(),
			TauntByTarget: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Taunt,
					TacticId:   h.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (h HaveVerbalBattleWithSomebodyTactic) IsTriggerPrepare() bool {
	return false
}

func (a HaveVerbalBattleWithSomebodyTactic) SetTriggerPrepare(triggerPrepare bool) {
}
