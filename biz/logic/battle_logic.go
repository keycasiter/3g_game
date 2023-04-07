package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics"
	"github.com/keycasiter/3g_game/biz/tactics/execute"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// req
type BattleLogicContextRequest struct {
	/** 队伍信息 **/
	// 出战队伍信息
	FightingTeam *vo.BattleTeam
	// 对战队伍信息
	EnemyTeam *vo.BattleTeam
}

// resp
type BattleLogicContextResponse struct {
}

//resp

// 对战上下文环境
type BattleLogicContext struct {
	/** DSL数据 **/
	//上下文
	Ctx context.Context
	// 入参
	ReqParam *BattleLogicContextRequest
	// 执行方法
	Funcs []func()
	// 执行错误
	RunErr error

	/** 回合对战全局变量 **/
	//对战战法全局holder
	TacticsParams *model.TacticsParams
	//回合结束标记
	BattleRoundEndFlag bool
}

func NewBattleLogicContext(ctx context.Context, req *BattleLogicContextRequest) *BattleLogicContext {
	runCtx := &BattleLogicContext{
		ReqParam: req,
	}
	//注入方法执行顺序
	runCtx.Funcs = []func(){
		//初始化元数据
		runCtx.initMetadata,
		//构建对战战法数据
		runCtx.buildBattleRoundParams,
		//对战准备阶段处理
		runCtx.processBattlePreparePhase,
		//对战对阵阶段处理
		runCtx.processBattleFightingPhase,
	}

	return runCtx
}

func (runCtx *BattleLogicContext) Run() error {
	for _, f := range runCtx.Funcs {
		f()
		if runCtx.RunErr != nil {
			hlog.CtxErrorf(runCtx.Ctx, "BattleLogicContext Func=%v Run err:%v", f, runCtx.RunErr)
			return runCtx.RunErr
		}
	}
	return nil
}

func (runCtx *BattleLogicContext) initMetadata() {

}

// 属性加点处理
func (runCtx *BattleLogicContext) handleAbilityAttrAddition(general *vo.BattleGeneral) {
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

// 武将等级处理
func (runCtx *BattleLogicContext) handleGeneralLevel(general *vo.BattleGeneral) {
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
func (runCtx *BattleLogicContext) handleGeneralTag(general *vo.BattleGeneral) {
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
		}
		hlog.CtxInfof(runCtx.Ctx, "[%s]发动战法【仙人】", general.BaseInfo.Name)
		hlog.CtxInfof(runCtx.Ctx, "[%s]是一名【仙人】，属性提高30%", general.BaseInfo.Name)
	}
}

// 建筑科技-阵营加成处理
func (runCtx *BattleLogicContext) handleBuildingTechGroupAddition(team *vo.BattleTeam) {
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
func (runCtx *BattleLogicContext) handleBuildingTechAttrAddition(team *vo.BattleTeam) {
	for _, general := range team.BattleGenerals {
		//武力加成
		general.BaseInfo.AbilityAttr.ForceBase = general.BaseInfo.AbilityAttr.ForceBase +
			team.BuildingTechAttrAddition.ForceAddition
		//智力加成
		general.BaseInfo.AbilityAttr.IntelligenceBase = general.BaseInfo.AbilityAttr.IntelligenceBase +
			team.BuildingTechAttrAddition.IntelligenceAddition
		//统率加成
		general.BaseInfo.AbilityAttr.CommandBase = general.BaseInfo.AbilityAttr.CommandBase +
			team.BuildingTechAttrAddition.CommandAddition
		//速度加成
		general.BaseInfo.AbilityAttr.SpeedBase = general.BaseInfo.AbilityAttr.SpeedBase +
			team.BuildingTechAttrAddition.SpeedAddition
	}
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-武】强化效果，【武力】属性提升了%d",
		team.BattleGenerals[0].BaseInfo.Name,
		int(team.BuildingTechAttrAddition.ForceAddition),
	)
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-谋】强化效果，【智力】属性提升了%d",
		team.BattleGenerals[0].BaseInfo.Name,
		int(team.BuildingTechAttrAddition.ForceAddition),
	)
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-防】强化效果，【统率】属性提升了%d",
		team.BattleGenerals[0].BaseInfo.Name,
		int(team.BuildingTechAttrAddition.CommandAddition),
	)
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【兵战-速】强化效果，【速度】属性提升了%d",
		team.BattleGenerals[0].BaseInfo.Name,
		int(team.BuildingTechAttrAddition.SpeedAddition),
	)
}

