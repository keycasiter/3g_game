package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 锦帆军
// 将弓兵进阶为气盖千夫的锦帆军：
// 部队普通攻击时，有45%概率使目标进入溃逃状态（伤害率64%，受武力影响），持续2回合
// 若目标已经溃逃则造成兵刃攻击（伤害率110%）并恢复伤害量的30%兵力；
// 若甘宁统领，提高友军6%会心
type JinFanArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JinFanArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JinFanArmyTactic) Prepare() {
	ctx := j.tacticsParams.Ctx
	currentGeneral := j.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		j.Name(),
	)

	//将弓兵进阶为气盖千夫的锦帆军：
	//部队普通攻击时，有45%概率使目标进入溃逃状态（伤害率64%，受武力影响），持续2回合
	//若目标已经溃逃则造成兵刃攻击（伤害率110%）并恢复伤害量的30%兵力；
	//若甘宁统领，提高友军6%会心
	pairGenerals := util.GetPairGeneralArr(j.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.GenerateRate(0.45) {
				sufferGeneral := j.tacticsParams.CurrentSufferGeneral
				//若目标已经溃逃则造成兵刃攻击（伤害率110%）并恢复伤害量的30%兵力；
				if util.DeBuffEffectContains(sufferGeneral, consts.DebuffEffectType_Escape) {
					weaponDmg := cast.ToInt64(pairGeneral.BaseInfo.AbilityAttr.ForceBase * 1.1)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: j.tacticsParams,
						AttackGeneral: triggerGeneral,
						SufferGeneral: sufferGeneral,
						DamageType:    consts.DamageType_Weapon,
						Damage:        weaponDmg,
						TacticName:    j.Name(),
					})
					//恢复
					resumeNum := cast.ToInt64(cast.ToFloat64(weaponDmg) * 0.3)
					util.ResumeSoldierNum(ctx, triggerGeneral, resumeNum)

				} else {
					//部队普通攻击时，有45%概率使目标进入溃逃状态（伤害率64%，受武力影响），持续2回合
					if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Escape, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     j.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						//注册消失效果
						util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_Escape,
								TacticId:   j.Id(),
							}) {
								dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 0.64)
								util.TacticDamage(&util.TacticDamageParam{
									TacticsParams: j.tacticsParams,
									AttackGeneral: triggerGeneral,
									SufferGeneral: revokeGeneral,
									DamageType:    consts.DamageType_Weapon,
									Damage:        dmg,
									TacticName:    j.Name(),
									EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Escape),
								})
							}

							return revokeResp
						})
					}
				}
			}

			return triggerResp
		})
		//若甘宁统领，提高友军6%会心
		if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.GanNing {
			util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
				EffectRate:     0.06,
				FromTactic:     j.Id(),
				ProduceGeneral: currentGeneral,
			})
		}
	}
}

func (j JinFanArmyTactic) Id() consts.TacticId {
	return consts.JinFanArmy
}

func (j JinFanArmyTactic) Name() string {
	return "锦帆军"
}

func (j JinFanArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (j JinFanArmyTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JinFanArmyTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JinFanArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (j JinFanArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (j JinFanArmyTactic) Execute() {
}

func (j JinFanArmyTactic) IsTriggerPrepare() bool {
	return false
}
