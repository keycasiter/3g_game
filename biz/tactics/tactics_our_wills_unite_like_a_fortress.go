package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 众志成城
// 战斗中，自身每回合有36%概率（受我军全体智力之和影响）治疗我军兵力最低武将（治疗率96%，受智力影响），并对敌军随机单体造成一次谋略伤害（伤害率96%，受智力影响）
type OurWillsUniteLikeAFortressTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a OurWillsUniteLikeAFortressTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a OurWillsUniteLikeAFortressTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//战斗中，自身每回合有36%概率（受我军全体智力之和影响）治疗我军兵力最低武将（治疗率96%，受智力影响），并对敌军随机单体造成一次谋略伤害（伤害率96%，受智力影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//自身每回合有36%概率（受我军全体智力之和影响）
		triggerRate := 0.36
		pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
		for _, general := range pairGenerals {
			triggerRate += general.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
		}
		if !util.GenerateRate(triggerRate) {
			return triggerResp
		}

		//治疗我军兵力最低武将（治疗率96%，受智力影响）
		lowestSoliderGeneral := util.GetPairLowestSoldierNumGeneral(a.tacticsParams, currentGeneral)
		resumeNum := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.96
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  a.tacticsParams,
			ProduceGeneral: triggerGeneral,
			SufferGeneral:  lowestSoliderGeneral,
			ResumeNum:      int64(resumeNum),
			TacticId:       a.Id(),
		})

		//并对敌军随机单体造成一次谋略伤害（伤害率96%，受智力影响）
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, a.tacticsParams)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     a.tacticsParams,
			AttackGeneral:     triggerGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: 0.96 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
			TacticId:          a.Id(),
			TacticName:        a.Name(),
		})

		return triggerResp
	})
}

func (a OurWillsUniteLikeAFortressTactic) Id() consts.TacticId {
	return consts.OurWillsUniteLikeAFortress
}

func (a OurWillsUniteLikeAFortressTactic) Name() string {
	return "众志成城"
}

func (a OurWillsUniteLikeAFortressTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a OurWillsUniteLikeAFortressTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a OurWillsUniteLikeAFortressTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a OurWillsUniteLikeAFortressTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a OurWillsUniteLikeAFortressTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a OurWillsUniteLikeAFortressTactic) Execute() {
}

func (a OurWillsUniteLikeAFortressTactic) IsTriggerPrepare() bool {
	return false
}
