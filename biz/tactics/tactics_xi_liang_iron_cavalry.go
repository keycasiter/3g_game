package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 西凉铁骑
// 将骑兵进阶为横行天下的西凉铁骑铁骑：
// 战斗前3回合，提高我军全体25%会心几率（触发时兵刃伤害提高100%）；
// 若马腾统领，则提高会心几率受速度影响
type XiLiangIronCavalryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (x XiLiangIronCavalryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	x.tacticsParams = tacticsParams
	x.triggerRate = 1.0
	return x
}

func (x XiLiangIronCavalryTactic) Prepare() {
	ctx := x.tacticsParams.Ctx
	currentGeneral := x.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		x.Name(),
	)

	//将骑兵进阶为横行天下的西凉铁骑铁骑：
	//战斗前3回合，提高我军全体25%会心几率（触发时兵刃伤害提高100%）；
	//若马腾统领，则提高会心几率受速度影响
	pairGenerals := util.GetPairGeneralArr(x.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		effectRate := 0.25
		if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.MaTeng {
			effectRate += currentGeneral.BaseInfo.AbilityAttr.SpeedBase / 100 / 100
		}

		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectRound:    3,
			FromTactic:     x.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_EnhanceWeapon,
					TacticId:   x.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (x XiLiangIronCavalryTactic) Id() consts.TacticId {
	return consts.XiLiangIronCavalry
}

func (x XiLiangIronCavalryTactic) Name() string {
	return "西凉铁骑"
}

func (x XiLiangIronCavalryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (x XiLiangIronCavalryTactic) GetTriggerRate() float64 {
	return x.triggerRate
}

func (x XiLiangIronCavalryTactic) SetTriggerRate(rate float64) {
	x.triggerRate = rate
}

func (x XiLiangIronCavalryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (x XiLiangIronCavalryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (x XiLiangIronCavalryTactic) Execute() {
}

func (x XiLiangIronCavalryTactic) IsTriggerPrepare() bool {
	return false
}
