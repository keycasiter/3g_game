package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//盛气凌敌
//战斗开始后前2回合，使敌军群体（2人）每回合都有90%的几率陷入缴械状态，无法进行普通攻击
type OverwhelmingTheEnemyWithVigourTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OverwhelmingTheEnemyWithVigourTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	o.tacticsParams = tacticsParams
	o.triggerRate = 1.0
	return o
}

func (o OverwhelmingTheEnemyWithVigourTactic) Prepare() {
	ctx := o.tacticsParams.Ctx
	currentGeneral := o.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		o.Name(),
	)
	//找到敌军2人
	enemyGenerals := util.GetEnemyGeneralsTwoArr(o.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		//施加效果
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_CancelWeapon, 0.9) {
			//注册消失效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerRound := params.CurrentRound

				//第三回合消失
				if triggerRound == consts.Battle_Round_Third {
					util.DebuffEffectWrapRemove(ctx, sufferGeneral, consts.DebuffEffectType_CancelWeapon)
				}

				return triggerResp
			})
		}
	}
}

func (o OverwhelmingTheEnemyWithVigourTactic) Id() consts.TacticId {
	return consts.OverwhelmingTheEnemyWithVigour
}

func (o OverwhelmingTheEnemyWithVigourTactic) Name() string {
	return "盛气凌敌"
}

func (o OverwhelmingTheEnemyWithVigourTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (o OverwhelmingTheEnemyWithVigourTactic) GetTriggerRate() float64 {
	return o.triggerRate
}

func (o OverwhelmingTheEnemyWithVigourTactic) SetTriggerRate(rate float64) {
	o.triggerRate = rate
}

func (o OverwhelmingTheEnemyWithVigourTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (o OverwhelmingTheEnemyWithVigourTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (o OverwhelmingTheEnemyWithVigourTactic) Execute() {
	return
}
