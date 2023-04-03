package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：镇扼防拒
// 战法描述：每回合有50%概率（受智力影响）使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响），
// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态
type SuppressChokesAndPreventRefusalsTactic struct {
	tacticsParams model.TacticsParams
}

func (s SuppressChokesAndPreventRefusalsTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	return s
}

func (s SuppressChokesAndPreventRefusalsTactic) Prepare() {
	for i := int(consts.Battle_Round_First); i <= int(consts.Battle_Round_Eighth); i++ {
		//每回合有50%概率
		//TODO（受智力影响）
		if !util.GenerateRate(0.5) {
			return
		}
		//使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响）
		//找到除当前战法执行外的副将
		viceGeneral := util.GetPairViceGeneralNotSelf(s.tacticsParams)
		//让这个副将援护友军
		generals := util.GetPairGeneralsNotSelf(s.tacticsParams, viceGeneral)
		for _, general := range generals {
			general.TacticsTriggerMap[consts.BattleAction_SufferAttack] = &SuppressChokesAndPreventRefusalsTactic{}
		}
	}
}

func (s SuppressChokesAndPreventRefusalsTactic) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) Execute() {
	return
}

func (s SuppressChokesAndPreventRefusalsTactic) Trigger() {

}

func (s SuppressChokesAndPreventRefusalsTactic) Id() int64 {
	return SuppressChokesAndPreventRefusals
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SuppressChokesAndPreventRefusalsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
