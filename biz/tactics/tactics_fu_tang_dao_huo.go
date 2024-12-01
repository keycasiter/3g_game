package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 赴汤蹈火
// 战斗第2-6回合，每回合有50%概率（受统率影响）使我军群体（2-3人）获得1次抵御，
// 并使其进入严密状态（严密：持续期间抵御次数可以叠加，且抵御有效时长延长8回合），持续2回合
type FuTangDaoHuoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a FuTangDaoHuoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.0
	return a
}

func (a FuTangDaoHuoTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	// 战斗第2-6回合，每回合有50%概率（受统率影响）使我军群体（2-3人）获得1次抵御，
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		currentRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if currentRound < consts.Battle_Round_Second {
			return triggerResp
		}

		triggerRate := 0.5 + triggerGeneral.BaseInfo.AbilityAttr.CommandBase/100/100
		if !util.GenerateRate(triggerRate) {
			return triggerResp
		}

		pairGenerals := util.GetPairGeneralsTwoOrThreeMap(triggerGeneral, a.tacticsParams)
		for _, general := range pairGenerals {
			util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
				EffectTimes: 1,
				FromTactic:  a.Id(),
			})

			// 并使其进入严密状态（严密：持续期间抵御次数可以叠加，且抵御有效时长延长8回合），持续2回合
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_YanMi, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     a.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    general,
					EffectType: consts.BuffEffectType_YanMi,
					TacticId:   a.Id(),
				})
			}
		}

		return triggerResp
	})

}

func (a FuTangDaoHuoTactic) Id() consts.TacticId {
	return consts.FuTangDaoHuo
}

func (a FuTangDaoHuoTactic) Name() string {
	return "赴汤蹈火"
}

func (a FuTangDaoHuoTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a FuTangDaoHuoTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a FuTangDaoHuoTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a FuTangDaoHuoTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a FuTangDaoHuoTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Apparatus,
	}
}

func (a FuTangDaoHuoTactic) Execute() {
}

func (a FuTangDaoHuoTactic) IsTriggerPrepare() bool {
	return false
}

func (a FuTangDaoHuoTactic) SetTriggerPrepare(triggerPrepare bool) {
}
