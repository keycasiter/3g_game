package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 驱散
// 解除敌军全体身上的增益效果，并提高自己28点智力，持续3回合
type DisperseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DisperseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.45
	return d
}

func (d DisperseTactic) Prepare() {
}

func (d DisperseTactic) Id() consts.TacticId {
	return consts.Disperse
}

func (d DisperseTactic) Name() string {
	return "驱散"
}

func (d DisperseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DisperseTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DisperseTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DisperseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DisperseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DisperseTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral
	//解除敌军全体身上的增益效果，并提高自己28点智力，持续3回合

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)

	//找到敌军全体
	enemyGenerals := util.GetEnemyGeneralArr(d.tacticsParams)
	for _, general := range enemyGenerals {
		util.BuffEffectClean(ctx, general)
	}

	//提高自己28点智力
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectRound: 3,
		EffectValue: 28,
		FromTactic:  d.Id(),
	}).IsSuccess {
		//消失效果
		util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
			Ctx:        ctx,
			General:    currentGeneral,
			EffectType: consts.BuffEffectType_IncrIntelligence,
			TacticId:   d.Id(),
		})
	}

}

func (d DisperseTactic) IsTriggerPrepare() bool {
	return false
}
