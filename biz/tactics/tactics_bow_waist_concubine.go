package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 弓腰姬
// 发动普通攻击前对敌军单体造成兵刃伤害（伤害率122%），自身拥有功能性增益状态时额外对其造成兵刃伤害（伤害率20%x状态数）并提高18武力，最多叠加5次
// 指挥 100%
type BowWaistConcubineTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BowWaistConcubineTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BowWaistConcubineTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_Attack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//发动普通攻击前对敌军单体造成兵刃伤害（伤害率122%）
		//找到敌军单体
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, b.tacticsParams)
		dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.22)

		//自身拥有功能性增益状态时额外对其造成兵刃伤害（伤害率20%x状态数）
		buffEffectNum := util.BuffEffectContainsNum(currentGeneral)
		if buffEffectNum > 0 {
			buffEffectDesc := ""
			for buffEffectType, _ := range currentGeneral.BuffEffectHolderMap {
				buffEffectDesc += fmt.Sprintf("[%v]", buffEffectType)
			}
			extDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * cast.ToFloat64(buffEffectNum) * 0.2)
			hlog.CtxInfof(ctx, "功能性增益状态数:%v , 功能性增益状态:%v , 造成额外伤害:%v", buffEffectNum, buffEffectDesc, extDmg)
			dmg += extDmg

			//并提高18武力，最多叠加5次
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
				EffectValue:    18,
				EffectTimes:    1,
				MaxEffectTimes: 5,
				FromTactic:     b.Id(),
			})
		}

		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: b.tacticsParams,
			AttackGeneral: triggerGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticName:    b.Name(),
			TacticId:      b.Id(),
		})

		return triggerResp
	})
}

func (b BowWaistConcubineTactic) Id() consts.TacticId {
	return consts.BowWaistConcubine
}

func (b BowWaistConcubineTactic) Name() string {
	return "弓腰姬"
}

func (b BowWaistConcubineTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BowWaistConcubineTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BowWaistConcubineTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BowWaistConcubineTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (b BowWaistConcubineTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BowWaistConcubineTactic) Execute() {

}

func (b BowWaistConcubineTactic) IsTriggerPrepare() bool {
	return false
}
