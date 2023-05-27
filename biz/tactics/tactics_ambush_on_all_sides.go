package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

//十面埋伏
//对有负面状态的敌军造成谋略攻击（伤害率96%，受智力影响），并对敌军群体（2人）施加禁疗（无法恢复兵力）
//及叛逃状态，每回合持续造成伤害（伤害率74%，受武力或智力最高一项影响，无视防御），持续2回合
type AmbushOnAllSidesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AmbushOnAllSidesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.35
	return a
}

func (a AmbushOnAllSidesTactic) Prepare() {
}

func (a AmbushOnAllSidesTactic) Id() consts.TacticId {
	return consts.AmbushOnAllSides
}

func (a AmbushOnAllSidesTactic) Name() string {
	return "十面埋伏"
}

func (a AmbushOnAllSidesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AmbushOnAllSidesTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AmbushOnAllSidesTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AmbushOnAllSidesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AmbushOnAllSidesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AmbushOnAllSidesTactic) Execute() {
	currentGeneral := a.tacticsParams.CurrentGeneral
	ctx := a.tacticsParams.Ctx

	//对有负面状态的敌军造成谋略攻击（伤害率96%，受智力影响），并对敌军群体（2人）施加禁疗（无法恢复兵力）
	//及叛逃状态，每回合持续造成伤害（伤害率74%，受武力或智力最高一项影响，无视防御），持续2回合
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	//找到敌军
	enemyGenerals := util.GetEnemyGeneralsTwoArr(a.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		//找到有负面状态的
		if util.DeBuffEffectContainsCheck(sufferGeneral) {
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.96)
			util.TacticDamage(&util.TacticDamageParam{
				TacticsParams: a.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: sufferGeneral,
				Damage:        dmg,
				TacticName:    a.Name(),
			})
		}
	}
	//对敌军群体（2人）施加禁疗（无法恢复兵力）
	twoEnemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, a.tacticsParams)
	for _, sufferGeneral := range twoEnemyGenerals {
		//施加禁疗效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
			EffectRate:  1.0,
			EffectRound: 2,
			FromTactic:  a.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				//效果消耗
				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_ProhibitionTreatment,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
		//施加叛逃效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Defect, &vo.EffectHolderParams{
			EffectRate:  1.0,
			EffectRound: 2,
			FromTactic:  a.Id(),
		}).IsSuccess {
			//注册触发效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral

				//叛逃状态，每回合持续造成伤害（伤害率74%，受武力或智力最高一项影响，无视防御），持续2回合
				//触发效果
				attr := util.GetGeneralHighestBetweenForceOrIntelligence(currentGeneral)
				dmg := cast.ToInt64(attr * 0.74)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams:  a.tacticsParams,
					AttackGeneral:  currentGeneral,
					SufferGeneral:  triggerGeneral,
					Damage:         dmg,
					TacticName:     a.Name(),
					EffectName:     fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
					IsIgnoreDefend: true,
				})

				//效果消耗
				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    triggerGeneral,
					EffectType: consts.DebuffEffectType_Defect,
					TacticId:   a.Id(),
				})

				return triggerResp
			})
		}
	}
}

func (a AmbushOnAllSidesTactic) IsTriggerPrepare() bool {
	return false
}
