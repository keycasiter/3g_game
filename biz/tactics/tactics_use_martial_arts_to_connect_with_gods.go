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

// 用武通神
// 战斗开始的第2、4、6、8回合，对敌军群体（2人）逐渐造成75%、105%、135%、165%谋略伤害（受智力影响）
type UseMartialArtsToConnectWithGodsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (u UseMartialArtsToConnectWithGodsTactic) IsTriggerPrepare() bool {
	return false
}

func (u UseMartialArtsToConnectWithGodsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	u.tacticsParams = tacticsParams
	u.triggerRate = 1.0
	return u
}

func (u UseMartialArtsToConnectWithGodsTactic) Prepare() {
	ctx := u.tacticsParams.Ctx
	currentGeneral := u.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		u.Name(),
	)

	//战斗开始的第2、4、6、8回合，对敌军群体（2人）逐渐造成75%、105%、135%、165%谋略伤害（受智力影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		switch triggerRound {
		case consts.Battle_Round_Second:
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				u.Name(),
				consts.BuffEffectType_UseMartialArtsToConnectWithGods_Prepare,
			)
			enemyGenerals := util.GetEnemyGeneralsTwoArr(u.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				dmg := cast.ToInt64(0.75 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: u.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: sufferGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Strategy,
					TacticName:    u.Name(),
				})
			}
		case consts.Battle_Round_Fourth:
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				u.Name(),
				consts.BuffEffectType_UseMartialArtsToConnectWithGods_Prepare,
			)
			enemyGenerals := util.GetEnemyGeneralsTwoArr(u.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				dmg := cast.ToInt64(1.05 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: u.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: sufferGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Strategy,
					TacticName:    u.Name(),
				})
			}
		case consts.Battle_Round_Sixth:
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				u.Name(),
				consts.BuffEffectType_UseMartialArtsToConnectWithGods_Prepare,
			)
			enemyGenerals := util.GetEnemyGeneralsTwoArr(u.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				dmg := cast.ToInt64(1.35 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: u.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: sufferGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Strategy,
					TacticName:    u.Name(),
				})
			}
		case consts.Battle_Round_Eighth:
			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				u.Name(),
				consts.BuffEffectType_UseMartialArtsToConnectWithGods_Prepare,
			)
			enemyGenerals := util.GetEnemyGeneralsTwoArr(u.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				dmg := cast.ToInt64(1.65 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: u.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: sufferGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Strategy,
					TacticName:    u.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (u UseMartialArtsToConnectWithGodsTactic) Id() consts.TacticId {
	return consts.UseMartialArtsToConnectWithGods
}

func (u UseMartialArtsToConnectWithGodsTactic) Name() string {
	return "用武通神"
}

func (u UseMartialArtsToConnectWithGodsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (u UseMartialArtsToConnectWithGodsTactic) GetTriggerRate() float64 {
	return u.triggerRate
}

func (u UseMartialArtsToConnectWithGodsTactic) SetTriggerRate(rate float64) {
	u.triggerRate = rate
}

func (u UseMartialArtsToConnectWithGodsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (u UseMartialArtsToConnectWithGodsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (u UseMartialArtsToConnectWithGodsTactic) Execute() {
	return
}
