package consts

// 负面效果
type DebuffEffectType int

const (
	//效果施加
	DebuffEffectType_Unknow                      DebuffEffectType = 0  //未知
	DebuffEffectType_Methysis                    DebuffEffectType = 1  //中毒
	DebuffEffectType_Firing                      DebuffEffectType = 2  //灼烧
	DebuffEffectType_Defect                      DebuffEffectType = 3  //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm                   DebuffEffectType = 4  //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos                       DebuffEffectType = 5  //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoStrategy                  DebuffEffectType = 6  //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth                  DebuffEffectType = 7  //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack                 DebuffEffectType = 8  //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove   DebuffEffectType = 9  //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove DebuffEffectType = 10 //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce    DebuffEffectType = 11 //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce  DebuffEffectType = 12 //造成谋略伤害减少
	DebuffEffectType_CanNotGeneralAttack         DebuffEffectType = 13 //无法普通攻击
	DebuffEffectType_CancelWeapon                DebuffEffectType = 14 //缴械（无法普通攻击）
	DebuffEffectType_DecrForce                   DebuffEffectType = 15 //降低武力
	DebuffEffectType_DecrIntelligence            DebuffEffectType = 16 //降低智力
	DebuffEffectType_DecrCommand                 DebuffEffectType = 17 //降低统率
	DebuffEffectType_DecrSpeed                   DebuffEffectType = 18 //降低速度

	//效果消失
	DebuffEffectType_Methysis_Disappear                    DebuffEffectType = 101 //中毒
	DebuffEffectType_Firing_Disappear                      DebuffEffectType = 102 //灼烧
	DebuffEffectType_Defect_Disappear                      DebuffEffectType = 103 //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm_Disappear                   DebuffEffectType = 104 //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos_Disappear                       DebuffEffectType = 105 //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoStrategy_Disappear                  DebuffEffectType = 106 //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth_Disappear                  DebuffEffectType = 107 //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack_Disappear                 DebuffEffectType = 108 //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove_Disappear   DebuffEffectType = 109 //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove_Disappear DebuffEffectType = 110 //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce_Disappear    DebuffEffectType = 111 //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce_Disappear  DebuffEffectType = 112 //造成谋略伤害减少
	DebuffEffectType_CanNotGeneralAttack_Disappear         DebuffEffectType = 113 //无法普通攻击
	DebuffEffectType_CancelWeapon_Disappear                DebuffEffectType = 114 //缴械（无法普通攻击）
	DebuffEffectType_DecrForce_Disappear                   DebuffEffectType = 115 //降低武力
	DebuffEffectType_DecrIntelligence_Disappear            DebuffEffectType = 116 //降低智力
	DebuffEffectType_DecrCommand_Disappear                 DebuffEffectType = 117 //降低统率
	DebuffEffectType_DecrSpeed_Disappear                   DebuffEffectType = 118 //降低速度
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
	}
	return ""
}

// 增益效果
type BuffEffectType int

const (
	//效果施加
	BuffEffectType_Unknow                      BuffEffectType = 0
	BuffEffectType_Evade                       BuffEffectType = 1  //规避
	BuffEffectType_EnhanceWeapon               BuffEffectType = 2  //会心
	BuffEffectType_EnhanceStrategy             BuffEffectType = 3  //奇谋
	BuffEffectType_GroupAttack                 BuffEffectType = 4  //群攻
	BuffEffectType_FirstAttack                 BuffEffectType = 5  //先攻
	BuffEffectType_Rest                        BuffEffectType = 6  //休整
	BuffEffectType_Defend                      BuffEffectType = 7  //抵御
	BuffEffectType_ContinuousAttack            BuffEffectType = 8  //连击
	BuffEffectType_StrikeBack                  BuffEffectType = 9  //反击
	BuffEffectType_TacticsTriggerImprove       BuffEffectType = 10 //战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove   BuffEffectType = 11 //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove BuffEffectType = 12 //造成谋略伤害增加
	BuffEffectType_SufferWeaponDamageDeduce    BuffEffectType = 13 //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce  BuffEffectType = 14 //受到谋略伤害减少
	BuffEffectType_IncrForce                   BuffEffectType = 15 //增加武力
	BuffEffectType_IncrIntelligence            BuffEffectType = 16 //增加智力
	BuffEffectType_IncrCommand                 BuffEffectType = 17 //增加统率
	BuffEffectType_IncrSpeed                   BuffEffectType = 18 //增加速度

	//效果消失
	BuffEffectType_Evade_Disappear                       BuffEffectType = 101 //规避
	BuffEffectType_EnhanceWeapon_Disappear               BuffEffectType = 102 //会心
	BuffEffectType_EnhanceStrategy_Disappear             BuffEffectType = 103 //奇谋
	BuffEffectType_GroupAttack_Disappear                 BuffEffectType = 104 //群攻
	BuffEffectType_FirstAttack_Disappear                 BuffEffectType = 105 //先攻
	BuffEffectType_Rest_Disappear                        BuffEffectType = 106 //休整
	BuffEffectType_Defend_Disappear                      BuffEffectType = 107 //抵御
	BuffEffectType_ContinuousAttack_Disappear            BuffEffectType = 108 //连击
	BuffEffectType_StrikeBack_Disappear                  BuffEffectType = 109 //反击
	BuffEffectType_TacticsTriggerImprove_Disappear       BuffEffectType = 110 //战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove_Disappear   BuffEffectType = 111 //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove_Disappear BuffEffectType = 112 //造成谋略伤害增加
	BuffEffectType_SufferWeaponDamageDeduce_Disappear    BuffEffectType = 113 //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce_Disappear  BuffEffectType = 114 //受到谋略伤害减少
	BuffEffectType_IncrForce_Disappear                   BuffEffectType = 115 //增加武力
	BuffEffectType_IncrIntelligence_Disappear            BuffEffectType = 116 //增加智力
	BuffEffectType_IncrCommand_Disappear                 BuffEffectType = 117 //增加统率
	BuffEffectType_IncrSpeed_Disappear                   BuffEffectType = 118 //增加速度
)

func (b BuffEffectType) String() string {
	switch b {
	case BuffEffectType_Evade:
		return "规避"
	case BuffEffectType_EnhanceWeapon:
		return "会心"
	}
	return ""
}
