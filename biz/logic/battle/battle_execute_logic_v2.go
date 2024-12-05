package battle

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/cache"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics"
	"github.com/keycasiter/3g_game/biz/tactics/execute"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	warbookExecute "github.com/keycasiter/3g_game/biz/warbook/execute"
	"github.com/spf13/cast"
)

// req
type BattleLogicV2ContextRequest struct {
	/** 队伍信息 **/
	// 出战队伍信息
	FightingTeam *vo.BattleTeam
	// 对战队伍信息
	EnemyTeam *vo.BattleTeam
	//用户uid
	Uid int64
}

// resp
type BattleLogicV2ContextResponse struct {
	//对战数据统计
	BattleResultStatistics *model.BattleResultStatistics
	//对战过程数据
	BattleProcessStatistics map[consts.BattlePhase]map[consts.BattleRound][]string
}

//resp

// 对战上下文环境
type BattleLogicV2Context struct {
	/** DSL数据 **/
	//上下文
	Ctx context.Context
	// 入参
	Req *BattleLogicV2ContextRequest
	// 出参
	Resp *BattleLogicV2ContextResponse
	// 执行方法
	Funcs []func()
	// 执行错误
	RunErr error

	/** 回合对战全局变量 **/
	//对战战法全局holder
	TacticsParams *model.TacticsParams
	//回合结束标记
	BattleRoundEndFlag bool

	/**对战数据统计**/
}

func NewBattleLogicV2Context(ctx context.Context, req *BattleLogicV2ContextRequest) *BattleLogicV2Context {
	runCtx := &BattleLogicV2Context{
		Ctx:  ctx,
		Req:  req,
		Resp: &BattleLogicV2ContextResponse{},
	}

	//注入方法执行顺序
	runCtx.Funcs = []func(){
		//构建对战战法数据
		runCtx.buildBattleRoundParams,
		//对战准备阶段处理
		runCtx.processBattlePreparePhase,
		//对战对阵阶段处理（1-8回合）
		runCtx.processBattleFightingPhase,
		//对战战报统计数据处理
		runCtx.processBattleReportStatistics,
	}

	return runCtx
}

func (runCtx *BattleLogicV2Context) Run() (*BattleLogicV2ContextResponse, error) {
	for _, f := range runCtx.Funcs {
		f()
		if runCtx.RunErr != nil {
			hlog.CtxErrorf(runCtx.Ctx, "BattleLogicV2Context Func=%v Run err:%v", f, runCtx.RunErr)
			return nil, runCtx.RunErr
		}
	}
	return runCtx.Resp, nil
}

// 属性加点处理
func (runCtx *BattleLogicV2Context) handleAbilityAttrAddition(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase + general.Addition.AbilityAttr.ForceBase
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase + general.Addition.AbilityAttr.IntelligenceBase
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase + general.Addition.AbilityAttr.CommandBase
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase + general.Addition.AbilityAttr.SpeedBase

}

// 兵书处理
func (runCtx *BattleLogicV2Context) handleWarBook(general *vo.BattleGeneral) {
	warbookExecute.NewWarBookExecutor(runCtx.Ctx, general, runCtx.TacticsParams).Execute()
}

// 武将等级处理
func (runCtx *BattleLogicV2Context) handleGeneralLevel(general *vo.BattleGeneral) {
	//武力加成
	general.BaseInfo.AbilityAttr.ForceBase =
		general.BaseInfo.AbilityAttr.ForceBase +
			general.BaseInfo.AbilityAttr.ForceRate*float64(general.Addition.GeneralLevel)
	//智力加成
	general.BaseInfo.AbilityAttr.IntelligenceBase =
		general.BaseInfo.AbilityAttr.IntelligenceBase +
			general.BaseInfo.AbilityAttr.IntelligenceRate*float64(general.Addition.GeneralLevel)
	//统率加成
	general.BaseInfo.AbilityAttr.CommandBase =
		general.BaseInfo.AbilityAttr.CommandBase +
			general.BaseInfo.AbilityAttr.CommandRate*float64(general.Addition.GeneralLevel)
	//速度加成
	general.BaseInfo.AbilityAttr.SpeedBase =
		general.BaseInfo.AbilityAttr.SpeedBase +
			general.BaseInfo.AbilityAttr.SpeedRate*float64(general.Addition.GeneralLevel)
}

// 武将标签处理
func (runCtx *BattleLogicV2Context) handleGeneralTag(general *vo.BattleGeneral) {
	//仙人：属性提高30%
	for _, generalTag := range general.BaseInfo.GeneralTag {
		switch consts.GeneralTag(generalTag) {
		case consts.GeneralTag_Celestials:
			//武力加成
			general.BaseInfo.AbilityAttr.ForceBase = general.BaseInfo.AbilityAttr.ForceBase * 1.3
			//智力加成
			general.BaseInfo.AbilityAttr.IntelligenceBase = general.BaseInfo.AbilityAttr.IntelligenceBase * 1.3
			//统率加成
			general.BaseInfo.AbilityAttr.CommandBase = general.BaseInfo.AbilityAttr.CommandBase * 1.3
			//速度加成
			general.BaseInfo.AbilityAttr.SpeedBase = general.BaseInfo.AbilityAttr.SpeedBase * 1.3

			hlog.CtxInfof(runCtx.Ctx, "[%s]发动战法【仙人】", general.BaseInfo.Name)
			hlog.CtxInfof(runCtx.Ctx, "[%s]是一名【仙人】，属性提高30%%", general.BaseInfo.Name)
		}
	}
}

