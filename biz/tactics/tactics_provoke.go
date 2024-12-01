package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 挑衅
// 嘲讽敌军全体使其攻击自己，持续1回合
// 主动，60%
type ProvokeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p ProvokeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.6
	return p
}

func (p ProvokeTactic) Prepare() {
}

func (p ProvokeTactic) Id() consts.TacticId {
	return consts.Provoke
}

func (p ProvokeTactic) Name() string {
	return "挑衅"
}

func (p ProvokeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (p ProvokeTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p ProvokeTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p ProvokeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p ProvokeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p ProvokeTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 嘲讽敌军全体使其攻击自己，持续1回合
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, p.tacticsParams)
	for _, general := range enemyGenerals {
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     p.Id(),
			TauntByTarget:  currentGeneral,
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Taunt,
					TacticId:   p.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (p ProvokeTactic) IsTriggerPrepare() bool {
	return false
}

func (a ProvokeTactic) SetTriggerPrepare(triggerPrepare bool) {
}
