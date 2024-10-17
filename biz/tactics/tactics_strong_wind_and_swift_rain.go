package tactics

import (
	"math"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 疾风骤雨
// 主动 40%
// 随机对敌军武将造成3-4次兵刃攻击（伤害率78%，每次提升6%），第3次和第4次攻击额外附带1回合禁疗状态
type StrongWindAndSwiftRainTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrongWindAndSwiftRainTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s StrongWindAndSwiftRainTactic) Prepare() {
}

func (s StrongWindAndSwiftRainTactic) Id() consts.TacticId {
	return consts.StrongWindAndSwiftRain
}

func (s StrongWindAndSwiftRainTactic) Name() string {
	return "疾风骤雨"
}

func (s StrongWindAndSwiftRainTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s StrongWindAndSwiftRainTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrongWindAndSwiftRainTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrongWindAndSwiftRainTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s StrongWindAndSwiftRainTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrongWindAndSwiftRainTactic) Execute() {
	currentGeneral := s.tacticsParams.CurrentGeneral
	ctx := s.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	//随机对敌军武将造成3-4次兵刃攻击（伤害率78%，每次提升6%），第3次和第4次攻击额外附带1回合禁疗状态
	attackTimes := cast.ToInt(math.Round(util.Random(3, 4)))
	improveRate := 0.06
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, s.tacticsParams)
	for i := 1; i <= attackTimes; i++ {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.78)
		if i > 1 {
			dmg = cast.ToInt64(cast.ToFloat64(dmg) * (1 + improveRate))
		}
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: s.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      s.Id(),
			TacticName:    s.Name(),
		})

		//禁疗
		if i == 3 || i == 4 {
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
				EffectRound: 1,
				FromTactic:  s.Id(),
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_ProhibitionTreatment,
						TacticId:   s.Id(),
					})

					return revokeResp
				})
			}
		}
	}
}

func (s StrongWindAndSwiftRainTactic) IsTriggerPrepare() bool {
	return false
}
