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

// 骑虎难下
// 当除自己之外的友军受到普通攻击时，有35概率对该友军造成当回合禁疗（无法恢复兵力）
// 并对敌军群体（2人）造成兵刃攻击（伤害率72%）
type RideTigerHardToGetOffTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RideTigerHardToGetOffTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RideTigerHardToGetOffTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	//当除自己之外的友军受到普通攻击时，有35概率对该友军造成当回合禁疗（无法恢复兵力）
	//并对敌军群体（2人）造成兵刃攻击（伤害率72%）

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)

	//找到友军2人
	pairGenerals := util.GetPairGeneralsTwoArr(r.tacticsParams)
	//注册效果
	for _, general := range pairGenerals {
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_RideTigerHardToGetOff_Prepare, &vo.EffectHolderParams{
			EffectRate: 1.0,
			FromTactic: r.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferGeneralAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral
				triggerRound := params.CurrentRound

				if !util.BuffEffectContains(triggerGeneral, consts.BuffEffectType_RideTigerHardToGetOff_Prepare) {
					return triggerResp
				}

				if !util.GenerateRate(0.35) {
					hlog.CtxInfof(ctx, "[%s]的「%v」因几率没有生效",
						triggerGeneral.BaseInfo.Name,
						consts.BuffEffectType_RideTigerHardToGetOff_Prepare,
					)
					return triggerResp
				}

				//当回合禁疗
				if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
					EffectRate: 1.0,
					FromTactic: r.Id(),
				}).IsSuccess {
					//注册下回合解除禁疗
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						revokeRound := params.CurrentRound

						if triggerRound+1 == revokeRound {
							util.DebuffEffectWrapRemove(ctx, revokeGeneral, consts.DebuffEffectType_ProhibitionTreatment, r.Id())
						}

						return revokeResp
					})
				}

				//并对敌军群体（2人）造成兵刃攻击（伤害率72%）
				enemyGenerals := util.GetEnemyGeneralsTwoArr(r.tacticsParams)
				for _, enemyGeneral := range enemyGenerals {
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.72)
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams: r.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: enemyGeneral,
						Damage:        dmg,
						DamageType:    consts.DamageType_Weapon,
						TacticName:    r.Name(),
						EffectName:    fmt.Sprintf("%v", consts.BuffEffectType_RideTigerHardToGetOff_Prepare),
					})
				}
				return triggerResp
			})
		}
	}
}

func (r RideTigerHardToGetOffTactic) Id() consts.TacticId {
	return consts.RideTigerHardToGetOff
}

func (r RideTigerHardToGetOffTactic) Name() string {
	return "骑虎难下"
}

func (r RideTigerHardToGetOffTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RideTigerHardToGetOffTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RideTigerHardToGetOffTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RideTigerHardToGetOffTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (r RideTigerHardToGetOffTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RideTigerHardToGetOffTactic) Execute() {
}

func (r RideTigerHardToGetOffTactic) IsTriggerPrepare() bool {
	return false
}
