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

// 江东小霸王
// 战斗中，普通攻击后有35%几率对目标再次发起猛攻（伤害率192%）并为我军单体恢复兵力（治疗率56%，受武力影响）
// 自身为主将时，战斗前2回合，使敌军单体首次受到伤害提高30%
// 被动，100%
type JiangdongLittleOverlordTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JiangdongLittleOverlordTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JiangdongLittleOverlordTactic) Prepare() {
	ctx := j.tacticsParams.Ctx
	currentGeneral := j.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		j.Name(),
	)
	// 战斗中，普通攻击后有35%几率对目标再次发起猛攻（伤害率192%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := j.tacticsParams.CurrentSufferGeneral

		if util.GenerateRate(0.35) {
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     j.tacticsParams,
				AttackGeneral:     triggerGeneral,
				SufferGeneral:     sufferGeneral,
				DamageType:        consts.DamageType_Weapon,
				DamageImproveRate: 1.92,
				TacticId:          j.Id(),
				TacticName:        j.Name(),
			})
			//并为我军单体恢复兵力（治疗率56%，受武力影响）
			pairGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, j.tacticsParams)
			resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 0.56)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  j.tacticsParams,
				ProduceGeneral: triggerGeneral,
				SufferGeneral:  pairGeneral,
				ResumeNum:      resumeNum,
				TacticId:       j.Id(),
			})
		}

		return triggerResp
	})
	// 自身为主将时，战斗前2回合，使敌军单体首次受到伤害提高30%
	if currentGeneral.IsMaster {
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, j.tacticsParams)
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.3,
			EffectRound:    2,
			FromTactic:     j.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
					TacticId:   j.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (j JiangdongLittleOverlordTactic) Id() consts.TacticId {
	return consts.JiangdongLittleOverlord
}

func (j JiangdongLittleOverlordTactic) Name() string {
	return "江东小霸王"
}

func (j JiangdongLittleOverlordTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (j JiangdongLittleOverlordTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JiangdongLittleOverlordTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JiangdongLittleOverlordTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (j JiangdongLittleOverlordTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (j JiangdongLittleOverlordTactic) Execute() {
}

func (j JiangdongLittleOverlordTactic) IsTriggerPrepare() bool {
	return false
}

func (a JiangdongLittleOverlordTactic) SetTriggerPrepare(triggerPrepare bool) {
}