// 建筑科技-阵营加成处理
func (runCtx *BattleLogicV2Context) handleBuildingTechGroupAddition(team *vo.BattleTeam) {
	//判断是否同阵营
	group := consts.Group_Unknow
	for _, general := range team.BattleGenerals {
		if group == consts.Group_Unknow {
			group = general.BaseInfo.Group
		} else if group != general.BaseInfo.Group {
			//非同阵营
			return
		}
	}
	//提取阵营加成
	var additionRate float64
	switch group {
	case consts.Group_WeiGuo:
		additionRate = team.BuildingTechGroupAddition.GroupWeiGuoRate
	case consts.Group_ShuGuo:
		additionRate = team.BuildingTechGroupAddition.GroupShuGuoRate
	case consts.Group_WuGuo:
		additionRate = team.BuildingTechGroupAddition.GroupWuGuoRate
	case consts.Group_QunXiong:
		additionRate = team.BuildingTechGroupAddition.GroupQunXiongRate
	}

	if additionRate == 0 {
		return
	}

	for _, general := range team.BattleGenerals {
		//武力加成
		general.BaseInfo.AbilityAttr.ForceBase = general.BaseInfo.AbilityAttr.ForceBase * (1 + additionRate)
		//智力加成
		general.BaseInfo.AbilityAttr.IntelligenceBase = general.BaseInfo.AbilityAttr.IntelligenceBase * (1 + additionRate)
		//统率加成
		general.BaseInfo.AbilityAttr.CommandBase = general.BaseInfo.AbilityAttr.CommandBase * (1 + additionRate)
		//速度加成
		general.BaseInfo.AbilityAttr.SpeedBase = general.BaseInfo.AbilityAttr.SpeedBase * (1 + additionRate)
	}
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【%s】强化效果，属性提升了%d%s",
		team.BattleGenerals[0].BaseInfo.Name,
		util.TranslateGroup(consts.Group(group)),
		int(additionRate),
		"%",
	)
}

// 建筑科技-属性加成处理
func (runCtx *BattleLogicV2Context) handleBuildingTechAttrAddition(team *vo.BattleTeam) {
	for _, general := range team.BattleGenerals {
		//武力加成
		if team.BuildingTechAttrAddition.ForceAddition > 0 {
			general.BaseInfo.AbilityAttr.ForceBase = general.BaseInfo.AbilityAttr.ForceBase +
				team.BuildingTechAttrAddition.ForceAddition
			hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-武】强化效果，【武力】属性提升了%d",
				team.BattleGenerals[0].BaseInfo.Name,
				int(team.BuildingTechAttrAddition.ForceAddition),
			)
		}
		//智力加成
		if team.BuildingTechAttrAddition.IntelligenceAddition > 0 {
			general.BaseInfo.AbilityAttr.IntelligenceBase = general.BaseInfo.AbilityAttr.IntelligenceBase +
				team.BuildingTechAttrAddition.IntelligenceAddition
			hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-谋】强化效果，【智力】属性提升了%d",
				team.BattleGenerals[0].BaseInfo.Name,
				int(team.BuildingTechAttrAddition.ForceAddition),
			)
		}
		//统率加成
		if team.BuildingTechAttrAddition.CommandAddition > 0 {
			general.BaseInfo.AbilityAttr.CommandBase = general.BaseInfo.AbilityAttr.CommandBase +
				team.BuildingTechAttrAddition.CommandAddition
			hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-防】强化效果，【统率】属性提升了%d",
				team.BattleGenerals[0].BaseInfo.Name,
				int(team.BuildingTechAttrAddition.CommandAddition),
			)
		}
		//速度加成
		if team.BuildingTechAttrAddition.SpeedAddition > 0 {
			general.BaseInfo.AbilityAttr.SpeedBase = general.BaseInfo.AbilityAttr.SpeedBase +
				team.BuildingTechAttrAddition.SpeedAddition
			hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-速】强化效果，【速度】属性提升了%d",
				team.BattleGenerals[0].BaseInfo.Name,
				int(team.BuildingTechAttrAddition.SpeedAddition),
			)
		}
	}
}

// 兵种适性处理
func (runCtx *BattleLogicV2Context) handleArmAbility(teamArmType consts.ArmType, general *vo.BattleGeneral) {
	armsAbility := consts.ArmsAbility_Unknow
	switch teamArmType {
	//骑兵
	case consts.ArmType_Cavalry:
		armsAbility = general.BaseInfo.ArmsAttr.Cavalry
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Cavalry, general.BaseInfo.AbilityAttr)
	//盾兵
	case consts.ArmType_Mauler:
		armsAbility = general.BaseInfo.ArmsAttr.Mauler
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Mauler, general.BaseInfo.AbilityAttr)
	//弓兵
	case consts.ArmType_Archers:
		armsAbility = general.BaseInfo.ArmsAttr.Archers
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Archers, general.BaseInfo.AbilityAttr)
	//枪兵
	case consts.ArmType_Spearman:
		armsAbility = general.BaseInfo.ArmsAttr.Spearman
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Spearman, general.BaseInfo.AbilityAttr)
	//器械
	case consts.ArmType_Apparatus:
		armsAbility = general.BaseInfo.ArmsAttr.Apparatus
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Apparatus, general.BaseInfo.AbilityAttr)
	}

	hlog.CtxInfof(runCtx.Ctx, "[%s]的【 %s 】兵种适性为【 %s 】，属性调整为原来的%s",
		general.BaseInfo.Name,
		util.TranslateArmType(teamArmType),
		util.TranslateArmsAbility(armsAbility),
		util.TranslateArmsAbilityAddition(armsAbility),
	)
}

