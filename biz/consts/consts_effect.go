package consts

// 负面效果
type DebuffEffectType int

var (
	//可以被清除的负面效果
	SupprtCleanDebuffEffectMap = map[DebuffEffectType]bool{
		DebuffEffectType_Methysis:            true,
		DebuffEffectType_Firing:              true,
		DebuffEffectType_Defect:              true,
		DebuffEffectType_Sandstorm:           true,
		DebuffEffectType_Chaos:               true,
		DebuffEffectType_NoStrategy:          true,
		DebuffEffectType_PoorHealth:          true,
		DebuffEffectType_WaterAttack:         true,
		DebuffEffectType_CancelWeapon:        true,
		DebuffEffectType_CanNotGeneralAttack: true,
		DebuffEffectType_CanNotActiveTactic:  true,
	}
	//可以被清除的正面效果
	SupprtCleanBuffEffectMap = map[BuffEffectType]bool{
		BuffEffectType_SuppressChokesAndPreventRefusals_Prepare: true,
	}
)

const (
	//效果施加
	DebuffEffectType_Unknow                          DebuffEffectType = iota //未知
	DebuffEffectType_Methysis                                                //中毒
	DebuffEffectType_StrongMethysis                                          //猛毒
	DebuffEffectType_Firing                                                  //灼烧
	DebuffEffectType_Firing_TengJia                                          //灼烧[藤甲额外伤害]
	DebuffEffectType_Defect                                                  //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Escape                                                  //溃逃（受武力影响，无视防御）
	DebuffEffectType_Sandstorm                                               //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos                                                   //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_Awe                                                     //震慑（无法行动）
	DebuffEffectType_NoStrategy                                              //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth                                              //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack                                             //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove                               //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove                             //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce                                //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce                              //造成谋略伤害减少
	DebuffEffectType_SufferResumeDeduce                                      //受到治疗降低
	DebuffEffectType_CanNotGeneralAttack                                     //无法普通攻击
	DebuffEffectType_CanNotActiveTactic                                      //无法发动主动战法
	DebuffEffectType_CancelWeapon                                            //缴械（无法普通攻击）
	DebuffEffectType_Taunt                                                   //嘲讽（强制攻击目标）
	DebuffEffectType_ProhibitionTreatment                                    //禁疗（无法恢复兵力）
	DebuffEffectType_FalseReport                                             //伪报（禁用指挥和被动战法，无视洞察）
	DebuffEffectType_Capture                                                 //捕获（无法行动和造成伤害、禁用指挥和被动战法，进入禁疗状态、无法被友方武将选中，且无法被净化）
	DebuffEffectType_Break                                                   //破坏（携带的装备失效）
	DebuffEffectType_Provoking                                               //挑拨（强迫目标释放的战法选择自己）
	DebuffEffectType_BeAttacked                                              //遇袭（行动滞后）
	DebuffEffectType_Flaw                                                    //破绽
	DebuffEffectType_ShockingFourRealms_Prepare                              //震骇四境[准备]
	DebuffEffectType_AggregateStoneIntoGold_Exchange                         //聚石成金[交换]

	DebuffEffectType_DecrForce        //降低武力
	DebuffEffectType_DecrIntelligence //降低智力
	DebuffEffectType_DecrCommand      //降低统率
	DebuffEffectType_DecrSpeed        //降低速度

	DebuffEffectType_TacticsActiveTriggerDecr  //主动战法发动率降低
	DebuffEffectType_TacticsPassiveTriggerDecr //被动战法发动率降低

	DebuffEffectType_CleverStrategyAndShrewdTactic //神机妙算
	DebuffEffectType_InterlockedStratagems         //铁索连环
	DebuffEffectType_FireJointVenture_BurningCamp  //火烧连营[焚营]

	DebuffEffectType_LectureField                                                 //舌战群儒
	DebuffEffectType_TigerAnger                                                   //虎嗔
	DebuffEffectType_TigerAnger_Prepare                                           //虎嗔[预备]
	DebuffEffectType_RampartsOfMetalsAndAMoatOfHotWaterTactic_CanNotGeneralAttack //金城汤池[无法普通攻击]
)

