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

func (GeneralDal) TableName() string {
	return "general"
}

type GeneralDal struct {
}

func NewGeneral() *GeneralDal {
	return &GeneralDal{}
}

func (g *GeneralDal) QueryGeneralList(ctx context.Context, condition *vo.QueryGeneralCondition) ([]*po.General, error) {
	list := make([]*po.General, 0)
	conn := DataBase.Model(&po.General{})

	//条件查询
	if len(condition.Ids) > 0 {
		conn.Where("id in (?)", condition.Ids)
	}
	if condition.Id > 0 {
		conn.Where("id = ?", condition.Id)
	}
	if strings.Trim(condition.Name, " ") != "" {
		conn.Where("name like ?", fmt.Sprintf("%%%s%%", condition.Name))
	}
	if condition.Gender > 0 {
		conn.Where("gender = ?", condition.Gender)
	}
	if condition.Control > 0 {
		conn.Where("control = ?", condition.Control)
	}
	if condition.Group > 0 {
		conn.Where("`group` = ?", condition.Group)
	}
	if condition.Quality > 0 {
		conn.Where("quality = ?", condition.Quality)
	}
	if len(condition.Tags) > 0 {
		conn.Where("json_contains(tag->'$',?,'$')", util.ToIntArrayString(condition.Tags))
	}
	if condition.IsSupportCollect > 0 {
		conn.Where("is_support_collect = ?", condition.IsSupportCollect)
	}
	if condition.IsSupportDynamics > 0 {
		conn.Where("is_support_dynamics = ?", condition.IsSupportDynamics)
	}

	if err := conn.Offset(condition.Offset).Limit(condition.Limit).Find(&list).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryGeneralList err:%v", err)
		return list, err
	}
	return list, nil
}
