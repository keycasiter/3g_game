package user

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/keycasiter/3g_game/biz/dal/cache"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/util"
	"golang.org/x/sync/errgroup"
)

// UserInfoDetail @Summary 查询用户信息详情
// @Description 查询用户信息详情
// @Tags 用户
// @Accept json
// @Produce json
// @Router /v1/user/detail [GET]
func UserInfoDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoDetailRequest
	var resp api.UserInfoDetailResponse
	resp.Meta = util.BuildSuccMeta()

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(hertzconsts.StatusBadRequest, err.Error())
		return
	}

	////1.获取登录用户的openId
	//hlog.CtxInfof(ctx, "GetUserWxOpenId Req:%s", util.ToJsonString(ctx, req))
	//requestUrl := fmt.Sprintf(conf.GetConfig().Wexin.GetUserOpenIdApiTemplateUrl,
	//	conf.GetConfig().Wexin.AppId,
	//	conf.GetConfig().Wexin.Secret,
	//	req.Code,
	//)
	//urlResp, err := http.Get(requestUrl)
	//if err != nil {
	//	hlog.CtxErrorf(ctx, "Http Get URL:%s err:%v", requestUrl, err)
	//	resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("err:%v", err))
	//	c.JSON(hertzconsts.StatusOK, resp)
	//	return
	//}
	//defer urlResp.Body.Close()
	//body, err := ioutil.ReadAll(urlResp.Body)
	//if err != nil {
	//	hlog.CtxErrorf(ctx, "Http Get URL:%s err:%v", requestUrl, err)
	//	resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("err:%v", err))
	//	c.JSON(hertzconsts.StatusOK, resp)
	//	return
	//}
	//respObj := &GetUserOpenIdResponse{}
	//err = json.Unmarshal(body, respObj)
	//if err != nil {
	//	hlog.CtxErrorf(ctx, "parse Object err:%v", err)
	//	resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("parse Object err:%v", err))
	//	c.JSON(hertzconsts.StatusOK, resp)
	//	return
	//}
	//hlog.CtxInfof(ctx, "GetUserWxOpenId Resp:%s", util.ToJsonString(ctx, respObj))
	//if respObj.OpenId == "" {
	//	hlog.CtxErrorf(ctx, "openId is empty")
	//	resp.Meta = util.BuildFailMetaWithMsg("微信openId为空")
	//	c.JSON(hertzconsts.StatusOK, resp)
	//	return
	//}

	//2.用户信息查询
	userInfo, err := mysql.NewUserInfo().QueryUserInfo(ctx, req.WxOpenId)
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserInfo err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("查询用户信息失败 err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//3.对战数据查询
	winRate := float64(0)
	highFreqTacticMap := make(map[int64]int64, 0)
	highFreqGeneralMap := make(map[int64]int64, 0)
	highFreqTeamMap := make(map[string]int64, 0)

	var eg errgroup.Group
	eg.SetLimit(4)
	//获取总胜率
	eg.Go(func() error {
		winRate, err = mysql.NewUserBattleRecord().QueryUserBattleWinRate(ctx, userInfo.Uid)
		if err != nil {
			hlog.CtxErrorf(ctx, "QueryUserBattleWinRate err:%v", err)
			return err
		}
		return nil
	})
	//获取使用的高频战法
	eg.Go(func() error {
		highFreqTacticMap, err = mysql.NewUserBattleRecord().QueryUserBattleHighFreqTacticStatistics(ctx, userInfo.Uid)
		if err != nil {
			hlog.CtxErrorf(ctx, "QueryUserBattleHighFreqTacticStatistics err:%v", err)
			return err
		}
		return nil
	})

	//获取使用的高频武将
	eg.Go(func() error {
		highFreqGeneralMap, err = mysql.NewUserBattleRecord().QueryUserBattleHighFreqUsedGeneralStatistics(ctx, userInfo.Uid)
		if err != nil {
			hlog.CtxErrorf(ctx, "QueryUserBattleHighFreqUsedGeneralStatistics err:%v", err)
			return err
		}
		return nil
	})

	//获取使用的高频阵容
	eg.Go(func() error {
		highFreqTeamMap, err = mysql.NewUserBattleRecord().QueryUserBattleHighFreqUsedTeamStatistics(ctx, userInfo.Uid)
		if err != nil {
			hlog.CtxErrorf(ctx, "QueryUserBattleHighFreqUsedTeamStatistics err:%v", err)
			return err
		}
		return nil
	})
	if err = eg.Wait(); err != nil {
		hlog.CtxErrorf(ctx, "parallel wait err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("查询用户信息失败 err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//组合resp
	resp.UserInfo = &api.UserInfo{
		Uid:       userInfo.Uid,
		NickName:  userInfo.NickName,
		AvatarUrl: userInfo.AvatarUrl,
		WxOpenId:  userInfo.WxOpenId,
		Level:     int64(userInfo.Level),
	}
	resp.BattleStatisticsInfo = &api.BattleStatisticsInfo{
		HighFreqGeneralList: buildHighFreqGeneralList(highFreqGeneralMap),
		HighFreqTacticsList: buildHighFreqTacticsList(highFreqTacticMap),
		HighFreqTeamList:    buildHighFreqTeamList(highFreqTeamMap),
		WinRate:             winRate,
	}

	c.JSON(hertzconsts.StatusOK, resp)
}

func buildHighFreqGeneralList(m map[int64]int64) []*api.GeneralRecord {
	generalRecords := make([]*api.GeneralRecord, 0)
	for generalId, times := range m {
		if generalCache, ok := cache.CacheGeneralMap[generalId]; ok {
			generalRecords = append(generalRecords, &api.GeneralRecord{
				General: &api.MetadataGeneral{
					Name:      generalCache.Name,
					Gender:    enum.Gender(generalCache.Gender),
					Group:     enum.Group(generalCache.Group),
					AvatarUrl: generalCache.AvatarUrl,
				},
				Times: times,
			})
		}
	}
	//按次数排序
	sort.SliceStable(generalRecords, func(i, j int) bool {
		return generalRecords[i].Times > generalRecords[j].Times
	})

	if len(generalRecords) > 3 {
		generalRecords = generalRecords[:3]
	}

	return generalRecords
}

func buildHighFreqTacticsList(m map[int64]int64) []*api.TacticsRecord {
	tacticsRecords := make([]*api.TacticsRecord, 0)
	for tacticId, times := range m {
		if tacticCache, ok := cache.CacheTacticMap[tacticId]; ok {
			tacticsRecords = append(tacticsRecords, &api.TacticsRecord{
				Tactics: &api.Tactics{
					Id:            tacticCache.Id,
					Name:          tacticCache.Name,
					TacticsSource: enum.TacticsSource(tacticCache.Source),
					Type:          enum.TacticsType(tacticCache.Type),
					Quality:       enum.TacticQuality(tacticCache.Quality),
				},
				Times: times,
			})
		}
	}
	//按次数排序
	sort.SliceStable(tacticsRecords, func(i, j int) bool {
		return tacticsRecords[i].Times > tacticsRecords[j].Times
	})

	if len(tacticsRecords) > 3 {
		tacticsRecords = tacticsRecords[:3]
	}

	return tacticsRecords
}

func buildHighFreqTeamList(m map[string]int64) []*api.TeamRecord {
	teamRecords := make([]*api.TeamRecord, 0)
	for generalIdsStr, times := range m {
		generalIds := make([]int64, 0)
		json.Unmarshal([]byte(generalIdsStr), &generalIds)

		//处理武将
		generalInfos := make([]*po.General, 0)
		for _, generalId := range generalIds {
			if generalCache, ok := cache.CacheGeneralMap[generalId]; ok {
				generalInfos = append(generalInfos, generalCache)
			}
		}
		teamRecords = append(teamRecords, &api.TeamRecord{
			BattleTeam: buildBattleTeam(generalInfos),
			Times:      times,
		})
	}
	//按次数排序
	sort.SliceStable(teamRecords, func(i, j int) bool {
		return teamRecords[i].Times > teamRecords[j].Times
	})

	if len(teamRecords) > 3 {
		teamRecords = teamRecords[:3]
	}

	return teamRecords
}

func buildBattleTeam(generalInfos []*po.General) *api.BattleTeam {
	battleTeam := &api.BattleTeam{}
	generalNames := make([]string, 0)
	for _, generalInfo := range generalInfos {
		battleTeam.BattleGenerals = append(battleTeam.BattleGenerals, &api.BattleGeneral{
			BaseInfo: &api.MetadataGeneral{
				Id:        generalInfo.Id,
				Name:      generalInfo.Name,
				Gender:    enum.Gender(generalInfo.Gender),
				Group:     enum.Group(generalInfo.Group),
				AvatarUrl: generalInfo.AvatarUrl,
			},
		})
		generalNames = append(generalNames, generalInfo.Name)
	}

	battleTeam.Name = strings.Join(generalNames, " ")

	return battleTeam
}
