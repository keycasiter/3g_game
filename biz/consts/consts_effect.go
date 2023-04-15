package consts

// 负面效果
type DebuffEffectType int

var (
	SupprtCleanDebuffEffectMap = map[DebuffEffectType]bool{
		DebuffEffectType_Methysis:     true,
		DebuffEffectType_Firing:       true,
		DebuffEffectType_Defect:       true,
		DebuffEffectType_Sandstorm:    true,
		DebuffEffectType_Chaos:        true,
		DebuffEffectType_NoStrategy:   true,
		DebuffEffectType_PoorHealth:   true,
		DebuffEffectType_WaterAttack:  true,
		DebuffEffectType_CancelWeapon: true,
	}

	SupprtCleanBuffEffectMap = map[BuffEffectType]bool{
		BuffEffectType_SuppressChokesAndPreventRefusals_Prepare: true,
		BuffEffectType_SeizeTheSoul:                             true,
	}
)

const (
	//效果施加
	DebuffEffectType_Unknow                      DebuffEffectType = iota //未知
	DebuffEffectType_Methysis                                            //中毒
	DebuffEffectType_Firing                                              //灼烧
	DebuffEffectType_Defect                                              //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm                                           //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos                                               //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoStrategy                                          //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth                                          //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack                                         //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove                           //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove                         //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce                            //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce                          //造成谋略伤害减少
	DebuffEffectType_CanNotGeneralAttack                                 //无法普通攻击
	DebuffEffectType_CancelWeapon                                        //缴械（无法普通攻击）

	DebuffEffectType_DecrForce        //降低武力
	DebuffEffectType_DecrIntelligence //降低智力
	DebuffEffectType_DecrCommand      //降低统率
	DebuffEffectType_DecrSpeed        //降低速度

	DebuffEffectType_CleverStrategyAndShrewdTactic //神机妙算
	DebuffEffectType_InterlockedStratagems         //铁索连环
	DebuffEffectType_SeizeTheSoul                  //夺魂挟魄

	DebuffEffectType_BraveAmbition_DecrForce        //（义胆雄心）降低武力
	DebuffEffectType_BraveAmbition_DecrIntelligence //（义胆雄心）降低智力
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
	case DebuffEffectType_Defect:
		return "叛逃"
	case DebuffEffectType_LaunchWeaponDamageDeduce:
		return "造成兵刃伤害降低"
	case DebuffEffectType_LaunchStrategyDamageDeduce:
		return "造成谋略伤害降低"
	case DebuffEffectType_InterlockedStratagems:
		return "铁索连环"
	case DebuffEffectType_CanNotGeneralAttack:
		return "无法攻击"
	case DebuffEffectType_SeizeTheSoul:
		return "夺魂挟魄"
	case DebuffEffectType_CleverStrategyAndShrewdTactic:
		return "神机妙算"
	case DebuffEffectType_BraveAmbition_DecrForce:
		return "武力降低"
	case DebuffEffectType_BraveAmbition_DecrIntelligence:
		return "智力降低"
	case DebuffEffectType_DecrForce:
		return "武力降低"
	case DebuffEffectType_DecrIntelligence:
		return "智力降低"
	case DebuffEffectType_DecrCommand:
		return "统率降低"
	case DebuffEffectType_DecrSpeed:
		return "速度降低"
	}
	return ""
}

// 增益效果
type BuffEffectType int

const (
	//效果施加
	BuffEffectType_Unknow                       BuffEffectType = iota
	BuffEffectType_Evade                                       //规避
	BuffEffectType_EnhanceWeapon                               //会心
	BuffEffectType_EnhanceStrategy                             //奇谋
	BuffEffectType_GroupAttack                                 //群攻
	BuffEffectType_FirstAttack                                 //先攻
	BuffEffectType_Rest                                        //休整
	BuffEffectType_Defend                                      //抵御
	BuffEffectType_ContinuousAttack                            //连击
	BuffEffectType_StrikeBack                                  //反击
	BuffEffectType_TacticsActiveTriggerImprove                 //主动战法发动率提升
	BuffEffectType_TacticsPassiveTriggerImprove                //被动战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove                   //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove                 //造成谋略伤害增加

	BuffEffectType_TacticsActiveTriggerDecr                 //主动战法发动率降低
	BuffEffectType_TacticsPassiveTriggerDecr                //被动战法发动率降低
	BuffEffectType_SufferWeaponDamageDeduce                 //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce               //受到谋略伤害减少
	BuffEffectType_IncrForce                                //增加武力
	BuffEffectType_IncrIntelligence                         //增加智力
	BuffEffectType_IncrCommand                              //增加统率
	BuffEffectType_IncrSpeed                                //增加速度
	BuffEffectType_EmergencyTreatment                       //急救
	BuffEffectType_Charming                                 //魅惑
	BuffEffectType_AppeaseArmyAndPeople_Prepare             //抚辑军民「预备」
	BuffEffectType_ThreeDaysOfSeparation_Prepare            //士别三日「预备」
	BuffEffectType_SeizeTheSoul                             //夺魂挟魄
	BuffEffectType_BraveAmbition_Prepare                    //义胆雄心「预备」
	BuffEffectType_HuangTianDangLi                          //黄天当立
	BuffEffectType_SuppressChokesAndPreventRefusals_Prepare //镇扼防拒「预备」
	BuffEffectType_Intervene                                //援护
)

func (b BuffEffectType) String() string {
	switch b {
	case BuffEffectType_LaunchWeaponDamageImprove:
		return "造成兵刃伤害增加"
	case BuffEffectType_LaunchStrategyDamageImprove:
		return "造成谋略伤害增加"
	case BuffEffectType_Evade:
		return "规避"
	case BuffEffectType_EnhanceWeapon:
		return "会心"
	case BuffEffectType_FirstAttack:
		return "先攻"
	case BuffEffectType_AppeaseArmyAndPeople_Prepare:
		return "抚辑军民[预备]"
	case BuffEffectType_ThreeDaysOfSeparation_Prepare:
		return "士别三日[预备]"
	case BuffEffectType_IncrForce:
		return "武力增加"
	case BuffEffectType_IncrIntelligence:
		return "智力增加"
	case BuffEffectType_IncrSpeed:
		return "速度增加"
	case BuffEffectType_IncrCommand:
		return "统率增加"
	case BuffEffectType_SufferWeaponDamageDeduce:
		return "受到兵刃伤害降低"
	case BuffEffectType_SufferStrategyDamageDeduce:
		return "受到谋略伤害降低"
	case BuffEffectType_BraveAmbition_Prepare:
		return "义胆雄心[预备]"
	case BuffEffectType_HuangTianDangLi:
		return "黄天当立"
	case BuffEffectType_Charming:
		return "魅惑"
	case BuffEffectType_SuppressChokesAndPreventRefusals_Prepare:
		return "镇扼防拒[预备]"
	case BuffEffectType_Intervene:
		return "援护"
	case BuffEffectType_Rest:
		return "休整"
	}
	return ""
}
