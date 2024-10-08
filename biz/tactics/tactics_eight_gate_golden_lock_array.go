package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 八门金锁阵
// 战斗前3回合，使敌军群体（2人）造成的伤害降低30%（受智力影响），
// 并使我军主将获得先攻状态（优先行动）
type EightGateGoldenLockArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e EightGateGoldenLockArrayTactic) IsTriggerPrepare() bool {
	return false
}

func (e EightGateGoldenLockArrayTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EightGateGoldenLockArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (e EightGateGoldenLockArrayTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EightGateGoldenLockArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 1.0
	return e
}

func (e EightGateGoldenLockArrayTactic) Prepare() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		e.Name(),
	)

	//战斗前3回合，使敌军群体（2人）造成的伤害降低30%（受智力影响）
	//找到敌军2人
	enemyGenerals := util.GetEnemyGeneralsTwoArr(e.tacticsParams)
	//造成的伤害降低30%
	//TODO（受智力影响）
	effectRate := 0.3 + (currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100)

	for _, sufferGeneral := range enemyGenerals {
		//效果施加
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate: effectRate,
			FromTactic: e.Id(),
		}).IsSuccess {
			hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害降低了%.2f%%",
				sufferGeneral.BaseInfo.Name,
				effectRate*100,
			)
		}

		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate: effectRate,
			FromTactic: e.Id(),
		}).IsSuccess {
			hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了%.2f%%",
				sufferGeneral.BaseInfo.Name,
				effectRate*100,
			)
		}
		//注册消失效果
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}

			//第四回合消失
			if params.CurrentRound == consts.Battle_Round_Fourth {
				triggerGeneral := params.CurrentGeneral
				if util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, e.Id()) {
					hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害提高了%.2f%%",
						triggerGeneral.BaseInfo.Name,
						effectRate*100,
					)
				}
				if util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, e.Id()) {
					hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提高了%.2f%%",
						triggerGeneral.BaseInfo.Name,
						effectRate*100,
					)
				}
			}
			return triggerResp
		})
	}
	//并使我军主将获得先攻状态（优先行动）
	//找到我军主将
	pairMasterGeneral := util.GetPairMasterGeneral(e.tacticsParams)
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_FirstAttack, &vo.EffectHolderParams{
		EffectRate: 1.0,
		FromTactic: e.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			if params.CurrentRound == consts.Battle_Round_Third {
				util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_FirstAttack, e.Id())
			}
			return triggerResp
		})
	}
}

func (e EightGateGoldenLockArrayTactic) Id() consts.TacticId {
	return consts.EightGateGoldenLockArray
}

func (e EightGateGoldenLockArrayTactic) Name() string {
	return "八门金锁阵"
}

func (e EightGateGoldenLockArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (e EightGateGoldenLockArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EightGateGoldenLockArrayTactic) Execute() {
	return
}
