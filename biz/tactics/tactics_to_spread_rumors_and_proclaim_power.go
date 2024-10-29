package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 传檄宣威
// 对随机敌军单体造成谋略伤攻击（伤害率165%，受智力影响），并缴械（无法进行普通攻击）2回合
type ToSpreadRumorsAndProclaimPowerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToSpreadRumorsAndProclaimPowerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t ToSpreadRumorsAndProclaimPowerTactic) Prepare() {
}

func (t ToSpreadRumorsAndProclaimPowerTactic) Id() consts.TacticId {
	return consts.ToSpreadRumorsAndProclaimPower
}

func (t ToSpreadRumorsAndProclaimPowerTactic) Name() string {
	return "传檄宣威"
}

func (t ToSpreadRumorsAndProclaimPowerTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ToSpreadRumorsAndProclaimPowerTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToSpreadRumorsAndProclaimPowerTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToSpreadRumorsAndProclaimPowerTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ToSpreadRumorsAndProclaimPowerTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToSpreadRumorsAndProclaimPowerTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//对随机敌军单体造成谋略伤攻击（伤害率165%，受智力影响），并缴械（无法进行普通攻击）2回合
	//伤害
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
	dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.65
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     t.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     enemyGeneral,
		DamageType:        consts.DamageType_Strategy,
		DamageImproveRate: dmgRate,
		TacticId:          t.Id(),
		TacticName:        t.Name(),
	})
	//施加效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_CancelWeapon,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t ToSpreadRumorsAndProclaimPowerTactic) IsTriggerPrepare() bool {
	return false
}
