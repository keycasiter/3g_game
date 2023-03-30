package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：抚辑军民
// 战法描述：战斗前3回合，使我军群体(2人)造成的伤害降低24%，
// 受到的伤害降低24%（受统率影响），
// 战斗第4回合时，恢复其兵力（治疗率126%，受智力影响）
type AppeaseArmyAndPeopleTactic struct {
	tacticsParams model.TacticsParams
}

func (a AppeaseArmyAndPeopleTactic) Init(tacticsParams model.TacticsParams) {
	a.tacticsParams = tacticsParams
}

func (a AppeaseArmyAndPeopleTactic) Id() int64 {
	return AppeaseArmyAndPeople
}

func (a AppeaseArmyAndPeopleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (a AppeaseArmyAndPeopleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a AppeaseArmyAndPeopleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AppeaseArmyAndPeopleTactic) Handle() {
	panic("implement me")
}
