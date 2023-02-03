package logic

import (
	"context"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mongodb"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestBattleLogicContext_Run(t *testing.T) {
	conf.InitConfig()
	mongodb.InitMongodb()

	arr := []string{"姜维", "庞统", "诸葛亮"}
	ctx := &mongodb.MetadataGeneralContext{}
	res, err := ctx.FindAll(context.Background(), bson.M{"name": bson.M{"$in": arr}})
	if err != nil {
		t.Fail()
	}
	//t.Logf("res:%v", util.ToJsonString(context.Background(), res))

	req := &BattleLogicContextRequest{
		FightingGenerals: []*vo.BattleGeneral{
			{
				ArmType:      consts.ArmType_Cavalry,
				BaseInfo:     nil,
				EquipTactics: nil,
				Addition:     nil,
			},
		},
		EnemyGenerals: nil,
	}
	runCtx := NewBattleLogicContext(context.Background(), req)
	runCtx.Run()
}