// 兵种适性处理
func (runCtx *BattleLogicContext) handleArmAbility(teamArmType consts.ArmType, general *vo.BattleGeneral) {
	armType := consts.ArmType_Unknow
	switch teamArmType {
	//骑兵
	case consts.ArmType_Cavalry:
		armType = consts.ArmType(general.BaseInfo.ArmsAttr.Cavalry)
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Cavalry, general.BaseInfo.AbilityAttr)
	//盾兵
	case consts.ArmType_Mauler:
		armType = consts.ArmType(general.BaseInfo.ArmsAttr.Mauler)
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Mauler, general.BaseInfo.AbilityAttr)
	//弓兵
	case consts.ArmType_Archers:
		armType = consts.ArmType(general.BaseInfo.ArmsAttr.Archers)
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Archers, general.BaseInfo.AbilityAttr)
	//枪兵
	case consts.ArmType_Spearman:
		armType = consts.ArmType(general.BaseInfo.ArmsAttr.Spearman)
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Spearman, general.BaseInfo.AbilityAttr)
	//器械
	case consts.ArmType_Apparatus:
		armType = consts.ArmType(general.BaseInfo.ArmsAttr.Apparatus)
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Apparatus, general.BaseInfo.AbilityAttr)
	}

	hlog.CtxInfof(runCtx.Ctx, "[%s]的【 %s 】兵种适性为【 %s 】，属性调整为原来的%s",
		general.BaseInfo.Name,
		util.TranslateArmType(teamArmType),
		util.TranslateArmsAbility(consts.ArmsAbility(armType)),
		util.TranslateArmsAbilityAddition(consts.ArmsAbility(armType)),
	)
}

// 对战准备阶段处理
func (runCtx *BattleLogicContext) processBattlePreparePhase() {
	tacticsParams := runCtx.TacticsParams

	/********************************************************/
	/*** 以下不受武将速度影响来执行，根据我方先/敌方后的顺序执行即可 ***/
	/********************************************************/

	//我方武将加成处理
	for _, general := range runCtx.ReqParam.FightingTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.FightingTeam, general)
		runCtx.handleTeamAddition(runCtx.ReqParam.FightingTeam)
	}

	//敌方武将加成处理
	for _, general := range runCtx.ReqParam.EnemyTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.EnemyTeam, general)
		runCtx.handleTeamAddition(runCtx.ReqParam.EnemyTeam)
	}

	//兵书处理：一兵书效果

	/****************************/
	/*** 以下受武将速度影响来执行 ***/
	/****************************/
	tacticsParams.AllGeneralArr = util.MakeGeneralsOrderBySpeed(tacticsParams.AllGeneralArr)

	//战法参数准备
	for _, currentGeneral := range tacticsParams.AllGeneralArr {
		//设置战法轮次属性
		runCtx.TacticsParams.CurrentRound = consts.Battle_Round_Prepare
		runCtx.TacticsParams.CurrentGeneral = currentGeneral

		//执行战法
		for _, tactic := range currentGeneral.EquipTactics {
			//战法发动顺序：1.被动 > 2.阵法 > 3.兵种 > 4.指挥 > 5.主动 > 6.普攻 > 7.突击
			//1.被动
			if _, ok := tactics.PassiveTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)
			}
			//2.阵法
			if _, ok := tactics.TroopsTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)
			}
			//3.兵种
			if _, ok := tactics.ArmTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)
			}
			//4.指挥
			if _, ok := tactics.CommandTacticsMap[tactic.Id]; ok {
				handler := tactics.TacticsHandlerMap[tactic.Id]
				execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)
			}
		}
	}

}

func (runCtx *BattleLogicContext) handleGeneralAddition(team *vo.BattleTeam, general *vo.BattleGeneral) {
	//1.国土效果
	//TODO 部队兵力不足可携带兵力上限的65%，国土效果不生效
	//2.士气加成
	//TODO 士气满不影响任何东西，不满100，则降低伤害，其余不影响；
	//看了下战报分析了下，每减少0.1士气，降低伤害比例是0.07%，如果0士气则降低70%伤害，其余不影响
	//3.兵种适性加成
	runCtx.handleArmAbility(team.ArmType, general)
	//4.武将标签的加成
	runCtx.handleGeneralTag(general)
	//武将加点加成
	runCtx.handleAbilityAttrAddition(general)
	//武将等级加成
	runCtx.handleGeneralLevel(general)
	//缘分加成
	//装备加成
	//特技加成
	//兵书加成
}

