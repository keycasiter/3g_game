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

func TestBattleLogicContext_Run(t *testing.T) {
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
