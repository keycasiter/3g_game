package battle

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/team"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/kr/pretty"
	"github.com/spf13/cast"
	"testing"
)

// 测试单局对战
func TestBattleLogicV2Context_Run(t *testing.T) {
	//太尉盾
	taiWeiDun := []*po.MetadataGeneral{
		//司马懿
		{
			Id:       int64(consts.SiMaYi),
			UniqueId: cast.ToString(10 + int64(consts.SiMaYi)),
			Name:     "司马懿",
			Group:    consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        83.16,
				ForceRate:        0,
				IntelligenceBase: 359.52,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      204.77,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        78.36,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
		//曹操
		{
			Id:       int64(consts.CaoCao),
			UniqueId: cast.ToString(10 + int64(consts.CaoCao)),
			Name:     "曹操",
			Group:    consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        135.84,
				ForceRate:        0,
				IntelligenceBase: 299.48,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      235.70,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        130.83,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
		//满宠
		{
			Id:       int64(consts.ManChong),
			UniqueId: cast.ToString(10 + int64(consts.ManChong)),
			Name:     "满宠",
			Group:    consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        112.66,
				ForceRate:        0,
				IntelligenceBase: 291.52,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      191.40,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        115.36,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
	}

	//麒麟弓
	qiLinGong := []*po.MetadataGeneral{
		//姜维
		{
			Id:       int64(consts.JiangWei),
			UniqueId: cast.ToString(20 + int64(consts.JiangWei)),
			Name:     "姜维",
			Group:    consts.Group_ShuGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        198.54,
				ForceRate:        0,
				IntelligenceBase: 291.89,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      192.35,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        124.87,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
		//庞统
		{
			Id:       int64(consts.PangTong),
			UniqueId: cast.ToString(20 + int64(consts.PangTong)),
			Name:     "庞统",
			Group:    consts.Group_ShuGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        58.96,
				ForceRate:        0,
				IntelligenceBase: 334.98,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      167.73,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        74.33,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
		//诸葛亮
		{
			Id:       int64(consts.ZhuGeLiang),
			UniqueId: cast.ToString(20 + int64(consts.ZhuGeLiang)),
			Name:     "诸葛亮",
			Group:    consts.Group_ShuGuo,
			AbilityAttr: &po.AbilityAttr{
				ForceBase:        62.13,
				ForceRate:        0,
				IntelligenceBase: 334.00,
				IntelligenceRate: 0,
				CharmBase:        0,
				CharmRate:        0,
				CommandBase:      236.46,
				CommandRate:      0,
				PoliticsBase:     0,
				PoliticsRate:     0,
				SpeedBase:        76.01,
				SpeedRate:        0,
			},
			ArmsAttr: &po.ArmsAttr{
				Cavalry:   consts.ArmsAbility_S,
				Mauler:    consts.ArmsAbility_S,
				Archers:   consts.ArmsAbility_S,
				Spearman:  consts.ArmsAbility_S,
				Apparatus: consts.ArmsAbility_S,
			},
		},
	}

	//司马懿战法
	siMaYiTactics := []*po.Tactics{
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
	}
	//曹操战法
	caoCaoTactics := []*po.Tactics{
		{
			Id:   consts.TraitorInTroubledTimes,
			Name: "乱世奸雄",
		},
		{
			Id:   consts.Charming,
			Name: "魅惑",
		},
		{
			Id:   consts.AppeaseArmyAndPeople,
			Name: "抚揖军民",
		},
	}
	//满宠战法
	manChongTactics := []*po.Tactics{
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
	}

	//姜维战法
	jiangweiTactics := []*po.Tactics{
		{
			Id:   consts.BraveAmbition,
			Name: "义胆雄心",
		},
		{
			Id:   consts.SeizeTheSoul,
			Name: "夺魂挟魄",
		},
		{
			Id:   consts.CupSnakeGhostCar,
			Name: "杯蛇鬼车",
		},
	}
	//庞统战法
	pangTongTactics := []*po.Tactics{
		{
			Id:   consts.InterlockedStratagems,
			Name: "连环计",
		},
		{
			Id:   consts.TaipingLaw,
			Name: "太平道法",
		},
		{
			Id:   consts.WuDangFlyArmy,
			Name: "无当飞军",
		},
	}
	//诸葛亮战法
	zhuGeLiangTactics := []*po.Tactics{
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
	}

	//加点
	addition := &vo.BattleGeneralAddition{
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

	//模拟对战双方属性
	//我方
	fightGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range taiWeiDun {
		general.GeneralBattleType = consts.GeneralBattleType_Fighting
		switch general.Name {
		case "司马懿":
			vo := &vo.BattleGeneral{
				IsMaster:     true,
				BaseInfo:     general,
				EquipTactics: siMaYiTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			fightGenerals = append(fightGenerals, vo)
		case "曹操":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: caoCaoTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			fightGenerals = append(fightGenerals, vo)
		case "满宠":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: manChongTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			fightGenerals = append(fightGenerals, vo)
		}

	}
	//敌人
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range qiLinGong {
		general.GeneralBattleType = consts.GeneralBattleType_Enemy
		switch general.Name {
		case "姜维":
			vo := &vo.BattleGeneral{
				IsMaster:     true,
				BaseInfo:     general,
				EquipTactics: jiangweiTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		case "庞统":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: pangTongTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		case "诸葛亮":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: zhuGeLiangTactics,
				Addition:     addition,
				SoldierNum:   10000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		}

	}

	//############ 模拟对战 ###########
	req := &BattleLogicV2ContextRequest{
		//出战队伍
		FightingTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Fighting,
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: fightGenerals,
		},
		//对战队伍
		EnemyTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Enemy,
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: enemyGenerals,
		},
	}
	fmt.Printf("req :%s \n", util.ToJsonString(context.Background(), req))
	runCtx := NewBattleLogicV2Context(context.Background(), req)
	runCtx.Run()
}

// 测试多局对战看平均数据
func TestBattleLogicV2Context_Run_Many(t *testing.T) {

	ctx := context.Background()

	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()

	//############ 1.填入对战阵容 ###########
	req := &BattleLogicV2ContextRequest{
		//出战队伍
		FightingTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Fighting,
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: team.GuanGuanZhang,
		},
		//对战队伍
		EnemyTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Enemy,
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: team.XiangxiangWuQi,
		},
	}
	//############ 配置队伍是敌是友 ###########
	for _, general := range req.FightingTeam.BattleGenerals {
		general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Fighting
	}
	for _, general := range req.EnemyTeam.BattleGenerals {
		general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Enemy
	}

	//############ 2.从数据库拉取阵容武将属性补充 ###########
	allGenerals := append(req.FightingTeam.BattleGenerals, req.EnemyTeam.BattleGenerals...)
	allGeneralIds := make([]int64, 0)
	for _, general := range allGenerals {
		allGeneralIds = append(allGeneralIds, general.BaseInfo.Id)
	}
	generalList, err := mysql.NewGeneral().QueryGeneralList(ctx, &vo.QueryGeneralCondition{
		Ids: allGeneralIds,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "run err:%v", err)
		return
	}
	generalMap := make(map[int64]*po.General, 0)
	for _, general := range generalList {
		generalMap[general.Id] = general
	}
	for _, general := range allGenerals {
		if generalDb, ok := generalMap[general.BaseInfo.Id]; ok {
			general.BaseInfo.Id = generalDb.Id
			general.BaseInfo.Name = generalDb.Name
			general.BaseInfo.Group = consts.Group(generalDb.Group)
			general.BaseInfo.UniqueId = cast.ToString(generalDb.Id)

			//属性
			abilityAttr := &po.AbilityAttrString{}
			util.ParseJsonObj(ctx, abilityAttr, generalDb.AbilityAttr)
			general.BaseInfo.AbilityAttr = &po.AbilityAttr{
				ForceBase:        cast.ToFloat64(abilityAttr.ForceBase),
				ForceRate:        cast.ToFloat64(abilityAttr.ForceRate),
				IntelligenceBase: cast.ToFloat64(abilityAttr.IntelligenceBase),
				IntelligenceRate: cast.ToFloat64(abilityAttr.IntelligenceRate),
				CharmBase:        cast.ToFloat64(abilityAttr.CharmBase),
				CharmRate:        cast.ToFloat64(abilityAttr.CharmRate),
				CommandBase:      cast.ToFloat64(abilityAttr.CommandBase),
				CommandRate:      cast.ToFloat64(abilityAttr.CommandRate),
				PoliticsBase:     cast.ToFloat64(abilityAttr.PoliticsBase),
				PoliticsRate:     cast.ToFloat64(abilityAttr.PoliticsRate),
				SpeedBase:        cast.ToFloat64(abilityAttr.SpeedBase),
				SpeedRate:        cast.ToFloat64(abilityAttr.SpeedRate),
			}

			//兵种适性
			armsAttr := &po.ArmsAttr{}
			util.ParseJsonObj(ctx, &armsAttr, generalDb.ArmAttr)
			general.BaseInfo.ArmsAttr = &po.ArmsAttr{
				Cavalry:   armsAttr.Cavalry,
				Mauler:    armsAttr.Mauler,
				Archers:   armsAttr.Archers,
				Spearman:  armsAttr.Spearman,
				Apparatus: armsAttr.Apparatus,
			}
		}
	}

	//模拟对战
	fmt.Printf("req :%s \n", util.ToJsonString(ctx, req))
	runCtx := NewBattleLogicV2Context(ctx, req)
	resp, err := runCtx.Run()
	if err != nil {
		hlog.CtxErrorf(ctx, "run err:%v", err)
		return
	}
	fmt.Printf("%v", resp)
	pretty.Logf("resp:%v", util.ToJsonString(ctx, resp))
}
