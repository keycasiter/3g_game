package lottery

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"math"
	"time"
)

// 武将抽卡
type GeneralLotteryLogic struct {
	Ctx   context.Context
	Req   *vo.GeneralLotteryRequest
	Resp  *vo.GeneralLotteryResponse
	Funcs []func()
	Err   error

	//**中间变量**
	//武将+概率卡池 map<武将ID,抽中概率>
	GeneralPool map[consts.General_Id]float64
	//当前卡池五星武将
	General5LevMap map[consts.General_Id]float64
	//概率分布 map<概率1万等分槽,武将ID>
	ChancePool map[int64]consts.General_Id
	//五星武将分布 map<概率等分槽,武将ID>
	Chance5LevPool map[int64]consts.General_Id
	//命中武将 map<武将ID,命中次数>
	HitGeneralMap map[consts.General_Id]int64
	//命中武将 arr
	HitGeneralArr []int64
	//连续未命中五星武将次数
	NotHit5LevGeneralNum int64
	//保底次数
	ProtectedMustHitNum int64
}

func NewGeneralLotteryLogic(ctx context.Context, req *vo.GeneralLotteryRequest) *GeneralLotteryLogic {
	g := &GeneralLotteryLogic{
		Ctx:            ctx,
		Req:            req,
		GeneralPool:    map[consts.General_Id]float64{},
		ChancePool:     map[int64]consts.General_Id{},
		HitGeneralMap:  map[consts.General_Id]int64{},
		HitGeneralArr:  make([]int64, 0),
		General5LevMap: map[consts.General_Id]float64{},
		Chance5LevPool: map[int64]consts.General_Id{},
	}
	//组合方法
	g.Funcs = []func(){
		//获取用户抽卡数据
		g.FetchUserGeneralLotteryInfo,
		//获取卡池
		g.GetPool,
		//配置卡池
		g.ConfigPool,
		//抽取卡池
		g.LotteryPool,
		//组合结果
		g.BuildResp,
	}
	return g
}

func (g *GeneralLotteryLogic) Run() (*vo.GeneralLotteryResponse, error) {
	for _, f := range g.Funcs {
		f()
		if g.Err != nil {
			return nil, g.Err
		}
	}
	return g.Resp, nil
}

func (g *GeneralLotteryLogic) FetchUserGeneralLotteryInfo() {
	info, err := mysql.NewUserGeneralLotteryInfo().QueryUserGeneralLotteryInfo(g.Ctx, g.Req.Uid, int64(g.Req.GeneralLottery))
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryUserGeneralLotteryInfo err:%v", err)
		g.Err = err
		return
	}
	g.NotHit5LevGeneralNum = info.NotHitLev5Times
}

func (g *GeneralLotteryLogic) GetPool() {
	pool := consts.GetGeneralPool(g.Req.GeneralLottery)
	if pool == nil {
		g.Err = errors.New("卡池不存在")
	}
	if len(pool) == 0 {
		g.Err = errors.New("卡池为空")
	}
	for generalId, rate := range pool {
		g.GeneralPool[generalId] = rate
	}
}

func (g *GeneralLotteryLogic) ConfigPool() {
	//**全部武将配置池**
	//概率分布下限边界
	lowLimit := int64(0)
	//概率分布上限边界
	upperLimit := int64(0)
	//步进最小单位, 0.01%
	step := int64(1)

	cursor := lowLimit
	for generalId, lotteryChance := range g.GeneralPool {
		rangeScope := cast.ToInt64(lotteryChance * 10000)
		upperLimit = cursor + rangeScope
		for i := cursor; i <= upperLimit; i += step {
			g.ChancePool[i] = generalId
		}
		cursor += rangeScope
	}
	//hlog.CtxInfof(g.Req.Ctx,"ChancePool:%s",util.ToJsonString(g.Req.Ctx,g.ChancePool))
	hlog.CtxInfof(g.Ctx, "[ConfigPool] lowLimit:%d , upperLimit:%d ,ChancePool size:%d", lowLimit, cursor, len(g.ChancePool))

	//**五星武将配置池**
	//概率分布下限边界
	lowLimit5Lev := int64(0)
	//概率分布上限边界
	upperLimit5Lev := int64(0)

	cursor5Lev := lowLimit5Lev
	//当前卡池五星武将整理
	for generalId, lotteryChance := range g.GeneralPool {
		//只需要五星武将
		if !consts.General5LevMap[generalId] {
			continue
		}

		g.General5LevMap[generalId] = lotteryChance
		rangeScope := cast.ToInt64(lotteryChance * 10000)
		upperLimit5Lev = cursor5Lev + rangeScope
		for i := cursor5Lev; i <= upperLimit5Lev; i += step {
			g.Chance5LevPool[i] = generalId
		}
		cursor5Lev += rangeScope
	}
	hlog.CtxInfof(g.Ctx, "[Config5LevPool] lowLimit:%d , upperLimit5Lev:%d ,Chance5LevPool size:%d", lowLimit5Lev, cursor5Lev, len(g.Chance5LevPool))

}

func (g *GeneralLotteryLogic) LotteryPool() {
	for i := int64(0); i < g.Req.RollTimes; i++ {
		//累计30抽不中走保底逻辑
		if g.NotHit5LevGeneralNum == 30 {
			//保底逻辑
			g.hitGeneralHandler(true)
			//累计清零
			g.NotHit5LevGeneralNum = 0
			//保底统计
			g.ProtectedMustHitNum++

			//同步抽卡记录
			err := syncNotHit5LevDataToDB(g.Ctx, g.Req.Uid, int64(g.Req.GeneralLottery), 0)
			if err != nil {
				hlog.CtxErrorf(g.Ctx, "syncNotHit5LevDataToDB err:%v", err)
				g.Err = err
				return
			}
		} else {
			g.hitGeneralHandler(false)
		}
	}
}

