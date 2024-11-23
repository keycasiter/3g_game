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

// 以寡敌众
// 战斗第2、4、6回合，恢复自身兵力（治疗率288%，受武力影响）并提高自身10%武力；
// 战斗第5回合，对敌军全体造成兵刃伤害（伤害率72%）
type YiGuaDiZhongTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a YiGuaDiZhongTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a YiGuaDiZhongTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗第2、4、6回合，恢复自身兵力（治疗率288%，受武力影响）并提高自身10%武力；
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound == consts.Battle_Round_Second ||
			triggerRound == consts.Battle_Round_Fourth ||
			triggerRound == consts.Battle_Round_Sixth {

			//恢复自身兵力（治疗率288%，受武力影响
			resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 2.88)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  a.tacticsParams,
				ProduceGeneral: currentGeneral,
				SufferGeneral:  triggerGeneral,
				ResumeNum:      resumeNum,
				TacticId:       a.Id(),
			})
			//并提高自身10%武力
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
				EffectRate:     0.1,
				FromTactic:     a.Id(),
				ProduceGeneral: currentGeneral,
			})
		}

		// 战斗第5回合，对敌军全体造成兵刃伤害（伤害率72%）
		if triggerRound == consts.Battle_Round_Fifth {
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, a.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     a.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 0.72,
					TacticId:          a.Id(),
					TacticName:        a.Name(),
				})
			}
		}

		return triggerResp
	})
}

func (a YiGuaDiZhongTactic) Id() consts.TacticId {
	return consts.YiGuaDiZhong
}

func (a YiGuaDiZhongTactic) Name() string {
	return "以寡敌众"
}

func (a YiGuaDiZhongTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a YiGuaDiZhongTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a YiGuaDiZhongTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a YiGuaDiZhongTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a YiGuaDiZhongTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a YiGuaDiZhongTactic) Execute() {
}

func (a YiGuaDiZhongTactic) IsTriggerPrepare() bool {
	return false
}