// 对战准备阶段处理
func (runCtx *BattleLogicV2Context) processBattlePreparePhase() {
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] processBattlePreparePhase...")

	tacticsParams := runCtx.TacticsParams

	//对战阶段
	tacticsParams.CurrentPhase = consts.Battle_Phase_Prepare

	/********************************************************/
	/*** 以下不受武将速度影响来执行，根据我方先/敌方后的顺序执行即可 ***/
	/********************************************************/
	hlog.CtxInfof(runCtx.Ctx, "<<<<<<<<<<<<<【准备阶段】>>>>>>>>>>>>>")

	//我方武将加成处理
	for _, general := range runCtx.Req.FightingTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.Req.FightingTeam, general)
	}
	//我方阵容加成处理
	runCtx.handleTeamAddition(runCtx.Req.FightingTeam)

	//敌方武将加成处理
	for _, general := range runCtx.Req.EnemyTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.Req.EnemyTeam, general)
	}
	//敌方阵容加成处理
	runCtx.handleTeamAddition(runCtx.Req.EnemyTeam)

	//准备阶段战法处理
	runCtx.handlePreparePhaseTactic(tacticsParams)
}

func (runCtx *BattleLogicV2Context) handlePreparePhaseTactic(tacticsParams *model.TacticsParams) {
	/****************************/
	/*** 以下受武将速度影响来执行 ***/
	/****************************/

	//速度从快到慢给武将排序
	tacticsParams.AllGeneralArr = util.MakeGeneralsOrderBySpeed(tacticsParams.AllGeneralArr)

	for _, currentGeneral := range tacticsParams.AllGeneralArr {
		//设置战法轮次属性
		runCtx.TacticsParams.CurrentRound = consts.Battle_Round_Prepare
		runCtx.TacticsParams.CurrentGeneral = currentGeneral

		//执行战法
		for _, tactic := range currentGeneral.EquipTactics {
			//战法发动顺序：1.被动 > 2.阵法 > 3.兵种 > 4.指挥 > 5.主动 > 6.普攻 > 7.突击
			//1.被动
			if _, ok := consts.PassiveTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				tacticHandler := handler.Init(tacticsParams)
				//发动率判断
				if !util.GenerateRate(tacticHandler.GetTriggerRate()) {
					hlog.CtxInfof(runCtx.Ctx, "【战法】%v，触发概率%v%，没有生效", tactic.Name, tacticHandler.GetTriggerRate())
					continue
				}
				//统计上报
				util.TacticReport(runCtx.TacticsParams,
					currentGeneral.BaseInfo.UniqueId,
					int64(tactic.Id),
					1,
					0,
					0,
				)

				//触发「发动被动战法开始」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_PassiveTactic]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}

				//战法执行
				execute.TacticsExecute(runCtx.Ctx, tacticHandler)

				//触发「发动被动战法结束」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_PassiveTacticEnd]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
			}
			//2.阵法
			if _, ok := consts.TroopsTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				tacticHandler := handler.Init(tacticsParams)
				//发动率判断
				if !util.GenerateRate(tacticHandler.GetTriggerRate()) {
					hlog.CtxInfof(runCtx.Ctx, "【战法】%v，触发概率%v%，没有生效", tactic.Name, tacticHandler.GetTriggerRate())
					continue
				}
				//统计上报
				util.TacticReport(runCtx.TacticsParams,
					currentGeneral.BaseInfo.UniqueId,
					int64(tactic.Id),
					1,
					0,
					0,
				)

				//触发「发动阵法开始」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_TroopsTactic]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}

				//战法执行
				execute.TacticsExecute(runCtx.Ctx, tacticHandler)

				//触发「发动阵法结束」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_TroopsTacticEnd]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
			}
			//3.兵种
			if _, ok := consts.ArmTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				tacticHandler := handler.Init(tacticsParams)
				//发动率判断
				if !util.GenerateRate(tacticHandler.GetTriggerRate()) {
					hlog.CtxInfof(runCtx.Ctx, "【战法】%v，触发概率%v%，没有生效", tactic.Name, tacticHandler.GetTriggerRate())
					continue
				}
				//统计上报
				util.TacticReport(runCtx.TacticsParams,
					currentGeneral.BaseInfo.UniqueId,
					int64(tactic.Id),
					1,
					0,
					0,
				)

				//触发「发动兵种开始」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_ArmTactic]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
				//战法执行
				execute.TacticsExecute(runCtx.Ctx, tacticHandler)

				//触发「发动兵种结束」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_ArmTacticEnd]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
			}
			//4.指挥
			if _, ok := consts.CommandTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				tacticHandler := handler.Init(tacticsParams)
				//发动率判断
				if !util.GenerateRate(tacticHandler.GetTriggerRate()) {
					hlog.CtxInfof(runCtx.Ctx, "【战法】%v，触发概率%v%，没有生效", tactic.Name, tacticHandler.GetTriggerRate())
					continue
				}
				//统计上报
				util.TacticReport(runCtx.TacticsParams,
					currentGeneral.BaseInfo.UniqueId,
					int64(tactic.Id),
					1,
					0,
					0,
				)
				//触发「发动指挥战法开始」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_CommandTactic]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}

				//战法执行
				execute.TacticsExecute(runCtx.Ctx, tacticHandler)

				//触发「发动指挥战法结束」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_CommandTacticEnd]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   consts.Battle_Round_Prepare,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
			}
		}
	}
}

