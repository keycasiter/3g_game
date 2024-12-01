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

// 横扫千军
// 对敌军全体造成100%兵刃伤害，若目标处于缴械或者计穷状态则有30%概率使目标处于震慑状态（无法行动），持续1回合
type SweepAwayTheMillionsOfEnemyTroopsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Prepare() {

}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Id() consts.TacticId {
	return consts.SweepAwayTheMillionsOfEnemyTroops
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Name() string {
	return "横扫千军"
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Execute() {
	//对敌军全体造成100%兵刃伤害，若目标处于缴械或者计穷状态则有30%概率使目标处于震慑状态（无法行动），持续1回合
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, s.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//伤害
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     s.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Weapon,
			DamageImproveRate: 1,
			TacticId:          s.Id(),
			TacticName:        s.Name(),
		})
		//若目标处于缴械或者计穷状态则有30%概率使目标处于震慑状态（无法行动），持续1回合
		if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_CancelWeapon) ||
			util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_NoStrategy) {

			if !util.GenerateRate(0.3) {
				return
			}
			//施加效果
			if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
				EffectRound: 1,
				FromTactic:  s.Id(),
			}).IsSuccess {
				//注册消失效果
				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemyGeneral,
					EffectType: consts.DebuffEffectType_Awe,
					TacticId:   s.Id(),
				})
			}
		}
	}
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) IsTriggerPrepare() bool {
	return false
}

func (a SweepAwayTheMillionsOfEnemyTroopsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
