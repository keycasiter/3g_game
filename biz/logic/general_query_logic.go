package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

type GeneralQueryLogic struct {
	Ctx  context.Context
	Req  api.GeneralQueryRequest
	Resp api.GeneralQueryResponse
}

func NewGeneralQueryLogic(ctx context.Context, req api.GeneralQueryRequest) *GeneralQueryLogic {
	return &GeneralQueryLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.GeneralQueryResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *GeneralQueryLogic) Handle() (api.GeneralQueryResponse, error) {
	//查询武将列表
	list, err := mysql.NewGeneral().QueryGeneralList(g.Ctx, buildQueryGeneralListReq(g.Req))
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryGeneralList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//查询战法列表
	tacticMap := make(map[int64]*po.Tactic, 0)
	tacticIds := make([]int64, 0)
	for _, general := range list {
		tacticIds = append(tacticIds, int64(general.SelfTacticId))
	}
	tactics, err := mysql.NewTactic().QueryTacticList(g.Ctx, &vo.QueryTacticCondition{
		Ids:    tacticIds,
		Offset: 0,
		Limit:  len(tacticIds),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryTacticList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}
	for _, tactic := range tactics {
		tacticMap[tactic.Id] = tactic
	}

	//组合resp
	resList := make([]*api.BattleGeneral, 0)
	for _, general := range list {
		//武将标签
		generalTags := make([]enum.GeneralTag, 0)
		tags := []string{}
		util.ParseJsonObj(g.Ctx, &tags, general.Tag)
		for _, tag := range tags {
			generalTags = append(generalTags, enum.GeneralTag(cast.ToInt64(tag)))
		}
		//武将属性
		abilityAttrPo := &vo.AbilityAttr{}
		util.ParseJsonObj(g.Ctx, abilityAttrPo, general.AbilityAttr)
		abilityAttr := api.AbilityAttr{
			ForceBase:        cast.ToFloat64(abilityAttrPo.ForceBase),
			ForceRate:        cast.ToFloat64(abilityAttrPo.ForceRate),
			IntelligenceBase: cast.ToFloat64(abilityAttrPo.IntelligenceBase),
			IntelligenceRate: cast.ToFloat64(abilityAttrPo.IntelligenceRate),
			CharmBase:        cast.ToFloat64(abilityAttrPo.CharmBase),
			CharmRate:        cast.ToFloat64(abilityAttrPo.CharmRate),
			CommandBase:      cast.ToFloat64(abilityAttrPo.CommandBase),
			CommandRate:      cast.ToFloat64(abilityAttrPo.CommandRate),
			PoliticsBase:     cast.ToFloat64(abilityAttrPo.PoliticsBase),
			PoliticsRate:     cast.ToFloat64(abilityAttrPo.PoliticsRate),
			SpeedBase:        cast.ToFloat64(abilityAttrPo.SpeedBase),
			SpeedRate:        cast.ToFloat64(abilityAttrPo.SpeedRate),
		}

		//兵种属性
		armsAttrPo := &po.ArmsAttr{}
		util.ParseJsonObj(g.Ctx, armsAttrPo, general.ArmAttr)
		armsAttr := api.ArmsAttr{
			Cavalry:   util.ArmsAbilityToEnum(armsAttrPo.Cavalry),
			Mauler:    util.ArmsAbilityToEnum(armsAttrPo.Mauler),
			Archers:   util.ArmsAbilityToEnum(armsAttrPo.Archers),
			Spearman:  util.ArmsAbilityToEnum(armsAttrPo.Spearman),
			Apparatus: util.ArmsAbilityToEnum(armsAttrPo.Apparatus),
		}

		//战法
		tactic := &api.Tactics{}
		if v, ok := tacticMap[int64(general.SelfTacticId)]; ok {
			if ok {
				tactic = &api.Tactics{
					Id:            v.Id,
					Name:          v.Name,
					TacticsSource: enum.TacticsSource(v.Source),
					Type:          enum.TacticsType(v.Type),
					Quality:       enum.TacticQuality(v.Quality),
				}
			}
		}

		resList = append(resList, &api.BattleGeneral{
			BaseInfo: &api.MetadataGeneral{
				Id:                general.Id,
				Name:              general.Name,
				Gender:            enum.Gender(general.Gender),
				Group:             enum.Group(general.Group),
				GeneralTag:        generalTags,
				AvatarUrl:         general.AvatarUrl,
				AbilityAttr:       &abilityAttr,
				ArmsAttr:          &armsAttr,
				SelfTactic:        tactic,
				GeneralQuality:    enum.GeneralQuality(general.Quality),
				IsSupportDynamics: general.IsSupportDynamics == 1,
				IsSupportCollect:  general.IsSupportCollect == 1,
			},
		})
	}
	g.Resp.GeneralList = resList

	return g.Resp, nil
}

func buildQueryGeneralListReq(req api.GeneralQueryRequest) *vo.QueryGeneralCondition {
	tags := make([]int, 0)
	for _, tag := range req.GetTags() {
		tags = append(tags, int(tag))
	}

	return &vo.QueryGeneralCondition{
		Id:                req.GetId(),
		Name:              req.GetName(),
		Gender:            int8(req.GetGender()),
		Control:           int32(req.GetControl()),
		Group:             int8(req.GetGroup()),
		Quality:           int8(req.GetQuality()),
		Tags:              tags,
		IsSupportDynamics: int8(req.GetIsSupportDynamics()),
		IsSupportCollect:  int8(req.GetIsSupportCollect()),
		Offset:            util.PageNoToOffset(req.GetPageNo(), req.GetPageSize()),
		Limit:             int(req.GetPageSize()),
		Ids:               req.Ids,
	}
}
