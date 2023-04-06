package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 火炽原燎
// 发动概率50%
// 对敌军群体(2-3人)施加灼烧状态，每回合持续造成伤害(伤害率56%，受智力影响)，持续2回合；
// 若目标已有灼烧状态则造成兵刃攻击(伤害率118%)
type BlazingWildfireTactic struct {
	tacticsParams model.TacticsParams
}

func (b BlazingWildfireTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	return b
}

func (b BlazingWildfireTactic) Prepare() {
	return
}

func (b BlazingWildfireTactic) Id() consts.TacticId {
	return consts.BlazingWildfire
}

func (b BlazingWildfireTactic) Name() string {
	return "火炽原燎"
}

func (b BlazingWildfireTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BlazingWildfireTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BlazingWildfireTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//发动概率50%
	if !util.GenerateRate(0.5) {
		return
	}

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		b.tacticsParams.CurrentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//对敌军群体(2-3人)施加灼烧状态，每回合持续造成伤害(伤害率56%，受智力影响)，持续2回合；
	//找到敌军2或3人
	hitIdMap := util.GetEnemyGeneralsTwoOrThreeMap(b.tacticsParams)
	enemyGenerals := util.GetEnemyGeneralArr(b.tacticsParams)
	for idx, sufferGeneral := range enemyGenerals {
		if _, ok := hitIdMap[int64(idx)]; ok {
			//若目标已有灼烧状态则造成兵刃攻击(伤害率118%)
			//判断当前被攻击武将是否有灼烧状态
			if _, ok := sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_Firing]; ok {
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.18)
				dmg, origin, remain := util.TacticDamage(b.tacticsParams, currentGeneral, sufferGeneral, dmg)

				hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d️)",
					sufferGeneral.BaseInfo.Name,
					currentGeneral.BaseInfo.Name,
					b.Name(),
					dmg,
					origin,
					remain,
				)
			} else {
				//施加灼烧状态
				sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_Firing] = 1.0
				sufferGeneral.DeBuffEffectCountMap[consts.DebuffEffectType_Firing][2] = 0.56 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase
				hlog.CtxInfof(ctx, "[%s]的「灼烧」效果已施加")
			}
		}
	}
}
