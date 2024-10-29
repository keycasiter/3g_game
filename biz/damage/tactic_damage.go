package damage

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

type TacticDamageLogic struct {
	param *TacticDamageParam
	funcs []func()
	err   error

	//中间变量
	damageNum        int64
	soldierNum       int64
	remainSoldierNum int64
	isEffect         bool
	isAvoidDamage    bool
}

func NewTacticDamageLogic(param *TacticDamageParam) *TacticDamageLogic {
	runCtx := &TacticDamageLogic{param: param}

	//伤害数据收集
	defer func() {
		runCtx.collectData()
	}()

	runCtx.funcs = []func(){
		//参数校验
		runCtx.checkParam,
		//特殊效果处理
		runCtx.specialEffectHandler,
		//前置触发器
		runCtx.triggerPreHandler,
		//伤害规避处理
		runCtx.avoidDamageHandler,
		//伤害效果处理
		runCtx.damageEffectHandler,
		//伤害分担处理
		runCtx.damageShareHandler,
		//伤害结算处理
		runCtx.damageCalculateHandler,
		//后置触发器
		runCtx.triggerPostHandler,
	}
	return runCtx
}

func (t *TacticDamageLogic) Process() (damageNum, soldierNum, remainSoldierNum int64, isEffect bool) {
	for _, f := range t.funcs {
		f()
		if t.err != nil {
			return 0, 0, 0, false
		}
		if t.isAvoidDamage {
			return 0, 0, 0, false
		}
	}
	return t.damageNum, t.soldierNum, t.remainSoldierNum, t.isEffect
}

func (t *TacticDamageLogic) checkParam() {
	if t.param.TacticId == 0 {
		panic(any("tacticId is nil"))
	}
	ctx := t.param.TacticsParams.Ctx

	//必填参数
	if t.param.AttackGeneral == nil ||
		t.param.SufferGeneral == nil ||
		t.param.DamageType == consts.DamageType_None ||
		(t.param.DamageImproveRate <= 0 && t.param.DamageDeduceRate <= 0) {
		hlog.CtxErrorf(ctx, "damage params err , attackGeneral:%s , sufferGeneral:%s ,damageType:%v",
			util.ToJsonString(ctx, t.param.AttackGeneral), util.ToJsonString(ctx, t.param.SufferGeneral), t.param.DamageType)
		panic(any("damage params err"))
	}
}

func (t *TacticDamageLogic) damageShareHandler() {
	param := t.param
	ctx := param.TacticsParams.Ctx
	tacticsParams := param.TacticsParams
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	damageType := param.DamageType
	tacticName := param.TacticName
	tacticId := param.TacticId
	isIgnoreDefend := param.IsIgnoreDefend

	//伤害分担
	if v, ok := util.BuffEffectGetAggrEffectRate(sufferGeneral, consts.BuffEffectType_ShareResponsibilityFor); ok {
		hlog.CtxInfof(ctx, "[%s]由于「%v」效果，本次攻击受到的伤害减少了%.2f%%",
			sufferGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
			v*100,
		)
		//分担伤害
		shareDmg := cast.ToInt64(cast.ToFloat64(t.damageNum) * v)
		//被分担后的伤害
		t.damageNum = cast.ToInt64(cast.ToFloat64(t.damageNum) * (1 - v))

		hlog.CtxInfof(ctx, "[%s]执行「%v」效果",
			sufferGeneral.ShareResponsibilityForByGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
		)
		if shareDmg > 0 {
			TacticDamage(&TacticDamageParam{
				TacticsParams:  tacticsParams,
				AttackGeneral:  attackGeneral,
				SufferGeneral:  sufferGeneral.ShareResponsibilityForByGeneral,
				TacticId:       tacticId,
				TacticName:     tacticName,
				IsIgnoreDefend: isIgnoreDefend,
				DamageType:     damageType,
			})
		}
	}
}

