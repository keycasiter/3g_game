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

// 破阵摧坚
// 准备1回合，使敌军群体（2人）统率、智力降低80点（受武力影响），持续2回合，并对其发动一次兵刃攻击（伤害率158%）
type BreakingThroughTheFormationAndDestroyingTheFirmTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.35
	return b
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Prepare() {
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheFormationAndDestroyingTheFirm
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Name() string {
	return "破阵摧坚"
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//准备1回合，使敌军群体（2人）统率、智力降低80点（受武力影响），持续2回合，并对其发动一次兵刃攻击（伤害率158%）
	b.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		//准备回合释放
		if currentRound+2 == triggerRound {
			b.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if b.isTriggered {
				return triggerResp
			} else {
				b.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)
			//找到敌军2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, b.tacticsParams)
			for _, general := range enemyGenerals {
				//降低属性
				v := 80 + triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100
				general.BaseInfo.AbilityAttr.CommandBase -= v
				general.BaseInfo.AbilityAttr.IntelligenceBase -= v
				hlog.CtxInfof(ctx, "[%]的统率降低了%.2f",
					general.BaseInfo.Name, v)

				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeRound := params.CurrentRound

					if currentRound+2 == revokeRound {
						general.BaseInfo.AbilityAttr.CommandBase -= v
						general.BaseInfo.AbilityAttr.IntelligenceBase -= v
						hlog.CtxInfof(ctx, "[%]的统率提高了%.2f",
							general.BaseInfo.Name, v)
					}

					return revokeResp
				})

				//攻击
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.58)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: b.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: general,
					Damage:        dmg,
					DamageType:    consts.DamageType_Weapon,
					TacticName:    b.Name(),
				})
			}
		}
		return triggerResp
	})
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
