package battle

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/team"
	"github.com/keycasiter/3g_game/biz/util"
)

// 测试单局对战
func TestBattleLogicV2Context_Run(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()

	//############ 模拟对战 ###########
	req := &BattleLogicV2ContextRequest{
		//出战队伍
		FightingTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Fighting,
			ArmType:        consts.ArmType_Archers,
			BattleGenerals: team.QunGong,
		},
		//对战队伍
		EnemyTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Enemy,
			ArmType:        consts.ArmType_Spearman,
			BattleGenerals: team.GuanGuanZhang,
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
			ArmType:        consts.ArmType_Mauler,
			BattleGenerals: team.QunGong,
		},
		//对战队伍
		EnemyTeam: &vo.BattleTeam{
			TeamType:       consts.TeamType_Enemy,
			ArmType:        consts.ArmType_Spearman,
			BattleGenerals: team.GuanGuanZhang,
		},
	}
	//############ 配置队伍是敌是友 ###########
	for _, general := range req.FightingTeam.BattleGenerals {
		general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Fighting
	}
	for _, general := range req.EnemyTeam.BattleGenerals {
		general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Enemy
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
	fmt.Printf("我方结果：%+v\n", resp.BattleResultStatistics.FightingTeam.BattleResult)
	for idx, general := range resp.BattleResultStatistics.FightingTeam.GeneralBattleStatisticsList {
		fmt.Printf("武将[%v]\n", req.FightingTeam.BattleGenerals[idx].BaseInfo.Name)
		for _, tactics := range general.TacticStatisticsList {
			fmt.Printf("战法[%v]触发[%v]杀敌[%v]恢复[%v]\n",
				tactics.TacticName,
				tactics.TriggerTimes,
				tactics.KillSoliderNum,
				tactics.ResumeSoliderNum,
			)
		}
	}
	fmt.Printf("敌方结果：%v\n", resp.BattleResultStatistics.EnemyTeam.BattleResult)
	for idx, general := range resp.BattleResultStatistics.EnemyTeam.GeneralBattleStatisticsList {
		fmt.Printf("武将[%v]\n", req.EnemyTeam.BattleGenerals[idx].BaseInfo.Name)
		for _, tactics := range general.TacticStatisticsList {
			fmt.Printf("战法[%v]触发[%v]杀敌[%v]恢复[%v]\n",
				tactics.TacticName,
				tactics.TriggerTimes,
				tactics.KillSoliderNum,
				tactics.ResumeSoliderNum,
			)
		}
	}
}

// 测试多局对战看平均数据
func TestBattleLogicV2Context_Run_Many_V2(t *testing.T) {

	ctx := context.Background()

	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()

	//############ 1.填入对战阵容 ###########
	teams := [][]*vo.BattleGeneral{
		team.QunGong,
		team.XiangxiangWuQi,
		team.GuanGuanZhang,
		team.TaiWeiDun,
		team.QiLinGong,
	}

	battleDatas := make([]string, 0)
	battleSize := 100
	//遍历所有阵容，和其他阵容对战，测评
	for _, fightingTeam := range teams {
		winSize := 0
		loseSize := 0
		drawSize := 0
		advantageDrawSize := 0
		inferiorityDrawSize := 0
		for _, enemyTeam := range teams {
			for i := 0; i < battleSize; i++ {
				req := &BattleLogicV2ContextRequest{
					//出战队伍
					FightingTeam: &vo.BattleTeam{
						TeamType:       consts.TeamType_Fighting,
						ArmType:        consts.ArmType_Mauler,
						BattleGenerals: fightingTeam,
					},
					//对战队伍
					EnemyTeam: &vo.BattleTeam{
						TeamType:       consts.TeamType_Enemy,
						ArmType:        consts.ArmType_Spearman,
						BattleGenerals: enemyTeam,
					},
				}
				//############ 配置队伍是敌是友 ###########
				for _, general := range req.FightingTeam.BattleGenerals {
					general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Fighting
					general.BaseInfo.UniqueId = util.GenerateUUID()
				}
				for _, general := range req.EnemyTeam.BattleGenerals {
					general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Enemy
					general.BaseInfo.UniqueId = util.GenerateUUID()
				}

				//模拟对战
				runCtx := NewBattleLogicV2Context(ctx, req)
				resp, err := runCtx.Run()
				if err != nil {
					hlog.CtxErrorf(ctx, "run err:%v", err)
					return
				}

				switch resp.BattleResultStatistics.FightingTeam.BattleResult {
				case consts.BattleResult_Win:
					winSize++
				case consts.BattleResult_Lose:
					loseSize++
				case consts.BattleResult_Draw:
					drawSize++
				case consts.BattleResult_Advantage_Draw:
					advantageDrawSize++
				case consts.BattleResult_Inferiority_Draw:
					inferiorityDrawSize++
				}
			}

			fightingTeamName := ""
			enemyTeamName := ""
			for _, general := range fightingTeam {
				fightingTeamName += general.BaseInfo.Name + " "
			}
			for _, general := range enemyTeam {
				enemyTeamName += general.BaseInfo.Name + " "
			}
			battleDatas = append(battleDatas,
				fmt.Sprintf("%v VS %v 胜:%v 平:%v 败:%v\n", fightingTeamName, enemyTeamName, winSize, drawSize, loseSize))
		}
	}
}
