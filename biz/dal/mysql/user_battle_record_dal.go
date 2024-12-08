package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/spf13/cast"
)

func (UserBattleRecordDal) TableName() string {
	return "user_battle_record"
}

type UserBattleRecordDal struct {
}

func NewUserBattleRecord() *UserBattleRecordDal {
	return &UserBattleRecordDal{}
}

func (g *UserBattleRecordDal) QueryUserBattleRecord(ctx context.Context, uid string, offset, limit int) ([]*po.UserBattleRecord, error) {
	userBattleRecords := make([]*po.UserBattleRecord, 0)
	conn := DataBase.Model(&po.UserBattleRecord{})

	conn.Where("uid = ?", uid)

	if limit > 0 {
		conn = conn.Offset(offset).Limit(limit)
	}

	if err := conn.Order("id desc").Find(&userBattleRecords).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return userBattleRecords, err
	}
	return userBattleRecords, nil
}

func (g *UserBattleRecordDal) CreateUserBattleRecord(ctx context.Context, userBattleRecord *po.UserBattleRecord) error {
	conn := DataBase.Model(&po.UserBattleRecord{})

	if err := conn.Create(userBattleRecord).Error; err != nil {
		hlog.CtxErrorf(ctx, "CreateUserBattleRecord err:%v", err)
		return err
	}
	return nil
}

// 查询用户对战总胜率
func (g *UserBattleRecordDal) QueryUserBattleResult(ctx context.Context, uid int64) (map[int64]int64, error) {
	type QueryResult struct {
		Result int64 `json:"result"`
		Cnt    int64 `json:"cnt"`
	}
	resList := make([]*QueryResult, 0)
	resMap := make(map[int64]int64, 0)

	sql := `select t.result, count(1) cnt
from (select json_extract(battle_record, '$.BattleResultStatistics.FightingTeam.BattleResult') as result
      from user_battle_record
      where uid = %v) t
group by t.result`
	sql = fmt.Sprintf(sql, uid)

	err := DataBase.Raw(sql).Find(&resList).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return resMap, err
	}
	for _, vo := range resList {
		resMap[vo.Result] = vo.Cnt
	}

	return resMap, nil
}

// 查询用户高频战法
func (g *UserBattleRecordDal) QueryUserBattleHighFreqTacticStatistics(ctx context.Context, uid int64) (map[int64]int64, error) {
	resultMap := make(map[int64]int64, 0)
	var tacticIds string
	sql := `select case when length(t.tactic_ids) > 0 then t.tactic_ids else '' end as tactic_ids
			from (select group_concat(regexp_replace(json_extract(json_extract(
																		  json_extract(battle_record,
																					   '$.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals'),
																		  '$[*].EquipTactics'), '$[*][*].id'), ']|\\[',
													 '')) as tactic_ids
				  from user_battle_record
				  where uid = %v) as t`
	sql = fmt.Sprintf(sql, uid)

	err := DataBase.Raw(sql).Find(&tacticIds).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return resultMap, err
	}

	if tacticIds != "" {
		for _, tacticId := range strings.Split(tacticIds, ",") {
			tacticIdInt := cast.ToInt64(strings.TrimSpace(tacticId))
			if times, ok := resultMap[tacticIdInt]; ok {
				resultMap[tacticIdInt] = times + 1
			} else {
				resultMap[tacticIdInt] = 1
			}
		}
	}

	return resultMap, nil
}

// 查询用户常用武将
func (g *UserBattleRecordDal) QueryUserBattleHighFreqUsedGeneralStatistics(ctx context.Context, uid int64) (map[int64]int64, error) {
	resultMap := make(map[int64]int64, 0)
	var generalIds string

	sql := `select case when length(t.general_ids) > 0 then t.general_ids else '' end as general_ids
			from (select group_concat(regexp_replace(json_extract(json_extract(
																		  json_extract(battle_record,
																					   '$.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals'),
																		  '$[*].BaseInfo'), '$[*].Id'),
													 ']|\\[', '')) as general_ids
				  from user_battle_record
				  where uid = %v) as t`
	sql = fmt.Sprintf(sql, uid)

	err := DataBase.Raw(sql).Find(&generalIds).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return resultMap, err
	}

	if generalIds != "" {
		for _, generalId := range strings.Split(generalIds, ",") {
			generalIdInt := cast.ToInt64(strings.TrimSpace(generalId))
			if times, ok := resultMap[generalIdInt]; ok {
				resultMap[generalIdInt] = times + 1
			} else {
				resultMap[generalIdInt] = 1
			}
		}
	}

	return resultMap, nil
}

// 查询用户常用阵容
func (g *UserBattleRecordDal) QueryUserBattleHighFreqUsedTeamStatistics(ctx context.Context, uid int64) (map[string]int64, error) {
	resultMap := make(map[string]int64, 0)
	type Result struct {
		Team string `json:"team"`
		Cnt  int64  `json:"cnt"`
	}
	resultList := make([]*Result, 0)

	sql := `select team, cnt
			from (select a.team, count(1) as cnt
				  from (select json_extract(
									   json_extract(battle_record,
																 '$.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals'),
									   '$[*].BaseInfo.Id'
								   ) as team
						from user_battle_record where uid = %v) as a
				  group by a.team) as b
			order by b.cnt desc`
	sql = fmt.Sprintf(sql, uid)

	err := DataBase.Raw(sql).Find(&resultList).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return resultMap, err
	}

	if len(resultList) > 0 {
		for _, vo := range resultList {
			if times, ok := resultMap[vo.Team]; ok {
				resultMap[vo.Team] = vo.Cnt + times
			} else {
				resultMap[vo.Team] = vo.Cnt
			}

		}
	}

	return resultMap, nil
}
