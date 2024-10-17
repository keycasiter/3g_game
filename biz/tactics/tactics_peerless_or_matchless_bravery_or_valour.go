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

// 勇冠三军
// 普通攻击之后，对攻击目标再次发起猛攻（伤害率180%），并使其进入震慑状态（无法行动），持续1回合
// 突击 30%
type PeerlessOrMatchlessBraveryOrValourTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.3
	return p
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Prepare() {
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Id() consts.TacticId {
	return consts.PeerlessOrMatchlessBraveryOrValour
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Name() string {
	return "勇冠三军"
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 普通攻击之后，对攻击目标再次发起猛攻（伤害率180%），并使其进入震慑状态（无法行动），持续1回合
	sufferGeneral := p.tacticsParams.CurrentSufferGeneral
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.8)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: p.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: sufferGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    p.Name(),
		TacticId:      p.Id(),
	})
	if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Awe,
				TacticId:   p.Id(),
			})

			return revokeResp
		})
	}
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) IsTriggerPrepare() bool {
	return false
}
