package lottery

import (
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"math"
)

//武将抽卡
type GeneralLotteryContext struct {
	Req   *GeneralLotteryRequest
	Resp  *GeneralLotteryResponse
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
	//连续未命中五星武将次数
	NotHit5LevGeneralNum int64
	//保底次数
	ProtectedMustHitNum int64
}

func NewGeneralLotteryContext(req *GeneralLotteryRequest) *GeneralLotteryContext {
	ctx := &GeneralLotteryContext{
		Req:            req,
		GeneralPool:    map[consts.General_Id]float64{},
		ChancePool:     map[int64]consts.General_Id{},
		HitGeneralMap:  map[consts.General_Id]int64{},
		General5LevMap: map[consts.General_Id]float64{},
		Chance5LevPool: map[int64]consts.General_Id{},
	}
	//组合方法
	ctx.Funcs = []func(){
		//获取卡池
		ctx.GetPool,
		//配置卡池
		ctx.ConfigPool,
		//抽取卡池
		ctx.LotteryPool,
		//组合结果
		ctx.BuildResp,
	}
	return ctx
}

func (g *GeneralLotteryContext) Run() (*GeneralLotteryResponse, error) {
	for _, f := range g.Funcs {
		f()
		if g.Err != nil {
			return nil, g.Err
		}
	}
	return g.Resp, nil
}

func (g *GeneralLotteryContext) GetPool() {
	pool := consts.GetGeneralPool(g.Req.GeneralLottery)
	if pool == nil {
		g.Err = errors.New("卡池不存在")
	}
	for generalId, rate := range pool {
		g.GeneralPool[generalId] = rate
	}
}

func (g *GeneralLotteryContext) ConfigPool() {
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
	hlog.CtxInfof(g.Req.Ctx, "[ConfigPool] lowLimit:%d , upperLimit:%d ,ChancePool size:%d", lowLimit, cursor, len(g.ChancePool))

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
	hlog.CtxInfof(g.Req.Ctx, "[Config5LevPool] lowLimit:%d , upperLimit5Lev:%d ,Chance5LevPool size:%d", lowLimit5Lev, cursor5Lev, len(g.Chance5LevPool))

}

func (g *GeneralLotteryContext) LotteryPool() {
	for i := int64(0); i < g.Req.RollTimes; i++ {
		//累计30抽不中走保底逻辑
		if g.NotHit5LevGeneralNum == 30 {
			//保底逻辑
			g.hitGeneralHandler(true)
			//累计清零
			g.NotHit5LevGeneralNum = 0
			//保底统计
			g.ProtectedMustHitNum++
		} else {
			g.hitGeneralHandler(false)
		}
	}
}

func (g *GeneralLotteryContext) hitGeneralHandler(isMustHit5Lev bool) {
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
			if _, okk := g.HitGeneralMap[generalId]; okk {
				g.HitGeneralMap[generalId]++
			} else {
				g.HitGeneralMap[generalId] = 1
			}
			//非五星武将计数
			if _, okkk := g.General5LevMap[generalId]; okkk {
				g.NotHit5LevGeneralNum = 0
			} else {
				g.NotHit5LevGeneralNum++
			}
			//hlog.CtxInfof(g.Req.Ctx, "[LotteryPool] random:%d , hit GeneralId:%d", random, int(generalId))
			break
		}
	}
}

func (g *GeneralLotteryContext) BuildResp() {
	//查询武将信息
	generalIds := make([]int64, 0)
	for generalId, _ := range g.HitGeneralMap {
		generalIds = append(generalIds, int64(generalId))
	}

	generals, err := mysql.NewGeneral().QueryGeneralList(g.Req.Ctx, &vo.QueryGeneralCondition{
		Ids:    generalIds,
		Offset: 0,
		Limit:  len(generalIds),
	})
	if err != nil {
		hlog.CtxErrorf(g.Req.Ctx, "QueryGeneralList err:%v", err)
		g.Err = err
		return
	}

	//整理resp
	generalLotteryList := make([]*GeneralLotteryInfo, 0)
	hit5LevGeneralNum := int64(0)
	for _, general := range generals {
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

		generalLotteryList = append(generalLotteryList, &GeneralLotteryInfo{
			GeneralInfo: general,
			HitNum:      hitNum,
			HitRate:     cast.ToFloat64(hitNum) / cast.ToFloat64(g.Req.RollTimes),
			LotteryRate: lotteryRate,
		})
	}
	g.Resp = &GeneralLotteryResponse{
		GeneralLotteryInfoList: generalLotteryList,
		ProtectedMustHitNum:    g.ProtectedMustHitNum,
		Hit5LevGeneralNum:      hit5LevGeneralNum,
	}
}
