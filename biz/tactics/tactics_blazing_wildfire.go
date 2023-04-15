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

// 火炽原燎
// 发动概率50%
// 对敌军群体(2-3人)施加灼烧状态，每回合持续造成伤害(伤害率56%，受智力影响)，持续2回合；
// 若目标已有灼烧状态则造成兵刃攻击(伤害率118%)
type BlazingWildfireTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BlazingWildfireTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BlazingWildfireTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BlazingWildfireTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BlazingWildfireTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.5
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

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//对敌军群体(2-3人)施加灼烧状态，每回合持续造成伤害(伤害率56%，受智力影响)，持续2回合；
	//找到敌军2或3人
	enemyGeneralMap := util.GetEnemyGeneralsTwoOrThreeMap(b.tacticsParams)
	for _, sufferGeneral := range enemyGeneralMap {
		//若目标已有灼烧状态则造成兵刃攻击(伤害率118%)
		//判断当前被攻击武将是否有灼烧状态
		if util.DeBuffEffectContains(sufferGeneral, consts.DebuffEffectType_Firing) {
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.18)
			util.TacticDamage(&util.TacticDamageParam{
				TacticsParams: b.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: sufferGeneral,
				Damage:        dmg,
				TacticName:    b.Name(),
			})

			//灼烧次数-1
			if !util.TacticsDebuffEffectCountWrapDecr(ctx, sufferGeneral, consts.DebuffEffectType_Firing, 1) {
				return
			}
		} else {
			//施加灼烧状态，每回合持续造成伤害(伤害率56%，受智力影响)，持续2回合
			if !util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_Firing, 1.0) {
				return
			}
			//次数
			if !util.TacticsDebuffEffectCountWrapIncr(ctx, sufferGeneral, consts.DebuffEffectType_Firing, 2, 2, false) {
				return
			}

			//注册伤害效果
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerGeneral := params.CurrentGeneral
				triggerResp := &vo.TacticsTriggerResult{}

				//剩余次数判断
				if !util.TacticsDebuffEffectCountWrapDecr(ctx, triggerGeneral, consts.DebuffEffectType_Firing, 1) {
					return triggerResp
				}

				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					triggerGeneral.BaseInfo.Name,
					b.Name(),
					consts.DebuffEffectType_Firing,
				)
				dmgNum := cast.ToInt64(0.56 * triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: b.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: triggerGeneral,
					Damage:        dmgNum,
					TacticName:    b.Name(),
				})

				return triggerResp
			})
		}
	}
}