func (b DebuffEffectType) String() string {
	switch b {
	case DebuffEffectType_AggregateStoneIntoGold_Exchange:
		return "聚石成金[交换]"
	case DebuffEffectType_Firing_TengJia:
		return "灼烧[藤甲额外伤害]"
	case DebuffEffectType_SufferResumeDeduce:
		return "受到治疗降低"
	case DebuffEffectType_ShockingFourRealms_Prepare:
		return "震骇四境[准备]"
	case DebuffEffectType_RampartsOfMetalsAndAMoatOfHotWaterTactic_CanNotGeneralAttack:
		return "金城汤池[无法普通攻击]"
	case DebuffEffectType_Flaw:
		return "破绽"
	case DebuffEffectType_StrongMethysis:
		return "猛毒"
	case DebuffEffectType_CanNotActiveTactic:
		return "无法发动主动战法"
	case DebuffEffectType_BeAttacked:
		return "遇袭"
	case DebuffEffectType_TigerAnger:
		return "虎嗔"
	case DebuffEffectType_TigerAnger_Prepare:
		return "虎嗔[预备]"
	case DebuffEffectType_SufferWeaponDamageImprove:
		return "受到兵刃伤害增加"
	case DebuffEffectType_SufferStrategyDamageImprove:
		return "受到谋略伤害增加"
	case DebuffEffectType_FireJointVenture_BurningCamp:
		return "火烧连营[焚营]"
	case DebuffEffectType_WaterAttack:
		return "水攻"
	case DebuffEffectType_Break:
		return "破坏"
	case DebuffEffectType_Sandstorm:
		return "沙暴"
	case DebuffEffectType_Awe:
		return "震慑"
	case DebuffEffectType_Escape:
		return "溃逃"
	case DebuffEffectType_ProhibitionTreatment:
		return "禁疗"
	case DebuffEffectType_TacticsPassiveTriggerDecr:
		return "被动战法发动率降低"
	case DebuffEffectType_TacticsActiveTriggerDecr:
		return "主动战法发动率降低"
	case DebuffEffectType_LectureField:
		return "舌战群儒"
	case DebuffEffectType_CancelWeapon:
		return "缴械"
	case DebuffEffectType_Taunt:
		return "嘲讽"
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
	case DebuffEffectType_CleverStrategyAndShrewdTactic:
		return "神机妙算"
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
	BuffEffectType_Unknow                              BuffEffectType = iota
	BuffEffectType_Evade                                              //规避
	BuffEffectType_EnhanceWeapon                                      //会心
	BuffEffectType_AttackHeart                                        //攻心
	BuffEffectType_EnhanceStrategy                                    //奇谋
	BuffEffectType_GroupAttack                                        //群攻
	BuffEffectType_FirstAttack                                        //先攻
	BuffEffectType_Rest                                               //休整
	BuffEffectType_Defend                                             //抵御
	BuffEffectType_ContinuousAttack                                   //连击
	BuffEffectType_StrikeBack                                         //反击
	BuffEffectType_Defection                                          //倒戈
	BuffEffectType_ShareResponsibilityFor                             //分担
	BuffEffectType_Insight                                            //洞察
	BuffEffectType_FightHard                                          //酣斗
	BuffEffectType_MustHit                                            //必中
	BuffEffectType_BreakFormation                                     //破阵
	BuffEffectType_TacticsActiveTriggerImprove                        //主动战法发动率提升
	BuffEffectType_TacticsActiveTriggerWithSelfImprove                //主动战法[自带]发动率提升
	BuffEffectType_TacticsActiveTriggerPrepareImprove                 //主动战法[准备战法]发动率提升
	BuffEffectType_TacticsActiveTriggerNoSelfImprove                  //主动战法[非自带]发动率提升
	BuffEffectType_TacticsActiveWithSelfDamageImprove                 //主动战法[自带]伤害提升
	BuffEffectType_TacticsActiveDamageImprove                         //主动战法伤害提升
	BuffEffectType_TacticsPassiveTriggerImprove                       //被动战法发动率提升
	BuffEffectType_TacticsAssaultTriggerImprove                       //突击战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove                          //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove                        //造成谋略伤害增加

	BuffEffectType_GeneralAttackDamageImprove                           //普通攻击伤害提升
	BuffEffectType_SufferWeaponDamageDeduce                             //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce                           //受到谋略伤害减少
	BuffEffectType_SufferActiveTacticDamageDeduce                       //受到[主动战法]伤害减少
	BuffEffectType_SufferAssaultTacticDamageDeduce                      //受到[突击战法]伤害减少
	BuffEffectType_IncrForce                                            //增加武力
	BuffEffectType_IncrIntelligence                                     //增加智力
	BuffEffectType_IncrCommand                                          //增加统率
	BuffEffectType_IncrSpeed                                            //增加速度
	BuffEffectType_EmergencyTreatment                                   //急救
	BuffEffectType_Charming                                             //魅惑
	BuffEffectType_AppeaseArmyAndPeople_Prepare                         //抚辑军民「预备」
	BuffEffectType_ThreeDaysOfSeparation_Prepare                        //士别三日「预备」
	BuffEffectType_BraveAmbition_Prepare                                //义胆雄心「预备」
	BuffEffectType_HuangTianDangLi                                      //黄天当立
	BuffEffectType_SuppressChokesAndPreventRefusals_Prepare             //镇扼防拒「预备」
	BuffEffectType_ClearEyedAndMalicious_Prepare                        //鹰视狼顾「预备」
	BuffEffectType_ClearEyedAndMalicious_ClearEyed_Prepare              //鹰视狼顾-鹰视「预备」
	BuffEffectType_UseMartialArtsToConnectWithGods_Prepare              //用武通神[预备]
	BuffEffectType_KillingMinisterInDream_Prepare                       //梦中弑臣[预备]
	BuffEffectType_BeFullyEquippedFor_Prepare                           //整装待发[预备]
	BuffEffectType_RideTigerHardToGetOff_Prepare                        //骑虎难下[预备]
	BuffEffectType_OutstandingTalent_Prepare                            //才器过人[预备]
	BuffEffectType_LowerBannersAndMuffleDrums_Prepare                   //偃旗息鼓[预备]
	BuffEffectType_Humility_Prepare                                     //谦让[预备]
	BuffEffectType_OnesResolveIsUnshaken_Prepare                        //矢志不移[预备]
	BuffEffectType_AbilityToRuleTheCountry_Prepare                      //经天纬地[预备]
	BuffEffectType_AncientEvilComes_Prepare                             //古之恶来[预备]
	BuffEffectType_CrowdMovesTenThousandCounts_Prepare                  //众动万计[预备]
	BuffEffectType_FireGodHeroStyle_Prepare                             //火神英风[预备]
	BuffEffectType_Intervene                                            //援护
	BuffEffectType_ActiveTactic_SkipPrepareRound                        //主动战法[跳过准备回合]
	BuffEffectType_GrazingArray                                         //掠阵
	BuffEffectType_HighWoodenPaddlesConnectedToTheCamp_Prepare          //高橹连营[预备]
	BuffEffectType_ThousandMileWalkingSingleRider_Prepare               //千里走单骑[预备]
	BuffEffectType_ImmunityChaos                                        //免疫混乱
	BuffEffectType_ImmunityCancelWeapon                                 //免疫缴械
	BuffEffectType_AccumulatePower                                      //蓄威
	BuffEffectType_Alert                                                //警戒
	BuffEffectType_ImmunitySandstorm                                    //免疫沙暴
	BuffEffectType_TigerIdiot_Locked                                    //虎痴[锁定]
	BuffEffectType_ThePostureOfAPhoenixWithAChickAndASweetPotato_Locked //鸱苕凤姿[锁定]
	BuffEffectType_TwelveWonderfulStrategies_Prepare                    //十二奇策[预备]
	BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare                 //以逸待劳[预备]
	BuffEffectType_WaitAtOnesEaseForTheFatigued_ImmunityControl         //以逸待劳[几率免疫控制]
)

func (b BuffEffectType) String() string {
	switch b {
	case BuffEffectType_SufferActiveTacticDamageDeduce:
		return "受到[主动战法]伤害减少"
	case BuffEffectType_SufferAssaultTacticDamageDeduce:
		return "受到[突击战法]伤害减少"
	case BuffEffectType_ThePostureOfAPhoenixWithAChickAndASweetPotato_Locked:
		return "鸱苕凤姿[锁定]"
	case BuffEffectType_WaitAtOnesEaseForTheFatigued_ImmunityControl:
		return "以逸待劳[几率免疫控制]"
	case BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare:
		return "以逸待劳[预备]"
	case BuffEffectType_TwelveWonderfulStrategies_Prepare:
		return "十二奇策[预备]"
	case BuffEffectType_TigerIdiot_Locked:
		return "虎痴[锁定]"
	case BuffEffectType_ImmunityCancelWeapon:
		return "免疫缴械"
	case BuffEffectType_ThousandMileWalkingSingleRider_Prepare:
		return "千里走单骑[预备]"
	case BuffEffectType_TacticsActiveDamageImprove:
		return "主动战法伤害提升"
	case BuffEffectType_TacticsActiveTriggerWithSelfImprove:
		return "主动战法[自带]发动率提升"
	case BuffEffectType_ImmunitySandstorm:
		return "免疫沙暴"
	case BuffEffectType_BreakFormation:
		return "破阵"
	case BuffEffectType_MustHit:
		return "必中"
	case BuffEffectType_TacticsActiveWithSelfDamageImprove:
		return "主动战法[自带]伤害提升"
	case BuffEffectType_TacticsAssaultTriggerImprove:
		return "被动战法发动率提升"
	case BuffEffectType_Alert:
		return "警戒"
	case BuffEffectType_AccumulatePower:
		return "蓄威"
	case BuffEffectType_ImmunityChaos:
		return "免疫混乱"
	case BuffEffectType_TacticsActiveTriggerNoSelfImprove:
		return "主动战法[非自带]发动率提升"
	case BuffEffectType_HighWoodenPaddlesConnectedToTheCamp_Prepare:
		return "高橹连营[预备]"
	case BuffEffectType_GrazingArray:
		return "掠阵"
	case BuffEffectType_TacticsActiveTriggerPrepareImprove:
		return "主动战法[准备战法]发动率提升"
	case BuffEffectType_FireGodHeroStyle_Prepare:
		return "火神英风[预备]"
	case BuffEffectType_ContinuousAttack:
		return "连击"
	case BuffEffectType_ActiveTactic_SkipPrepareRound:
		return "主动战法[跳过准备回合]"
	case BuffEffectType_CrowdMovesTenThousandCounts_Prepare:
		return "众动万计[预备]"
	case BuffEffectType_FightHard:
		return "酣斗"
	case BuffEffectType_GeneralAttackDamageImprove:
		return "普通攻击伤害提升"
	case BuffEffectType_Defend:
		return "抵御"
	case BuffEffectType_AncientEvilComes_Prepare:
		return "古之恶来[预备]"
	case BuffEffectType_AttackHeart:
		return "攻心"
	case BuffEffectType_AbilityToRuleTheCountry_Prepare:
		return "经天纬地[预备]"
	case BuffEffectType_GroupAttack:
		return "群攻"
	case BuffEffectType_OnesResolveIsUnshaken_Prepare:
		return "矢志不移[预备]"
	case BuffEffectType_Humility_Prepare:
		return "谦让[预备]"
	case BuffEffectType_LowerBannersAndMuffleDrums_Prepare:
		return "偃旗息鼓[预备]"
	case BuffEffectType_OutstandingTalent_Prepare:
		return "才器过人[预备]"
	case BuffEffectType_RideTigerHardToGetOff_Prepare:
		return "骑虎难下[预备]"
	case BuffEffectType_BeFullyEquippedFor_Prepare:
		return "整装待发[预备]"
	case BuffEffectType_TacticsActiveTriggerImprove:
		return "主动战法发动率提升"
	case BuffEffectType_StrikeBack:
		return "反击"
	case BuffEffectType_Defection:
		return "倒戈"
	case BuffEffectType_ShareResponsibilityFor:
		return "分担"
	case BuffEffectType_KillingMinisterInDream_Prepare:
		return "梦中弑臣[预备]"
	case BuffEffectType_UseMartialArtsToConnectWithGods_Prepare:
		return "用武通神[预备]"
	case BuffEffectType_Insight:
		return "洞察"
	case BuffEffectType_LaunchWeaponDamageImprove:
		return "造成兵刃伤害增加"
	case BuffEffectType_LaunchStrategyDamageImprove:
		return "造成谋略伤害增加"
	case BuffEffectType_Evade:
		return "规避"
	case BuffEffectType_EnhanceWeapon:
		return "会心"
	case BuffEffectType_EnhanceStrategy:
		return "奇谋"
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
	case BuffEffectType_EmergencyTreatment:
		return "急救"
	case BuffEffectType_ClearEyedAndMalicious_Prepare:
		return "鹰视狼顾[预备]"
	case BuffEffectType_ClearEyedAndMalicious_ClearEyed_Prepare:
		return "鹰视狼顾-鹰视[预备]"
	}
	return ""
}
