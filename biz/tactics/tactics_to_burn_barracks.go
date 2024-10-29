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

// 焚辎营垒
// 对敌军群体（2人）造成谋略伤害（伤害率146%，受智力影响）并使其进入禁疗状态，持续1回合
type ToBurnBarracksTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToBurnBarracksTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.4
	return t
}

func (t ToBurnBarracksTactic) Prepare() {
}

func (t ToBurnBarracksTactic) Id() consts.TacticId {
	return consts.ToBurnBarracks
}

func (t ToBurnBarracksTactic) Name() string {
	return "焚辎营垒"
}

func (t ToBurnBarracksTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ToBurnBarracksTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToBurnBarracksTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToBurnBarracksTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ToBurnBarracksTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToBurnBarracksTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 对敌军群体（2人）造成谋略伤害（伤害率146%，受智力影响）并使其进入禁疗状态，持续1回合
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, t.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//伤害
		dmgRate := enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.46
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     t.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          t.Id(),
			TacticName:        t.Name(),
		})
		//效果
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: enemyGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_ProhibitionTreatment,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (t ToBurnBarracksTactic) IsTriggerPrepare() bool {
	return false
}
