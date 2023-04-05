package logic

import (
	"context"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mongodb"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/kr/pretty"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestBattleLogicContext_Run_DataFromMongoDB(t *testing.T) {
	conf.InitConfig()
	mongodb.InitMongodb()

	//麒麟弓
	arr := []string{"姜维", "庞统", "诸葛亮"}
	ctx := &mongodb.MetadataGeneralContext{}
	qiLinGongGenerals, err := ctx.FindAll(context.Background(), bson.M{"name": bson.M{"$in": arr}})
	if err != nil {
		t.Fail()
	}
	pretty.Log(qiLinGongGenerals)
	//装配战法
	tactics := []*po.Tactics{
		{
			Id:   0,
			Name: "",
		},
		{
			Id:   0,
			Name: "",
		},
		{
			Id:   0,
			Name: "",
		},
	}
	//加点
	addition := &vo.BattleGeneralAddition{
		AbilityAttr: po.AbilityAttr{
			ForceBase:        50,
			IntelligenceBase: 0,
			CharmBase:        0,
			CommandBase:      0,
			PoliticsBase:     0,
			SpeedBase:        0,
		},
		GeneralLevel:     50,
		GeneralStarLevel: 5,
		Predestination:   0,
	}

	//模拟对战双方属性
	fightGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range qiLinGongGenerals {
		fightGenerals = append(fightGenerals, &vo.BattleGeneral{
			BaseInfo:     general,
			EquipTactics: tactics,
			Addition:     addition,
		})
	}
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range qiLinGongGenerals {
		enemyGenerals = append(enemyGenerals, &vo.BattleGeneral{
			BaseInfo:     general,
			EquipTactics: tactics,
			Addition:     addition,
		})
	}

	//############ 模拟对战 ###########
	req := &BattleLogicContextRequest{
		//出战队伍
		FightingTeam: &vo.BattleTeam{
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: fightGenerals,
		},
		//对战队伍
		EnemyTeam: &vo.BattleTeam{
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: enemyGenerals,
		},
	}
	runCtx := NewBattleLogicContext(context.Background(), req)
	runCtx.Run()
}

func TestBattleLogicContext_Run_DataFromMock(t *testing.T) {
	//太尉盾
	taiWeiDun := []*po.MetadataGeneral{
		//司马懿
		{
			Id:       int64(consts.SiMaYi),
			UniqueId: 10 + int64(consts.SiMaYi),
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
			UniqueId: 10 + int64(consts.CaoCao),
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
			UniqueId: 10 + int64(consts.ManChong),
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

	//太尉盾
	taiWeiDun2 := []*po.MetadataGeneral{
		//司马懿
		{
			Id:       int64(consts.SiMaYi),
			UniqueId: 20 + int64(consts.SiMaYi),
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
			UniqueId: 20 + int64(consts.CaoCao),
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
			UniqueId: 20 + int64(consts.ManChong),
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
				SoldierNum:   30000,
			}
			fightGenerals = append(fightGenerals, vo)
		case "曹操":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: caoCaoTactics,
				Addition:     addition,
				SoldierNum:   30000,
			}
			fightGenerals = append(fightGenerals, vo)
		case "满宠":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: manChongTactics,
				Addition:     addition,
				SoldierNum:   30000,
			}
			fightGenerals = append(fightGenerals, vo)
		}

	}
	//敌人
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range taiWeiDun2 {
		general.GeneralBattleType = consts.GeneralBattleType_Enemy
		switch general.Name {
		case "司马懿":
			vo := &vo.BattleGeneral{
				IsMaster:     true,
				BaseInfo:     general,
				EquipTactics: siMaYiTactics,
				Addition:     addition,
				SoldierNum:   30000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		case "曹操":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: caoCaoTactics,
				Addition:     addition,
				SoldierNum:   30000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		case "满宠":
			vo := &vo.BattleGeneral{
				BaseInfo:     general,
				EquipTactics: manChongTactics,
				Addition:     addition,
				SoldierNum:   30000,
			}
			enemyGenerals = append(enemyGenerals, vo)
		}

	}

	//############ 模拟对战 ###########
	req := &BattleLogicContextRequest{
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
	runCtx := NewBattleLogicContext(context.Background(), req)
	runCtx.Run()
}