func (runCtx *BattleLogicV2Context) handleGeneralAddition(team *vo.BattleTeam, general *vo.BattleGeneral) {
	//**基础属性加成**
	//武将加点加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 武将加点加成")
	runCtx.handleAbilityAttrAddition(general)
	//武将等级加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 武将等级加成")
	runCtx.handleGeneralLevel(general)
	//兵书加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 兵书加成")
	runCtx.handleWarBook(general)

	//剧本特殊效果加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 剧本特殊效果加成")
	//士气加成 (看了下战报分析了下，每减少0.1士气，降低伤害比例是0.07%，如果0士气则降低70%伤害，其余不影响)
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 士气加成")
	//兵种适性加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 兵种适性加成")
	runCtx.handleArmAbility(team.ArmType, general)
	//武将标签加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 武将标签加成")
	runCtx.handleGeneralTag(general)

	//城池建筑加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 城池建筑加成")
	//缘分加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 缘分加成")
	//装备加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 装备加成")
	//特技加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 特技加成")
	//兵书加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleGeneralAddition 兵书加成")
}

func (runCtx *BattleLogicV2Context) handleTeamAddition(team *vo.BattleTeam) {
	//兵战-科技加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleTeamAddition 兵战-科技加成")
	runCtx.handleBuildingTechAttrAddition(team)
	//协力-科技加成
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleTeamAddition 协力-科技加成")
	runCtx.handleBuildingTechGroupAddition(team)
	//兵种-科技加成 TODO
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] handleTeamAddition 兵种-科技加成")
}

// 对战战报统计数据处理
func (runCtx *BattleLogicV2Context) processBattleReportStatistics() {
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] processBattleReportStatistics...")

	//我军
	fightingGeneralsStatisticsList := make([]*model.GeneralBattleStatistics, 0)
	for _, general := range runCtx.Req.FightingTeam.BattleGenerals {
		//战法统计
		tacticStatisticsList := make([]*model.TacticStatistics, 0)
		tacticStatisticsMap, ok := runCtx.TacticsParams.BattleTacticStatisticsMap[general.BaseInfo.UniqueId]

		for _, tactic := range general.EquipTactics {
			tacticStatistics, okk := tacticStatisticsMap[int64(tactic.Id)]
			if ok && okk {
				tacticStatisticsList = append(tacticStatisticsList, &model.TacticStatistics{
					TacticId:              tacticStatistics.TacticId,
					TacticName:            tacticStatistics.TacticName,
					TriggerTimes:          tacticStatistics.TriggerTimes,
					KillSoliderNum:        tacticStatistics.KillSoliderNum,
					ResumeSoliderNum:      tacticStatistics.ResumeSoliderNum,
					RoundTriggerTimes:     tacticStatistics.RoundTriggerTimes,
					RoundKillSoliderNum:   tacticStatistics.RoundKillSoliderNum,
					RoundResumeSoliderNum: tacticStatistics.RoundResumeSoliderNum,
				})
			} else {
				tacticCache, okkk := cache.CacheTacticMap[int64(tactic.Id)]
				if !okkk {
					continue
				}
				tacticStatisticsList = append(tacticStatisticsList, &model.TacticStatistics{
					TacticId:   tacticCache.Id,
					TacticName: tacticCache.Name,
				})
			}
		}
		//普攻统计
		generalAttackStatistics := &model.TacticStatistics{}
		if attackStatistics, ok := runCtx.TacticsParams.BattleAttackStatisticsMap[general.BaseInfo.UniqueId]; ok {
			generalAttackStatistics.KillSoliderNum = attackStatistics.KillSoliderNum
			generalAttackStatistics.ResumeSoliderNum = attackStatistics.ResumeSoliderNum
			generalAttackStatistics.TriggerTimes = attackStatistics.TriggerTimes
			generalAttackStatistics.RoundTriggerTimes = attackStatistics.RoundTriggerTimes
			generalAttackStatistics.RoundKillSoliderNum = attackStatistics.RoundKillSoliderNum
			generalAttackStatistics.RoundResumeSoliderNum = attackStatistics.RoundResumeSoliderNum
		}

		fightingGeneralsStatisticsList = append(fightingGeneralsStatisticsList, &model.GeneralBattleStatistics{
			TacticStatisticsList:    tacticStatisticsList,
			GeneralAttackStatistics: generalAttackStatistics,
			RoundRemainSoliderNum:   general.RoundRemainSoliderNum,
		})
	}

	//敌军
	enemyGeneralsStatisticsList := make([]*model.GeneralBattleStatistics, 0)
	for _, general := range runCtx.Req.EnemyTeam.BattleGenerals {
		//战法统计
		tacticStatisticsList := make([]*model.TacticStatistics, 0)
		tacticStatisticsMap, ok := runCtx.TacticsParams.BattleTacticStatisticsMap[general.BaseInfo.UniqueId]

		for _, tactic := range general.EquipTactics {
			tacticStatistics, okk := tacticStatisticsMap[int64(tactic.Id)]
			if ok && okk {
				tacticStatisticsList = append(tacticStatisticsList, &model.TacticStatistics{
					TacticId:              tacticStatistics.TacticId,
					TacticName:            tacticStatistics.TacticName,
					TriggerTimes:          tacticStatistics.TriggerTimes,
					KillSoliderNum:        tacticStatistics.KillSoliderNum,
					ResumeSoliderNum:      tacticStatistics.ResumeSoliderNum,
					RoundTriggerTimes:     tacticStatistics.RoundTriggerTimes,
					RoundKillSoliderNum:   tacticStatistics.RoundKillSoliderNum,
					RoundResumeSoliderNum: tacticStatistics.RoundResumeSoliderNum,
				})
			} else {
				tacticCache, okkk := cache.CacheTacticMap[int64(tactic.Id)]
				if !okkk {
					continue
				}
				tacticStatisticsList = append(tacticStatisticsList, &model.TacticStatistics{
					TacticId:   tacticCache.Id,
					TacticName: tacticCache.Name,
				})
			}
		}
		//普攻统计
		generalAttackStatistics := &model.TacticStatistics{}
		if attackStatistics, ok := runCtx.TacticsParams.BattleAttackStatisticsMap[general.BaseInfo.UniqueId]; ok {
			generalAttackStatistics.KillSoliderNum = attackStatistics.KillSoliderNum
			generalAttackStatistics.ResumeSoliderNum = attackStatistics.ResumeSoliderNum
			generalAttackStatistics.TriggerTimes = attackStatistics.TriggerTimes
			generalAttackStatistics.RoundTriggerTimes = attackStatistics.RoundTriggerTimes
			generalAttackStatistics.RoundKillSoliderNum = attackStatistics.RoundKillSoliderNum
			generalAttackStatistics.RoundResumeSoliderNum = attackStatistics.RoundResumeSoliderNum
		}

		enemyGeneralsStatisticsList = append(enemyGeneralsStatisticsList, &model.GeneralBattleStatistics{
			TacticStatisticsList:    tacticStatisticsList,
			GeneralAttackStatistics: generalAttackStatistics,
			RoundRemainSoliderNum:   general.RoundRemainSoliderNum,
		})
	}

	runCtx.Resp.BattleResultStatistics = &model.BattleResultStatistics{
		//我军
		FightingTeam: &model.TeamBattleStatistics{
			BattleTeam: &vo.BattleTeam{
				TeamType:       runCtx.Req.FightingTeam.TeamType,
				ArmType:        runCtx.Req.FightingTeam.ArmType,
				BattleGenerals: runCtx.makeGeneralInfos(runCtx.Req.FightingTeam.BattleGenerals),
			},
			BattleResult:                util.JudgeBattleResult(runCtx.Req.FightingTeam, runCtx.Req.EnemyTeam),
			GeneralBattleStatisticsList: fightingGeneralsStatisticsList,
		},
		//敌军
		EnemyTeam: &model.TeamBattleStatistics{
			BattleTeam: &vo.BattleTeam{
				TeamType:       runCtx.Req.EnemyTeam.TeamType,
				ArmType:        runCtx.Req.EnemyTeam.ArmType,
				BattleGenerals: runCtx.makeGeneralInfos(runCtx.Req.EnemyTeam.BattleGenerals),
			},
			BattleResult:                util.JudgeBattleResult(runCtx.Req.EnemyTeam, runCtx.Req.FightingTeam),
			GeneralBattleStatisticsList: enemyGeneralsStatisticsList,
		},
	}
	//TODO 处理对战过程数据
	battleProcessStatistics := make(map[consts.BattlePhase]map[consts.BattleRound][]string, 0)

	runCtx.Resp.BattleProcessStatistics = battleProcessStatistics
}

