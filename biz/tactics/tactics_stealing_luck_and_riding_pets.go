package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 窃幸乘宠
// 我方主将恢复兵力且自身不为主将时，降低其20%治疗量，自身会恢复降低的治疗量，奇数回合对敌军群体（2人）造成谋略伤害（伤害率90%，受智力影响），
// 额外对其中智力低于自身的单位造成谋略伤害（伤害率120%，受智力影响）
// 指挥，100%
type StealingLuckAndRidingPetsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StealingLuckAndRidingPetsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s StealingLuckAndRidingPetsTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 我方主将恢复兵力且自身不为主将时，降低其20%治疗量，自身会恢复降低的治疗量，奇数回合对敌军群体（2人）造成谋略伤害（伤害率90%，受智力影响），
	// 额外对其中智力低于自身的单位造成谋略伤害（伤害率120%，受智力影响）
	pairMasterGeneral := util.GetPairMasterGeneral(currentGeneral, s.tacticsParams)
	if pairMasterGeneral.BaseInfo.Id != currentGeneral.BaseInfo.Id {
		util.DebuffEffectWrapSet(ctx, pairMasterGeneral, consts.DebuffEffectType_SufferResumeDeduce, &vo.EffectHolderParams{
			EffectRate:     0.2,
			FromTactic:     s.Id(),
			ProduceGeneral: currentGeneral,
		})
	}
	//注册效果器
	util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_SufferResume, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		resumeNum := cast.ToInt64(cast.ToFloat64(params.CurrentResume) * 0.2)
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  s.tacticsParams,
			ProduceGeneral: currentGeneral,
			SufferGeneral:  currentGeneral,
			ResumeNum:      resumeNum,
			TacticId:       s.Id(),
		})
		return triggerResp
	})

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		//奇数回合
		if triggerRound%2 != 0 {
			//敌军2人
			dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.9
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, s.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//伤害
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     nil,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticId:          s.Id(),
					TacticName:        s.Name(),
				})
				//额外对其中智力低于自身的单位造成谋略伤害（伤害率120%，受智力影响）
				if enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase < triggerGeneral.BaseInfo.AbilityAttr.IntelligenceRate {
					strategyDmgRate := enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.2
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     s.tacticsParams,
						AttackGeneral:     triggerGeneral,
						SufferGeneral:     enemyGeneral,
						DamageType:        consts.DamageType_Strategy,
						DamageImproveRate: strategyDmgRate,
						TacticId:          s.Id(),
						TacticName:        s.Name(),
					})
				}
			}
		}

		return triggerResp
	})
}

func (s StealingLuckAndRidingPetsTactic) Id() consts.TacticId {
	return consts.StealingLuckAndRidingPets
}

func (s StealingLuckAndRidingPetsTactic) Name() string {
	return "窃幸乘宠"
}

func (s StealingLuckAndRidingPetsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s StealingLuckAndRidingPetsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StealingLuckAndRidingPetsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StealingLuckAndRidingPetsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s StealingLuckAndRidingPetsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StealingLuckAndRidingPetsTactic) Execute() {

}

func (s StealingLuckAndRidingPetsTactic) IsTriggerPrepare() bool {
	return false
}

func (a StealingLuckAndRidingPetsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
