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

// 大戟士
// 将枪兵进阶为横冲直撞的大戟士：
// 我军全体武力提升14点，进行普通攻击时，有35%几率对敌军单体造成兵刃伤害（伤害率122%）
// 若张合统领，则发动几率提高为40%
type GreatHalberdWarriorTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GreatHalberdWarriorTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GreatHalberdWarriorTactic) Prepare() {

}

func (g GreatHalberdWarriorTactic) Id() consts.TacticId {
	return consts.GreatHalberdWarrior
}

func (g GreatHalberdWarriorTactic) Name() string {
	return "大戟士"
}

func (g GreatHalberdWarriorTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (g GreatHalberdWarriorTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GreatHalberdWarriorTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GreatHalberdWarriorTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (g GreatHalberdWarriorTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (g GreatHalberdWarriorTactic) Execute() {
	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)
	// 将枪兵进阶为横冲直撞的大戟士：
	// 我军全体武力提升14点，进行普通攻击时，有35%几率对敌军单体造成兵刃伤害（伤害率122%）
	// 若张合统领，则发动几率提高为40%

	//找到我军全体
	pairGenerals := util.GetPairGeneralArr(currentGeneral, g.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//施加效果
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    14,
			FromTactic:     g.Id(),
			ProduceGeneral: currentGeneral,
		})
		//效果
		triggerRate := 0.35
		if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.ZhangHe {
			triggerRate = 0.4
		}
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			//找到敌军单体
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, g.tacticsParams)

			if util.GenerateRate(triggerRate) {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     g.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 1.22,
					TacticId:          g.Id(),
					TacticName:        g.Name(),
				})
			}

			return triggerResp
		})
	}
}

func (g GreatHalberdWarriorTactic) IsTriggerPrepare() bool {
	return false
}

func (a GreatHalberdWarriorTactic) SetTriggerPrepare(triggerPrepare bool) {
}
