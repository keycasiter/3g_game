package battle

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
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

	battleSize := 100

	winSize := 0
	loseSize := 0
	drawSize := 0
	advantageDrawSize := 0
	inferiorityDrawSize := 0
	fightingGeneralsData := make([]*vo.BattleGeneral, 0)
	enemyGeneralsData := make([]*vo.BattleGeneral, 0)

	for i := 0; i < battleSize; i++ {
		fightingGenerals, err := util.DeepCopyBattleGenerals(team.QunGong)
		if err != nil {
			t.Failed()
		}
		enemyGenerals, err := util.DeepCopyBattleGenerals(team.GuanGuanZhang)
		if err != nil {
			t.Failed()
		}

		req := &BattleLogicV2ContextRequest{
			//出战队伍
			FightingTeam: &vo.BattleTeam{
				TeamType:       consts.TeamType_Fighting,
				ArmType:        consts.ArmType_Archers,
				BattleGenerals: fightingGenerals,
			},
			//对战队伍
			EnemyTeam: &vo.BattleTeam{
				TeamType:       consts.TeamType_Enemy,
				ArmType:        consts.ArmType_Spearman,
				BattleGenerals: enemyGenerals,
			},
		}
		//############ 配置默认值 ###########
		for _, general := range req.FightingTeam.BattleGenerals {
			general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Fighting
			general.Addition.GeneralLevel = 50
		}
		for _, general := range req.EnemyTeam.BattleGenerals {
			general.BaseInfo.GeneralBattleType = consts.GeneralBattleType_Enemy
			general.Addition.GeneralLevel = 50
		}

		//模拟对战
		runCtx := NewBattleLogicV2Context(ctx, req)
		resp, err := runCtx.Run()
		if err != nil {
			hlog.CtxErrorf(ctx, "run err:%v", err)
			return
		}

		//战报结果收集
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

		fightingGeneralsData = buildBattleResultData(resp.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals, fightingGeneralsData)
		enemyGeneralsData = buildBattleResultData(resp.BattleResultStatistics.EnemyTeam.BattleTeam.BattleGenerals, enemyGeneralsData)
	}

	//打印
	fmt.Printf("胜:%v 平:%v 优平:%v 劣平:%v 败:%v\n", winSize, drawSize, advantageDrawSize, inferiorityDrawSize, loseSize)
	for _, data := range fightingGeneralsData {
		fmt.Printf("\n%v\n损失:%v , 累计伤害:%v , 普攻:%v , 恢复:%v",
			data.BaseInfo.Name,
			data.LossSoldierNum/int64(battleSize),
			data.AccumulateTotalDamageNum/int64(battleSize),
			data.AccumulateAttackDamageNum/int64(battleSize),
			data.AccumulateTotalResumeNum/int64(battleSize),
		)
		//战法
		fmt.Println("\n【战法触发】")
		for k, v := range data.TacticAccumulateTriggerMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
		fmt.Println("\n【战法伤害】")
		for k, v := range data.TacticAccumulateDamageMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
		fmt.Println("\n【战法恢复】")
		for k, v := range data.TacticAccumulateResumeMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
	}
	for _, data := range enemyGeneralsData {
		fmt.Printf("\n%v\n损失:%v , 累计伤害:%v , 普攻:%v , 恢复:%v",
			data.BaseInfo.Name,
			data.LossSoldierNum/int64(battleSize),
			data.AccumulateTotalDamageNum/int64(battleSize),
			data.AccumulateAttackDamageNum/int64(battleSize),
			data.AccumulateTotalResumeNum/int64(battleSize),
		)
		//战法
		fmt.Println("\n【战法触发】")
		for k, v := range data.TacticAccumulateTriggerMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
		fmt.Println("\n【战法伤害】")
		for k, v := range data.TacticAccumulateDamageMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
		fmt.Println("\n【战法恢复】")
		for k, v := range data.TacticAccumulateResumeMap {
			fmt.Printf("%v:%v ", k, v/int64(battleSize))
		}
	}
}

func buildBattleResultData(generals []*vo.BattleGeneral, generalsData []*vo.BattleGeneral) []*vo.BattleGeneral {
	for idx, general := range generals {
		if len(generalsData) > idx {
			tmpGeneralData := generalsData[idx]
			tmpGeneralData.SoldierNum += general.SoldierNum
			tmpGeneralData.LossSoldierNum += general.LossSoldierNum
			tmpGeneralData.AccumulateTotalDamageNum += general.AccumulateTotalDamageNum
			tmpGeneralData.AccumulateAttackDamageNum += general.AccumulateAttackDamageNum
			tmpGeneralData.AccumulateTotalResumeNum += general.AccumulateTotalResumeNum
			for k, v := range general.TacticAccumulateDamageMap {
				if add, ok := tmpGeneralData.TacticAccumulateDamageMap[k]; ok {
					tmpGeneralData.TacticAccumulateDamageMap[k] = v + add
				} else {
					tmpGeneralData.TacticAccumulateDamageMap[k] = v
				}
			}
			for k, v := range general.TacticAccumulateTriggerMap {
				if add, ok := tmpGeneralData.TacticAccumulateTriggerMap[k]; ok {
					tmpGeneralData.TacticAccumulateTriggerMap[k] = v + add
				} else {
					tmpGeneralData.TacticAccumulateTriggerMap[k] = v
				}
			}
			for k, v := range general.TacticAccumulateResumeMap {
				if add, ok := tmpGeneralData.TacticAccumulateResumeMap[k]; ok {
					tmpGeneralData.TacticAccumulateResumeMap[k] = v + add
				} else {
					tmpGeneralData.TacticAccumulateResumeMap[k] = v
				}
			}
		} else {
			generalsData = append(generalsData, &vo.BattleGeneral{
				BaseInfo: &po.MetadataGeneral{
					Name: general.BaseInfo.Name,
				},
				SoldierNum:                 general.SoldierNum,
				LossSoldierNum:             general.LossSoldierNum,
				AccumulateTotalDamageNum:   general.AccumulateTotalDamageNum,
				AccumulateAttackDamageNum:  general.AccumulateAttackDamageNum,
				AccumulateTotalResumeNum:   general.AccumulateTotalResumeNum,
				TacticAccumulateTriggerMap: general.TacticAccumulateTriggerMap,
				TacticAccumulateResumeMap:  general.TacticAccumulateResumeMap,
				TacticAccumulateDamageMap:  general.TacticAccumulateDamageMap,
			})
		}
	}
	return generalsData
}
