package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 勇烈持重
// 受到伤害时，有35%概率净化自己负面效果，同时使随机敌军单体进入震慑状态，持续1回合，该效果每回合最多触发一次
// 被动 100%
type BraveAndResoluteTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	isTriggered   bool
}

func (b BraveAndResoluteTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BraveAndResoluteTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	//每回合清理触发判定
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		b.isTriggered = false

		return triggerResp
	})

	//注册触发器
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		attackGeneral := params.AttackGeneral

		//该效果每回合最多触发一次
		if b.isTriggered {
			return triggerResp
		}

		//施加效果
		b.isTriggered = true
		if util.DebuffEffectWrapSet(ctx, attackGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
			EffectRound: 1,
			FromTactic:  b.Id(),
		}).IsSuccess {
			//注册消失效果
			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    attackGeneral,
				EffectType: consts.DebuffEffectType_Awe,
				TacticId:   b.Id(),
			})
		}

		return triggerResp
	})
}

func (b BraveAndResoluteTactic) Id() consts.TacticId {
	return consts.BraveAndResolute
}

func (b BraveAndResoluteTactic) Name() string {
	return "勇烈持重"
}

func (b BraveAndResoluteTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BraveAndResoluteTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BraveAndResoluteTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BraveAndResoluteTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BraveAndResoluteTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveAndResoluteTactic) Execute() {
}

func (b BraveAndResoluteTactic) IsTriggerPrepare() bool {
	return false
}
