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

// 草船借箭
// 主动战法 65%
// 移除我军群体(2-3人)负面效果，并使我军群体(2人)获得急救状态，每次受到伤害时有70%几率回复一定兵力（伤害量的28%，受统率影响）
// 持续2回合，该战法发动后会进入1回合冷却
type BorrowArrowsWithThatchedBoatsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BorrowArrowsWithThatchedBoatsTactic) IsTriggerPrepare() bool {
	return false
}

func (b BorrowArrowsWithThatchedBoatsTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BorrowArrowsWithThatchedBoatsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BorrowArrowsWithThatchedBoatsTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BorrowArrowsWithThatchedBoatsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.65
	return b
}

func (b BorrowArrowsWithThatchedBoatsTactic) Prepare() {
	return
}

func (b BorrowArrowsWithThatchedBoatsTactic) Id() consts.TacticId {
	return consts.BorrowArrowsWithThatchedBoats
}

func (b BorrowArrowsWithThatchedBoatsTactic) Name() string {
	return "草船借箭"
}

func (b BorrowArrowsWithThatchedBoatsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (b BorrowArrowsWithThatchedBoatsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BorrowArrowsWithThatchedBoatsTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//判断是否冷却
	if cnt, ok := currentGeneral.TacticsFrozenMap[b.Id()]; ok {
		if cnt > 0 {
			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果生效中",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)
			return
		}
	}
	//该战法发动后会进入1回合冷却
	if !util.TacticFrozenWrapSet(currentGeneral, b.Id(), 1, 1, false) {
		return
	}
	//注册冷却效果消失
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		revokeResp := &vo.TacticsTriggerResult{}
		revokeRound := params.CurrentRound
		revokeGeneral := params.CurrentGeneral
		//1回合冷却，下下回合冷却结束
		if currentRound+2 == revokeRound {
			if !util.TacticFrozenWrapRemove(revokeGeneral, b.Id()) {
				hlog.CtxErrorf(ctx, "TacticFrozenWrapRemove err")
				panic("TacticFrozenWrapRemove err")
			}

			hlog.CtxInfof(ctx, "[%s]的「%s[冷却]」效果已消失",
				currentGeneral.BaseInfo.Name,
				b.Name(),
			)
		}
		return revokeResp
	})

	//移除我军群体(2-3人)负面效果
	twoOrThreeGenerals := util.GetPairGeneralsTwoOrThreeMap(b.tacticsParams)
	for _, general := range twoOrThreeGenerals {
		util.DebuffEffectClean(ctx, general)
	}
	//并使我军群体(2人)获得急救状态，每次受到伤害时有70%几率回复一定兵力（伤害量的28%，受统率影响）,持续2回合
	twoGenerals := util.GetPairGeneralsTwoArr(b.tacticsParams)
	for _, general := range twoGenerals {
		//施加急救效果，持续2回合
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_EmergencyTreatment, &vo.EffectHolderParams{
			EffectTimes: 2,
			FromTactic:  b.Id(),
		}).IsSuccess {
			//注册触发效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerGeneral := params.CurrentGeneral
				triggerResp := &vo.TacticsTriggerResult{}
				//效果消耗
				if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    triggerGeneral,
					EffectType: consts.BuffEffectType_EmergencyTreatment,
					TacticId:   b.Id(),
				}) {
					if !util.GenerateRate(0.7) {
						hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「急救」效果因几率没有生效",
							triggerGeneral.BaseInfo.Name,
							b.Name(),
						)
						return triggerResp
					} else {
						hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「急救」效果",
							triggerGeneral.BaseInfo.Name,
							b.Name(),
						)
						// TODO 受统率影响
						resumeNum := cast.ToInt64(cast.ToFloat64(params.CurrentDamage) * 0.28)
						finalResumeNum, holdNum, finalNum := util.ResumeSoldierNum(triggerGeneral, resumeNum)
						hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
							triggerGeneral.BaseInfo.Name,
							finalResumeNum,
							holdNum,
							finalNum,
						)
					}
				}

				//效果消失
				if util.BuffEffectOfTacticIsDeplete(triggerGeneral, consts.BuffEffectType_EmergencyTreatment, b.Id()) {
					util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_EmergencyTreatment, b.Id())
				}
				return triggerResp
			})
		}
	}
	return
}
