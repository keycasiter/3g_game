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

// 四面楚歌
// 准备1回合，对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率144%，受智力影响），持续2回合
type BeBesiegedOnAllSidesTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BeBesiegedOnAllSidesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.5
	return b
}

func (b BeBesiegedOnAllSidesTactic) Prepare() {
}

func (b BeBesiegedOnAllSidesTactic) Id() consts.TacticId {
	return consts.BeBesiegedOnAllSides
}

func (b BeBesiegedOnAllSidesTactic) Name() string {
	return "四面楚歌"
}

func (b BeBesiegedOnAllSidesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BeBesiegedOnAllSidesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeBesiegedOnAllSidesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeBesiegedOnAllSidesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BeBesiegedOnAllSidesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeBesiegedOnAllSidesTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//准备1回合，对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率144%，受智力影响），持续2回合
	b.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			b.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)

			//找到敌军2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, b.tacticsParams)
			//中毒效果设置
			for _, enemyGeneral := range enemyGenerals {
				util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Methysis, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  b.Id(),
				})
				//注册伤害效果
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					triggerDmgGeneral := params.CurrentGeneral
					triggerDmgResp := &vo.TacticsTriggerResult{}

					//效果消耗
					if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    triggerGeneral,
						EffectType: consts.DebuffEffectType_Methysis,
						TacticId:   b.Id(),
					}) {
						hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
							triggerDmgGeneral.BaseInfo.Name,
							b.Name(),
							consts.DebuffEffectType_Methysis,
						)
						//伤害
						dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.44)
						for _, enemyGeneral := range enemyGenerals {
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: b.tacticsParams,
								AttackGeneral: triggerGeneral,
								SufferGeneral: enemyGeneral,
								Damage:        dmg,
								DamageType:    consts.DamageType_Strategy,
								TacticName:    b.Name(),
							})
						}
					}

					//效果消失
					if util.DeBuffEffectOfTacticIsDeplete(triggerDmgGeneral, consts.DebuffEffectType_Methysis, b.Id()) {
						util.DebuffEffectWrapRemove(ctx, triggerDmgGeneral, consts.DebuffEffectType_Methysis, b.Id())
					}

					return triggerDmgResp
				})
			}
		}

		return triggerResp
	})
}

func (b BeBesiegedOnAllSidesTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
