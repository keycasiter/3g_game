package execute

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法执行器
func TacticsExecute(ctx context.Context, tactics tactics.Tactics, tacticsParams model.TacticsParams) {
	//初始化当前回合战法参数
	tactics.Init(tacticsParams)

	//通过战法触发概率判断是否触发
	isTrigger := util.GenerateRate(tactics.TriggerRate())
	if isTrigger {
		//增益结算
		switch tactics.BuffEffect() {
		case consts.BuffEffectType_Evade:
			tacticsParams.CurrentGeneral.BuffEffectMap[consts.BuffEffectType_Evade] = tactics.BuffEffectRate()

			hlog.CtxInfof(ctx, "武将[%s]获取规避效果", tacticsParams.CurrentGeneral.BaseInfo.Name)
		}

		//减益结算
		tactics.DebuffEffect()

		//伤害结算
		if tactics.DamageRate() > 0 || tactics.DamageNum() > 0 {
			//伤害类型计算
			switch tactics.DamageType() {
			//兵刃伤害
			case consts.DamageType_Weapon:

			//谋略伤害
			case consts.DamageType_Strategy:
			}
		}
	}

	//是否可以普攻
	if tactics.IsGeneralAttack() {

	}
}
