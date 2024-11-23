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

// 虎侯
// 战斗中，友军群体发动或受到普通攻击时，自身有45%概率使我军全体统率提升15点，可叠加5次，持续2回合，并对敌军单体造成兵刃伤害（伤害率66%）
type HuHouTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a HuHouTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a HuHouTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//战斗中，友军群体发动或受到普通攻击时，自身有45%概率使我军全体统率提升15点，可叠加5次，持续2回合，并对敌军单体造成兵刃伤害（伤害率66%）
	pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		triggerFunc := func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.GenerateRate(0.45) {
				for _, general := range pairGenerals {
					if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
						EffectValue:    15,
						EffectRound:    2,
						EffectTimes:    1,
						MaxEffectTimes: 5,
					}).IsSuccess {
						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    general,
							EffectType: consts.BuffEffectType_IncrCommand,
							TacticId:   a.Id(),
						})
					}
				}
				//并对敌军单体造成兵刃伤害（伤害率66%）
				enemeyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, a.tacticsParams)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     a.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     enemeyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 0.66,
					TacticId:          a.Id(),
					TacticName:        a.Name(),
				})
			}

			return triggerResp
		}

		//发动普通攻击
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, triggerFunc)
		//受到普通攻击
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferGeneralAttackEnd, triggerFunc)
	}
}

func (a HuHouTactic) Id() consts.TacticId {
	return consts.HuHou
}

func (a HuHouTactic) Name() string {
	return "虎侯"
}

func (a HuHouTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a HuHouTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a HuHouTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a HuHouTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a HuHouTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a HuHouTactic) Execute() {
}

func (a HuHouTactic) IsTriggerPrepare() bool {
	return false
}
