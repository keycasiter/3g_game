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

// 卧薪尝胆
// 对敌军群体（2人）发动一次兵刃攻击（伤害率96%），
// 并有25几率（根据自身连击、洞察、先攻、必中、破阵的状态数，每多一种提高10%几率）造成震慑（无法行动）状态，持续1回合
type SleepOnTheBrushwoodAndTasteTheGallTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.45
	return s
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Prepare() {
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Id() consts.TacticId {
	return consts.SleepOnTheBrushwoodAndTasteTheGall
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Name() string {
	return "卧薪尝胆"
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	//对敌军群体（2人）发动一次兵刃攻击（伤害率96%），
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, s.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.96)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: s.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      s.Id(),
			TacticName:    s.Name(),
		})

		//并有25几率（根据自身连击、洞察、先攻、必中、破阵的状态数，每多一种提高10%几率）造成震慑（无法行动）状态，持续1回合
		triggerRate := 0.25
		buffs := []consts.BuffEffectType{
			consts.BuffEffectType_ContinuousAttack,
			consts.BuffEffectType_Insight,
			consts.BuffEffectType_FirstAttack,
			consts.BuffEffectType_MustHit,
			consts.BuffEffectType_BreakFormation,
		}
		for _, buff := range buffs {
			if util.BuffEffectContains(currentGeneral, buff) {
				triggerRate += 0.1
			}
		}
		if util.GenerateRate(triggerRate) {
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
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
						EffectType: consts.DebuffEffectType_Awe,
						TacticId:   s.Id(),
					})

					return revokeResp
				})
			}
		}
	}

}

func (s SleepOnTheBrushwoodAndTasteTheGallTactic) IsTriggerPrepare() bool {
	return false
}
