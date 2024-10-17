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

// 伪书相间
// 对敌军单体造成谋略伤害（伤害率206%，受智力影响）
// 若目标处于混乱状态则使目标对其友军单体发动攻击（伤害率186%，类型取决于目标武力、智力较高的一项），
// 否则施加混乱（攻击和战法无差别选择目标）状态，持续1回合
type FakeBooksAlternateWithEachOtherTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FakeBooksAlternateWithEachOtherTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.45
	return f
}

func (f FakeBooksAlternateWithEachOtherTactic) Prepare() {

}

func (f FakeBooksAlternateWithEachOtherTactic) Id() consts.TacticId {
	return consts.FakeBooksAlternateWithEachOther
}

func (f FakeBooksAlternateWithEachOtherTactic) Name() string {
	return "伪书相间"
}

func (f FakeBooksAlternateWithEachOtherTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FakeBooksAlternateWithEachOtherTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FakeBooksAlternateWithEachOtherTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FakeBooksAlternateWithEachOtherTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FakeBooksAlternateWithEachOtherTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FakeBooksAlternateWithEachOtherTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	//对敌军单体造成谋略伤害（伤害率206%，受智力影响）
	//找到敌军单体
	enemeyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, f.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 2.06)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: f.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemeyGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      f.Id(),
		TacticName:    f.Name(),
	})
	//若目标处于混乱状态则使目标对其友军单体发动攻击（伤害率186%，类型取决于目标武力、智力较高的一项），
	if util.DeBuffEffectContains(enemeyGeneral, consts.DebuffEffectType_Chaos) {
		//找到该目标友军
		enemyPairGeneral := util.GetPairOneGeneralNotSelf(f.tacticsParams, enemeyGeneral)
		attrType, attrVal := util.GetGeneralHighestBetweenForceOrIntelligence(enemeyGeneral)
		chaosDmg := cast.ToInt64(attrVal * 1.86)

		//伤害类型
		dmgType := consts.DamageType_Strategy
		if attrType == consts.AbilityAttr_Force {
			dmgType = consts.DamageType_Weapon
		}

		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: f.tacticsParams,
			AttackGeneral: enemeyGeneral,
			SufferGeneral: enemyPairGeneral,
			DamageType:    dmgType,
			Damage:        chaosDmg,
			TacticId:      f.Id(),
			TacticName:    f.Name(),
		})
	} else {
		//否则施加混乱（攻击和战法无差别选择目标）状态，持续1回合
		if util.DebuffEffectWrapSet(ctx, enemeyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    enemeyGeneral,
				EffectType: consts.DebuffEffectType_Chaos,
				TacticId:   f.Id(),
			})
		}
	}
}

func (f FakeBooksAlternateWithEachOtherTactic) IsTriggerPrepare() bool {
	return false
}
