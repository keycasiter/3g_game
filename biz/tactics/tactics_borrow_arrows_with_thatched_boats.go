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

//草船借箭
//主动战法 65%
//移除我军群体(2-3人)负面效果，并使我军群体(2人)获得急救状态，每次受到伤害时有70%几率回复一定兵力（伤害量的28%，受统率影响）
//持续2回合，该战法发动后会进入1回合冷却
type BorrowArrowsWithThatchedBoatsTactic struct {
	tacticsParams model.TacticsParams
}

func (b BorrowArrowsWithThatchedBoatsTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
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

	//65%概率
	if !util.GenerateRate(0.65) {
		return
	}
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		b.tacticsParams.CurrentGeneral.BaseInfo.Name,
		b.Name(),
	)
	//移除我军群体(2-3人)负面效果
	twoOrThreeGenerals := util.GetPairGeneralsTwoOrThreeMap(b.tacticsParams)
	for _, general := range twoOrThreeGenerals {
		util.DebuffEffectClean(ctx, general)
	}
	//并使我军群体(2人)获得急救状态，每次受到伤害时有70%几率回复一定兵力（伤害量的28%，受统率影响）,持续2回合
	twoGenerals := util.GetPairGeneralsTwoArr(b.tacticsParams)
	for _, general := range twoGenerals {
		//施加急救效果
		general.BuffEffectHolderMap[consts.BuffEffectType_EmergencyTreatment] = 1.0
		//持续2回合
		general.BuffEffectCountMap[consts.BuffEffectType_EmergencyTreatment][2] = 1.0
		hlog.CtxInfof(ctx, "[%s]的「急救」状态已施加", general.BaseInfo.Name)

		//注册触发效果
		util.TacticsTriggerWrapSet(general,
			consts.BattleAction_SufferAttack,
			func(params vo.TacticsTriggerParams) {
				if !util.GenerateRate(0.7) {
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「急救」效果因几率没有生效",
						general.BaseInfo.Name,
						b.Name(),
					)
					return
				} else {
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「急救」效果",
						general.BaseInfo.Name,
						b.Name(),
					)
					// TODO 受统率影响
					resumeNum := cast.ToInt64(cast.ToFloat64(params.CurrentDamage) * 0.28)
					hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d️️)",
						general.BaseInfo.Name,
						resumeNum,
						general.SoldierNum,
						general.SoldierNum+resumeNum,
					)
					general.SoldierNum += resumeNum
				}
			},
		)
	}

	//该战法发动后会进入1回合冷却
	currentGeneral.TacticsFrozenMap[b.Id()] = 1
	return
}