package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"gorm.io/gorm/clause"
)

func (JymGoodsDal) TableName() string {
	return "jym_goods"
}

type JymGoodsDal struct {
}

func NewJymGoods() *JymGoodsDal {
	return &JymGoodsDal{}
}

func (g *JymGoodsDal) BatchSaveJymGoods(ctx context.Context, pos []*po.JymGoods) error {
	err := DataBase.Model(&po.JymGoods{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(pos, 10).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "[DB] batchSaveJymGoods CreateInBatches err:%v", err)
		return err
	}
	return nil
}

//根据战法查询
func (g JymGoodsDal) QueryJymGoodsBySkill(ctx context.Context, skillNames []string) ([]*po.JymGoods, error) {
	sql := "select item.goods_url from jym_goods item," +
		"JSON_TABLE(goods_detail, '$.apiData.itemLingxiRoleDetail.s3RoleCustomizeInfo.skills[*]' COLUMNS (\n " +
		" skillId int PATH '$.skillId',\n        " +
		" name VARCHAR(255) PATH '$.name'\n   " +
		" )) AS skill_table\n"
	if len(skillNames) > 0 {
		sql += " where "
		for i, name := range skillNames {
			sql += "skill_table.name =\"" + name + "\""
			if len(skillNames) > 1 && i < len(skillNames)-1 {
				sql += " and "
			}
		}
	}
	resList := make([]*po.JymGoods, 0)
	err := DataBase.Model(&po.JymGoods{}).Exec(sql).Find(&resList).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "[DB] QueryJymGoodsBySkill err:%v", err)
		return nil, err
	}
	return resList, nil
}

//根据武将查询
func (g JymGoodsDal) QueryJymGoodsByHero(ctx context.Context, skillNames []string) ([]*po.JymGoods, error) {
	sql := "select item.goods_url,hero_table.*\n" +
		"from jym_goods item\n        " +
		",JSON_TABLE(goods_detail, '$.apiData.itemLingxiRoleDetail.s3RoleCustomizeInfo.heros[*]' COLUMNS (\n " +
		"id int PATH '$.id',\n" +
		"name VARCHAR(255) PATH '$.name',\n" +
		"star int PATH '$.star',\n" +
		"is_awake bool PATH '$.isAwake',\n" +
		"camp int PATH '$.camp',\n" +
		"stage int PATH '$.stage',\n" +
		"isUnlockTalent bool PATH '$.isUnlockTalent',\n" +
		"isUnlockTalent3 bool PATH '$.isUnlockTalent3'\n" +
		")) AS hero_table"
	if len(skillNames) > 0 {
		sql += " where "
		for i, name := range skillNames {
			sql += "skill_table.name =\"" + name + "\""
			if len(skillNames) > 1 && i < len(skillNames)-1 {
				sql += " and "
			}
		}
	}
	resList := make([]*po.JymGoods, 0)
	err := DataBase.Model(&po.JymGoods{}).Exec(sql).Find(&resList).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "[DB] QueryJymGoodsBySkill err:%v", err)
		return nil, err
	}
	return resList, nil
}
