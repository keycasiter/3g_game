package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 金城汤池
// 无法发动普通攻击（无法被净化），每2回合轮流执行：
// 治疗我军群体（2人，治疗率98%，受智力影响）；
// 对敌军群体造成兵刃伤害（伤害率102%）及灼烧状态，每回合持续造成伤害（伤害率72%，受智力影响），持续1回合
// 指挥 100%
type RampartsOfMetalsAndAMoatOfHotWaterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral
	executeFlag := true

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	// 无法发动普通攻击（无法被净化），每2回合轮流执行：
	util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_RampartsOfMetalsAndAMoatOfHotWaterTactic_CanNotGeneralAttack, &vo.EffectHolderParams{
		FromTactic:     r.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound%2 != 0 {
			if executeFlag {
				// 治疗我军群体（2人，治疗率98%，受智力影响）；
				pairGenerals := util.GetPairGeneralsTwoOrThreeMap(currentGeneral, r.tacticsParams)
				for _, pairGeneral := range pairGenerals {
					resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.98)
					util.ResumeSoldierNum(&util.ResumeParams{
						Ctx:            ctx,
						TacticsParams:  r.tacticsParams,
						ProduceGeneral: currentGeneral,
						SufferGeneral:  pairGeneral,
						ResumeNum:      resumeNum,
						TacticId:       r.Id(),
					})
				}

				//切换执行标志
				executeFlag = false
			} else {
				// 对敌军群体造成兵刃伤害（伤害率102%）及灼烧状态，每回合持续造成伤害（伤害率72%，受智力影响），持续1回合
				enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, r.tacticsParams)
				for _, general := range enemyGenerals {
					//伤害
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     r.tacticsParams,
						AttackGeneral:     triggerGeneral,
						SufferGeneral:     general,
						DamageType:        consts.DamageType_Weapon,
						DamageImproveRate: 1.02,
						TacticId:          r.Id(),
						TacticName:        r.Name(),
					})
					//效果
					if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
						EffectRound:    1,
						FromTactic:     r.Id(),
						ProduceGeneral: triggerGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_Firing,
								TacticId:   r.Id(),
							}) {
								fireDmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.72
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:     r.tacticsParams,
									AttackGeneral:     currentGeneral,
									SufferGeneral:     revokeGeneral,
									DamageType:        consts.DamageType_Strategy,
									DamageImproveRate: fireDmgRate,
									TacticId:          r.Id(),
									TacticName:        r.Name(),
									EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
								})
							}

							return revokeResp
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Id() consts.TacticId {
	return consts.RampartsOfMetalsAndAMoatOfHotWater
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Name() string {
	return "金城汤池"
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Execute() {
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) IsTriggerPrepare() bool {
	return false
}

func (a RampartsOfMetalsAndAMoatOfHotWaterTactic) SetTriggerPrepare(triggerPrepare bool) {
}
