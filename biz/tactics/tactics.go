package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type Tactics interface {
	// Init 初始化
	Init(tacticsParams model.TacticsParams)

	/** 基础属性 **/

	//战法ID
	Id() int64
	//战法来源
	TacticsSource() consts.TacticsSource
	//战法类型
	TacticsType() consts.TacticsType
	//支持兵种
	SupportArmTypes() []consts.ArmType

	//战法处理
	Handle()
}
