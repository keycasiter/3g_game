package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 临机制胜
// 对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率120%，受智力影响），持续2回合
// 若敌军已有中毒状态，则使其随机获得灼烧（受智力影响）、叛逃（受武力活智力最高一项影响，无视防御）、沙暴（受智力影响）状态中的一种
// 每回合持续造成伤害（伤害率120%），持续2回合，该战法发动后进入1回合冷却
type SeizeTheOpportunityToWinTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SeizeTheOpportunityToWinTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.55
	return s
}

func (s SeizeTheOpportunityToWinTactic) Prepare() {
}

func (s SeizeTheOpportunityToWinTactic) Id() consts.TacticId {
	return consts.SeizeTheOpportunityToWin
}

func (s SeizeTheOpportunityToWinTactic) Name() string {
	return "临机制胜"
}

func (s SeizeTheOpportunityToWinTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SeizeTheOpportunityToWinTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SeizeTheOpportunityToWinTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SeizeTheOpportunityToWinTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SeizeTheOpportunityToWinTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SeizeTheOpportunityToWinTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	currentRound := s.tacticsParams.CurrentRound

	//判断是否冷却
	if ok := currentGeneral.TacticFrozenMap[s.Id()]; ok {
		hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效，无法发动",
			currentGeneral.BaseInfo.Name,
			s.Name(),
		)
		return
	}

	currentGeneral.TacticFrozenMap[s.Id()] = true
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	//注册冷却效果消失
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeRound := params.CurrentRound

		//1回合冷却，下下回合冷却结束
		if currentRound+2 == revokeRound {
			currentGeneral.TacticFrozenMap[s.Id()] = false

			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果已消失",
				currentGeneral.BaseInfo.Name,
				s.Name(),
			)
		}
		return revokeResp
	})

	//对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率120%，受智力影响），持续2回合
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, s.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//若敌军已有中毒状态，则使其随机获得灼烧（受智力影响）、叛逃（受武力活智力最高一项影响，无视防御）、沙暴（受智力影响）状态中的一种
		//每回合持续造成伤害（伤害率120%），持续2回合，该战法发动后进入1回合冷却
		if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Methysis) {
			debuffs := []consts.DebuffEffectType{
				consts.DebuffEffectType_Firing,
				consts.DebuffEffectType_Defect,
				consts.DebuffEffectType_Sandstorm,
			}
			hitIdx := util.GenerateHitOneIdx(len(debuffs))
			debuff := debuffs[hitIdx]
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, debuff, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     s.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: debuff,
						TacticId:   s.Id(),
					}) {
						switch debuff {
						case consts.DebuffEffectType_Firing:
							dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.2)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: s.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: revokeGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        dmg,
								TacticId:      s.Id(),
								TacticName:    s.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
							})
						case consts.DebuffEffectType_Defect:
							_, val := util.GetGeneralHighestBetweenForceOrIntelligence(currentGeneral)
							dmg := cast.ToInt64(val * 1.2)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams:  s.tacticsParams,
								AttackGeneral:  currentGeneral,
								SufferGeneral:  revokeGeneral,
								DamageType:     consts.DamageType_Strategy,
								Damage:         dmg,
								TacticId:       s.Id(),
								TacticName:     s.Name(),
								EffectName:     fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
								IsIgnoreDefend: true,
							})
						case consts.DebuffEffectType_Sandstorm:
							dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.2)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: s.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: revokeGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        dmg,
								TacticId:      s.Id(),
								TacticName:    s.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Sandstorm),
							})
						}
					}

					return revokeResp
				})
			}
		} else {
			//未有中毒效果
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Methysis, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     s.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Methysis,
						TacticId:   s.Id(),
					}) {
						dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.2)
						util.TacticDamage(&util.TacticDamageParam{
							TacticsParams: s.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: revokeGeneral,
							DamageType:    consts.DamageType_Strategy,
							Damage:        dmg,
							TacticId:      s.Id(),
							TacticName:    s.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Methysis),
						})
					}

					return revokeResp
				})
			}
		}
	}
}

func (s SeizeTheOpportunityToWinTactic) IsTriggerPrepare() bool {
	return false
}
