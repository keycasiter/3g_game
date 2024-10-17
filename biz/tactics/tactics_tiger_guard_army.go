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

// 虎卫军
// 将盾兵进阶为善固疆场的虎卫军：
// 战斗中，我军主将即将受到普通攻击时，副将提高12武力，最多提高5次
// 并会分别对攻击者造成兵刃伤害（伤害率72%，受各自损失兵力影响，最多提高40%），每回合最多触发1次
// 若典韦或许褚统领，自身统率提高50
type TigerGuardArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerGuardArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerGuardArmyTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	triggerRoundHolderMap := map[consts.General_Id]map[consts.BattleRound]bool{}

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 将盾兵进阶为善固疆场的虎卫军：
	// 战斗中，我军主将即将受到普通攻击时，副将提高12武力，最多提高5次
	pairMasterGeneral := util.GetPairMasterGeneral(t.tacticsParams)
	util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		attackGeneral := params.AttackGeneral
		triggerRound := params.CurrentRound

		viceGenerals := util.GetPairViceGenerals(t.tacticsParams)
		for _, viceGeneral := range viceGenerals {
			util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
				EffectValue:    12,
				EffectTimes:    1,
				MaxEffectTimes: 5,
				FromTactic:     t.Id(),
				ProduceGeneral: currentGeneral,
			})

			// 并会分别对攻击者造成兵刃伤害（伤害率72%，受各自损失兵力影响，最多提高40%），每回合最多触发1次
			if !triggerRoundHolderMap[consts.General_Id(viceGeneral.BaseInfo.Id)][triggerRound] {
				dmgRate := 0.72
				dmgImproveRate := cast.ToFloat64(viceGeneral.LossSoldierNum / 100 / 100)
				if dmgImproveRate > 0.4 {
					dmgImproveRate = 0.4
				}
				dmgRate += dmgImproveRate

				dmg := cast.ToInt64(viceGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: viceGeneral,
					SufferGeneral: attackGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      t.Id(),
					TacticName:    t.Name(),
				})
				triggerRoundHolderMap[consts.General_Id(viceGeneral.BaseInfo.Id)][triggerRound] = true
			}
		}

		return triggerResp
	})
	// 若典韦或许褚统领，自身统率提高50
	if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.XuChu ||
		consts.General_Id(currentGeneral.BaseInfo.Id) == consts.DianWei {
		util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectValue:    50,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		})
	}
}

func (t TigerGuardArmyTactic) Id() consts.TacticId {
	return consts.TigerGuardArmy
}

func (t TigerGuardArmyTactic) Name() string {
	return "虎卫军"
}

func (t TigerGuardArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TigerGuardArmyTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerGuardArmyTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerGuardArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (t TigerGuardArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Mauler,
	}
}

func (t TigerGuardArmyTactic) Execute() {
}

func (t TigerGuardArmyTactic) IsTriggerPrepare() bool {
	return false
}
