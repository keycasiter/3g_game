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

// 抬棺决战
// 准备1回合，移除敌军群体（2人）的增益效果，随后造成兵刃攻击（伤害率255%）
// 主动，30%
type TheBattleOfCarryingCoffinTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t TheBattleOfCarryingCoffinTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.3
	return t
}

func (t TheBattleOfCarryingCoffinTactic) Prepare() {
}

func (t TheBattleOfCarryingCoffinTactic) Id() consts.TacticId {
	return consts.TheBattleOfCarryingCoffin
}

func (t TheBattleOfCarryingCoffinTactic) Name() string {
	return "抬棺决战"
}

func (t TheBattleOfCarryingCoffinTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TheBattleOfCarryingCoffinTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheBattleOfCarryingCoffinTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheBattleOfCarryingCoffinTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheBattleOfCarryingCoffinTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheBattleOfCarryingCoffinTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	currentRound := t.tacticsParams.CurrentRound

	t.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			t.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if t.isTriggered {
				return triggerResp
			} else {
				t.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				triggerGeneral.BaseInfo.Name,
				t.Name(),
			)
			// 准备1回合，移除敌军群体（2人）的增益效果
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, t.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				util.BuffEffectClean(ctx, enemyGeneral)
				//随后造成兵刃攻击（伤害率255%）
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 2.55)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      t.Id(),
					TacticName:    t.Name(),
				})
			}
		}
		return triggerResp
	})
}

func (t TheBattleOfCarryingCoffinTactic) IsTriggerPrepare() bool {
	return false
}
