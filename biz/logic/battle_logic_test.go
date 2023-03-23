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
			Name:  "司马懿",
			Group: consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
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
				SpeedBase:        50,
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
			Name:  "曹操",
			Group: consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
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
				SpeedBase:        40,
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
			Name:  "满宠",
			Group: consts.Group_WeiGuo,
			AbilityAttr: &po.AbilityAttr{
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
				SpeedBase:        60,
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
	//装配战法
	taiWeiDunTactics := []*po.Tactics{
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
		fightGenerals = append(fightGenerals, &vo.BattleGeneral{
			BaseInfo:     general,
			EquipTactics: taiWeiDunTactics,
			Addition:     addition,
		})
	}
	//敌人
	taiWeiDun2 := make([]*po.MetadataGeneral, 0)
	taiWeiDun2 = taiWeiDun2[0 : len(taiWeiDun)-1]
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range taiWeiDun2 {
		general.GeneralBattleType = consts.GeneralBattleType_Enemy
		enemyGenerals = append(enemyGenerals, &vo.BattleGeneral{
			BaseInfo:     general,
			EquipTactics: taiWeiDunTactics,
			Addition:     addition,
		})
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
