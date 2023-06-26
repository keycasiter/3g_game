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

// 江东猛虎
// 对敌军群体2人造成126%兵刃伤害，并嘲讽，持续2回合
// 主动，50%
type JiangdongTigerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JiangdongTigerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 0.5
	return j
}

func (j JiangdongTigerTactic) Prepare() {

}

func (j JiangdongTigerTactic) Id() consts.TacticId {
	return consts.JiangdongTiger
}

func (j JiangdongTigerTactic) Name() string {
	return "江东猛虎"
}

func (j JiangdongTigerTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (j JiangdongTigerTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JiangdongTigerTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JiangdongTigerTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (j JiangdongTigerTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (j JiangdongTigerTactic) Execute() {
	ctx := j.tacticsParams.Ctx
	currentGeneral := j.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		j.Name(),
	)

	// 对敌军群体2人造成126%兵刃伤害，并嘲讽，持续2回合
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, j.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.26)
		util.TacticDamage(&util.TacticDamageParam{
			TacticsParams: j.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: enemyGeneral,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      j.Id(),
			TacticName:    j.Name(),
		})
		//嘲讽
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     j.Id(),
			TauntByTarget:  currentGeneral,
			ProduceGeneral: nil,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Taunt,
					TacticId:   j.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (j JiangdongTigerTactic) IsTriggerPrepare() bool {
	return false
}
