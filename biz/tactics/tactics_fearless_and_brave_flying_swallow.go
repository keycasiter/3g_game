package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 轻勇飞燕
// 对敌军单体造成兵刃攻击（伤害率84%），随机释放2-4次
// 主动，40%
type FearlessAndBraveFlyingSwallowTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FearlessAndBraveFlyingSwallowTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.4
	return f
}

func (f FearlessAndBraveFlyingSwallowTactic) Prepare() {
}

func (f FearlessAndBraveFlyingSwallowTactic) Id() consts.TacticId {
	return consts.FearlessAndBraveFlyingSwallow
}

func (f FearlessAndBraveFlyingSwallowTactic) Name() string {
	return "轻勇飞燕"
}

func (f FearlessAndBraveFlyingSwallowTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FearlessAndBraveFlyingSwallowTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FearlessAndBraveFlyingSwallowTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FearlessAndBraveFlyingSwallowTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FearlessAndBraveFlyingSwallowTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FearlessAndBraveFlyingSwallowTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	// 对敌军单体造成兵刃攻击（伤害率84%），随机释放2-4次
	// 找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, f.tacticsParams)
	//随机释放
	random := []int{2, 3, 4}
	hitIdx := util.GenerateHitOneIdx(len(random))
	for i := 0; i < random[hitIdx]; i++ {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.84)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: f.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticName:    f.Name(),
			TacticId:      f.Id(),
		})
	}
}

func (f FearlessAndBraveFlyingSwallowTactic) IsTriggerPrepare() bool {
	return false
}
