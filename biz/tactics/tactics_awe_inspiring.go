package tactics

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 威风凛凛
// 对敌军单体造成嘲讽（强迫目标普通攻击自己）或缴械（无法进行普通攻击）状态，持续2回合，有75%概率额外释放一次
type AweInspiringTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AweInspiringTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.45
	return a
}

func (a AweInspiringTactic) Prepare() {
}

func (a AweInspiringTactic) Id() consts.TacticId {
	return consts.AweInspiring
}

func (a AweInspiringTactic) Name() string {
	return "威风凛凛"
}

func (a AweInspiringTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AweInspiringTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AweInspiringTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AweInspiringTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AweInspiringTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AweInspiringTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	//对敌军单体造成嘲讽（强迫目标普通攻击自己）或缴械（无法进行普通攻击）状态，持续2回合，有75%概率额外释放一次

	//找到敌军单体
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, a.tacticsParams)
	hitIdx := util.GenerateHitOneIdx(len(enemyGenerals))
	enemyGeneral := enemyGenerals[hitIdx]

	//施加效果
	a.executeTactic(ctx, enemyGeneral, currentGeneral)

	//抛除已施加效果的敌军单体
	newEnemyGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range enemyGenerals {
		if general.BaseInfo.Id != enemyGeneral.BaseInfo.Id {
			newEnemyGenerals = append(newEnemyGenerals, general)
		}
	}

	//有75%概率额外释放一次
	if util.GenerateRate(0.75) && len(newEnemyGenerals) > 0 {
		hitNextIdx := util.GenerateHitOneIdx(len(newEnemyGenerals))
		nextEnemyGeneral := newEnemyGenerals[hitNextIdx]
		a.executeTactic(ctx, nextEnemyGeneral, currentGeneral)
	}
}

func (a AweInspiringTactic) executeTactic(ctx context.Context, enemyGeneral *vo.BattleGeneral, currentGeneral *vo.BattleGeneral) {
	//嘲讽
	if util.GenerateRate(0.5) {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:   2,
			FromTactic:    a.Id(),
			TauntByTarget: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    enemyGeneral,
				EffectType: consts.DebuffEffectType_Taunt,
				TacticId:   a.Id(),
			})
		}
	} else {
		//缴械
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
			EffectRound: 2,
			FromTactic:  a.Id(),
		}).IsSuccess {
			//注册消失效果
			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    enemyGeneral,
				EffectType: consts.DebuffEffectType_CancelWeapon,
				TacticId:   a.Id(),
			})
		}
	}
}

func (a AweInspiringTactic) IsTriggerPrepare() bool {
	return false
}

func (a AweInspiringTactic) SetTriggerPrepare(triggerPrepare bool) {
}
