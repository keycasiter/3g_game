package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：熯天炽地
// 战法描述：准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
// 主动战法 发动率35%
type TheSkyIsBlazingTactic struct {
	tacticsParams *model.TacticsParams
}

func (t TheSkyIsBlazingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSkyIsBlazingTactic) TriggerRate() float64 {
	return 0.35
}

func (t TheSkyIsBlazingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t TheSkyIsBlazingTactic) Prepare() {

}

func (t TheSkyIsBlazingTactic) Name() string {
	return "熯天炽地"
}

func (t TheSkyIsBlazingTactic) Execute() {
	//准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
	// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。

	//找到敌军全体
	enemyGenerals := util.GetEnemyGeneralArr(t.tacticsParams)
	for _, general := range enemyGenerals {
		//准备1回合,持续2回合
		//TODO 受智力影响
		util.DebuffEffectWrapSet(general,
			consts.DebuffEffectType_Firing, 1.02,
		)
		util.DebuffEffectWrapSet(general, consts.DebuffEffectType_Firing, 0.72)
		util.DebuffEffectWrapSet(general, consts.DebuffEffectType_Firing, 0.72)
	}
}

func (t TheSkyIsBlazingTactic) Trigger() {
	return
}

func (t TheSkyIsBlazingTactic) Id() consts.TacticId {
	return consts.TheSkyIsBlazing
}

func (t TheSkyIsBlazingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSkyIsBlazingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
