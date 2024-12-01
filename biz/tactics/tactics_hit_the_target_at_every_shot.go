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

// 百步穿杨
// 准备1回合，提高自身25%会心几率，持续2回合，随后对敌军全体造成兵刃伤害（伤害率180%），若目标处于控制状态，则该次兵刃攻击更为强力（伤害率240%）
// 主动 35%
type HitTheTargetAtEveryShotTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (h HitTheTargetAtEveryShotTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.35
	return h
}

func (h HitTheTargetAtEveryShotTactic) Prepare() {
}

func (h HitTheTargetAtEveryShotTactic) Id() consts.TacticId {
	return consts.HitTheTargetAtEveryShot
}

func (h HitTheTargetAtEveryShotTactic) Name() string {
	return "百步穿杨"
}

func (h HitTheTargetAtEveryShotTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HitTheTargetAtEveryShotTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HitTheTargetAtEveryShotTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HitTheTargetAtEveryShotTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HitTheTargetAtEveryShotTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HitTheTargetAtEveryShotTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral
	currentRound := h.tacticsParams.CurrentRound

	// 准备1回合，提高自身25%会心几率，持续2回合，随后对敌军全体造成兵刃伤害（伤害率180%），若目标处于控制状态，则该次兵刃攻击更为强力（伤害率240%）
	h.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			h.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if h.isTriggered {
				return triggerResp
			} else {
				h.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				h.Name(),
			)
			//提高会心
			util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
				EffectRate:     0.2,
				EffectRound:    2,
				FromTactic:     h.Id(),
				ProduceGeneral: currentGeneral,
			})
			//找到敌军全体
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, h.tacticsParams)

			for _, enemyGeneral := range enemyGenerals {
				dmgRate := 1.8
				if util.DeBuffEffectContainsControl(enemyGeneral) {
					dmgRate = 2.4
				}

				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     h.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: dmgRate,
					TacticId:          h.Id(),
					TacticName:        h.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (h HitTheTargetAtEveryShotTactic) IsTriggerPrepare() bool {
	return h.isTriggerPrepare
}

func (a HitTheTargetAtEveryShotTactic) SetTriggerPrepare(triggerPrepare bool) {
}