func (runCtx *BattleLogicV2Context) makeGeneralInfos(battleGenerals []*vo.BattleGeneral) []*vo.BattleGeneral {
	//补充武将信息
	for _, general := range battleGenerals {
		generalCache := cache.CacheGeneralMap[general.BaseInfo.Id]
		general.BaseInfo.AvatarUri = generalCache.AvatarUrl
	}
	return battleGenerals
}

// 对战对阵阶段处理
func (runCtx *BattleLogicV2Context) processBattleFightingPhase() {
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] processBattleFightingPhase...")

	//对战阶段
	runCtx.TacticsParams.CurrentPhase = consts.Battle_Phase_Fighting

	hlog.CtxInfof(runCtx.Ctx, "<<<<<<<<<<<<<【对战阶段】>>>>>>>>>>>>>")

	//最多8回合
	currentRound := consts.Battle_Round_Unknow
	for i := 0; i < int(consts.Battle_Round_Eighth); i++ {
		//回合开始判断回合停止标签
		if runCtx.BattleRoundEndFlag {
			break
		}
		//重置每回合的中间数据
		runCtx.TacticsParams.CurrentResumeNum = 0
		runCtx.TacticsParams.CurrentDamageNum = 0
		//对战回合增加、设置
		currentRound++
		runCtx.TacticsParams.CurrentRound = currentRound

		//回合剩余兵力
		runCtx.handleRoundRemainSoliderNum()

		//对战回合处理
		runCtx.processBattleFightingRound(currentRound)
	}

	//打印每队战况
	hlog.CtxInfof(runCtx.Ctx, "******************《 战报总结 》********************")
	hlog.CtxInfof(runCtx.Ctx, "战斗结束 , 结束时回合数：%d", currentRound)
	for _, general := range runCtx.Req.FightingTeam.BattleGenerals {
		hlog.CtxInfof(runCtx.Ctx, "[%s]兵力[%d]", general.BaseInfo.Name, general.SoldierNum)
	}
	for _, general := range runCtx.Req.EnemyTeam.BattleGenerals {
		hlog.CtxInfof(runCtx.Ctx, "[%s]兵力[%d]", general.BaseInfo.Name, general.SoldierNum)
	}
}

func (runCtx *BattleLogicV2Context) handleRoundRemainSoliderNum() {
	//我方
	for _, general := range runCtx.TacticsParams.FightingTeam.BattleGenerals {
		if m, ok := general.RoundRemainSoliderNum[runCtx.TacticsParams.CurrentPhase]; ok {
			m[runCtx.TacticsParams.CurrentRound] = general.SoldierNum
		} else {
			newM := make(map[consts.BattleRound]int64, 0)
			newM[runCtx.TacticsParams.CurrentRound] = general.SoldierNum
			general.RoundRemainSoliderNum[runCtx.TacticsParams.CurrentPhase] = newM
		}
	}
	//敌方
	for _, general := range runCtx.TacticsParams.EnemyTeam.BattleGenerals {
		if m, ok := general.RoundRemainSoliderNum[runCtx.TacticsParams.CurrentPhase]; ok {
			m[runCtx.TacticsParams.CurrentRound] = general.SoldierNum
		} else {
			newM := make(map[consts.BattleRound]int64, 0)
			newM[runCtx.TacticsParams.CurrentRound] = general.SoldierNum
			general.RoundRemainSoliderNum[runCtx.TacticsParams.CurrentPhase] = newM
		}
	}
}

