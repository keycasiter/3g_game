package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 一举歼灭
// 准备1回合，对敌军单体造成一次兵刃攻击（伤害率335%），若目标处于沙暴状态，则额外使目标进入混乱（攻击和战法无差别选择目标）状态，持续2回合
type AnnihilateInOneFellSwoopTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (a AnnihilateInOneFellSwoopTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.4
	return a
}

func (a AnnihilateInOneFellSwoopTactic) Prepare() {

}

func (a AnnihilateInOneFellSwoopTactic) Id() consts.TacticId {
	return consts.AnnihilateInOneFellSwoop
}

func (a AnnihilateInOneFellSwoopTactic) Name() string {
	return "一举歼灭"
}

func (a AnnihilateInOneFellSwoopTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a AnnihilateInOneFellSwoopTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AnnihilateInOneFellSwoopTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AnnihilateInOneFellSwoopTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AnnihilateInOneFellSwoopTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AnnihilateInOneFellSwoopTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound

	//准备1回合，对敌军单体造成一次兵刃攻击（伤害率335%），若目标处于沙暴状态，则额外使目标进入混乱（攻击和战法无差别选择目标）状态，持续2回合
	a.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			a.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if a.isTriggered {
				return triggerResp
			} else {
				a.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				a.Name(),
			)

			//找到敌军1人
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, a.tacticsParams)
			//对敌军单体造成一次兵刃攻击（伤害率335%）
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 3.35)
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams: a.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: enemyGeneral,
				DamageType:    consts.DamageType_Weapon,
				Damage:        dmg,
				TacticName:    a.Name(),
				TacticId:      a.Id(),
			})
			//若目标处于沙暴状态，则额外使目标进入混乱（攻击和战法无差别选择目标）状态，持续2回合
			if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Sandstorm) {
				util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  a.Id(),
				})
				//注册消失效果
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeGeneral := params.CurrentGeneral
					revokeResp := &vo.TacticsTriggerResult{}

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Chaos,
						TacticId:   a.Id(),
					})
					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (a AnnihilateInOneFellSwoopTactic) IsTriggerPrepare() bool {
	return a.isTriggerPrepare
}
