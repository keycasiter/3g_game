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
)

// 神上使
// 主动 50%
// 准备1回合，对敌军全体造成溃逃状态，每回合持续造成伤害（伤害率68%，受武力影响），持续2回合
type DivineEnvoyTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (d DivineEnvoyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.5
	return d
}

func (d DivineEnvoyTactic) Prepare() {
}

func (d DivineEnvoyTactic) Id() consts.TacticId {
	return consts.DivineEnvoy
}

func (d DivineEnvoyTactic) Name() string {
	return "神上使"
}

func (d DivineEnvoyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DivineEnvoyTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DivineEnvoyTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DivineEnvoyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DivineEnvoyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DivineEnvoyTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral
	currentRound := d.tacticsParams.CurrentRound

	//准备1回合，对敌军全体造成溃逃状态，每回合持续造成伤害（伤害率68%，受武力影响），持续2回合
	d.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			d.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if d.isTriggered {
				return triggerResp
			} else {
				d.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				d.Name(),
			)
			//对敌军全体造成溃逃状态，每回合持续造成伤害（伤害率68%，受武力影响），持续2回合
			enemyGenerals := util.GetEnemyGeneralArr(d.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//溃逃状态
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Escape, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     d.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Escape,
							TacticId:   d.Id(),
						}) {
							dmgRate := triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100/100 + 0.68
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     d.tacticsParams,
								AttackGeneral:     triggerGeneral,
								SufferGeneral:     enemyGeneral,
								DamageType:        consts.DamageType_Weapon,
								DamageImproveRate: dmgRate,
								TacticName:        d.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Escape),
							})
						}

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (d DivineEnvoyTactic) IsTriggerPrepare() bool {
	return d.isTriggerPrepare
}
