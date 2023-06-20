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

// 铁骑驱驰
// 战斗前2回合，使敌军全体处于遇袭状态（行动滞后），我军全体发动突击战法后，降低普通攻击目标15%统率，持续3回合，可叠加
type IronHorseDriveTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i IronHorseDriveTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 1.0
	return i
}

func (i IronHorseDriveTactic) Prepare() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)

	//战斗前2回合，使敌军全体处于遇袭状态（行动滞后），我军全体发动突击战法后，降低普通攻击目标15%统率，持续3回合，可叠加
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, i.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		//遇袭
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_BeAttacked, &vo.EffectHolderParams{
			EffectRound:    3,
			FromTactic:     i.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_BeAttacked,
					TacticId:   i.Id(),
				})

				return revokeResp
			})
		}
	}

	//我军全体发动突击战法后，降低普通攻击目标15%统率，持续3回合，可叠加
	pairGenerals := util.GetPairGeneralArr(i.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AssaultTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			sufferGeneral := i.tacticsParams.CurrentSufferGeneral

			decrVal := cast.ToInt64(sufferGeneral.BaseInfo.AbilityAttr.CommandBase * 0.15)
			if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
				EffectValue:    decrVal,
				EffectRound:    3,
				FromTactic:     i.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				//消失效果
				util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_DecrCommand,
						TacticId:   i.Id(),
					})

					return revokeResp
				})
			}

			return triggerResp
		})
	}
}

func (i IronHorseDriveTactic) Id() consts.TacticId {
	return consts.IronHorseDrive
}

func (i IronHorseDriveTactic) Name() string {
	return "铁骑驱驰"
}

func (i IronHorseDriveTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (i IronHorseDriveTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IronHorseDriveTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IronHorseDriveTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (i IronHorseDriveTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (i IronHorseDriveTactic) Execute() {
}

func (i IronHorseDriveTactic) IsTriggerPrepare() bool {
	return false
}
