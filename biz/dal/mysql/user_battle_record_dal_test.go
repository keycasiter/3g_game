package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/keycasiter/3g_game/biz/conf"
)

func TestUserBattleRecordDal_QueryUserBattleWinRate(t *testing.T) {
	conf.InitConfig()
	InitMysql()
	ctx := context.Background()
	winRate, err := NewUserBattleRecord().QueryUserBattleHighFreqUsedGeneralStatistics(ctx, 3268589903)
	if err != nil {
		t.Failed()
	}
	fmt.Printf("%v", winRate)
}

func TestUserBattleRecordDal_QueryUserBattleHighFreqUsedGeneralStatistics(t *testing.T) {
	conf.InitConfig()
	InitMysql()
	ctx := context.Background()
	resultMap, err := NewUserBattleRecord().QueryUserBattleHighFreqUsedGeneralStatistics(ctx, 3268589903)
	if err != nil {
		t.Failed()
	}
	fmt.Printf("%v", resultMap)
}

func TestUserBattleRecordDal_QueryUserBattleHighFreqUsedTeamStatistics(t *testing.T) {
	conf.InitConfig()
	InitMysql()
	ctx := context.Background()
	resultMap, err := NewUserBattleRecord().QueryUserBattleHighFreqUsedTeamStatistics(ctx, 3268589903)
	if err != nil {
		t.Failed()
	}
	fmt.Printf("%v", resultMap)
}

func TestUserBattleRecordDal_QueryUserBattleHighFreqTacticStatistics(t *testing.T) {
	conf.InitConfig()
	InitMysql()
	ctx := context.Background()
	resultMap, err := NewUserBattleRecord().QueryUserBattleHighFreqTacticStatistics(ctx, 3268589903)
	if err != nil {
		t.Failed()
	}
	fmt.Printf("%v", resultMap)
}
