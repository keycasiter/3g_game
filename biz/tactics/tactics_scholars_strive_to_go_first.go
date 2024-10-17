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

// 士争先赴
// 提高自带主动战法20%伤害，成功发动自带主动战法前，50%概率对敌方群体（2～3人）造成兵刃伤害（伤害率120%）
// 被动，100%
type ScholarsStriveToGoFirstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ScholarsStriveToGoFirstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s ScholarsStriveToGoFirstTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	//提高自带主动战法20%伤害
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveWithSelfDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.2,
		FromTactic:     s.Id(),
		ProduceGeneral: currentGeneral,
	})
	//成功发动自带主动战法前，50%概率对敌方群体（2～3人）造成兵刃伤害（伤害率120%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.5) {
			enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(s.tacticsParams)
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.2)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: s.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      s.Id(),
					TacticName:    s.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (s ScholarsStriveToGoFirstTactic) Id() consts.TacticId {
	return consts.ScholarsStriveToGoFirst
}

func (s ScholarsStriveToGoFirstTactic) Name() string {
	return "士争先赴"
}

func (s ScholarsStriveToGoFirstTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s ScholarsStriveToGoFirstTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ScholarsStriveToGoFirstTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ScholarsStriveToGoFirstTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s ScholarsStriveToGoFirstTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ScholarsStriveToGoFirstTactic) Execute() {
}

func (s ScholarsStriveToGoFirstTactic) IsTriggerPrepare() bool {
	return false
}
