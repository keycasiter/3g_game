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

// 临战先登
// 对敌军群体（2人）造成兵刃攻击（伤害率150%），之后自己进入虚弱状态，持续1回合
// 主动，100%
type ToAscendBeforeBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToAscendBeforeBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ToAscendBeforeBattleTactic) Prepare() {

}

func (t ToAscendBeforeBattleTactic) Id() consts.TacticId {
	return consts.ToAscendBeforeBattle
}

func (t ToAscendBeforeBattleTactic) Name() string {
	return "临战先登"
}

func (t ToAscendBeforeBattleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t ToAscendBeforeBattleTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToAscendBeforeBattleTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToAscendBeforeBattleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ToAscendBeforeBattleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToAscendBeforeBattleTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//对敌军群体（2人）造成兵刃攻击（伤害率150%），之后自己进入虚弱状态，持续1回合
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, t.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.5)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: t.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      t.Id(),
			TacticName:    t.Name(),
		})
	}
	//虚弱效果
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_PoorHealth,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t ToAscendBeforeBattleTactic) IsTriggerPrepare() bool {
	return false
}
