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

// 连环计
// 发动概率35%
// 准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，
// 并发动谋略攻击（伤害率156%，受智力影响）
type InterlockedStratagemsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
}

func (i InterlockedStratagemsTactic) IsTriggerPrepare() bool {
	return i.isTriggerPrepare
}

func (i InterlockedStratagemsTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i InterlockedStratagemsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i InterlockedStratagemsTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i InterlockedStratagemsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.35
	return i
}

func (i InterlockedStratagemsTactic) Prepare() {
	return
}

func (i InterlockedStratagemsTactic) Id() consts.TacticId {
	return consts.InterlockedStratagems
}

func (i InterlockedStratagemsTactic) Name() string {
	return "连环计"
}

func (i InterlockedStratagemsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InterlockedStratagemsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InterlockedStratagemsTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral
	currentRound := i.tacticsParams.CurrentRound

	i.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)

	//准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，
	//获取敌军全体
	allEnemyGenerals := util.GetEnemyGeneralMap(i.tacticsParams)
	//注册发动效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		//准备一回合
		if !(currentRound+1 == params.CurrentRound) {
			return triggerResp
		}

		i.isTriggerPrepare = false
		hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
			currentGeneral.BaseInfo.Name,
			i.Name(),
		)
		for _, sufferGeneral := range allEnemyGenerals {
			//施加效果
			if !util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_InterlockedStratagems, &vo.EffectHolderParams{
				EffectRate: 1.0,
				FromTactic: i.Id(),
			}).IsSuccess {
				continue
			}
			//注册铁锁连环效果
			registerFunc := func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				lockGeneral := params.CurrentGeneral

				if !util.DeBuffEffectContains(lockGeneral, consts.DebuffEffectType_InterlockedStratagems) {
					return triggerResp
				}
				//释放效果
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					lockGeneral.BaseInfo.Name,
					i.Name(),
					consts.DebuffEffectType_InterlockedStratagems,
				)
				//找到2个队友进行伤害反弹
				pairGenerals := util.GetPairGeneralsTwoArrByGeneral(lockGeneral, i.tacticsParams)
				for _, reboundGeneral := range pairGenerals {
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.15)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams:          i.tacticsParams,
						AttackGeneral:          currentGeneral,
						SufferGeneral:          reboundGeneral,
						Damage:                 dmg,
						DamageType:             consts.DamageType_Strategy,
						TacticName:             i.Name(),
						EffectName:             fmt.Sprintf("%v", consts.DebuffEffectType_InterlockedStratagems),
						IsBanInterLockedEffect: true,
					})
				}
				return triggerResp
			}

			//被战法攻击
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferTroopsTactic, registerFunc)
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferArmTactic, registerFunc)
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferCommandTactic, registerFunc)
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferAssaultTactic, registerFunc)
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferPassiveTactic, registerFunc)
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferActiveTactic, registerFunc)
			//被普攻
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_SufferGeneralAttack, registerFunc)

			//效果消失注册
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeRound := params.CurrentRound
				//持续2回合
				if currentRound+3 == revokeRound {
					util.DebuffEffectWrapRemove(ctx, sufferGeneral, consts.DebuffEffectType_InterlockedStratagems, i.Id())
				}
				return triggerResp
			})
		}
		//并发动谋略攻击（伤害率156%，受智力影响）
		for _, sufferGeneral := range allEnemyGenerals {
			dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.56)
			util.TacticDamage(&util.TacticDamageParam{
				TacticsParams: i.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: sufferGeneral,
				Damage:        dmgNum,
				DamageType:    consts.DamageType_Strategy,
				TacticName:    i.Name(),
			})
		}
		return triggerResp
	})
}
