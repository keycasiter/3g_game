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
	triggerRate   float64
}

func (t ThreeDaysOfSeparationTactic) IsTriggerPrepare() bool {
	return false
}

func (t ThreeDaysOfSeparationTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThreeDaysOfSeparationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThreeDaysOfSeparationTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
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
				dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.8)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: sufferGeneral,
					Damage:        dmgNum,
					DamageType:    consts.DamageType_Strategy,
					TacticName:    t.Name(),
				})
			}
		}
		return triggerResp
	})
	//战斗前3回合，获得30%概率规避效果
	hlog.CtxInfof(ctx, "[%s]的规避率提高了30.00%%",
		currentGeneral.BaseInfo.Name,
	)
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Evade, &vo.EffectHolderParams{
		EffectRate: 0.3,
		FromTactic: t.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			//第四回合
			if params.CurrentRound == consts.Battle_Round_Fourth {
				util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_Evade, t.Id())
				hlog.CtxInfof(ctx, "[%s]的规避率降低了30.00%%",
					currentGeneral.BaseInfo.Name,
				)
			}
			return triggerResp
		})

		//战斗前3回合，无法进行普通攻击
		if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_CanNotGeneralAttack, &vo.EffectHolderParams{
			FromTactic: t.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral
				//第四回合
				if params.CurrentRound == consts.Battle_Round_Fourth {
					util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_CanNotGeneralAttack, t.Id())
				}
				return triggerResp
			})
		}
	}
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