// 每回合对战处理
func (runCtx *BattleLogicV2Context) processBattleFightingRound(currentRound consts.BattleRound) {
	tacticsParams := runCtx.TacticsParams

	//判断先攻战法生效
	//判断是否有武将本回合有先攻战法生效
	//for _, general := range allGenerals {
	//	for _, tactic := range general.EquipTactics {
	//	}
	//}
	hlog.CtxInfof(runCtx.Ctx, "战斗回合：===========【%d】===========", currentRound)
	//每一轮重新排序，按速度执行
	tacticsParams.AllGeneralArr = util.MakeGeneralsOrderBySpeed(tacticsParams.AllGeneralArr)

	for _, currentGeneral := range tacticsParams.AllGeneralArr {
		//设置战法轮次属性
		runCtx.TacticsParams.CurrentGeneral = currentGeneral
		//普攻次数(默认一次)
		attackCanCnt := 1

		//武将回合前置处理器
		if !runCtx.GeneralRoundPreProcessor(tacticsParams) {
			continue
		}

		//打印当前执行队伍、武将、速度
		hlog.CtxInfof(runCtx.Ctx, "[%v][%s]开始行动",
			currentGeneral.BaseInfo.GeneralBattleType,
			currentGeneral.BaseInfo.Name,
		)
		hlog.CtxInfof(runCtx.Ctx, "[%v]兵力: %v",
			currentGeneral.BaseInfo.Name,
			currentGeneral.SoldierNum,
		)

		//是否可以行动
		if !util.IsCanBeginAction(runCtx.Ctx, currentGeneral) {
			continue
		}

		//触发「开始行动」触发器
		if funcs, ok := currentGeneral.TacticsTriggerMap[consts.BattleAction_BeginAction]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   currentRound,
					CurrentGeneral: currentGeneral,
					CurrentDamage:  tacticsParams.CurrentDamageNum,
					CurrentResume:  tacticsParams.CurrentResumeNum,
				}
				f(params)
			}
		}

		//战法发动顺序：1.主动 > 2.普攻 > 3.突击
		//按装配顺序执行主动战法
		for tacticSeq, tactic := range currentGeneral.EquipTactics {
			//1.主动
			if _, ok := consts.ActiveTacticsMap[tactic.Id]; ok {
				//战法参数设置
				runCtx.TacticsParams.TacticsType = consts.TacticsType_Active
				handler := tactics.TacticsHandlerMap[tactic.Id]
				tacticHandler := handler.Init(tacticsParams)

				//主动战法拦截
				if !util.IsCanActiveTactic(runCtx.Ctx, currentGeneral) {
					goto activeTacticFlag
				}

				//触发「发动主动战法前」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_ActiveTactic]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   currentRound,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						triggerResp := f(params)
						if triggerResp.IsTerminate {
							hlog.CtxInfof(runCtx.Ctx, "[%s]的【%s】被阻止了",
								currentGeneral.BaseInfo.Name,
								tactic.Name,
							)
							break
						}
					}
				}

				//已发动准备战法跳过
				if tacticHandler.IsTriggerPrepare() {
					hlog.CtxInfof(runCtx.Ctx, "准备战法:【%v】跳过", tactic.Name)
					continue
				}

				//发动率判断
				triggerRate := tacticHandler.GetTriggerRate()
				//发动率提升[主动战法]
				if effectParams, okk := util.BuffEffectGet(currentGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove); okk {
					for _, param := range effectParams {
						triggerRate += param.TriggerRate
						hlog.CtxInfof(runCtx.Ctx, "[%s]由于【%s】的影响，主动战法发动率提升%.2f%%",
							currentGeneral.BaseInfo.Name,
							param.FromTactic,
							param.TriggerRate*100,
						)
					}
				}
				//发动率提升[非自带]
				if effectParams, okk := util.BuffEffectGet(currentGeneral, consts.BuffEffectType_TacticsActiveTriggerNoSelfImprove); okk {
					if tacticSeq > 0 {
						for _, param := range effectParams {
							triggerRate += param.TriggerRate
							hlog.CtxInfof(runCtx.Ctx, "[%s]由于【%s】的影响，主动战法[非自带]发动率提升%.2f%%",
								currentGeneral.BaseInfo.Name,
								param.FromTactic,
								param.TriggerRate*100,
							)
						}
					}
				}
				//发动率提升[自带]
				if effectParams, okk := util.BuffEffectGet(currentGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove); okk {
					if tacticSeq == 0 {
						for _, param := range effectParams {
							triggerRate += param.TriggerRate
							hlog.CtxInfof(runCtx.Ctx, "[%s]由于【%s】的影响，主动战法[非自带]发动率提升%.2f%%",
								currentGeneral.BaseInfo.Name,
								param.FromTactic,
								param.TriggerRate*100,
							)
						}
					}
				}
				//发动率提升[准备战法]
				if consts.ActivePrepareTacticsMap[tactic.Id] {
					if effectParams, okk := util.BuffEffectGet(currentGeneral, consts.BuffEffectType_TacticsActiveTriggerPrepareImprove); okk {
						for _, param := range effectParams {
							triggerRate += param.TriggerRate
							hlog.CtxInfof(runCtx.Ctx, "[%s]由于【%s】的影响，主动战法[准备]发动率提升%.2f%%",
								currentGeneral.BaseInfo.Name,
								param.FromTactic,
								param.TriggerRate*100,
							)
						}
					}
				}
				//发动率降低
				if effectParams, okk := util.DeBuffEffectGet(currentGeneral, consts.DebuffEffectType_TacticsActiveTriggerDecr); okk {
					for _, param := range effectParams {
						triggerRate -= param.TriggerRate
						hlog.CtxInfof(runCtx.Ctx, "[%s]由于【%s】主动战法发动率降低%.2f%%",
							currentGeneral.BaseInfo.Name,
							param.FromTactic,
							param.TriggerRate*100,
						)
					}
				}
				if !util.GenerateRate(triggerRate) {
					continue
				}
				//战法执行
				execute.TacticsExecute(runCtx.Ctx, tacticHandler)

				//触发「发动主动战法后」触发器
				if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_ActiveTacticEnd]; okk {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:   currentRound,
							CurrentGeneral: currentGeneral,
							CurrentDamage:  0,
							CurrentTactic:  tacticHandler,
						}
						f(params)
					}
				}
				//武将清理
				util.RemoveGeneralWhenSoldierNumIsEmpty(tacticsParams)
				//战法后置处理器
				if !runCtx.TacticPostProcessor(tacticsParams) {
					runCtx.BattleRoundEndFlag = true
					goto exitFlag
				}
			}
		}
	activeTacticFlag:
		//2.普攻
		//连击效果设置
		if util.BuffEffectContains(currentGeneral, consts.BuffEffectType_ContinuousAttack) {
			attackCanCnt++
		}

		if attackCanCnt > 0 {
			for i := 0; i <= attackCanCnt; i++ {
				//2.1 普通攻击拦截
				if !util.IsCanGeneralAttack(runCtx.Ctx, currentGeneral) {
					goto AttactTacticFlag
				}

				//找到普攻目标
				sufferGeneral := util.GetEnemyOneGeneral(currentGeneral, tacticsParams)
				tacticsParams.CurrentSufferGeneral = sufferGeneral
				//普通攻击触发器
				if funcs, ok := currentGeneral.TacticsTriggerMap[consts.BattleAction_Attack]; ok {
					for _, f := range funcs {
						params := &vo.TacticsTriggerParams{
							CurrentRound:        currentRound,
							CurrentGeneral:      currentGeneral,
							AttackGeneral:       currentGeneral,
							SufferAttackGeneral: sufferGeneral,
						}
						f(params)
					}
				}

				//发起攻击
				damage.AttackDamage(tacticsParams, currentGeneral, sufferGeneral, 0)

				//群攻效果
				effectParams, ok := util.BuffEffectGet(currentGeneral, consts.BuffEffectType_GroupAttack)
				if ok {
					dmgRate := float64(0)
					for _, effectParam := range effectParams {
						dmgRate += effectParam.EffectRate
					}

					//群攻效果攻击敌军其他武将
					otherEnemyGenerals := util.GetEnemyGeneralsNotSelfByGeneral(sufferGeneral, tacticsParams)
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
					for _, enemyGeneral := range otherEnemyGenerals {
						damage.AttackDamage(tacticsParams, currentGeneral, enemyGeneral, dmg)
					}
				}
				//普攻次数减一
				attackCanCnt--

				//武将清理
				util.RemoveGeneralWhenSoldierNumIsEmpty(tacticsParams)

				//战法后置处理器
				if !runCtx.TacticPostProcessor(tacticsParams) {
					runCtx.BattleRoundEndFlag = true
					goto exitFlag
				}

				//3.突击战法
				for _, tactic := range currentGeneral.EquipTactics {
					if _, ok := consts.AssaultTacticsMap[tactic.Id]; ok {
						//战法参数设置
						runCtx.TacticsParams.TacticsType = consts.TacticsType_Assault

						handler := tactics.TacticsHandlerMap[tactic.Id]
						tacticHandler := handler.Init(tacticsParams)
						//发动率判断
						triggerRate := runCtx.getAssaultTriggerRate(tacticHandler, currentGeneral)
						if !util.GenerateRate(triggerRate) {
							continue
						}

						//触发「发动突击战法开始」触发器
						if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_AssaultTactic]; okk {
							for _, f := range funcs {
								params := &vo.TacticsTriggerParams{
									CurrentRound:   currentRound,
									CurrentGeneral: currentGeneral,
									CurrentDamage:  0,
									CurrentTactic:  tacticHandler,
								}
								f(params)
							}
						}

						//战法执行
						execute.TacticsExecute(runCtx.Ctx, tacticHandler)

						//触发「发动突击战法后」触发器
						if funcs, okk := currentGeneral.TacticsTriggerMap[consts.BattleAction_AssaultTacticEnd]; okk {
							for _, f := range funcs {
								params := &vo.TacticsTriggerParams{
									CurrentRound:   currentRound,
									CurrentGeneral: currentGeneral,
									CurrentDamage:  0,
									CurrentTactic:  tacticHandler,
								}
								f(params)
							}
						}

						//武将清理
						util.RemoveGeneralWhenSoldierNumIsEmpty(tacticsParams)

						//战法后置处理器
						if !runCtx.TacticPostProcessor(tacticsParams) {
							runCtx.BattleRoundEndFlag = true
							goto exitFlag
						}
					}
				}
			}
		}
	AttactTacticFlag:

		//触发「结束行动」触发器
		if funcs, ok := currentGeneral.TacticsTriggerMap[consts.BattleAction_EndAction]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   currentRound,
					CurrentGeneral: currentGeneral,
				}
				f(params)
			}
		}

		//武将回合结束处理器
		runCtx.GeneralRoundPostProcessor(tacticsParams)
	}

