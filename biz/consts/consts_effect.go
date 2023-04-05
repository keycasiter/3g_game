package consts

// 负面效果
type DebuffEffectType int

const (
	//效果施加
	DebuffEffectType_Unknow                        DebuffEffectType = 0  //未知
	DebuffEffectType_Methysis                      DebuffEffectType = 1  //中毒
	DebuffEffectType_Firing                        DebuffEffectType = 2  //灼烧
	DebuffEffectType_Defect                        DebuffEffectType = 3  //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm                     DebuffEffectType = 4  //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos                         DebuffEffectType = 5  //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoStrategy                    DebuffEffectType = 6  //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth                    DebuffEffectType = 7  //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack                   DebuffEffectType = 8  //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove     DebuffEffectType = 9  //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove   DebuffEffectType = 10 //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce      DebuffEffectType = 11 //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce    DebuffEffectType = 12 //造成谋略伤害减少
	DebuffEffectType_CanNotGeneralAttack           DebuffEffectType = 13 //无法普通攻击
	DebuffEffectType_CancelWeapon                  DebuffEffectType = 14 //缴械（无法普通攻击）
	DebuffEffectType_DecrForce                     DebuffEffectType = 15 //降低武力
	DebuffEffectType_DecrIntelligence              DebuffEffectType = 16 //降低智力
	DebuffEffectType_DecrCommand                   DebuffEffectType = 17 //降低统率
	DebuffEffectType_DecrSpeed                     DebuffEffectType = 18 //降低速度
	DebuffEffectType_CleverStrategyAndShrewdTactic DebuffEffectType = 19 //神机妙算
	DebuffEffectType_InterlockedStratagems         DebuffEffectType = 20 //铁索连环
)

func (b DebuffEffectType) String() string {
	switch b {
	case DebuffEffectType_Methysis:
		return "中毒"
	case DebuffEffectType_Firing:
		return "灼烧"
	case DebuffEffectType_PoorHealth:
		return "虚弱"
	case DebuffEffectType_NoStrategy:
		return "计穷"
	case DebuffEffectType_Chaos:
		return "混乱"
	case DebuffEffectType_LaunchWeaponDamageDeduce:
		return "造成兵刃伤害降低"
	case DebuffEffectType_LaunchStrategyDamageDeduce:
		return "造成谋略伤害降低"
	case DebuffEffectType_InterlockedStratagems:
		return "铁索连环"
	case DebuffEffectType_CanNotGeneralAttack:
		return "无法攻击"
	}
	return ""
}

// 增益效果
type BuffEffectType int

const (
	//效果施加
	BuffEffectType_Unknow                        BuffEffectType = 0
	BuffEffectType_Evade                         BuffEffectType = 1  //规避
	BuffEffectType_EnhanceWeapon                 BuffEffectType = 2  //会心
	BuffEffectType_EnhanceStrategy               BuffEffectType = 3  //奇谋
	BuffEffectType_GroupAttack                   BuffEffectType = 4  //群攻
	BuffEffectType_FirstAttack                   BuffEffectType = 5  //先攻
	BuffEffectType_Rest                          BuffEffectType = 6  //休整
	BuffEffectType_Defend                        BuffEffectType = 7  //抵御
	BuffEffectType_ContinuousAttack              BuffEffectType = 8  //连击
	BuffEffectType_StrikeBack                    BuffEffectType = 9  //反击
	BuffEffectType_TacticsTriggerImprove         BuffEffectType = 10 //战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove     BuffEffectType = 11 //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove   BuffEffectType = 12 //造成谋略伤害增加
	BuffEffectType_SufferWeaponDamageDeduce      BuffEffectType = 13 //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce    BuffEffectType = 14 //受到谋略伤害减少
	BuffEffectType_IncrForce                     BuffEffectType = 15 //增加武力
	BuffEffectType_IncrIntelligence              BuffEffectType = 16 //增加智力
	BuffEffectType_IncrCommand                   BuffEffectType = 17 //增加统率
	BuffEffectType_IncrSpeed                     BuffEffectType = 18 //增加速度
	BuffEffectType_EmergencyTreatment            BuffEffectType = 19 //急救
	BuffEffectType_Charming                      BuffEffectType = 20 //魅惑
	BuffEffectType_AppeaseArmyAndPeople_Prepare  BuffEffectType = 21 //抚辑军民「预备」
	BuffEffectType_ThreeDaysOfSeparation_Prepare BuffEffectType = 22 //士别三日「预备」
	BuffEffectType_SeizeTheSoul                  BuffEffectType = 23 //夺魂挟魄
)

func (b BuffEffectType) String() string {
	switch b {
	case BuffEffectType_Evade:
		return "规避"
	case BuffEffectType_EnhanceWeapon:
		return "会心"
	case BuffEffectType_FirstAttack:
		return "先攻"
	case BuffEffectType_AppeaseArmyAndPeople_Prepare:
		return "抚辑军民「预备」"
	case BuffEffectType_ThreeDaysOfSeparation_Prepare:
		return "士别三日「预备」"
	case BuffEffectType_IncrForce:
		return "武力增加"
	case BuffEffectType_IncrIntelligence:
		return "智力增加"
	case BuffEffectType_IncrSpeed:
		return "速度增加"
	case BuffEffectType_IncrCommand:
		return "统率增加"
	}
	return ""
}
