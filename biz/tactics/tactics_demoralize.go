package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//挫锐
//战斗前3回合，使敌军单体进入虚弱状态，造成伤害时有65%几率无法造成伤害
type DemoralizeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DemoralizeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DemoralizeTactic) Prepare() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)

	//找到敌人单体
	enemyGeneral := util.GetEnemyOneGeneral(d.tacticsParams)
	//施加效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_PoorHealth, 0.9) {
		//注册消失效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound

			//第四回合消失
			if triggerRound == consts.Battle_Round_Fourth {
				util.DebuffEffectWrapRemove(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon)
			}

			return triggerResp
		})
	}

}

func (d DemoralizeTactic) Id() consts.TacticId {
	return consts.Demoralize
}

func (d DemoralizeTactic) Name() string {
	return "挫锐"
}

func (d DemoralizeTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DemoralizeTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DemoralizeTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DemoralizeTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (d DemoralizeTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DemoralizeTactic) Execute() {
	return
}