exitFlag:
}

// 获取突击战法触发率
func (runCtx *BattleLogicV2Context) getAssaultTriggerRate(tacticHandler _interface.Tactics, currentGeneral *vo.BattleGeneral) float64 {
	triggerRate := tacticHandler.GetTriggerRate()
	improveRate, _ := util.BuffEffectGetAggrTriggerRate(currentGeneral, consts.BuffEffectType_TacticsAssaultTriggerImprove)
	decrRate, _ := util.DeBuffEffectGetAggrTriggerRate(currentGeneral, consts.DebuffEffectType_TacticsAssaultTriggerDecr)
	triggerRate += improveRate - decrRate
	return triggerRate
}

// 战法执行前置处理器
func (runCtx BattleLogicV2Context) TacticPreProcessor(tacticsParams *model.TacticsParams) bool {
	//兵力校验
	if tacticsParams.CurrentGeneral.SoldierNum == 0 {
		return false
	}

	return true
}

// 武将回合执行前置处理器
func (runCtx BattleLogicV2Context) TacticPostProcessor(tacticsParams *model.TacticsParams) bool {
	//主将校验
	masterGeneralCnt := 0
	for _, general := range tacticsParams.AllGeneralMap {
		if general.IsMaster {
			masterGeneralCnt++
		}
	}
	if masterGeneralCnt != 2 {
		return false
	}

	return true
}

