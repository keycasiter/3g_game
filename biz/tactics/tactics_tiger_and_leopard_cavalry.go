package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 虎豹骑
// 将骑兵进阶为天下骁锐的虎豹骑：
// 我军全体提高40武力，战斗前3回合，我军全体突击战法发动率提高10%，
// 若曹纯统领时，提升的发动概率额外受武力影响
type TigerAndLeopardCavalryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerAndLeopardCavalryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerAndLeopardCavalryTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 将骑兵进阶为天下骁锐的虎豹骑：
	// 我军全体提高40武力，战斗前3回合，我军全体突击战法发动率提高10%，
	// 若曹纯统领时，提升的发动概率额外受武力影响
	pairGenerals := util.GetPairGeneralArr(currentGeneral, t.tacticsParams)
	//武力
	for _, pairGeneral := range pairGenerals {
		//我军全体提高40武力
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    40,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		})
		//战斗前3回合，我军全体突击战法发动率提高10%
		triggerRate := 0.1
		//若曹纯统领时，提升的发动概率额外受武力影响
		if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.CaoChun {
			triggerRate += currentGeneral.BaseInfo.AbilityAttr.ForceBase / 100 / 100
		}
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsAssaultTriggerImprove, &vo.EffectHolderParams{
			TriggerRate:    triggerRate,
			EffectRound:    3,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsAssaultTriggerImprove,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (t TigerAndLeopardCavalryTactic) Id() consts.TacticId {
	return consts.TigerAndLeopardCavalry
}

func (t TigerAndLeopardCavalryTactic) Name() string {
	return "虎豹骑"
}

func (t TigerAndLeopardCavalryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TigerAndLeopardCavalryTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerAndLeopardCavalryTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerAndLeopardCavalryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (t TigerAndLeopardCavalryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (t TigerAndLeopardCavalryTactic) Execute() {
}

func (t TigerAndLeopardCavalryTactic) IsTriggerPrepare() bool {
	return false
}
