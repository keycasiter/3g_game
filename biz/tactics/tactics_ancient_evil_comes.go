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

// 古之恶来
// 我军主将即将受到普通攻击时，自身会对攻击者进行一次猛击（伤害率80%）并使其造成兵刃伤害降低18%，持续1回合，
// 随后为我军主将承担此次普通攻击
type AncientEvilComesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AncientEvilComesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.0
	return a
}

func (a AncientEvilComesTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	//我军主将即将受到普通攻击时，自身会对攻击者进行一次猛击（伤害率80%）并使其造成兵刃伤害降低18%，持续1回合，
	//随后为我军主将承担此次普通攻击

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//找到我军主将
	masterGeneral := util.GetPairMasterGeneral(currentGeneral, a.tacticsParams)
	util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_AncientEvilComes_Prepare, &vo.EffectHolderParams{
		EffectRate:     1.0,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	})

	util.TacticsTriggerWrapRegister(masterGeneral, consts.BattleAction_SufferGeneralAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.SufferAttackGeneral
		attackGeneral := params.AttackGeneral

		hlog.CtxInfof(ctx, "[%s]执行来自[%s]的「%v」效果",
			triggerGeneral.BaseInfo.Name,
			currentGeneral.BaseInfo.Name,
			consts.BuffEffectType_AncientEvilComes_Prepare,
		)

		//猛击（伤害率80%）
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     a.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     attackGeneral,
			DamageImproveRate: 0.8,
			DamageType:        consts.DamageType_Weapon,
			TacticName:        a.Name(),
		})
		//造成兵刃伤害降低18%，持续1回合
		if util.DebuffEffectWrapSet(ctx, attackGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:  0.18,
			EffectRound: 1,
			FromTactic:  a.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(attackGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeGeneral := params.CurrentGeneral
				revokeResp := &vo.TacticsTriggerResult{}

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
					TacticId:   a.Id(),
				})

				return revokeResp
			})
		}
		//随后为我军主将承担此次普通攻击
		masterGeneral.HelpByGeneral = currentGeneral

		return triggerResp
	})
}

func (a AncientEvilComesTactic) Id() consts.TacticId {
	return consts.AncientEvilComes
}

func (a AncientEvilComesTactic) Name() string {
	return "古之恶来"
}

func (a AncientEvilComesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AncientEvilComesTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AncientEvilComesTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AncientEvilComesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a AncientEvilComesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AncientEvilComesTactic) Execute() {

}

func (a AncientEvilComesTactic) IsTriggerPrepare() bool {
	return false
}

func (a AncientEvilComesTactic) SetTriggerPrepare(triggerPrepare bool) {
}