// 武将回合执行前置处理器
func (runCtx BattleLogicV2Context) GeneralRoundPreProcessor(tacticsParams *model.TacticsParams) bool {
	//兵力为0直接退出
	if tacticsParams.CurrentGeneral.SoldierNum == 0 {
		return false
	}
	return true
}

// 战法执行后置处理器
func (runCtx BattleLogicV2Context) GeneralRoundPostProcessor(tacticsParams *model.TacticsParams) {
}

func (runCtx *BattleLogicV2Context) buildBattleRoundParams() {
	hlog.CtxDebugf(runCtx.Ctx, "[BattleLogicV2Context] buildBattleRoundParams...")

	tacticsParams := &model.TacticsParams{}
	tacticsParams.CurrentRound = consts.Battle_Round_Unknow

	//对阵队伍信息
	tacticsParams.FightingTeam = runCtx.Req.FightingTeam
	tacticsParams.EnemyTeam = runCtx.Req.EnemyTeam

	//武将信息转map/arr，方便后续直接调用
	tacticsParams.FightingGeneralMap = make(map[string]*vo.BattleGeneral, 0)
	tacticsParams.EnemyGeneralMap = make(map[string]*vo.BattleGeneral, 0)
	tacticsParams.AllGeneralMap = make(map[string]*vo.BattleGeneral, 0)
	for _, general := range runCtx.Req.FightingTeam.BattleGenerals {
		tacticsParams.FightingGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralArr = append(tacticsParams.AllGeneralArr, general)
	}
	for _, general := range runCtx.Req.EnemyTeam.BattleGenerals {
		tacticsParams.EnemyGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralArr = append(tacticsParams.AllGeneralArr, general)
	}

	//初始化增益效果/负面效果
	for _, general := range tacticsParams.AllGeneralArr {
		if general.BuffEffectHolderMap == nil {
			general.BuffEffectHolderMap = map[consts.BuffEffectType][]*vo.EffectHolderParams{}
		}
		if general.DeBuffEffectHolderMap == nil {
			general.DeBuffEffectHolderMap = map[consts.DebuffEffectType][]*vo.EffectHolderParams{}
		}
		if general.TacticsTriggerMap == nil {
			general.TacticsTriggerMap = map[consts.BattleAction][]func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult{}
		}
		if general.BuffEffectCountMap == nil {
			general.BuffEffectCountMap = map[consts.BuffEffectType]int64{}
		}
		if general.DeBuffEffectCountMap == nil {
			general.DeBuffEffectCountMap = map[consts.DebuffEffectType]int64{}
		}

		if general.TacticFrozenMap == nil {
			general.TacticFrozenMap = map[consts.TacticId]bool{}
		}
		if general.TacticAccumulateDamageMap == nil {
			general.TacticAccumulateDamageMap = map[consts.TacticId]int64{}
		}
		if general.TacticAccumulateResumeMap == nil {
			general.TacticAccumulateResumeMap = map[consts.TacticId]int64{}
		}
		if general.TacticAccumulateTriggerMap == nil {
			general.TacticAccumulateTriggerMap = map[consts.TacticId]int64{}
		}
		if general.RoundRemainSoliderNum == nil {
			general.RoundRemainSoliderNum = map[consts.BattlePhase]map[consts.BattleRound]int64{}
		}
		if general.WarbookAccumulateResumeMap == nil {
			general.WarbookAccumulateResumeMap = map[consts.WarBookDetailType]int64{}
		}
		if general.WarbookAccumulateTriggerMap == nil {
			general.WarbookAccumulateTriggerMap = map[consts.WarBookDetailType]int64{}
		}
		if general.WarbookAccumulateDamageMap == nil {
			general.WarbookAccumulateDamageMap = map[consts.WarBookDetailType]int64{}
		}
	}

	//对战过程统计
	if tacticsParams.BattleProcessStatisticsMap == nil {
		tacticsParams.BattleProcessStatisticsMap = map[consts.BattlePhase]map[consts.BattleRound][]string{}
	}
	//对战战法统计
	if tacticsParams.BattleTacticStatisticsMap == nil {
		tacticsParams.BattleTacticStatisticsMap = map[string]map[int64]*model.TacticStatistics{}
	}
	//对战普攻统计
	if tacticsParams.BattleAttackStatisticsMap == nil {
		tacticsParams.BattleAttackStatisticsMap = map[string]*model.TacticStatistics{}
	}

	runCtx.TacticsParams = tacticsParams
}
