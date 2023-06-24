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

// 震骇四境
// 发动2次对敌军单体的兵刃攻击（伤害率178%），分别造成使其首次受到兵刃伤害提高30%（受武力影响）及计穷状态，持续1回合，每次目标独立选择
// 主动，35%
type ShockingFourRealmsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ShockingFourRealmsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.35
	return s
}

func (s ShockingFourRealmsTactic) Prepare() {

}

func (s ShockingFourRealmsTactic) Id() consts.TacticId {
	return consts.ShockingFourRealms
}

func (s ShockingFourRealmsTactic) Name() string {
	return "震骇四境"
}

func (s ShockingFourRealmsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s ShockingFourRealmsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ShockingFourRealmsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ShockingFourRealmsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s ShockingFourRealmsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ShockingFourRealmsTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 发动2次对敌军单体的兵刃攻击（伤害率178%），分别造成使其首次受到兵刃伤害提高30%（受武力影响）及计穷状态，持续1回合，每次目标独立选择
	for i := 0; i < 2; i++ {
		//发动2次对敌军单体的兵刃攻击（伤害率178%）
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, s.tacticsParams)
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.78)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: s.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      s.Id(),
			TacticName:    s.Name(),
		})
		//分别造成使其首次受到兵刃伤害提高30%（受武力影响）及计穷状态，持续1回合，每次目标独立选择
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ShockingFourRealms_Prepare, &vo.EffectHolderParams{
			EffectRate:     0.3 + (currentGeneral.BaseInfo.AbilityAttr.ForceBase / 100 / 100),
			EffectRound:    1,
			FromTactic:     s.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_ShockingFourRealms_Prepare,
					TacticId:   s.Id(),
				})

				return revokeResp
			})
		}
		//计穷状态，持续1回合，每次目标独立选择
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     s.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_NoStrategy,
					TacticId:   s.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (s ShockingFourRealmsTactic) IsTriggerPrepare() bool {
	return false
}