func (t *TacticDamageLogic) damageEffectHandler() {
	param := t.param
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	isIgnoreDefend := param.IsIgnoreDefend

	//兵力为空
	if sufferGeneral.SoldierNum == 0 {
		t.damageNum = 0
		t.soldierNum = 0
		t.remainSoldierNum = 0
		t.isEffect = false
		return
	}

	//****伤害提升效果****
	improveDmgRate := t.param.DamageImproveRate
	//自带主动战法伤害提升
	if attackGeneral.EquipTactics[0].Id == param.TacticId && consts.ActiveTacticsMap[attackGeneral.EquipTactics[0].Id] {
		if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_TacticsActiveWithSelfDamageImprove); ok {
			for _, effectParam := range effectParams {
				improveDmgRate += effectParam.EffectRate
			}
		}
	}
	//主动战法伤害提升
	if consts.ActiveTacticsMap[attackGeneral.EquipTactics[0].Id] {
		if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_TacticsActiveDamageImprove); ok {
			for _, effectParam := range effectParams {
				improveDmgRate += effectParam.EffectRate
			}
		}
	}
	//造成伤害提升：兵刃/谋略
	switch param.DamageType {
	case consts.DamageType_Weapon:
		if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove); ok {
			for _, effectParam := range effectParams {
				improveDmgRate += effectParam.EffectRate
			}
		}
	case consts.DamageType_Strategy:
		if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove); ok {
			for _, effectParam := range effectParams {
				improveDmgRate += effectParam.EffectRate
			}
		}
	}
	//藤甲兵灼烧效果
	if effectParams, ok := util.DeBuffEffectGet(sufferGeneral, consts.DebuffEffectType_Firing_TengJia); ok {
		if consts.FireTacticsMap[t.param.TacticId] {
			for _, effectParam := range effectParams {
				improveDmgRate += effectParam.EffectRate
			}
		}
	}

	//****伤害减少效果****
	//自带主动战法伤害减少
	deduceEffectRate := t.param.DamageDeduceRate
	if sufferGeneral.EquipTactics[0].Id == param.TacticId && consts.ActiveTacticsMap[sufferGeneral.EquipTactics[0].Id] {
		if effectParams, ok := util.BuffEffectGet(sufferGeneral, consts.BuffEffectType_SufferActiveTacticDamageDeduce); ok {
			for _, effectParam := range effectParams {
				deduceEffectRate += effectParam.EffectRate
			}
		}
	}
	//突击战法伤害减少
	if consts.AssaultTacticsMap[param.TacticId] {
		if effectParams, ok := util.BuffEffectGet(sufferGeneral, consts.BuffEffectType_SufferAssaultTacticDamageDeduce); ok {
			for _, effectParam := range effectParams {
				deduceEffectRate += effectParam.EffectRate
			}
		}
	}
	//造成伤害减少：兵刃/谋略
	switch param.DamageType {
	case consts.DamageType_Weapon:
		if effectParams, ok := util.BuffEffectGet(sufferGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce); ok {
			for _, effectParam := range effectParams {
				deduceEffectRate += effectParam.EffectRate
			}
		}
	case consts.DamageType_Strategy:
		if effectParams, ok := util.BuffEffectGet(sufferGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce); ok {
			for _, effectParam := range effectParams {
				deduceEffectRate += effectParam.EffectRate
			}
		}
	}

	//伤害减免最大90%
	if deduceEffectRate > 0.9 {
		deduceEffectRate = 0.9
	}

	//**伤害计算公式**
	//最终伤害 = 保底伤害 +（兵力基础伤害+属性差×等级差系数）
	//兵刃 保底伤害 +（兵力基础伤害+(武力-统率)×属性差）x 变量
	//谋略  保底伤害 +（兵力基础伤害+(智力-智力)×属性差）x 变量
	//普通攻击  保底伤害 +（兵力基础伤害+(武力-统率)×属性差）x 1
	attackAttr := float64(0)
	defendAttr := float64(0)
	switch param.DamageType {
	//兵刃伤害
	case consts.DamageType_Weapon:
		attackAttr = t.param.AttackGeneral.BaseInfo.AbilityAttr.ForceBase
		defendAttr = t.param.SufferGeneral.BaseInfo.AbilityAttr.CommandBase
	//谋略伤害
	case consts.DamageType_Strategy:
		attackAttr = t.param.AttackGeneral.BaseInfo.AbilityAttr.IntelligenceBase
		defendAttr = t.param.SufferGeneral.BaseInfo.AbilityAttr.IntelligenceBase
	}
	//是否无视防御
	if isIgnoreDefend {
		defendAttr = 0
	}

	t.damageNum =
		//保底伤害
		GetMinimumGuaranteeDmg(t.param.AttackGeneral.SoldierNum) +
			cast.ToInt64(
				//兵力基础伤害
				(cast.ToFloat64(GetBaseDmg(t.param.AttackGeneral.SoldierNum))+
					//属性差
					(attackAttr-defendAttr)*1.44)*
					//变量
					(1+improveDmgRate-deduceEffectRate))

	//最终伤害随机值
	t.damageNum = FluctuateDamage(t.damageNum)

	//特殊战法效果处理
	//刀出如霆
	if t.param.TacticId == consts.TheKnifeLikeThunderbolt {
		//由敌军全部武将平分
		pairGenerals := util.GetPairGeneralsNotSelf(t.param.TacticsParams, t.param.SufferGeneral)
		t.damageNum = t.damageNum / int64(1+len(pairGenerals))
	}

	//****伤害计算兜底****
	//兜底伤害为负的情况
	if t.damageNum < 0 {
		t.damageNum = 0
	}

	//兜底伤害大于剩余兵力情况
	t.soldierNum = sufferGeneral.SoldierNum
	if t.damageNum >= t.soldierNum {
		t.damageNum = t.soldierNum
	}
}

