package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

//req
type BattleLogicContextRequest struct {
	/** 队伍信息 **/
	// 出战队伍信息
	FightingTeam *vo.BattleTeam
	// 对战队伍信息
	EnemyTeam *vo.BattleTeam
}

//resp
type BattleLogicContextResponse struct {
}

//resp

// 对战上下文环境
type BattleLogicContext struct {
	//上下文
	Ctx context.Context
	// 入参
	ReqParam *BattleLogicContextRequest
	// 执行方法
	Funcs []func()
	// 执行错误
	RunErr error
}

func NewBattleLogicContext(ctx context.Context, req *BattleLogicContextRequest) *BattleLogicContext {
	runCtx := &BattleLogicContext{
		ReqParam: req,
	}
	//注入方法执行顺序
	runCtx.Funcs = []func(){
		//初始化元数据
		runCtx.initMetadata,
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

//属性加点处理
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

//武将等级处理
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

//武将标签处理
func (runCtx *BattleLogicContext) handleGeneralTag(general *vo.BattleGeneral) {
	//仙人：属性提高30%
	for _, generalTag := range general.BaseInfo.GeneralTag {
		switch generalTag {
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

//建筑科技-阵营加成处理
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
		general.BaseInfo.AbilityAttr.ForceBase = general.BaseInfo.AbilityAttr.ForceBase * additionRate
		//智力加成
		general.BaseInfo.AbilityAttr.IntelligenceBase = general.BaseInfo.AbilityAttr.IntelligenceBase * additionRate
		//统率加成
		general.BaseInfo.AbilityAttr.CommandBase = general.BaseInfo.AbilityAttr.CommandBase * additionRate
		//速度加成
		general.BaseInfo.AbilityAttr.SpeedBase = general.BaseInfo.AbilityAttr.SpeedBase * additionRate
	}
	hlog.CtxInfof(runCtx.Ctx, "[%s]队获得【%s】强化效果，属性提升了%d%s",
		team.BattleGenerals[0].BaseInfo.Name,
		util.TranslateGroup(consts.Group(group)),
		int(additionRate),
		"%",
	)
}

//建筑科技-属性加成处理
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

//兵种适性处理
func (runCtx *BattleLogicContext) handleArmAbility(teamArmType consts.ArmType, general *vo.BattleGeneral) {
	armType := consts.ArmType_Unknow
	switch teamArmType {
	//骑兵
	case consts.ArmType_Cavalry:
		armType = general.BaseInfo.ArmsAttr.Cavalry
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Cavalry, general.BaseInfo.AbilityAttr)
	//盾兵
	case consts.ArmType_Mauler:
		armType = general.BaseInfo.ArmsAttr.Mauler
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Mauler, general.BaseInfo.AbilityAttr)
	//弓兵
	case consts.ArmType_Archers:
		armType = general.BaseInfo.ArmsAttr.Archers
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Archers, general.BaseInfo.AbilityAttr)
	//枪兵
	case consts.ArmType_Spearman:
		armType = general.BaseInfo.ArmsAttr.Spearman
		util.CalGeneralArmAbility(general.BaseInfo.ArmsAttr.Spearman, general.BaseInfo.AbilityAttr)
	//器械
	case consts.ArmType_Apparatus:
		armType = general.BaseInfo.ArmsAttr.Apparatus
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
	//出战武将加成处理
	for _, general := range runCtx.ReqParam.FightingTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.FightingTeam, general)
		runCtx.handleTeamAddition(runCtx.ReqParam.FightingTeam)
	}

	//对战武将加成处理
	for _, general := range runCtx.ReqParam.EnemyTeam.BattleGenerals {
		runCtx.handleGeneralAddition(runCtx.ReqParam.EnemyTeam, general)
		runCtx.handleTeamAddition(runCtx.ReqParam.EnemyTeam)
	}

	//hlog.CtxInfof(runCtx.Ctx, "fighting team => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.FightingTeam))
	//hlog.CtxInfof(runCtx.Ctx, "enemy team => %s", util.ToJsonString(runCtx.Ctx, runCtx.ReqParam.EnemyTeam))
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

}
