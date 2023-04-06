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
}

func (e EightGateGoldenLockArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
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
		sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] += effectRate
		hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害降低了%.2f%%",
			sufferGeneral.BaseInfo.Name,
			effectRate*100,
		)
		sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] += effectRate
		hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了%.2f%%",
			sufferGeneral.BaseInfo.Name,
			effectRate*100,
		)
		//注册消失效果
		util.TacticsTriggerWrapSet(sufferGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) {
			//第四回合消失
			if e.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
				sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] -= effectRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					sufferGeneral.BaseInfo.Name,
					consts.DebuffEffectType_LaunchStrategyDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]造成的兵刃伤害提高了%d%%",
					sufferGeneral.BaseInfo.Name,
					effectRate*100,
				)
				sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchStrategyDamageDeduce] -= effectRate
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					sufferGeneral.BaseInfo.Name,
					consts.DebuffEffectType_LaunchStrategyDamageDeduce,
				)
				hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提高了%d%%",
					sufferGeneral.BaseInfo.Name,
					effectRate*100,
				)
			}
		})
	}
	//并使我军主将获得先攻状态（优先行动）
	//找到我军主将
	pairMasterGeneral := util.GetPairMasterGeneral(e.tacticsParams)
	pairMasterGeneral.BuffEffectHolderMap[consts.BuffEffectType_FirstAttack] = 1.0
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		pairMasterGeneral.BaseInfo.Name,
		consts.BuffEffectType_FirstAttack,
	)
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
