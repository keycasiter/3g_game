package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 锦帆百翎
// 战斗中，提高自身50%会心几率及20%会心伤害；
// 自身为主将时，提高友军群体（2人）5%会心几率及10%会心伤害
// 被动，100%
type JinfanArmyHundredFeathersTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JinfanArmyHundredFeathersTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JinfanArmyHundredFeathersTactic) Prepare() {
	ctx := j.tacticsParams.Ctx
	currentGeneral := j.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		j.Name(),
	)
	// 战斗中，提高自身50%会心几率及20%会心伤害；
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
		TriggerRate:    0.5,
		EffectRate:     0.2,
		FromTactic:     j.Id(),
		ProduceGeneral: currentGeneral,
	})
	// 自身为主将时，提高友军群体（2人）5%会心几率及10%会心伤害
	if currentGeneral.IsMaster {
		pairGenerals := util.GetPairGeneralsTwoArr(currentGeneral, j.tacticsParams)
		for _, pairGeneral := range pairGenerals {
			util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
				TriggerRate:    0.05,
				EffectRate:     0.1,
				FromTactic:     j.Id(),
				ProduceGeneral: currentGeneral,
			})
		}
	}
}

func (j JinfanArmyHundredFeathersTactic) Id() consts.TacticId {
	return consts.JinfanArmyHundredFeathers
}

func (j JinfanArmyHundredFeathersTactic) Name() string {
	return "锦帆百翎"
}

func (j JinfanArmyHundredFeathersTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (j JinfanArmyHundredFeathersTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JinfanArmyHundredFeathersTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JinfanArmyHundredFeathersTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (j JinfanArmyHundredFeathersTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (j JinfanArmyHundredFeathersTactic) Execute() {

}

func (j JinfanArmyHundredFeathersTactic) IsTriggerPrepare() bool {
	return false
}

func (a JinfanArmyHundredFeathersTactic) SetTriggerPrepare(triggerPrepare bool) {
}
