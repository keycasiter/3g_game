package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：乱世奸雄
// 战法描述：战斗中，使友军群体(2人)造成伤害提高16%（受智力影响），
// 自己受到伤害降低18%（受智力影响），如果自己为主将，副将造成伤害时，会为主将恢复其伤害量10%的兵力
type TraitorInTroubledTimesTactic struct {
	tacticsParams model.TacticsParams
}

func (t TraitorInTroubledTimesTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	return t
}

func (t TraitorInTroubledTimesTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) Name() string {
	return "乱世奸雄"
}

func (t TraitorInTroubledTimesTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) Trigger() {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}
