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
)

// 矢志不移
// 前2回合获得群攻（伤害率100%）状态，但只有50%概率发动群攻（普通攻击时对目标同部队其他武将造成伤害）效果，
// 第3回合起，每回合提高15武力
type OnesResolveIsUnshakenTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OnesResolveIsUnshakenTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	o.tacticsParams = tacticsParams
	o.triggerRate = 1.0
	return o
}

func (o OnesResolveIsUnshakenTactic) Prepare() {
	ctx := o.tacticsParams.Ctx
	currentGeneral := o.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		o.Name(),
	)

	// 前2回合获得群攻（伤害率100%）状态，但只有50%概率发动群攻（普通攻击时对目标同部队其他武将造成伤害）效果，
	//施加效果
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
		EffectRate: 1.0,
		FromTactic: o.Id(),
	}).IsSuccess {
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
			currentGeneral.BaseInfo.Name,
			consts.BuffEffectType_OnesResolveIsUnshaken_Prepare,
		)
		//注册效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRound := params.CurrentRound
			sufferGeneral := o.tacticsParams.CurrentSufferGeneral

			if triggerRound > consts.Battle_Round_Second {
				return triggerResp
			}

			if util.GenerateRate(0.5) {
				//找到被攻击者队友
				enemyGenerals := util.GetPairGeneralsNotSelf(o.tacticsParams, sufferGeneral)
				for _, enemyGeneral := range enemyGenerals {
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
						triggerGeneral.BaseInfo.Name,
						o.Name(),
						consts.BuffEffectType_GroupAttack,
					)
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     o.tacticsParams,
						AttackGeneral:     triggerGeneral,
						SufferGeneral:     enemyGeneral,
						DamageType:        consts.DamageType_Weapon,
						DamageImproveRate: 1.0,
						TacticId:          o.Id(),
						TacticName:        o.Name(),
						EffectName:        fmt.Sprintf("%v", consts.BuffEffectType_GroupAttack),
					})

				}
			} else {
				hlog.CtxInfof(ctx, "[%s]来自[%s]【%s】的「%v」效果因几率没有生效",
					triggerGeneral.BaseInfo.Name,
					triggerGeneral.BaseInfo.Name,
					o.Name(),
					consts.BuffEffectType_GroupAttack,
				)
			}

			return triggerResp
		})
	}

	// 第3回合起，每回合提高15武力
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound >= consts.Battle_Round_Third {
			hlog.CtxInfof(ctx, "[%s]的武力提高了15点",
				triggerGeneral.BaseInfo.Name,
			)
			triggerGeneral.BaseInfo.AbilityAttr.ForceBase += 15
		}

		return triggerResp
	})
}

func (o OnesResolveIsUnshakenTactic) Id() consts.TacticId {
	return consts.OnesResolveIsUnshaken
}

func (o OnesResolveIsUnshakenTactic) Name() string {
	return "矢志不移"
}

func (o OnesResolveIsUnshakenTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (o OnesResolveIsUnshakenTactic) GetTriggerRate() float64 {
	return o.triggerRate
}

func (o OnesResolveIsUnshakenTactic) SetTriggerRate(rate float64) {
	o.triggerRate = rate
}

func (o OnesResolveIsUnshakenTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (o OnesResolveIsUnshakenTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (o OnesResolveIsUnshakenTactic) Execute() {
}

func (o OnesResolveIsUnshakenTactic) IsTriggerPrepare() bool {
	return false
}
