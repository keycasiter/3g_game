package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 青州兵
// 将枪兵进阶为冲坚毁锐的青州兵：
// 战斗2回合，使我军群体（2人）受到普通攻击时对攻击者进行一次反击（伤害率72%），
// 第3回合，回合开始时依次为我军全体恢复兵力，优先完全恢复我军兵力最低单体，再恢复我军其他单体
// （总治疗率180%，受武力影响，额外受敌军造成伤害影响）
// 若曹操统领，治疗效果额外受统率影响
type QingZhouSoldierTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (q QingZhouSoldierTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	q.tacticsParams = tacticsParams
	q.triggerRate = 1.0
	return q
}

func (q QingZhouSoldierTactic) Prepare() {
	ctx := q.tacticsParams.Ctx
	currentGeneral := q.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		q.Name(),
	)
	//战斗2回合，使我军群体（2人）受到普通攻击时对攻击者进行一次反击（伤害率72%）
	pairGenerals := util.GetPairGeneralsTwoArrByGeneral(currentGeneral, q.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.SufferAttackGeneral
			triggerRound := params.CurrentRound
			attackGeneral := params.AttackGeneral

			if triggerRound <= consts.Battle_Round_Second {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     q.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     attackGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 0.72,
					TacticId:          q.Id(),
					TacticName:        q.Name(),
					EffectName:        fmt.Sprintf("%v", consts.BuffEffectType_StrikeBack),
				})
			}

			return triggerResp
		})
	}
	//第3回合，回合开始时依次为我军全体恢复兵力，优先完全恢复我军兵力最低单体，再恢复我军其他单体（总治疗率180%，受武力影响，额外受敌军造成伤害影响）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Third {
			resumeRate := 1.8
			// 若曹操统领，治疗效果额外受统率影响
			if consts.General_Id(currentGeneral.BaseInfo.Id) == consts.CaoCao {
				resumeRate += currentGeneral.BaseInfo.AbilityAttr.CommandBase / 100 / 100
			}
			//TODO 额外受敌军造成伤害影响

			//优先完全恢复我军兵力最低单体
			allPairGenerals := util.GetPairGeneralArr(currentGeneral, q.tacticsParams)
			general := util.GetLowestSoliderNumGeneral(allPairGenerals)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  q.tacticsParams,
				ProduceGeneral: general,
				SufferGeneral:  general,
				ResumeNum:      general.LossSoldierNum,
				TacticId:       q.Id(),
			})
			//再恢复我军其他单体
			for _, pairGeneral := range allPairGenerals {
				if pairGeneral.BaseInfo.Id != general.BaseInfo.Id {
					resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * resumeRate)
					util.ResumeSoldierNum(&util.ResumeParams{
						Ctx:            ctx,
						TacticsParams:  q.tacticsParams,
						ProduceGeneral: triggerGeneral,
						SufferGeneral:  pairGeneral,
						ResumeNum:      resumeNum,
						TacticId:       q.Id(),
					})
				}
			}
		}

		return triggerResp
	})
}

func (q QingZhouSoldierTactic) Id() consts.TacticId {
	return consts.QingZhouSoldier
}

func (q QingZhouSoldierTactic) Name() string {
	return "青州兵"
}

func (q QingZhouSoldierTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (q QingZhouSoldierTactic) GetTriggerRate() float64 {
	return q.triggerRate
}

func (q QingZhouSoldierTactic) SetTriggerRate(rate float64) {
	q.triggerRate = rate
}

func (q QingZhouSoldierTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (q QingZhouSoldierTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (q QingZhouSoldierTactic) Execute() {
}

func (q QingZhouSoldierTactic) IsTriggerPrepare() bool {
	return false
}

func (a QingZhouSoldierTactic) SetTriggerPrepare(triggerPrepare bool) {
}