func (runCtx *BattleLogicContext) handleTeamAddition(team *vo.BattleTeam) {
	//1.兵战-科技加成
	runCtx.handleBuildingTechAttrAddition(team)
	//2.协力-科技加成
	runCtx.handleBuildingTechGroupAddition(team)
	//3.兵种-科技加成 TODO
}

// 对战对阵阶段处理
func (runCtx *BattleLogicContext) processBattleFightingPhase() {
	//最多8回合
	currentRound := consts.Battle_Round_Unknow
	for i := 0; i < int(consts.Battle_Round_Eighth); i++ {
		if runCtx.BattleRoundEndFlag {
			break
		}
		currentRound++
		runCtx.processBattleFightingRound(currentRound)
	}

	//打印每队战况
	hlog.CtxInfof(runCtx.Ctx, "******************《 战报 》********************")
	hlog.CtxInfof(runCtx.Ctx, "战斗结束 , 结束时回合数：%d", currentRound)
	for _, general := range runCtx.ReqParam.FightingTeam.BattleGenerals {
		hlog.CtxInfof(runCtx.Ctx, "[%s]兵力[%d]", general.BaseInfo.Name, general.SoldierNum)
	}
	for _, general := range runCtx.ReqParam.EnemyTeam.BattleGenerals {
		hlog.CtxInfof(runCtx.Ctx, "[%s]兵力[%d]", general.BaseInfo.Name, general.SoldierNum)
	}
}

// 每回合对战处理
func (runCtx *BattleLogicContext) processBattleFightingRound(currentRound consts.BattleRound) {
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
		runCtx.TacticsParams.CurrentRound = currentRound
		runCtx.TacticsParams.CurrentGeneral = currentGeneral

		//武将回合前置处理器
		if !runCtx.GeneralRoundPreProcessor(tacticsParams) {
			continue
		}

		//打印当前执行队伍、武将、速度
		hlog.CtxInfof(runCtx.Ctx, "[%v][%s]开始行动",
			currentGeneral.BaseInfo.GeneralBattleType,
			currentGeneral.BaseInfo.Name,
		)
		//触发「开始行动」战法效果
		if funcs, ok := currentGeneral.TacticsTriggerMap[consts.BattleAction_BeginAction]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   currentRound,
					CurrentGeneral: currentGeneral,
					CurrentDamage:  0,
				}
				f(params)
			}
		}

		//战法发动顺序：1.主动 > 2.普攻 > 3.突击
		//按装配顺序执行主动战法
		for _, tactic := range currentGeneral.EquipTactics {
			//1.主动
			if _, ok := tactics.ActiveTacticsMap[tactic.Id]; ok {
				//主动战法拦截
				if _, ok := currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_NoStrategy]; ok {
					hlog.CtxInfof(runCtx.Ctx, "武将[%s]处于[负面]计穷状态，无法发动主动战法", currentGeneral.BaseInfo.Name)
				}
				if _, ok := currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_PoorHealth]; ok {
					hlog.CtxInfof(runCtx.Ctx, "武将[%s]处于[负面]虚弱状态，无法发动主动战法", currentGeneral.BaseInfo.Name)
				}

				handler := tactics.TacticsHandlerMap[tactic.Id]
				execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)

				//武将清理
				util.RemoveGeneralWhenSoldierNumIsEmpty(tacticsParams)
				//战法后置处理器
				if !runCtx.TacticPostProcessor(tacticsParams) {
					runCtx.BattleRoundEndFlag = true
					goto exitFlag
				}
			}
		}
		//2.普攻
		//普攻次数(默认一次)
		attackCanCnt := 1
		if attackCanCnt > 0 {
			//2.1 普通攻击拦截
			if _, ok := currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_PoorHealth]; ok {
				hlog.CtxInfof(runCtx.Ctx, "武将[%s]处于「%v」状态，无法普通攻击",
					currentGeneral.BaseInfo.Name,
					consts.DebuffEffectType_PoorHealth,
				)
				return
			}
			if _, ok := currentGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_CanNotGeneralAttack]; ok {
				hlog.CtxInfof(runCtx.Ctx, "武将[%s]处于「%v」状态，无法普通攻击",
					currentGeneral.BaseInfo.Name,
					consts.DebuffEffectType_CanNotGeneralAttack,
				)
				return
			}
			//2.2 触发兵书效果
			//TODO
			//2.3 触发战法效果
			//负面效果
			//增益效果
			//战法触发器
			if funcs, ok := currentGeneral.TacticsTriggerMap[consts.BattleAction_Attack]; ok {
				for _, f := range funcs {
					params := &vo.TacticsTriggerParams{}
					f(params)
				}
			}

			//找到普攻目标
			sufferGeneral := util.GetEnemyOneGeneral(tacticsParams)
			//发起攻击
			util.AttackDamage(tacticsParams, currentGeneral, sufferGeneral)
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
				if _, ok := tactics.AssaultTacticsMap[tactic.Id]; ok {
					handler := tactics.TacticsHandlerMap[tactic.Id]
					execute.TacticsExecute(runCtx.Ctx, handler, tacticsParams)

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

		//武将回合结束处理器
		runCtx.GeneralRoundPostProcessor(tacticsParams)
	}

