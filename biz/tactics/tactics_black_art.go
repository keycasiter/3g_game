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
	"github.com/spf13/cast"
)

// 妖术
// 准备1回合，对敌军全体施加沙暴状态，每回合持续造成伤害（伤害率72%，受智力影响）并使自己获得1次抵御，可以免疫伤害，持续2回合
type BlackArtTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (b BlackArtTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BlackArtTactic) Prepare() {

}

func (b BlackArtTactic) Id() consts.TacticId {
	return consts.BlackArt
}

func (b BlackArtTactic) Name() string {
	return "妖术"
}

func (b BlackArtTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BlackArtTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BlackArtTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BlackArtTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BlackArtTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BlackArtTactic) Execute() {

	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	//准备1回合，对敌军全体施加沙暴状态，每回合持续造成伤害（伤害率72%，受智力影响）并使自己获得1次抵御，可以免疫伤害，持续2回合
	b.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

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

			//对敌军全体施加沙暴状态，每回合持续造成伤害（伤害率72%，受智力影响）并使自己获得1次抵御，可以免疫伤害，持续2回合
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, b.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				//沙暴状态
				if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Sandstorm, &vo.EffectHolderParams{
					EffectRate:  1.0,
					EffectRound: 2,
					FromTactic:  b.Id(),
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						triggerDmgResp := &vo.TacticsTriggerResult{}
						triggerDmgGeneral := params.CurrentGeneral

						//沙暴伤害
						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    triggerDmgGeneral,
							EffectType: consts.DebuffEffectType_Sandstorm,
							TacticId:   0,
						}) {
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     b.tacticsParams,
								AttackGeneral:     currentGeneral,
								SufferGeneral:     triggerGeneral,
								DamageType:        consts.DamageType_Strategy,
								DamageImproveRate: 0.72,
								TacticId:          b.Id(),
								TacticName:        b.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Sandstorm),
							})
						}

						return triggerDmgResp
					})
				}
			}
			//发动者自己获得1次抵御
			util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
				EffectRate:     1.0,
				EffectTimes:    1,
				MaxEffectTimes: cast.ToInt64(consts.INT_MAX),
				FromTactic:     b.Id(),
				ProduceGeneral: currentGeneral,
			})
		}

		return triggerResp
	})
}

func (b BlackArtTactic) IsTriggerPrepare() bool {
	return false
}

func (a BlackArtTactic) SetTriggerPrepare(triggerPrepare bool) {
}
