package team

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

//通用加点
var addition = &vo.BattleGeneralAddition{
	AbilityAttr: po.AbilityAttr{
		ForceBase:        0,
		ForceRate:        0,
		IntelligenceBase: 0,
		IntelligenceRate: 0,
		CharmBase:        0,
		CharmRate:        0,
		CommandBase:      0,
		CommandRate:      0,
		PoliticsBase:     0,
		PoliticsRate:     0,
		SpeedBase:        0,
		SpeedRate:        0,
	},
	GeneralLevel:     0,
	GeneralStarLevel: 0,
	Predestination:   0,
}

var (
	//荀关枪
	//皇冠枪
	//渡江枪
	//麒麟弓
	QiLinGong = []*vo.BattleGeneral{
		//姜维
		{
			IsMaster: true,
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.JiangWei),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.BraveAmbition,
					Name: "义胆雄心",
				},
				{
					Id:   consts.SeizeTheSoul,
					Name: "夺魂挟魄",
				},
				{
					Id:   consts.BlazingWildfire,
					Name: "火炽原燎",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
		//庞统
		{
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.PangTong),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.CleverStrategyAndShrewdTactics,
					Name: "神机妙算",
				},
				{
					Id:   consts.EightGateGoldenLockArray,
					Name: "八门金锁阵",
				},
				{
					Id:   consts.BorrowArrowsWithThatchedBoats,
					Name: "草船借箭",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
		//诸葛亮
		{
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.ZhuGeLiang),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.CleverStrategyAndShrewdTactics,
					Name: "神机妙算",
				},
				{
					Id:   consts.EightGateGoldenLockArray,
					Name: "八门金锁阵",
				},
				{
					Id:   consts.Curettage,
					Name: "刮骨疗毒",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
	}
	//太尉盾
	TaiWeiDun = []*vo.BattleGeneral{
		//司马懿
		{
			IsMaster: true,
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.SiMaYi),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.ClearEyedAndMalicious,
					Name: "鹰视狼顾",
				},
				{
					Id:   consts.ThreeDaysOfSeparation,
					Name: "士别三日",
				},
				{
					Id:   consts.TheSkyIsBlazing,
					Name: "熯天炽地",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
		//曹操
		{
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.CaoCao),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.TraitorInTroubledTimes,
					Name: "乱世奸雄",
				},
				{
					Id:   consts.Charming,
					Name: "魅惑",
				},
				{
					Id:   consts.TengjiaSoldier,
					Name: "藤甲兵",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
		//满宠
		{
			BaseInfo: &po.MetadataGeneral{
				Id: int64(consts.ManChong),
			},
			EquipTactics: []*po.Tactics{
				{
					Id:   consts.SuppressChokesAndPreventRefusals,
					Name: "镇扼防拒",
				},
				{
					Id:   consts.FrontalVectorArray,
					Name: "锋矢阵",
				},
				{
					Id:   consts.Curettage,
					Name: "刮骨疗毒",
				},
			},
			Addition:   addition,
			SoldierNum: 10000,
		},
	}
	//富贵骑
	//群弓
	//香香吴骑
	//荣光枪
	//社稷弓
	//北伐枪
	//关关赵
	//关关张
	//吴枪
	//姬关枪
	//诗诗盾
	//山河盾
	//渊骑
	//三仙盾
	//忠义枪
	//忠义骑
	//魏法骑
	//孙太鲁
	//蜀智
	//SP蜀智
	//五虎枪
	//三势贾
	//三势吕
	//三势陆
	//肉弓
	//桃园盾
	//都督弓
	//孙权形一弓
	//爆头骑
)