func (t *TacticDamageLogic) avoidDamageHandler() {
	param := t.param
	ctx := param.TacticsParams.Ctx
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral

	//是否可以造成伤害
	if !util.IsCanDamage(ctx, attackGeneral) {
		t.isAvoidDamage = true
		t.damageNum = 0
		return
	}

	//是否可以规避
	if util.IsCanEvade(ctx, sufferGeneral) {
		t.isAvoidDamage = true
		t.damageNum = 0
		return
	}

	//抵御效果判断
	if effectParams, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Defend]; ok {
		effectType := consts.BuffEffectType_Defend
		for idx, effectParam := range effectParams {
			if effectParam.EffectTimes > 0 {
				effectParam.EffectTimes--
				hlog.CtxInfof(ctx, "[%s]来自【%v】的「%v」效果，本次免疫伤害",
					sufferGeneral.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)

				//清除
				if effectParam.EffectTimes == 0 {
					sufferGeneral.BuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					//如果该效果绑定参数结构体为空，则顺便移除该效果
					if len(sufferGeneral.BuffEffectHolderMap[effectType]) == 0 {
						delete(sufferGeneral.BuffEffectHolderMap, effectType)
						hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
							sufferGeneral.BaseInfo.Name,
							effectType,
						)
					} else {
						hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
							sufferGeneral.BaseInfo.Name,
							effectParam.FromTactic,
							effectType,
						)
					}
				}

				//伤害抵御
				t.isAvoidDamage = true
				t.damageNum = 0
			}
		}
	}
}

func (t *TacticDamageLogic) specialEffectHandler() {
	param := t.param
	attackGeneral := param.AttackGeneral
	tacticName := param.TacticName

	//触发器禁用开关
	if tacticName == "连环计" && param.IsBanInterLockedEffect {
		return
	}

	//** 特殊战法效果 **
	//虎痴效果
	if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_TigerIdiot_Locked); ok {
		if len(effectParams) > 0 {
			effectParam := effectParams[0]
			param.SufferGeneral = effectParam.LockingTarget
		}
	}
	//弓腰姬
	if t.param.TacticId == consts.BowWaistConcubine {
		//自身拥有功能性增益状态时额外对其造成兵刃伤害（伤害率20%x状态数）
		buffEffectNum := util.BuffEffectContainsNum(attackGeneral)
		if buffEffectNum > 0 {
			buffEffectDesc := ""
			for buffEffectType, _ := range attackGeneral.BuffEffectHolderMap {
				buffEffectDesc += fmt.Sprintf("[%v]", buffEffectType)
			}
			t.param.DamageImproveRate += cast.ToFloat64(buffEffectNum) * 0.2
		}
	}
}

