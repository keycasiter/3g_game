package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 燕人咆哮
// 战斗第2、4回合，对敌军全体造成兵刃攻击（伤害率104%）
// 若目标处于缴械状态，则额外使目标统率降低50%，持续2回合
// 自身为主将时，降低统率效果额外对计穷状态的目标生效
// 被动，100%
type YanPeopleRoarTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (y YanPeopleRoarTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	y.tacticsParams = tacticsParams
	y.triggerRate = 1.0
	return y
}

func (y YanPeopleRoarTactic) Prepare() {
	ctx := y.tacticsParams.Ctx
	currentGeneral := y.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		y.Name(),
	)

	// 战斗第2、4回合，对敌军全体造成兵刃攻击（伤害率104%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound == consts.Battle_Round_Second ||
			triggerRound == consts.Battle_Round_Fourth {
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, y.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.04)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: y.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      y.Id(),
					TacticName:    y.Name(),
				})
				// 若目标处于缴械状态，则额外使目标统率降低50%，持续2回合
				// 自身为主将时，降低统率效果额外对计穷状态的目标生效
				check := util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_CancelWeapon)
				if currentGeneral.IsMaster && !check {
					check = util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_NoStrategy)
				}
				if check {
					decrVal := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.CommandBase * 0.5)
					if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
						EffectRound:    2,
						EffectValue:    decrVal,
						FromTactic:     y.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_DecrCommand,
								TacticId:   y.Id(),
							})

							return revokeResp
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (y YanPeopleRoarTactic) Id() consts.TacticId {
	return consts.YanPeopleRoar
}

func (y YanPeopleRoarTactic) Name() string {
	return "燕人咆哮"
}

func (y YanPeopleRoarTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (y YanPeopleRoarTactic) GetTriggerRate() float64 {
	return y.triggerRate
}

func (y YanPeopleRoarTactic) SetTriggerRate(rate float64) {
	y.triggerRate = rate
}

func (y YanPeopleRoarTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (y YanPeopleRoarTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (y YanPeopleRoarTactic) Execute() {
}

func (y YanPeopleRoarTactic) IsTriggerPrepare() bool {
	return false
}
