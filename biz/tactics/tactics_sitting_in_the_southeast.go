package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 坐断东南
// 战斗中，自身及友军单体成功发动普通攻击后，自身有75%几率随机获得连击、洞察、先攻、必中、破阵状态的一种，持续2回合
// 自身为主将时，优先获得不同的状态
// 指挥，100%
type SittingIntheSoutheastTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SittingIntheSoutheastTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SittingIntheSoutheastTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	effectGenerals := make([]*vo.BattleGeneral, 0)

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 战斗中，自身及友军单体成功发动普通攻击后，自身有75%几率随机获得连击、洞察、先攻、必中、破阵状态的一种，持续2回合
	// 自身为主将时，优先获得不同的状态
	pairGeneral := util.GetPairOneGeneralNotSelf(s.tacticsParams, currentGeneral)
	effectGenerals = append(effectGenerals, pairGeneral)
	effectGenerals = append(effectGenerals, currentGeneral)

	for _, effectGeneral := range effectGenerals {
		util.TacticsTriggerWrapRegister(effectGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			buffs := []consts.BuffEffectType{
				consts.BuffEffectType_ContinuousAttack,
				consts.BuffEffectType_Insight,
				consts.BuffEffectType_FirstAttack,
				consts.BuffEffectType_MustHit,
				consts.BuffEffectType_BreakFormation,
			}
			if util.GenerateRate(0.75) {
				if currentGeneral.IsMaster {
					notContainBuffs := make([]consts.BuffEffectType, 0)
					for _, buff := range buffs {
						if !util.BuffEffectContains(currentGeneral, buff) {
							notContainBuffs = append(notContainBuffs, buff)
						}
					}
					if len(notContainBuffs) > 0 {
						notHitIdx := util.GenerateHitOneIdx(len(notContainBuffs))
						util.BuffEffectWrapSet(ctx, currentGeneral, notContainBuffs[notHitIdx], &vo.EffectHolderParams{
							EffectRound:    2,
							FromTactic:     s.Id(),
							ProduceGeneral: triggerGeneral,
						})
					}
				} else {
					hitIdx := util.GenerateHitOneIdx(len(buffs))
					util.BuffEffectWrapSet(ctx, currentGeneral, buffs[hitIdx], &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     s.Id(),
						ProduceGeneral: triggerGeneral,
					})
				}
			}

			return triggerResp
		})
	}
}

func (s SittingIntheSoutheastTactic) Id() consts.TacticId {
	return consts.SittingIntheSoutheast
}

func (s SittingIntheSoutheastTactic) Name() string {
	return "坐断东南"
}

func (s SittingIntheSoutheastTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SittingIntheSoutheastTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SittingIntheSoutheastTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SittingIntheSoutheastTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SittingIntheSoutheastTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SittingIntheSoutheastTactic) Execute() {

}

func (s SittingIntheSoutheastTactic) IsTriggerPrepare() bool {
	return false
}
