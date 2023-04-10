package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 无当飞军
// 将弓兵进阶为矢不虚发的无当飞军：
// 我军全体统率、速度提高22点，首回合对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率80%，受智力影响），持续3回合
// 若王平统领，对敌军全体施加中毒状态，但伤害率降低（伤害率66%，受智力影响）
type WuDangFlyArmyTactic struct {
	tacticsParams *model.TacticsParams
}

func (w WuDangFlyArmyTactic) TriggerRate() float64 {
	return 1.00
}

func (w WuDangFlyArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	return w
}

func (w WuDangFlyArmyTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	//我军全体统率、速度提高22点
	//找我我军全体
	pairGenerals := util.GetPairGeneralArr(w.tacticsParams)
	for _, general := range pairGenerals {
		general.BaseInfo.AbilityAttr.CommandBase += 22
		hlog.CtxInfof(ctx, "[%s]的统率提高了22",
			general.BaseInfo.Name,
		)
		general.BaseInfo.AbilityAttr.SpeedBase += 22
		hlog.CtxInfof(ctx, "[%s]的速度提高了22",
			general.BaseInfo.Name,
		)
	}
	//首回合对敌军群体（2人）施加中毒状态，每回合持续造成伤害（伤害率80%，受智力影响），持续3回合
	//找到敌军2人
	enemyGenerals := util.GetEnemyGeneralsTwoArr(w.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		//持续3回合
		if !util.TacticsDebuffEffectCountWrapIncr(ctx, sufferGeneral, consts.DebuffEffectType_Methysis, 3, 3, false) {
			return
		}
		//施加中毒效果
		sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_Methysis] = 1.0
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
			sufferGeneral.BaseInfo.Name,
			consts.DebuffEffectType_Methysis,
		)
		//注册效果
		util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerGeneral := params.CurrentGeneral
			triggerResp := &vo.TacticsTriggerResult{}

			if !util.TacticsDebuffEffectCountWrapDecr(triggerGeneral, consts.DebuffEffectType_Methysis, 1) {
				//次数不足
				return triggerResp
			}
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				w.Name(),
				consts.DebuffEffectType_Methysis,
			)
			dmgNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.8)
			dmg, origin, remain, isEffect := util.TacticDamage(w.tacticsParams, currentGeneral, triggerGeneral, dmgNum, consts.BattleAction_SufferArmTactic)
			if !isEffect {
				return triggerResp
			}
			hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的「%v」效果，损失了兵力%d(%d↘%d)",
				triggerGeneral.BaseInfo.Name,
				currentGeneral.BaseInfo.Name,
				w.Name(),
				consts.DebuffEffectType_Methysis,
				dmg,
				origin,
				remain,
			)
			return triggerResp
		})
	}

	//TODO 若王平统领，对敌军全体施加中毒状态，但伤害率降低（伤害率66%，受智力影响）
}

func (w WuDangFlyArmyTactic) Id() consts.TacticId {
	return consts.WuDangFlyArmy
}

func (w WuDangFlyArmyTactic) Name() string {
	return "无当飞军"
}

func (w WuDangFlyArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (w WuDangFlyArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (w WuDangFlyArmyTactic) Execute() {
	return
}
