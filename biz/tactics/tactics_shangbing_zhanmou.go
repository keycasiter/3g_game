package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 上兵伐谋
// 分别对敌军兵力最低、武力最高、智力最低的武将发动一次谋略攻击（伤害率128%，受智力影响）
type ShangbingZhanmouTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d ShangbingZhanmouTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.5
	return d
}

func (d ShangbingZhanmouTactic) Prepare() {

}

func (d ShangbingZhanmouTactic) Id() consts.TacticId {
	return consts.ShangbingZhanmou
}

func (d ShangbingZhanmouTactic) Name() string {
	return "上兵伐谋"
}

func (d ShangbingZhanmouTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (d ShangbingZhanmouTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d ShangbingZhanmouTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d ShangbingZhanmouTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d ShangbingZhanmouTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Archers,
		consts.ArmType_Apparatus,
	}
}

func (d ShangbingZhanmouTactic) Execute() {
	currentGeneral := d.tacticsParams.CurrentGeneral
	// 分别对敌军兵力最低、武力最高、智力最低的武将发动一次谋略攻击（伤害率128%，受智力影响）

	//敌军兵力最低
	enemyHighestSoliderGeneral := util.GetEnemyOneGeneralByHighestSolider(currentGeneral, d.tacticsParams)
	//武力最高
	enemyLowestForceGeneral := util.GetEnemyOneGeneralByLowestAttr(currentGeneral, d.tacticsParams, consts.AbilityAttr_Force)
	//智力最低
	enemyLowestIntelligenceGeneral := util.GetEnemyOneGeneralByLowestAttr(currentGeneral, d.tacticsParams, consts.AbilityAttr_Intelligence)

	for _, enemyGeneral := range []*vo.BattleGeneral{
		enemyHighestSoliderGeneral,
		enemyLowestForceGeneral,
		enemyLowestIntelligenceGeneral,
	} {
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     d.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: 1.28,
			TacticId:          d.Id(),
			TacticName:        d.Name(),
		})
	}
}

func (d ShangbingZhanmouTactic) IsTriggerPrepare() bool {
	return false
}

func (a ShangbingZhanmouTactic) SetTriggerPrepare(triggerPrepare bool) {
}
