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

// 刚烈不屈
// 战斗中，使自己统率提高38点，受到兵刃伤害时有40%几率对敌军群体（2人）造成兵刃伤害（伤害率84%）
// 被动，100%
type StrongAndUnyieldingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrongAndUnyieldingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s StrongAndUnyieldingTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 战斗中，使自己统率提高38点，受到兵刃伤害时有40%几率对敌军群体（2人）造成兵刃伤害（伤害率84%）
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    38,
		FromTactic:     s.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferWeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.4) {
			//找到敌军2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, s.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     s.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 0.84,
					TacticId:          s.Id(),
					TacticName:        s.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (s StrongAndUnyieldingTactic) Id() consts.TacticId {
	return consts.StrongAndUnyielding
}

func (s StrongAndUnyieldingTactic) Name() string {
	return "刚烈不屈"
}

func (s StrongAndUnyieldingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s StrongAndUnyieldingTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrongAndUnyieldingTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrongAndUnyieldingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s StrongAndUnyieldingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrongAndUnyieldingTactic) Execute() {
}

func (s StrongAndUnyieldingTactic) IsTriggerPrepare() bool {
	return false
}

func (a StrongAndUnyieldingTactic) SetTriggerPrepare(triggerPrepare bool) {
}
