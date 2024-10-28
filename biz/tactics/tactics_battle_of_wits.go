package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 斗智
// 准备1回合，对敌军群体（2人）造成谋略攻击（伤害率155%）
type BattleOfWitsTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (b BattleOfWitsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.3
	return b
}

func (b BattleOfWitsTactic) Prepare() {
}

func (b BattleOfWitsTactic) Id() consts.TacticId {
	return consts.BattleOfWits
}

func (b BattleOfWitsTactic) Name() string {
	return "斗智"
}

func (b BattleOfWitsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BattleOfWitsTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BattleOfWitsTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BattleOfWitsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BattleOfWitsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BattleOfWitsTactic) Execute() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	b.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//准备1回合，对敌军群体（2人）造成谋略攻击（伤害率155%）
	//注册发动效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		//准备回合释放
		if currentRound+2 == triggerRound {
			b.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if b.isTriggered {
				return triggerResp
			} else {
				b.isTriggered = true
			}

			//敌军群体（2人）
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, b.tacticsParams)
			for _, sufferGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     b.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     sufferGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: 1.55,
					TacticId:          b.Id(),
					TacticName:        b.Name(),
				})
			}
		}
		return triggerResp
	})
}

func (b BattleOfWitsTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
