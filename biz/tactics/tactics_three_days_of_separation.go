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
	tacticsParams model.TacticsParams
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t ThreeDaysOfSeparationTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	currentRound := t.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_ThreeDaysOfSeparation_Prepare,
	)
	util.TacticsTriggerWrapSet(currentGeneral,
		consts.BattleAction_Attack,
		func(params vo.TacticsTriggerParams) {
			//第四回合
			if currentRound == consts.Battle_Round_Fourth {
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					currentGeneral.BaseInfo.Name,
					t.Name(),
					consts.BuffEffectType_ThreeDaysOfSeparation_Prepare,
				)
				currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase += 68
				hlog.CtxInfof(ctx, "[%s]的智力提高了68")

				//并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
				//找到敌军全体
				enemyGenerals := util.GetEnemyGeneralArr(t.tacticsParams)
				for _, sufferGeneral := range enemyGenerals {
					//TODO 受智力影响
					dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.8)
					finalDmg, originNum, remaindNum := util.TacticDamage(ctx, currentGeneral, sufferGeneral, dmgNum)
					hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d️️️)",
						sufferGeneral.BaseInfo.Name,
						currentGeneral.BaseInfo.Name,
						t.Name(),
						finalDmg,
						originNum,
						remaindNum,
					)
				}
			}
		})
	//战斗前3回合，获得30%概率规避效果
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_Evade,
	)
	hlog.CtxInfof(ctx, "[%s]的规避率提高了30.00%",
		currentGeneral.BaseInfo.Name,
	)
	currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade] += 0.3
	util.TacticsTriggerWrapSet(currentGeneral, consts.BattleAction_Attack, func(params vo.TacticsTriggerParams) {
		//第四回合
		if currentRound == consts.Battle_Round_Fourth {
			currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade] -= 0.3
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				currentGeneral.BaseInfo.Name,
				consts.BuffEffectType_Evade,
			)
			hlog.CtxInfof(ctx, "[%s]的规避率降低了30.00%")
		}
	})

	//战斗前3回合，无法进行普通攻击
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.DebuffEffectType_CanNotGeneralAttack,
	)
	currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_CanNotGeneralAttack] = 1.0
	util.TacticsTriggerWrapSet(currentGeneral, consts.BattleAction_Attack, func(params vo.TacticsTriggerParams) {
		//第四回合
		if currentRound == consts.Battle_Round_Fourth {
			delete(currentGeneral.DeBuffEffectHolderMap, consts.DebuffEffectType_CanNotGeneralAttack)

			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				currentGeneral.BaseInfo.Name,
				consts.DebuffEffectType_CanNotGeneralAttack,
			)
		}
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
