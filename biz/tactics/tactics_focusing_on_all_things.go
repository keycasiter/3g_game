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

// 垂心万物
// 战斗中
// - 奇数回合有70%概率(受智力影响)使我军武力最高的武将造成兵刃伤害提高20%（受智力影响，持续到其行动前）
// 若目标不处于连击状态则令其尝试对敌军单体发动一次普通攻击
// 否则自身对敌军单体造成谋略伤害（伤害率188%，受智力影响）
// - 偶数回合有70%概率（受智力影响）恢复我军群体（2人）兵力（治疗率86%，受智力影响）
// 指挥 100%
type FocusingOnAllThingsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FocusingOnAllThingsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FocusingOnAllThingsTactic) Prepare() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	// - 奇数回合有70%概率(受智力影响)使我军武力最高的武将造成兵刃伤害提高20%（受智力影响，持续到其行动前）
	// 若目标不处于连击状态则令其尝试对敌军单体发动一次普通攻击
	// 否则自身对敌军单体造成谋略伤害（伤害率188%，受智力影响）
	// - 偶数回合有70%概率（受智力影响）恢复我军群体（2人）兵力（治疗率86%，受智力影响）

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//奇数回合
		if triggerRound%2 != 0 {
			triggerRate := 0.7 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			if util.GenerateRate(triggerRate) {
				general := util.GetPairGeneralWhoIsHighestForce(f.tacticsParams)
				//施加效果
				effectRate := 0.2 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     effectRate,
					EffectRound:    1,
					FromTactic:     f.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//消失效果
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    general,
						EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
						TacticId:   f.Id(),
					})
				}

				//找到敌军单体
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(general, f.tacticsParams)
				//若目标不处于连击状态则令其尝试对敌军单体发动一次普通攻击
				if !util.BuffEffectContains(general, consts.BuffEffectType_ContinuousAttack) {
					util.AttackDamage(f.tacticsParams, general, enemyGeneral, 0)
				} else {
					//否则自身对敌军单体造成谋略伤害（伤害率188%，受智力影响）
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.88)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: f.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: enemyGeneral,
						DamageType:    consts.DamageType_Strategy,
						Damage:        dmg,
						TacticName:    f.Name(),
						TacticId:      f.Id(),
					})
				}
			}
		}
		//偶数回合有70%概率（受智力影响）恢复我军群体（2人）兵力（治疗率86%，受智力影响）
		if triggerRound%2 == 0 {
			triggerRate := 0.7 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			if util.GenerateRate(triggerRate) {
				pairGenerals := util.GetPairGeneralsTwoArrByGeneral(currentGeneral, f.tacticsParams)
				for _, general := range pairGenerals {
					resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.86)
					util.ResumeSoldierNum(&util.ResumeParams{
						Ctx:            ctx,
						TacticsParams:  f.tacticsParams,
						ProduceGeneral: currentGeneral,
						SufferGeneral:  general,
						ResumeNum:      resumeNum,
						TacticId:       f.Id(),
					})
				}
			}
		}
		return triggerResp
	})
}

func (f FocusingOnAllThingsTactic) Id() consts.TacticId {
	return consts.FocusingOnAllThings
}

func (f FocusingOnAllThingsTactic) Name() string {
	return "垂心万物"
}

func (f FocusingOnAllThingsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FocusingOnAllThingsTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FocusingOnAllThingsTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FocusingOnAllThingsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (f FocusingOnAllThingsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FocusingOnAllThingsTactic) Execute() {

}

func (f FocusingOnAllThingsTactic) IsTriggerPrepare() bool {
	return false
}
