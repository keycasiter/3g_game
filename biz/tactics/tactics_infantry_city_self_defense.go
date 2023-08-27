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

// 婴城自守
// 恢复我军群体（2人）兵力（治疗率92%，受智力影响），
// 并使其获得休整状态（每回合恢复一次兵力，治疗率62%），持续1回合
type InfantryCitySelfDefenseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InfantryCitySelfDefenseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.5
	return i
}

func (i InfantryCitySelfDefenseTactic) Prepare() {
}

func (i InfantryCitySelfDefenseTactic) Id() consts.TacticId {
	return consts.InfantryCitySelfDefense
}

func (i InfantryCitySelfDefenseTactic) Name() string {
	return "婴城自守"
}

func (i InfantryCitySelfDefenseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (i InfantryCitySelfDefenseTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i InfantryCitySelfDefenseTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i InfantryCitySelfDefenseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InfantryCitySelfDefenseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InfantryCitySelfDefenseTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)
	// 恢复我军群体（2人）兵力（治疗率92%，受智力影响），
	// 并使其获得休整状态（每回合恢复一次兵力，治疗率62%），持续1回合
	pairGenerals := util.GetPairGeneralsTwoArr(i.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		resumeNum := cast.ToInt64(pairGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.92)
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  i.tacticsParams,
			ProduceGeneral: pairGeneral,
			SufferGeneral:  pairGeneral,
			ResumeNum:      resumeNum,
			TacticId:       i.Id(),
		})
		//施加状态
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     i.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Rest,
					TacticId:   i.Id(),
				}) {
					restResumeNum := cast.ToInt64(revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.62)
					util.ResumeSoldierNum(&util.ResumeParams{
						Ctx:            ctx,
						TacticsParams:  i.tacticsParams,
						ProduceGeneral: revokeGeneral,
						SufferGeneral:  revokeGeneral,
						ResumeNum:      restResumeNum,
						TacticId:       i.Id(),
					})
				}

				return revokeResp
			})
		}
	}
}

func (i InfantryCitySelfDefenseTactic) IsTriggerPrepare() bool {
	return false
}