func (t *TacticDamageLogic) triggerPreHandler() {
	param := t.param
	tacticsParams := param.TacticsParams
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	damageType := param.DamageType

	// 「伤害开始」触发器
	if funcs, okk := attackGeneral.TacticsTriggerMap[consts.BattleAction_Damage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	// 「遭受伤害开始」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferDamage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	battleAction := consts.BattleAction_StrategyDamage
	if damageType == consts.DamageType_Weapon {
		battleAction = consts.BattleAction_WeaponDamage
	}
	if funcs, okk := attackGeneral.TacticsTriggerMap[battleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	sufferBattleAction := consts.BattleAction_SufferStrategyDamage
	if damageType == consts.DamageType_Weapon {
		sufferBattleAction = consts.BattleAction_SufferWeaponDamage
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[sufferBattleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_Tactic]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferTactic]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}

	//被伤害效果触发器
	//映射转换
	sufferEffectTriggerMapping := map[consts.TacticsType]consts.BattleAction{
		consts.TacticsType_Active:        consts.BattleAction_SufferActiveTactic,
		consts.TacticsType_Passive:       consts.BattleAction_SufferPassiveTactic,
		consts.TacticsType_Assault:       consts.BattleAction_SufferAssaultTactic,
		consts.TacticsType_Arm:           consts.BattleAction_SufferArmTactic,
		consts.TacticsType_Command:       consts.BattleAction_SufferCommandTactic,
		consts.TacticsType_TroopsTactics: consts.BattleAction_SufferTroopsTactic,
	}
	action := sufferEffectTriggerMapping[tacticsParams.TacticsType]
	if funcs, ok := sufferGeneral.TacticsTriggerMap[action]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
}

func (t *TacticDamageLogic) damageCalculateHandler() {
	//伤害结算
	t.param.SufferGeneral.SoldierNum -= t.damageNum
	t.remainSoldierNum = t.param.SufferGeneral.SoldierNum
}

func (t *TacticDamageLogic) triggerPostHandler() {
	param := t.param
	tacticsParams := param.TacticsParams
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	damageType := param.DamageType

	// 「伤害开始」触发器
	if funcs, okk := attackGeneral.TacticsTriggerMap[consts.BattleAction_DamageEnd]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	// 「遭受伤害结束」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferDamageEnd]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
				DamageType:     damageType,
			}
			f(params)
		}
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_TacticEnd]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//「战法攻击后」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferTacticEnd]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//「发动兵刃/谋略伤害结束」触发器
	battleAction := consts.BattleAction_StrategyDamageEnd
	if damageType == consts.DamageType_Weapon {
		battleAction = consts.BattleAction_WeaponDamageEnd
	}
	if funcs, okk := attackGeneral.TacticsTriggerMap[battleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	sufferBattleAction := consts.BattleAction_SufferStrategyDamageEnd
	if damageType == consts.DamageType_Weapon {
		sufferBattleAction = consts.BattleAction_SufferWeaponDamageEnd
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[sufferBattleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//被伤害效果后触发器
	//战法伤害触发器
	sufferEffectTriggerMapping := map[consts.TacticsType]consts.BattleAction{
		consts.TacticsType_Active:        consts.BattleAction_SufferActiveTacticEnd,
		consts.TacticsType_Passive:       consts.BattleAction_SufferPassiveTacticEnd,
		consts.TacticsType_Assault:       consts.BattleAction_SufferAssaultTacticEnd,
		consts.TacticsType_Arm:           consts.BattleAction_SufferArmTacticEnd,
		consts.TacticsType_Command:       consts.BattleAction_SufferCommandTacticEnd,
		consts.TacticsType_TroopsTactics: consts.BattleAction_SufferTroopsTacticEnd,
	}
	action := sufferEffectTriggerMapping[tacticsParams.TacticsType]
	if funcs, ok := sufferGeneral.TacticsTriggerMap[action]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
}

// 战法伤害数据收集
// @attack 攻击武将
// @suffer 被攻击武将
// @damage 伤害量
// @return 实际伤害/原兵力/剩余兵力
func (t *TacticDamageLogic) collectData() {
	param := t.param
	ctx := t.param.TacticsParams.Ctx
	tacticsParams := param.TacticsParams
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	tacticName := param.TacticName
	effectName := param.EffectName

	//统计数据
	attackGeneral.AccumulateTotalDamageNum += t.damageNum
	attackGeneral.TacticAccumulateDamageMap[param.TacticId] += t.damageNum
	attackGeneral.TacticAccumulateTriggerMap[param.TacticId] += 1
	sufferGeneral.LossSoldierNum += t.damageNum

	//计数
	switch param.DamageType {
	case consts.DamageType_Weapon:
		attackGeneral.ExecuteWeaponAttackNum++
		sufferGeneral.SufferExecuteWeaponAttackNum++
	case consts.DamageType_Strategy:
		attackGeneral.ExecuteStrategyAttackNum++
		sufferGeneral.SufferExecuteStrategyAttackNum++

		//统计上报
		util.TacticReport(tacticsParams,
			attackGeneral.BaseInfo.UniqueId,
			int64(param.TacticId),
			1,
			t.damageNum,
			0,
		)

		if effectName == "" {
			hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d)",
				sufferGeneral.BaseInfo.Name,
				attackGeneral.BaseInfo.Name,
				tacticName,
				t.damageNum,
				t.soldierNum,
				t.remainSoldierNum,
			)
		} else {
			hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】「%v」的伤害，损失了兵力%d(%d↘%d)",
				sufferGeneral.BaseInfo.Name,
				attackGeneral.BaseInfo.Name,
				tacticName,
				effectName,
				t.damageNum,
				t.soldierNum,
				t.remainSoldierNum,
			)
		}

		return
	}

}
func TacticDamage(param *TacticDamageParam) (damageNum, soldierNum, remainSoldierNum int64, isEffect bool) {
	return NewTacticDamageLogic(param).Process()
}