exitFlag:
}

// 战法执行前置处理器
func (runCtx BattleLogicContext) TacticPreProcessor(tacticsParams *model.TacticsParams) bool {
	//兵力校验
	if tacticsParams.CurrentGeneral.SoldierNum == 0 {
		return false
	}

	return true
}

// 武将回合执行前置处理器
func (runCtx BattleLogicContext) TacticPostProcessor(tacticsParams *model.TacticsParams) bool {
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
func (runCtx BattleLogicContext) GeneralRoundPreProcessor(tacticsParams *model.TacticsParams) bool {
	//兵力为0直接退出
	if tacticsParams.CurrentGeneral.SoldierNum == 0 {
		return false
	}
	return true
}

// 战法执行后置处理器
func (runCtx BattleLogicContext) GeneralRoundPostProcessor(tacticsParams *model.TacticsParams) {
	//所有冷却战法-1
	for _, general := range tacticsParams.AllGeneralMap {
		for tacticId, cnt := range general.TacticsFrozenMap {
			if cnt > 0 {
				general.TacticsFrozenMap[tacticId]--
			}
		}
	}
	return
}

func (runCtx *BattleLogicContext) buildBattleRoundParams() {
	tacticsParams := &model.TacticsParams{}
	tacticsParams.CurrentRound = consts.Battle_Round_Unknow

	//武将信息转map/arr，方便后续直接调用
	tacticsParams.FightingGeneralMap = make(map[int64]*vo.BattleGeneral, 0)
	tacticsParams.EnemyGeneralMap = make(map[int64]*vo.BattleGeneral, 0)
	tacticsParams.AllGeneralMap = make(map[int64]*vo.BattleGeneral, 0)
	for _, general := range runCtx.ReqParam.FightingTeam.BattleGenerals {
		tacticsParams.FightingGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralArr = append(tacticsParams.AllGeneralArr, general)
	}
	for _, general := range runCtx.ReqParam.EnemyTeam.BattleGenerals {
		tacticsParams.EnemyGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralMap[general.BaseInfo.UniqueId] = general
		tacticsParams.AllGeneralArr = append(tacticsParams.AllGeneralArr, general)
	}

	//初始化增益效果/负面效果
	for _, general := range tacticsParams.AllGeneralArr {
		if general.BuffEffectHolderMap == nil {
			general.BuffEffectHolderMap = map[consts.BuffEffectType]float64{}
		}
		if general.DeBuffEffectHolderMap == nil {
			general.DeBuffEffectHolderMap = map[consts.DebuffEffectType]float64{}
		}
		if general.BuffEffectCountMap == nil {
			general.BuffEffectCountMap = map[consts.BuffEffectType]int64{}
		}
		if general.DeBuffEffectCountMap == nil {
			general.DeBuffEffectCountMap = map[consts.DebuffEffectType]int64{}
		}
		if general.TacticsTriggerMap == nil {
			general.TacticsTriggerMap = map[consts.BattleAction][]func(params *vo.TacticsTriggerParams){}
		}
		if general.TacticsFrozenMap == nil {
			general.TacticsFrozenMap = map[consts.TacticId]int64{}
		}
	}

	runCtx.TacticsParams = tacticsParams
}
