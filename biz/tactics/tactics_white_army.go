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

// 白毦兵
// 将枪兵进阶为攻无不破的白毦兵：
// 我军全体战斗中普通攻击后有45%概率对攻击目标再次发起一次谋略攻击（伤害率110%，受智力影响）
// 若陈到统领，则谋略攻击更为强力（伤害率130%，受智力影响）
type WhiteArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WhiteArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WhiteArmyTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)
	//我军全体战斗中普通攻击后有45%概率对攻击目标再次发起一次谋略攻击（伤害率110%，受智力影响）
	pairGenerals := util.GetPairGeneralArr(currentGeneral, w.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			attackGeneral := params.AttackGeneral

			if util.GenerateRate(0.45) {
				//若陈到统领，则谋略攻击更为强力（伤害率130%，受智力影响）
				dmgRate := 1.1
				if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.ChenDao {
					dmgRate = 1.3
				}

				dmgRate = triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + dmgRate
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     w.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     attackGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticId:          w.Id(),
					TacticName:        w.Name(),
				})
			}

			return triggerResp
		})
	}
}

func (w WhiteArmyTactic) Id() consts.TacticId {
	return consts.WhiteArmy
}

func (w WhiteArmyTactic) Name() string {
	return "白毦兵"
}

func (w WhiteArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WhiteArmyTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WhiteArmyTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WhiteArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (w WhiteArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (w WhiteArmyTactic) Execute() {
}

func (w WhiteArmyTactic) IsTriggerPrepare() bool {
	return false
}
