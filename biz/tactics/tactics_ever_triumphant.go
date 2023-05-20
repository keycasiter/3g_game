package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 所向披靡
// 准备1回合，对敌军全体发动一次兵刃攻击(伤害率246%)
type EverTriumphantTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (e EverTriumphantTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 0.3
	return e
}

func (e EverTriumphantTactic) Prepare() {

}

func (e EverTriumphantTactic) Id() consts.TacticId {
	return consts.EverTriumphant
}

func (e EverTriumphantTactic) Name() string {
	return "所向披靡"
}

func (e EverTriumphantTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (e EverTriumphantTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EverTriumphantTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EverTriumphantTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (e EverTriumphantTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EverTriumphantTactic) Execute() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral
	currentRound := e.tacticsParams.CurrentRound

	//准备1回合，对敌军全体发动一次兵刃攻击(伤害率246%)
	e.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		e.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if currentRound+1 == triggerRound {
			//准备回合释放
			e.isTriggerPrepare = false
			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				e.Name(),
			)
			//找到敌军全体
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, e.tacticsParams)
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 2.46)
			for _, enemyGeneral := range enemyGenerals {
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: e.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					Damage:        dmg,
					TacticName:    e.Name(),
				})
			}
		}

		return triggerResp
	})

}

func (e EverTriumphantTactic) IsTriggerPrepare() bool {
	return e.isTriggerPrepare
}
