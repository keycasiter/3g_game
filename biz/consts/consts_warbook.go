package consts

// 兵书类型
type WarBookType int

const (
	WarBookType_Fighting          WarBookType = 1 //作战
	WarBookType_TruthAndFalsehood WarBookType = 2 //虚实
	WarBookType_MilitaryForm      WarBookType = 3 //军形
	WarBookType_NineChanges       WarBookType = 4 //九变
)

// 兵书枚举
type WarBookDetailType int

const (
	//作战
	WarBookDetailType_TheOddAndTheRightCoexist WarBookDetailType = 1 //奇正相生
	WarBookDetailType_BraveButNotBrave         WarBookDetailType = 2 //蛮勇非勇
	WarBookDetailType_NotBraveWillDie          WarBookDetailType = 3 //不勇则死
	WarBookDetailType_MilitaryAbility          WarBookDetailType = 4 //武略
	WarBookDetailType_VictoriousBattle         WarBookDetailType = 5 //胜战
	WarBookDetailType_PersistentSpirit         WarBookDetailType = 6 //执锐
	WarBookDetailType_MilitaryStrategy         WarBookDetailType = 7 //文韬
	WarBookDetailType_HideKnife                WarBookDetailType = 8 //藏刀
	//虚实
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 9 //大谋不谋
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 10 //以治击乱
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 11 //攻其不备
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //鬼谋
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //妙算
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //将威
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //神机
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //占卜
	//军形
	//九变
)
