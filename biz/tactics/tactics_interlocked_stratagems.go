package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 连环计
// 发动概率35%
// 准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，
// 并发动谋略攻击（伤害率156%，受智力影响）
type InterlockedStratagemsTactic struct {
	tacticsParams *model.TacticsParams
}

func (i InterlockedStratagemsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	return i
}

func (i InterlockedStratagemsTactic) Prepare() {
	return
}

func (i InterlockedStratagemsTactic) Id() consts.TacticId {
	return consts.InterlockedStratagems
}

func (i InterlockedStratagemsTactic) Name() string {
	return "连环计"
}

func (i InterlockedStratagemsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i InterlockedStratagemsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i InterlockedStratagemsTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	//发动概率35%
	if !util.GenerateRate(0.35) {
		return
	} else {
		hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
			currentGeneral.BaseInfo.Name,
			i.Name(),
		)
		//准备一回合，对敌军全体释放铁索连环，使其任一目标受到伤害时会反馈15%（受智力影响）伤害给其他单位，持续2回合，

		//并发动谋略攻击（伤害率156%，受智力影响）
	}
}