func (g *GeneralLotteryLogic) hitGeneralHandler(isMustHit5Lev bool) {
	randomUpperLimit := float64(len(g.ChancePool))
	lotteryPool := g.ChancePool
	if isMustHit5Lev {
		randomUpperLimit = float64(len(g.Chance5LevPool))
		lotteryPool = g.Chance5LevPool
	}

	//考虑slot空白边界，不命中重新抽
	for {
		//生成随机数，按照等分池大小size来生成随机数进行命中
		random := cast.ToInt64(fmt.Sprintf("%.0f", math.Round(util.Random(0, randomUpperLimit))))
		if generalId, ok := lotteryPool[random]; ok {
			g.HitGeneralArr = append(g.HitGeneralArr, int64(generalId))
			if _, okk := g.HitGeneralMap[generalId]; okk {
				g.HitGeneralMap[generalId]++
			} else {
				g.HitGeneralMap[generalId] = 1
			}
			//非五星武将计数
			if _, okkk := g.General5LevMap[generalId]; okkk {
				g.NotHit5LevGeneralNum = 0
				err := syncNotHit5LevDataToDB(g.Ctx, g.Req.Uid, int64(g.Req.GeneralLottery), 0)
				if err != nil {
					hlog.CtxErrorf(g.Ctx, "syncNotHit5LevDataToDB err:%v", err)
					g.Err = err
					return
				}
			} else {
				g.NotHit5LevGeneralNum++
				err := syncNotHit5LevDataToDB(g.Ctx, g.Req.Uid, int64(g.Req.GeneralLottery), g.NotHit5LevGeneralNum)
				if err != nil {
					hlog.CtxErrorf(g.Ctx, "syncNotHit5LevDataToDB err:%v", err)
					g.Err = err
					return
				}
			}
			//hlog.CtxInfof(g.Req.Ctx, "[LotteryPool] random:%d , hit GeneralId:%d", random, int(generalId))
			break
		}
	}
}

func syncNotHit5LevDataToDB(ctx context.Context, uid string, cardPoolId int64, num int64) error {
	//查询用户卡池抽奖信息
	info, err := mysql.NewUserGeneralLotteryInfo().QueryUserGeneralLotteryInfo(ctx, uid, cardPoolId)
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserGeneralLotteryInfo err:%v", err)
		return err
	}
	//不存在
	if info == nil || info.Uid == "" {
		nowTime := time.Now()
		err = mysql.NewUserGeneralLotteryInfo().CreateUserGeneralLotteryInfo(ctx, &po.UserGeneralLotteryInfo{
			Uid:             uid,
			CardPoolId:      cardPoolId,
			NotHitLev5Times: num,
			CreatedAt:       nowTime,
			UpdatedAt:       nowTime,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "CreateUserGeneralLotteryInfo err:%v", err)
			return err
		}
	} else {
		//存在
		err = mysql.NewUserGeneralLotteryInfo().UpdateUserGeneralLotteryInfo(ctx, &po.UserGeneralLotteryInfo{
			Uid:             uid,
			CardPoolId:      cardPoolId,
			NotHitLev5Times: num,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "UpdateUserGeneralLotteryInfo err:%v", err)
			return err
		}
	}

	return nil
}

func (g *GeneralLotteryLogic) BuildResp() {
	//查询武将信息
	generalInfos := make([]*po.General, 0)
	generalIds := make([]int64, 0)
	for _, generalId := range g.HitGeneralArr {
		generalIds = append(generalIds, generalId)
	}

	//整理resp
	generalLotteryList := make([]*vo.GeneralLotteryInfo, 0)
	duplicateGeneral := make(map[int64]bool, 0)
	hit5LevGeneralNum := int64(0)
	for _, general := range generalInfos {
		//命中次数
		hitNum := int64(0)
		if currentHitNum, ok := g.HitGeneralMap[consts.General_Id(general.Id)]; ok {
			hitNum = currentHitNum

			//5星将
			if _, okk := g.General5LevMap[consts.General_Id(general.Id)]; okk {
				hit5LevGeneralNum += currentHitNum
			}
		}
		//配置概率
		lotteryRate := float64(0)
		if currentLotteryRate, ok := g.GeneralPool[consts.General_Id(general.Id)]; ok {
			lotteryRate = currentLotteryRate
		}
		//去重
		if duplicateGeneral[general.Id] {
			continue
		} else {
			duplicateGeneral[general.Id] = true
		}

		generalLotteryList = append(generalLotteryList, &vo.GeneralLotteryInfo{
			GeneralInfo: general,
			HitNum:      hitNum,
			HitRate:     cast.ToFloat64(hitNum) / cast.ToFloat64(g.Req.RollTimes),
			LotteryRate: lotteryRate,
		})
	}
	g.Resp = &vo.GeneralLotteryResponse{
		GeneralLotteryInfoList: generalLotteryList,
		ProtectedMustHitNum:    g.ProtectedMustHitNum,
		Hit5LevGeneralNum:      hit5LevGeneralNum,
		NotHitLev5Times:        g.NotHit5LevGeneralNum,
	}
}
