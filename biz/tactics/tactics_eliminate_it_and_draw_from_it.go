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

// 绝其汲道
// 准备1回合，对敌军群体（2-3人）造成一次兵刃攻击（伤害率162%），使其进入禁疗状态（无法恢复兵力），持续1回合
type EliminateItAndDrawFromItTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (e EliminateItAndDrawFromItTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 0.5
	return e
}

func (e EliminateItAndDrawFromItTactic) Prepare() {

}

func (e EliminateItAndDrawFromItTactic) Id() consts.TacticId {
	return consts.EliminateItAndDrawFromIt
}

func (e EliminateItAndDrawFromItTactic) Name() string {
	return "绝其汲道"
}

func (e EliminateItAndDrawFromItTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (e EliminateItAndDrawFromItTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EliminateItAndDrawFromItTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EliminateItAndDrawFromItTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (e EliminateItAndDrawFromItTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EliminateItAndDrawFromItTactic) Execute() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral
	currentRound := e.tacticsParams.CurrentRound

	//准备1回合，对敌军群体（2-3人）造成一次兵刃攻击（伤害率162%），使其进入禁疗状态（无法恢复兵力），持续1回合
	e.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		e.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			e.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if e.isTriggered {
				return triggerResp
			} else {
				e.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				e.Name(),
			)

			//找到敌军2人~3人
			enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(e.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//造成一次兵刃攻击（伤害率162%）
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     e.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 1.62,
					TacticId:          e.Id(),
					TacticName:        e.Name(),
				})
				//使其进入禁疗状态（无法恢复兵力），持续1回合
				util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
					EffectRound: 1,
					FromTactic:  e.Id(),
				})
				//注册消失效果
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revoekGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revoekGeneral,
						EffectType: consts.DebuffEffectType_ProhibitionTreatment,
						TacticId:   e.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (e EliminateItAndDrawFromItTactic) IsTriggerPrepare() bool {
	return e.isTriggerPrepare
}
