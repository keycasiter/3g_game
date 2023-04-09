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

// 战法名称：士别三日
// 战法描述：战斗前3回合，无法进行普通攻击但获得30%概率规避效果
// 第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
type ThreeDaysOfSeparationTactic struct {
	tacticsParams *model.TacticsParams
}

func (t ThreeDaysOfSeparationTactic) TriggerRate() float64 {
	return 1.00
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t ThreeDaysOfSeparationTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_ThreeDaysOfSeparation_Prepare,
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		//第四回合
		if params.CurrentRound == consts.Battle_Round_Fourth {
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				currentGeneral.BaseInfo.Name,
				t.Name(),
				consts.BuffEffectType_ThreeDaysOfSeparation_Prepare,
			)
			currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase += 68
			hlog.CtxInfof(ctx, "[%s]的智力提高了68", currentGeneral.BaseInfo.Name)

			//并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
			//找到敌军全体
			enemyGenerals := util.GetEnemyGeneralArr(t.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				//TODO 受智力影响
				dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.8)
				finalDmg, originNum, remaindNum, isEffect := util.TacticDamage(t.tacticsParams, currentGeneral, sufferGeneral, dmgNum, consts.BattleAction_SufferPassiveTactic)
				if !isEffect {
					return triggerResp
				}
				hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d)",
					sufferGeneral.BaseInfo.Name,
					currentGeneral.BaseInfo.Name,
					t.Name(),
					finalDmg,
					originNum,
					remaindNum,
				)
			}
		}
		return triggerResp
	})
	//战斗前3回合，获得30%概率规避效果
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_Evade,
	)
	hlog.CtxInfof(ctx, "[%s]的规避率提高了30.00%%",
		currentGeneral.BaseInfo.Name,
	)
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade] += 0.3
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		//第四回合
		if params.CurrentRound == consts.Battle_Round_Fourth {
			currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade] -= 0.3
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				currentGeneral.BaseInfo.Name,
				consts.BuffEffectType_Evade,
			)
			hlog.CtxInfof(ctx, "[%s]的规避率降低了30.00%%",
				currentGeneral.BaseInfo.Name,
			)
		}
		return triggerResp
	})

	//战斗前3回合，无法进行普通攻击
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.DebuffEffectType_CanNotGeneralAttack,
	)
	currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_CanNotGeneralAttack] = 1.0
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		//第四回合
		if params.CurrentRound == consts.Battle_Round_Fourth {
			delete(currentGeneral.DeBuffEffectHolderMap, consts.DebuffEffectType_CanNotGeneralAttack)

			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				currentGeneral.BaseInfo.Name,
				consts.DebuffEffectType_CanNotGeneralAttack,
			)
		}
		return triggerResp
	})
}

func (t ThreeDaysOfSeparationTactic) Execute() {
	return
}

func (t ThreeDaysOfSeparationTactic) Trigger() {
	return
}

func (t ThreeDaysOfSeparationTactic) Name() string {
	return "士别三日"
}

func (t ThreeDaysOfSeparationTactic) Id() consts.TacticId {
	return consts.ThreeDaysOfSeparation
}

func (t ThreeDaysOfSeparationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThreeDaysOfSeparationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
