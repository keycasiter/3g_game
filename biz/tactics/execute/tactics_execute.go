package execute

import (
	"github.com/keycasiter/3g_game/biz/tactics"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法执行器
func TacticsExecute(tactics tactics.Tactics, tacticsParams model.TacticsParams) {
	//初始化当前回合战法参数
	tactics.Init(tacticsParams)

	//通过战法触发概率判断是否触发
	isTrigger := util.GenerateRate(tactics.TriggerRate())
	if isTrigger {
		//增益结算
		tactics.BuffEffect()

		//减益结算
		tactics.DebuffEffect()

		//伤害结算
		if tactics.DamageRate() > 0 || tactics.DamageNum() > 0 {

		}
	}

	//是否可以普攻
	if tactics.IsGeneralAttack() {

	}
}
