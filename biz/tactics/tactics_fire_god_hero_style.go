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

// 火神英风
// 战斗前2回合，对敌军全体武将分别有15%概率（处于灼烧状态的目标为70%概率）发动兵刃攻击（伤害率142%），
// 战斗第3回合时，使我军群体（2～3人）武力提升20（受武力影响）且自身每回合行动前会治疗我军群体（2人，治疗率158%，受武力影响），持续3回合；
// 自身为主将时，兵刃攻击触发概率提升10%
// 指挥 100%
type FireGodHeroStyleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FireGodHeroStyleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FireGodHeroStyleTactic) Prepare() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	// 战斗前2回合，对敌军全体武将分别有15%概率（处于灼烧状态的目标为70%概率）发动兵刃攻击（伤害率142%），
	//找到敌军全体
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, f.tacticsParams)
	for _, general := range enemyGenerals {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound

			//战斗前2回合
			if triggerRound <= consts.Battle_Round_Second {
				triggerRate := 0.15
				if util.DeBuffEffectContains(general, consts.DebuffEffectType_Firing) {
					triggerRate = 0.7
				}
				//自身为主将时，兵刃攻击触发概率提升10%
				if currentGeneral.IsMaster {
					triggerRate += 0.1
				}

				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.42)
				if util.GenerateRate(triggerRate) {
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: f.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: general,
						DamageType:    consts.DamageType_Weapon,
						Damage:        dmg,
						TacticName:    f.Name(),
					})
				}
			}

			return triggerResp
		})
	}
	// 战斗第3回合时，使我军群体（2～3人）武力提升20（受武力影响）且自身每回合行动前会治疗我军群体（2人，治疗率158%，受武力影响），持续3回合；
	//找到我军群体2～3人
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}

		if triggerRound == consts.Battle_Round_Third {
			pairGenerals := util.GetPairGeneralsTwoOrThreeMap(f.tacticsParams)
			for _, general := range pairGenerals {
				//武力提升
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
					EffectValue:    20,
					EffectRound:    3,
					FromTactic:     f.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//消失效果
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_IncrForce,
							TacticId:   f.Id(),
						})

						return revokeResp
					})
				}
			}
		}
		return triggerResp
	})

	//自身每回合行动前会治疗我军群体（2人，治疗率158%，受武力影响），持续3回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_FireGodHeroStyle_Prepare, &vo.EffectHolderParams{
		EffectRound:    3,
		FromTactic:     f.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_FireGodHeroStyle_Prepare,
				TacticId:   f.Id(),
			}) {
				//找到我军群体2人
				pairGenerals := util.GetPairGeneralsTwoArrByGeneral(revokeGeneral, f.tacticsParams)
				for _, general := range pairGenerals {
					resumeNum := cast.ToInt64(general.BaseInfo.AbilityAttr.ForceBase * 1.58)
					util.ResumeSoldierNum(ctx, general, resumeNum)
				}
			}

			return revokeResp
		})
	}
}

func (f FireGodHeroStyleTactic) Id() consts.TacticId {
	return consts.FireGodHeroStyle
}

func (f FireGodHeroStyleTactic) Name() string {
	return "火神英风"
}

func (f FireGodHeroStyleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FireGodHeroStyleTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FireGodHeroStyleTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FireGodHeroStyleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (f FireGodHeroStyleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FireGodHeroStyleTactic) Execute() {
}

func (f FireGodHeroStyleTactic) IsTriggerPrepare() bool {
	return false
}
