package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 枪舞如风
// 使自身及友军单体获得2次防御，持续1回合，并使自身本回合发动普通攻击，对目标造成兵刃伤害（伤害率240%）及掠阵状态：
// 掠阵状态叠加两次时，移除掠阵状态并使自身提高40点武力，可叠加；
// 若与关兴同时出战，则友军单体必须选择关兴
// 主动，35%
type GunDanceLikeTheWindTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GunDanceLikeTheWindTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.35
	return g
}

func (g GunDanceLikeTheWindTactic) Prepare() {

}

func (g GunDanceLikeTheWindTactic) Id() consts.TacticId {
	return consts.GunDanceLikeTheWind
}

func (g GunDanceLikeTheWindTactic) Name() string {
	return "枪舞如风"
}

func (g GunDanceLikeTheWindTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (g GunDanceLikeTheWindTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GunDanceLikeTheWindTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GunDanceLikeTheWindTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (g GunDanceLikeTheWindTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GunDanceLikeTheWindTactic) Execute() {
	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)

	// 使自身及友军单体获得2次防御，持续1回合，并使自身本回合发动普通攻击，对目标造成兵刃伤害（伤害率240%）及掠阵状态：
	// 掠阵状态叠加两次时，移除掠阵状态并使自身提高40点武力，可叠加；
	// 若与关兴同时出战，则友军单体必须选择关兴
	//找到友军单体和自身
	generals := make([]*vo.BattleGeneral, 0)
	pairGeneral := util.GetPairOneGeneral(g.tacticsParams, currentGeneral)
	generals = append(generals, pairGeneral)
	generals = append(generals, currentGeneral)

	for _, general := range generals {
		//施加效果
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
			EffectTimes:    2,
			EffectRound:    1,
			FromTactic:     g.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_Defend,
					TacticId:   g.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (g GunDanceLikeTheWindTactic) IsTriggerPrepare() bool {
	return false
}
