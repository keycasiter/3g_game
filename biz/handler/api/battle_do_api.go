// Code generated by hertz generator.

package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/logic/battle"
	api "github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/common"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/kr/pretty"
	"github.com/spf13/cast"
)

// @title 三战配将侯小程序
// @version 1.0
// @description 三战配将侯小程序

// @contact.name 3g_game
// @contact.url keycasiter@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 3game.tech
// @BasePath /
// @schemes http

// BattleDo @Summary 发起模拟对战
// @Description 发起模拟对战
// @Tags 模拟对战
// @Accept json
// @Produce json
// @@Success 0
// @Failure
// @Router /v1/battle/do [POST]
func BattleDo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.BattleDoRequest
	resp := new(api.BattleDoResponse)
	resp.Meta = util.BuildSuccMeta()

	err = c.BindAndValidate(&req)
	//日志打印
	hlog.CtxInfof(ctx, "BattleExecute Req:%s", util.ToJsonString(ctx, req))

	if err != nil {
		hlog.CtxErrorf(ctx, "BattleLogicContext Run Err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg("未知错误")
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//参数校验
	if err := checkParam(req); err != nil {
		hlog.CtxWarnf(ctx, "BattleLogicContext Param Err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg("未知错误")
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//逻辑执行
	serviceResp, err := battle.NewBattleLogicContext(ctx, buildBattleDoRequest(ctx, req)).Run()
	if err != nil {
		hlog.CtxErrorf(ctx, "BattleLogicContext Run Err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg("未知错误")
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//组合resp
	copier.Copy(resp, serviceResp)
	buildResponse(resp, serviceResp)

	c.JSON(hertzconsts.StatusOK, resp)
	//日志打印
	pretty.Logf("resp:%s", util.ToJsonString(ctx, resp))
}

func buildResponse(resp *api.BattleDoResponse, serviceResp *battle.BattleLogicContextResponse) {
	resp.Meta = &common.Meta{
		StatusCode: enum.ResponseCode_Success,
		StatusMsg:  "成功",
	}
	//组合过程数据
	resp.BattleProcessStatistics = makeBattleProcessStatistics(serviceResp)
	//组合统计数据
	resp.BattleResultStatistics = makeBattleResultStatistics(serviceResp)
}

// 战报数据
func makeBattleProcessStatistics(serviceResp *battle.BattleLogicContextResponse) map[int64]map[int64][]string {
	battleProcessStatistics := make(map[int64]map[int64][]string, 0)
	for battlePhase, battleRoundStatisticsMap := range serviceResp.BattleProcessStatistics {
		m := make(map[int64][]string, 0)
		for round, strings := range battleRoundStatisticsMap {
			m[cast.ToInt64(round)] = strings
		}
		battleProcessStatistics[int64(battlePhase)] = m
	}
	return battleProcessStatistics
}

// 对战统计数据
func makeBattleResultStatistics(serviceResp *battle.BattleLogicContextResponse) *api.BattleResultStatistics {
	return &api.BattleResultStatistics{
		//我军
		FightingTeam: &api.TeamBattleStatistics{
			BattleTeam: &api.BattleTeam{
				TeamType:       enum.TeamType(serviceResp.BattleResultStatistics.FightingTeam.BattleTeam.TeamType),
				ArmType:        enum.ArmType(serviceResp.BattleResultStatistics.FightingTeam.BattleTeam.ArmType),
				BattleGenerals: makeBattleGenerals(serviceResp.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals),
				SoliderNum:     makeSoliderNum(serviceResp.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals),
				RemainNum:      makeRemainNum(serviceResp.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals),
			},
			BattleResult:                int64(serviceResp.BattleResultStatistics.FightingTeam.BattleResult),
			GeneralBattleStatisticsList: makeGeneralBattleStatisticsList(serviceResp.BattleResultStatistics.FightingTeam.GeneralBattleStatisticsList),
		},
		//敌军
		EnemyTeam: &api.TeamBattleStatistics{
			BattleTeam: &api.BattleTeam{
				TeamType:       enum.TeamType(serviceResp.BattleResultStatistics.EnemyTeam.BattleTeam.TeamType),
				ArmType:        enum.ArmType(serviceResp.BattleResultStatistics.EnemyTeam.BattleTeam.ArmType),
				BattleGenerals: makeBattleGenerals(serviceResp.BattleResultStatistics.EnemyTeam.BattleTeam.BattleGenerals),
				SoliderNum:     makeSoliderNum(serviceResp.BattleResultStatistics.EnemyTeam.BattleTeam.BattleGenerals),
				RemainNum:      makeRemainNum(serviceResp.BattleResultStatistics.EnemyTeam.BattleTeam.BattleGenerals),
			},
			BattleResult:                int64(serviceResp.BattleResultStatistics.EnemyTeam.BattleResult),
			GeneralBattleStatisticsList: makeGeneralBattleStatisticsList(serviceResp.BattleResultStatistics.EnemyTeam.GeneralBattleStatisticsList),
		},
	}
}

func makeGeneralBattleStatisticsList(statisticsList []*model.GeneralBattleStatistics) []*api.GeneralBattleStatistics {
	resList := make([]*api.GeneralBattleStatistics, 0)
	for _, statistics := range statisticsList {
		tacticStatisticsList := make([]*api.TacticStatistics, 0)
		//战法统计
		for _, tacticStatistics := range statistics.TacticStatisticsList {
			tacticStatisticsList = append(tacticStatisticsList, &api.TacticStatistics{
				TacticId:         tacticStatistics.TacticId,
				TacticName:       tacticStatistics.TacticName,
				TacticQuality:    tacticStatistics.TacticQuality,
				TriggerTimes:     tacticStatistics.TriggerTimes,
				KillSoliderNum:   tacticStatistics.KillSoliderNum,
				ResumeSoliderNum: tacticStatistics.ResumeSoliderNum,
			})
		}
		//普攻
		attackStatistics := &api.TacticStatistics{
			TriggerTimes:     statistics.GeneralAttackStatistics.TriggerTimes,
			KillSoliderNum:   statistics.GeneralAttackStatistics.KillSoliderNum,
			ResumeSoliderNum: statistics.GeneralAttackStatistics.ResumeSoliderNum,
		}
		resList = append(resList, &api.GeneralBattleStatistics{
			TacticStatisticsList:    tacticStatisticsList,
			GeneralAttackStatistics: attackStatistics,
		})
	}

	return resList
}

func makeSoliderNum(battleGenerals []*vo.BattleGeneral) int64 {
	teamSoliderNum := int64(0)
	for _, general := range battleGenerals {
		teamSoliderNum += general.InitSoldierNum
	}
	return teamSoliderNum
}

func makeRemainNum(battleGenerals []*vo.BattleGeneral) int64 {
	teamRemainNum := int64(0)
	for _, general := range battleGenerals {
		teamRemainNum += general.SoldierNum
	}
	return teamRemainNum
}

func makeBattleGenerals(battleGenerals []*vo.BattleGeneral) []*api.BattleGeneral {
	resList := make([]*api.BattleGeneral, 0)
	for _, general := range battleGenerals {
		resList = append(resList, &api.BattleGeneral{
			BaseInfo: &api.MetadataGeneral{
				Id:          general.BaseInfo.Id,
				Name:        general.BaseInfo.Name,
				Group:       enum.Group(general.BaseInfo.Group),
				GeneralTag:  makeGeneralTag(general.BaseInfo.GeneralTag),
				AvatarUrl:   general.BaseInfo.AvatarUri,
				AbilityAttr: makeAbilityAttr(general),
				ArmsAttr:    makeArmsAttr(general),
			},
			IsMaster:   general.IsMaster,
			SoldierNum: general.InitSoldierNum,
			RemainNum:  general.SoldierNum,
		})
	}
	return resList
}

func makeAbilityAttr(general *vo.BattleGeneral) *api.AbilityAttr {
	return &api.AbilityAttr{
		ForceBase:        cast.ToString(general.BaseInfo.AbilityAttr.ForceBase),
		IntelligenceBase: cast.ToString(general.BaseInfo.AbilityAttr.IntelligenceBase),
		CommandBase:      cast.ToString(general.BaseInfo.AbilityAttr.CommandBase),
		SpeedBase:        cast.ToString(general.BaseInfo.AbilityAttr.SpeedBase),
	}
}

func makeArmsAttr(general *vo.BattleGeneral) *api.ArmsAttr {
	return &api.ArmsAttr{
		Cavalry:   util.ArmsAbilityToEnum(general.BaseInfo.ArmsAttr.Cavalry),
		Mauler:    util.ArmsAbilityToEnum(general.BaseInfo.ArmsAttr.Mauler),
		Archers:   util.ArmsAbilityToEnum(general.BaseInfo.ArmsAttr.Archers),
		Spearman:  util.ArmsAbilityToEnum(general.BaseInfo.ArmsAttr.Spearman),
		Apparatus: util.ArmsAbilityToEnum(general.BaseInfo.ArmsAttr.Apparatus),
	}
}

func buildBattleDoRequest(ctx context.Context, req api.BattleDoRequest) *battle.BattleLogicContextRequest {
	//查询武将信息
	generalIds := make([]int64, 0)
	for _, general := range req.FightingTeam.BattleGenerals {
		generalIds = append(generalIds, general.BaseInfo.Id)
	}
	for _, general := range req.EnemyTeam.BattleGenerals {
		generalIds = append(generalIds, general.BaseInfo.Id)
	}
	generals, err := mysql.NewGeneral().QueryGeneralList(ctx, &vo.QueryGeneralCondition{
		Ids:    generalIds,
		Offset: 0,
		Limit:  len(generalIds),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryGeneralList err:%v", err)
		return nil
	}
	//整理成map
	generalInfoMap := make(map[int64]*po.General, 0)
	for _, general := range generals {
		generalInfoMap[general.Id] = general
	}

	//我方信息
	fightingTeamGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range req.FightingTeam.BattleGenerals {
		//我方武将信息（从db获取）
		generalInfo := &po.General{}
		if info, ok := generalInfoMap[general.BaseInfo.Id]; ok {
			generalInfo = info
		} else {
			panic(any(fmt.Sprintf("武将信息不存在:%d", general.BaseInfo.Id)))
		}
		//武将标签
		generalTags := make([]consts.GeneralTag, 0)
		generalTagsArr := []string{}
		util.ParseJsonObj(ctx, &generalTagsArr, generalInfo.Tag)
		for _, tag := range generalTagsArr {
			generalTags = append(generalTags, consts.GeneralTag(cast.ToInt(tag)))
		}
		//佩戴战法(自带战法+选择战法)
		equipTactics := make([]*po.Tactics, 0)
		for _, tactic := range general.EquipTactics {
			equipTactics = append(equipTactics, &po.Tactics{
				Id:            consts.TacticId(tactic.Id),
				Name:          fmt.Sprintf("%v", consts.TacticId(tactic.Id)),
				TacticsSource: consts.TacticsSource(tactic.TacticsSource),
				Type:          consts.TacticsType(tactic.Type),
			})
		}
		//武将信息
		fightingTeamGenerals = append(fightingTeamGenerals, &vo.BattleGeneral{
			//基础信息
			BaseInfo: &po.MetadataGeneral{
				//ID，从参数获取
				Id: general.BaseInfo.Id,
				//姓名，从db获取
				Name: generalInfo.Name,
				//阵营，从db获取
				Group: consts.Group(generalInfo.Group),
				//标签，从db获取
				GeneralTag: generalTags,
				//能力值，从db获取
				AbilityAttr: buildAbilityAttr(ctx, generalInfo),
				//兵种适性，从参数获取
				ArmsAttr: buildArmsAttr(general),
				//我军/敌军，从参数获取
				GeneralBattleType: consts.GeneralBattleType(general.BaseInfo.GeneralBattleType),
				//本局对战唯一ID，系统每次生成
				UniqueId: util.GenerateUUID(),
			},
			//战法
			EquipTactics: equipTactics,
			//加点
			Addition: &vo.BattleGeneralAddition{
				//属性加点，从参数获取
				AbilityAttr: po.AbilityAttr{
					ForceBase:        cast.ToFloat64(general.Addition.AbilityAttr.ForceBase),
					IntelligenceBase: cast.ToFloat64(general.Addition.AbilityAttr.IntelligenceBase),
					CommandBase:      cast.ToFloat64(general.Addition.AbilityAttr.CommandBase),
					SpeedBase:        cast.ToFloat64(general.Addition.AbilityAttr.SpeedBase),
				},
				//武将等级，从参数获取
				GeneralLevel: consts.GeneralLevel(general.Addition.GeneralLevel),
				//武将红度，从参数获取
				GeneralStarLevel: consts.GeneralStarLevel(general.Addition.GeneralStarLevel),
			},
			//是否主将，从参数获取
			IsMaster: general.IsMaster,
			//携带兵力，从参数获取
			SoldierNum: general.SoldierNum,
			//原始兵力，从参数获取
			InitSoldierNum: general.SoldierNum,
		})
	}

	//敌方信息
	enemyTeamGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range req.EnemyTeam.BattleGenerals {
		//我方武将信息（从db获取）
		generalInfo := &po.General{}
		if info, ok := generalInfoMap[general.BaseInfo.Id]; ok {
			generalInfo = info
		} else {
			panic(any(fmt.Sprintf("武将信息不存在:%d", general.BaseInfo.Id)))
		}
		//武将标签
		generalTags := make([]consts.GeneralTag, 0)
		generalTagsArr := []string{}
		util.ParseJsonObj(ctx, &generalTagsArr, generalInfo.Tag)
		for _, tag := range generalTagsArr {
			generalTags = append(generalTags, consts.GeneralTag(cast.ToInt(tag)))
		}
		//佩戴战法(自带战法+选择战法)
		equipTactics := make([]*po.Tactics, 0)
		for _, tactic := range general.EquipTactics {
			equipTactics = append(equipTactics, &po.Tactics{
				Id:            consts.TacticId(tactic.Id),
				Name:          fmt.Sprintf("%v", tactic.Id),
				TacticsSource: consts.TacticsSource(tactic.TacticsSource),
				Type:          consts.TacticsType(tactic.Type),
			})
		}
		//武将信息
		enemyTeamGenerals = append(enemyTeamGenerals, &vo.BattleGeneral{
			//基础信息
			BaseInfo: &po.MetadataGeneral{
				//ID，从参数获取
				Id: general.BaseInfo.Id,
				//姓名，从db获取
				Name: generalInfo.Name,
				//阵营，从db获取
				Group: consts.Group(generalInfo.Group),
				//标签，从db获取
				GeneralTag: generalTags,
				//能力值，从db获取
				AbilityAttr: buildAbilityAttr(ctx, generalInfo),
				//兵种适性，从参数获取
				ArmsAttr: buildArmsAttr(general),
				//我军/敌军，从参数获取
				GeneralBattleType: consts.GeneralBattleType(general.BaseInfo.GeneralBattleType),
				//本局对战唯一ID，系统每次生成
				UniqueId: util.GenerateUUID(),
			},
			//战法
			EquipTactics: equipTactics,
			//加点
			Addition: &vo.BattleGeneralAddition{
				//属性加点，从参数获取
				AbilityAttr: po.AbilityAttr{
					ForceBase:        cast.ToFloat64(general.Addition.AbilityAttr.ForceBase),
					IntelligenceBase: cast.ToFloat64(general.Addition.AbilityAttr.IntelligenceBase),
					CommandBase:      cast.ToFloat64(general.Addition.AbilityAttr.CommandBase),
					SpeedBase:        cast.ToFloat64(general.Addition.AbilityAttr.SpeedBase),
				},
				//武将等级，从参数获取
				GeneralLevel: consts.GeneralLevel(general.Addition.GeneralLevel),
				//武将红度，从参数获取
				GeneralStarLevel: consts.GeneralStarLevel(general.Addition.GeneralStarLevel),
			},
			//是否主将，从参数获取
			IsMaster: general.IsMaster,
			//携带兵力，从参数获取
			SoldierNum: general.SoldierNum,
			//原始兵力，从参数获取
			InitSoldierNum: general.SoldierNum,
		})
	}

	//组装
	serviceReq := &battle.BattleLogicContextRequest{
		//我方
		FightingTeam: &vo.BattleTeam{
			TeamType:                  consts.TeamType(req.FightingTeam.TeamType),
			ArmType:                   consts.ArmType(req.FightingTeam.ArmType),
			BattleGenerals:            fightingTeamGenerals,
			BuildingTechAttrAddition:  makeBuildingTechAttrAddition(req.FightingTeam),
			BuildingTechGroupAddition: makeBuildingTechGroupAddition(req.FightingTeam),
		},
		//敌方
		EnemyTeam: &vo.BattleTeam{
			TeamType:                  consts.TeamType(req.EnemyTeam.TeamType),
			ArmType:                   consts.ArmType(req.EnemyTeam.ArmType),
			BattleGenerals:            enemyTeamGenerals,
			BuildingTechAttrAddition:  makeBuildingTechAttrAddition(req.EnemyTeam),
			BuildingTechGroupAddition: makeBuildingTechGroupAddition(req.EnemyTeam),
		},
	}
	return serviceReq
}

func makeBuildingTechGroupAddition(team *api.BattleTeam) vo.BuildingTechGroupAddition {
	if team.BuildingTechGroupAddition == nil {
		return vo.BuildingTechGroupAddition{}
	}
	return vo.BuildingTechGroupAddition{
		GroupWeiGuoRate:   cast.ToFloat64(team.BuildingTechGroupAddition.GroupWuGuoRate),
		GroupShuGuoRate:   cast.ToFloat64(team.BuildingTechGroupAddition.GroupShuGuoRate),
		GroupWuGuoRate:    cast.ToFloat64(team.BuildingTechGroupAddition.GroupWuGuoRate),
		GroupQunXiongRate: cast.ToFloat64(team.BuildingTechGroupAddition.GroupQunXiongRate),
	}
}

func makeGeneralTag(tags []consts.GeneralTag) []enum.GeneralTag {
	resList := make([]enum.GeneralTag, 0)
	for _, tag := range tags {
		resList = append(resList, enum.GeneralTag(tag))
	}
	return resList
}

func makeBuildingTechAttrAddition(team *api.BattleTeam) vo.BuildingTechAttrAddition {
	if team.BuildingTechAttrAddition == nil {
		return vo.BuildingTechAttrAddition{}
	}

	return vo.BuildingTechAttrAddition{
		ForceAddition:        cast.ToFloat64(team.BuildingTechAttrAddition.ForceAddition),
		IntelligenceAddition: cast.ToFloat64(team.BuildingTechAttrAddition.IntelligenceAddition),
		CommandAddition:      cast.ToFloat64(team.BuildingTechAttrAddition.CommandAddition),
		SpeedAddition:        cast.ToFloat64(team.BuildingTechAttrAddition.SpeedAddition),
	}
}

func buildArmsAttr(general *api.BattleGeneral) *po.ArmsAttr {
	return &po.ArmsAttr{
		Cavalry:   consts.ArmsAbility(cast.ToString(general.BaseInfo.ArmsAttr.Cavalry)),
		Mauler:    consts.ArmsAbility(cast.ToString(general.BaseInfo.ArmsAttr.Mauler)),
		Archers:   consts.ArmsAbility(cast.ToString(general.BaseInfo.ArmsAttr.Archers)),
		Spearman:  consts.ArmsAbility(cast.ToString(general.BaseInfo.ArmsAttr.Spearman)),
		Apparatus: consts.ArmsAbility(cast.ToString(general.BaseInfo.ArmsAttr.Apparatus)),
	}
}

func buildAbilityAttr(ctx context.Context, generalInfo *po.General) *po.AbilityAttr {
	attr := &po.AbilityAttrString{}
	util.ParseJsonObj(ctx, attr, generalInfo.AbilityAttr)
	return &po.AbilityAttr{
		ForceBase:        cast.ToFloat64(attr.ForceBase),
		ForceRate:        cast.ToFloat64(attr.ForceRate),
		IntelligenceBase: cast.ToFloat64(attr.IntelligenceBase),
		IntelligenceRate: cast.ToFloat64(attr.IntelligenceRate),
		CommandBase:      cast.ToFloat64(attr.CommandBase),
		CommandRate:      cast.ToFloat64(attr.CommandRate),
		SpeedBase:        cast.ToFloat64(attr.SpeedBase),
		SpeedRate:        cast.ToFloat64(attr.SpeedRate),
	}
}

func checkParam(req api.BattleDoRequest) error {
	//队伍
	if req.FightingTeam == nil {
		return errors.New("我军队伍不能为空")
	}
	if req.EnemyTeam == nil {
		return errors.New("敌军队伍不能为空")
	}
	//兵种
	if req.FightingTeam.ArmType == enum.ArmType_Unknow {
		return errors.New("我军队伍兵种不能为空")
	}
	if req.EnemyTeam.ArmType == enum.ArmType_Unknow {
		return errors.New("敌军队伍兵种不能为空")
	}
	// 武将数量
	if len(req.FightingTeam.BattleGenerals) == 0 {
		return errors.New("我军武将数量不能为空")
	}
	if len(req.EnemyTeam.BattleGenerals) == 0 {
		return errors.New("敌军武将数量不能为空")
	}
	return nil
}