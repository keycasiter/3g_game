package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"strings"
)

func (TeamDal) TableName() string {
	return "team"
}

type TeamDal struct {
}

func NewTeam() *TeamDal {
	return &TeamDal{}
}

func (g *TeamDal) QueryTeamList(ctx context.Context, condition *vo.QueryTeamCondition) ([]*po.Team, error) {
	list := make([]*po.Team, 0)
	conn := DataBase.Model(&po.Team{})

	//条件查询
	if condition.Id > 0 {
		conn.Where("id = ?", condition.Id)
	}
	if strings.Trim(condition.Name, " ") != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}
	if condition.Group > 0 {
		conn.Where("group = ?", condition.Group)
	}
	if len(condition.GeneralIds) > 0 {
		conn.Where("json_contains(general_ids->'$',?,'$')", util.ToIntArrayString(condition.GeneralIds))
	}

	if err := conn.Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryTeamList err:%v", err)
		return list, err
	}
	return list, nil
}
