package consts

//***抽卡卡池***

type GeneralLotteryPool int

const (
	Unknown GeneralLotteryPool = iota
	//S1千军竞发
	S1QianJunJingFa
	//S1兵连祸结
	S1BingLianHuoJie
	//S1群雄割据
	S1QunXiongGeJu
	//S1盛食厉兵
	S1ChengShiLiBing
	//S1英雄集结
	S1YingXiongJiJie
	//S1赛季名将
	S1SaiJiMingJiang
	//S1汉室衰微
	S1HanShiShuaiWei
	//S1典藏卡包
	S1DianCangKaBao
	//S1霸业卡包
	S1BaYeKaBao
	//S1赛季初始大卡池
	S1SaiJiChuShiDaKaChi
	//S1赛季最终大卡池
	S1SaiJiZuiZhongDaKaChi
	//S2兵连祸结
	S2BingLianHuoJie
	//S2群雄割据
	S2QunXiongGeJu
	//S2英雄集结
	S2YingXiongJiJie
	//S2盛食厉兵
	S2ChengShiLiBing
	//S2汉室衰微
	S2HanShiShuaiWei
	//S2赛季名将
	S2SaiJiMingJiang
	//S2典藏卡包
	S2DianCangKaBao
	//S2霸业卡包
	S2BaYeKaBao
	//S2赛季初始大卡池
	S2SaiJiChuShiDaKaChi
	//S2赛季最终大卡池
	S2SaiJiZuiZhongDaKaChi
	//S3纵横天下
	S3ZongHengTianXia
	//S3蜀魂汉将
	S3ShuHunHanJiang
	//S3群英荟萃
	S3QunYingHuiCui
	//S3吴越猛士
	S3WuYueMengShi
	//S3魏武雄兵
	S3WeiWuXiongBing
	//S3乱世纷争
	S3LuanShiFenZheng
	//S3赛季名将
	S3SaiJiMingJiang
	//S3典藏卡包
	S3DianCangKaBao
	//S3赛季初始大卡池
	S3SaiJiChuShiDaKaChi
	//S3赛季最终大卡池
	S3SaiJiZuiZhongDaKaChi
	//PK赛季名将1
	PKSaiJiMingCheng1
	//PK赛季名将2
	PKSaiJiMingJiang2
	//PK赛季名将3
	PKSaiJiMingCheng3
	//PK赛季初始大卡池
	PKSaiJiChuShiDaKaChi
	//PK赛季最终大卡池
	PKSaiJiZuiZhongDaKaChi
	//PK动如雷霆
	PKDongRuLeiTing
	//PK其疾如风
	PKQiJiRuFeng
	//PK其徐如林
	PKQiXuRuLin
	//PK难知如阴
	PKNanZhiRuYin
	//PK不动如山
	PKBuDongRuShan
	//PK侵略如火
	PKQinLveRuHuo
	//PK霸业奖励1
	PKBaYeJiangLi1
	//PK霸业奖励2
	PKBaYeJiangLi2
	//PK霸业奖励3
	PKBaYeJiangLi3
	//PK典藏卡包
	PKDianCangKaBao
)

func (g GeneralLotteryPool) String() string {
	switch g {
	//S1千军竞发
	case S1QianJunJingFa:
		return "S1千军竞发"
	//S1兵连祸结
	case S1BingLianHuoJie:
		return "S1兵连祸结"
	//S1群雄割据
	case S1QunXiongGeJu:
		return "S1群雄割据"
	//S1盛食厉兵
	case S1ChengShiLiBing:
		return "S1盛食厉兵"
	//S1英雄集结
	case S1YingXiongJiJie:
		return "S1英雄集结"
	//S1赛季名将
	case S1SaiJiMingJiang:
		return "S1赛季名将"
		//S1汉室衰微
	case S1HanShiShuaiWei:
		return "S1汉室衰微"
	//S1典藏卡包
	case S1DianCangKaBao:
		return "S1典藏卡包"
	//S1霸业卡包
	case S1BaYeKaBao:
		return "S1霸业卡包"
	//S1赛季初始大卡池
	case S1SaiJiChuShiDaKaChi:
		return "S1赛季初始大卡池"
	//S1赛季最终大卡池
	case S1SaiJiZuiZhongDaKaChi:
		return "S1赛季最终大卡池"
		//S2兵连祸结
	case S2BingLianHuoJie:
		return "S2兵连祸结"
		//S2群雄割据
	case S2QunXiongGeJu:
		return "S2群雄割据"
	//S2英雄集结
	case S2YingXiongJiJie:
		return "S2英雄集结"
	//S2盛食厉兵
	case S2ChengShiLiBing:
		return "S2盛食厉兵"
	//S2汉室衰微
	case S2HanShiShuaiWei:
		return "S2汉室衰微"
	//S2赛季名将
	case S2SaiJiMingJiang:
		return "S2赛季名将"
	//S2典藏卡包
	case S2DianCangKaBao:
		return "S2典藏卡包"
	//S2霸业卡包
	case S2BaYeKaBao:
		return "S2霸业卡包"
	//S2赛季初始大卡池
	case S2SaiJiChuShiDaKaChi:
		return "S2赛季初始大卡池"
	//S2赛季最终大卡池
	case S2SaiJiZuiZhongDaKaChi:
		return "S2赛季最终大卡池"
	//S3纵横天下
	case S3ZongHengTianXia:
		return "S3纵横天下"
	//S3蜀魂汉将
	case S3ShuHunHanJiang:
		return "S3蜀魂汉将"
	//S3群英荟萃
	case S3QunYingHuiCui:
		return "S3群英荟萃"
	//S3吴越猛士
	case S3WuYueMengShi:
		return "S3吴越猛士"
	//S3魏武雄兵
	case S3WeiWuXiongBing:
		return "S3魏武雄兵"
	//S3乱世纷争
	case S3LuanShiFenZheng:
		return "S3乱世纷争"
	//S3赛季名将
	case S3SaiJiMingJiang:
		return "S3赛季名将"
	//S3典藏卡包
	case S3DianCangKaBao:
		return "S3典藏卡包"
	//S3赛季初始大卡池
	case S3SaiJiChuShiDaKaChi:
		return "S3赛季初始大卡池"
	//S3赛季最终大卡池
	case S3SaiJiZuiZhongDaKaChi:
		return "S3赛季最终大卡池"
	//PK赛季名将1
	case PKSaiJiMingCheng1:
		return "PK赛季名将1"
	//PK赛季名将2
	case PKSaiJiMingJiang2:
		return "PK赛季名将2"
	//PK赛季名将3
	case PKSaiJiMingCheng3:
		return "PK赛季名将3"
	//PK赛季初始大卡池
	case PKSaiJiChuShiDaKaChi:
		return "PK赛季初始大卡池"
	//PK赛季最终大卡池
	case PKSaiJiZuiZhongDaKaChi:
		return "PK赛季最终大卡池"
	//PK动如雷霆
	case PKDongRuLeiTing:
		return "PK动如雷霆"
	//PK其疾如风
	case PKQiJiRuFeng:
		return "PK其疾如风"
	//PK其徐如林
	case PKQiXuRuLin:
		return "PK其徐如林"
	//PK难知如阴
	case PKNanZhiRuYin:
		return "PK难知如阴"
	//PK不动如山
	case PKBuDongRuShan:
		return "PK不动如山"
	//PK侵略如火
	case PKQinLveRuHuo:
		return "PK侵略如火"
	//PK霸业奖励1
	case PKBaYeJiangLi1:
		return "PK霸业奖励1"
	//PK霸业奖励2
	case PKBaYeJiangLi2:
		return "PK霸业奖励2"
	//PK霸业奖励3
	case PKBaYeJiangLi3:
		return "PK霸业奖励3"
	//PK典藏卡包
	case PKDianCangKaBao:
		return "PK典藏卡包"
	}
	return "未知卡包"
}

var GeneralLotteryPoolArr = []GeneralLotteryPool{
	//S1千军竞发
	S1QianJunJingFa,
	//S1兵连祸结
	S1BingLianHuoJie,
	//S1群雄割据
	S1QunXiongGeJu,
	//S1盛食厉兵
	S1ChengShiLiBing,
	//S1英雄集结
	S1YingXiongJiJie,
	//S1赛季名将
	S1SaiJiMingJiang,
	//S1汉室衰微
	S1HanShiShuaiWei,
	//S1典藏卡包
	S1DianCangKaBao,
	//S1霸业卡包
	S1BaYeKaBao,
	//S1赛季初始大卡池
	S1SaiJiChuShiDaKaChi,
	//S1赛季最终大卡池
	S1SaiJiZuiZhongDaKaChi,
	//S2兵连祸结
	S2BingLianHuoJie,
	//S2群雄割据
	S2QunXiongGeJu,
	//S2英雄集结
	S2YingXiongJiJie,
	//S2盛食厉兵
	S2ChengShiLiBing,
	//S2汉室衰微
	S2HanShiShuaiWei,
	//S2赛季名将
	S2SaiJiMingJiang,
	//S2典藏卡包
	S2DianCangKaBao,
	//S2霸业卡包
	S2BaYeKaBao,
	//S2赛季初始大卡池
	S2SaiJiChuShiDaKaChi,
	//S2赛季最终大卡池
	S2SaiJiZuiZhongDaKaChi,
	//S3纵横天下
	S3ZongHengTianXia,
	//S3蜀魂汉将
	S3ShuHunHanJiang,
	//S3群英荟萃
	S3QunYingHuiCui,
	//S3吴越猛士
	S3WuYueMengShi,
	//S3魏武雄兵
	S3WeiWuXiongBing,
	//S3乱世纷争
	S3LuanShiFenZheng,
	//S3赛季名将
	S3SaiJiMingJiang,
	//S3典藏卡包
	S3DianCangKaBao,
	//S3赛季初始大卡池
	S3SaiJiChuShiDaKaChi,
	//S3赛季最终大卡池
	S3SaiJiZuiZhongDaKaChi,
	//PK赛季名将1
	PKSaiJiMingCheng1,
	//PK赛季名将2
	PKSaiJiMingJiang2,
	//PK赛季名将3
	PKSaiJiMingCheng3,
	//PK赛季初始大卡池
	PKSaiJiChuShiDaKaChi,
	//PK赛季最终大卡池
	PKSaiJiZuiZhongDaKaChi,
	//PK动如雷霆
	PKDongRuLeiTing,
	//PK其疾如风
	PKQiJiRuFeng,
	//PK其徐如林
	PKQiXuRuLin,
	//PK难知如阴
	PKNanZhiRuYin,
	//PK不动如山
	PKBuDongRuShan,
	//PK侵略如火
	PKQinLveRuHuo,
	//PK霸业奖励1
	PKBaYeJiangLi1,
	//PK霸业奖励2
	PKBaYeJiangLi2,
	//PK霸业奖励3
	PKBaYeJiangLi3,
	//PK典藏卡包
	PKDianCangKaBao,
}

var GeneralLotteryPoolMap = map[GeneralLotteryPool]map[General_Id]float64{
	//S1千军竞发
	S1QianJunJingFa: S1QianJunJingFaPool,
	//S1兵连祸结
	S1BingLianHuoJie: S1BingLianHuoJiePool,
	//S1群雄割据
	S1QunXiongGeJu: S1QunXiongGeJuPool,
	//S1盛食厉兵
	S1ChengShiLiBing: S1ChengShiLiBingPool,
	//S1英雄集结
	S1YingXiongJiJie: S1YingXiongJiJiePool,
	//S1赛季名将
	S1SaiJiMingJiang: S1SaiJiMingJiangPool,
	//S1汉室衰微
	S1HanShiShuaiWei: S1HanShiShuaiWeiPool,
	//S1典藏卡包
	S1DianCangKaBao: S1DianCangKaBaoPool,
	//S1霸业卡包
	S1BaYeKaBao: S1BaYeKaBaoPool,
	//S1赛季初始大卡池
	S1SaiJiChuShiDaKaChi: S1SaiJiChuShiDaKaChiPool,
	//S1赛季最终大卡池
	S1SaiJiZuiZhongDaKaChi: S1SaiJiZuiZhongDaKaChiPool,
	//S2兵连祸结
	S2BingLianHuoJie: S2BingLianHuoJiePool,
	//S2群雄割据
	S2QunXiongGeJu: S2QunXiongGeJuPool,
	//S2英雄集结
	S2YingXiongJiJie: S2YingXiongJiJiePool,
	//S2盛食厉兵
	S2ChengShiLiBing: S2ChengShiLiBingPool,
	//S2汉室衰微
	S2HanShiShuaiWei: S2HanShiShuaiWeiPool,
	//S2赛季名将
	S2SaiJiMingJiang: S2SaiJiMingJiangPool,
	//S2典藏卡包
	S2DianCangKaBao: S2DianCangKaBaoPool,
	//S2霸业卡包
	S2BaYeKaBao: S2BaYeKaBaoPool,
	//S2赛季初始大卡池
	S2SaiJiChuShiDaKaChi: S2SaiJiChuShiDaKaChiPool,
	//S2赛季最终大卡池
	S2SaiJiZuiZhongDaKaChi: S2SaiJiZuiZhongDaKaChiPool,
	//S3纵横天下
	S3ZongHengTianXia: S3ZongHengTianXiaPool,
	//S3蜀魂汉将
	S3ShuHunHanJiang: S3ShuHunHanJiangPool,
	//S3群英荟萃
	S3QunYingHuiCui: S3QunYingHuiCuiPool,
	//S3吴越猛士
	S3WuYueMengShi: S3WuYueMengShiPool,
	//S3魏武雄兵
	S3WeiWuXiongBing: S3WeiWuXiongBingPool,
	//S3乱世纷争
	S3LuanShiFenZheng: S3LuanShiFenZhengPool,
	//S3赛季名将
	S3SaiJiMingJiang: S3SaiJiMingJiangPool,
	//S3典藏卡包
	S3DianCangKaBao: S3DianCangKaBaoPool,
	//S3赛季初始大卡池
	S3SaiJiChuShiDaKaChi: S3SaiJiChuShiDaKaChiPool,
	//S3赛季最终大卡池
	S3SaiJiZuiZhongDaKaChi: S3SaiJiZuiZhongDaKaChiPool,
	//PK赛季名将1
	PKSaiJiMingCheng1: PKSaiJiMingCheng1Pool,
	//PK赛季名将2
	PKSaiJiMingJiang2: PKSaiJiMingJiang2Pool,
	//PK赛季名将3
	PKSaiJiMingCheng3: PKSaiJiMingCheng3Pool,
	//PK赛季初始大卡池
	PKSaiJiChuShiDaKaChi: PKSaiJiChuShiDaKaChiPool,
	//PK赛季最终大卡池
	PKSaiJiZuiZhongDaKaChi: PKSaiJiZuiZhongDaKaChiPool,
	//PK动如雷霆
	PKDongRuLeiTing: PKDongRuLeiTingPool,
	//PK其疾如风
	PKQiJiRuFeng: PKQiJiRuFengPool,
	//PK其徐如林
	PKQiXuRuLin: PKQiXuRuLinPool,
	//PK难知如阴
	PKNanZhiRuYin: PKNanZhiRuYinPool,
	//PK不动如山
	PKBuDongRuShan: PKBuDongRuShanPool,
	//PK侵略如火
	PKQinLveRuHuo: PKQinLveRuHuoPool,
	//PK霸业奖励1
	PKBaYeJiangLi1: PKBaYeJiangLi1Pool,
	//PK霸业奖励2
	PKBaYeJiangLi2: PKBaYeJiangLi2Pool,
	//PK霸业奖励3
	PKBaYeJiangLi3: PKBaYeJiangLi3Pool,
	//PK典藏卡包
	PKDianCangKaBao: PKDianCangKaBaoPool,
}

//根据【卡池枚举】获得【卡池信息】
func GetGeneralPool(generalLotteryPool GeneralLotteryPool) map[General_Id]float64 {
	if pool, ok := GeneralLotteryPoolMap[generalLotteryPool]; ok {
		return pool
	}
	return nil
}

//***卡池***

//S1千军竞发
var S1QianJunJingFaPool = map[General_Id]float64{}

//S1兵连祸结
var S1BingLianHuoJiePool = map[General_Id]float64{}

//S1群雄割据
var S1QunXiongGeJuPool = map[General_Id]float64{}

//S1盛食厉兵
var S1ChengShiLiBingPool = map[General_Id]float64{}

//S1英雄集结
var S1YingXiongJiJiePool = map[General_Id]float64{}

//S1赛季名将
var S1SaiJiMingJiangPool = map[General_Id]float64{}

//S1汉室衰微
var S1HanShiShuaiWeiPool = map[General_Id]float64{}

//S1典藏卡包
var S1DianCangKaBaoPool = map[General_Id]float64{}

//S1霸业卡包
var S1BaYeKaBaoPool = map[General_Id]float64{}

//S1赛季初始大卡池
var S1SaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//S1赛季最终大卡池
var S1SaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

//S2兵连祸结
var S2BingLianHuoJiePool = map[General_Id]float64{}

//S2群雄割据
var S2QunXiongGeJuPool = map[General_Id]float64{}

//S2英雄集结
var S2YingXiongJiJiePool = map[General_Id]float64{}

//S2盛食厉兵
var S2ChengShiLiBingPool = map[General_Id]float64{}

//S2汉室衰微
var S2HanShiShuaiWeiPool = map[General_Id]float64{}

//S2赛季名将
var S2SaiJiMingJiangPool = map[General_Id]float64{}

//S2典藏卡包
var S2DianCangKaBaoPool = map[General_Id]float64{}

//S2霸业卡包
var S2BaYeKaBaoPool = map[General_Id]float64{
	LuSu:       0.02,
	HaoZhao:    0.02,
	LingTong:   0.02,
	JiaXu:      0.02,
	JiangWan:   0.02,
	MuLuDaWang: 0.02,
	YangXiu:    0.088,
	DongYun:    0.088,
	ChenWu:     0.088,
	ZhuHuan:    0.088,
	ZhangXiu:   0.088,
	ChenLin:    0.088,
	LiuFeng:    0.088,
	LiaoHua:    0.088,
	BuLianShi:  0.088,
	JiangQin:   0.088,
}

//S2赛季初始大卡池
var S2SaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//S2赛季最终大卡池
var S2SaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

//S3纵横天下
var S3ZongHengTianXiaPool = map[General_Id]float64{
	ZhangFei:      0.0033,
	MengHuo:       0.0033,
	ZhuRongFuRen:  0.0033,
	WuTuGu:        0.0066,
	GanNing:       0.0066,
	PangDe:        0.0066,
	DengAi:        0.0066,
	ZhongHui:      0.0066,
	LingTong:      0.0066,
	MuLuDaWang:    0.0066,
	GuanHai:       0.04,
	HuangFuSong:   0.04,
	ZhangManCheng: 0.04,
	ZhangBao:      0.04,
	ZhangLiang:    0.04,
	JiLing:        0.04,
	ZhangYan:      0.04,
	ZhangRen:      0.04,
	MaTie:         0.04,
	CheZhou:       0.1168,
	BianXi:        0.1168,
	PanFeng:       0.1168,
	SongXian:      0.1168,
	ZhuGeZhan:     0.1168,
}

//S3蜀魂汉将
var S3ShuHunHanJiangPool = map[General_Id]float64{}

//S3群英荟萃
var S3QunYingHuiCuiPool = map[General_Id]float64{}

//S3吴越猛士
var S3WuYueMengShiPool = map[General_Id]float64{}

//S3魏武雄兵
var S3WeiWuXiongBingPool = map[General_Id]float64{}

//S3乱世纷争
var S3LuanShiFenZhengPool = map[General_Id]float64{}

//S3赛季名将
var S3SaiJiMingJiangPool = map[General_Id]float64{}

//S3典藏卡包
var S3DianCangKaBaoPool = map[General_Id]float64{
	CaoCao:        0.0233,
	LvBu:          0.0233,
	ZhangLiang:    0.0233,
	SiMaYi:        0.0233,
	SunShangXiang: 0.0233,
	GuanYu:        0.0465,
	LvMeng:        0.0465,
	MaChao:        0.0465,
	HuangZhong:    0.0465,
	ZhangFei:      0.0465,
	XiaHouDun:     0.0465,
	YanLiang:      0.0465,
	WenChou:       0.0465,
	HuangGai:      0.0465,
	TaiShiCi:      0.0465,
	PangDe:        0.0465,
	MengHuo:       0.0465,
	GuoJia:        0.0465,
	SunJian:       0.0465,
	LingTong:      0.0465,
	DengAi:        0.0465,
	DiaoChan:      0.0465,
	HuangYueYing:  0.0465,
	LvLingQi:      0.0465,
}

//S3赛季初始大卡池
var S3SaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//S3赛季最终大卡池
var S3SaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

//PK赛季名将1
var PKSaiJiMingCheng1Pool = map[General_Id]float64{
	JiangWei:     0.0029,
	YuanShu:      0.0029,
	ZhouTai:      0.0029,
	GuanYinPing:  0.0029,
	SPYuanShao:   0.0029,
	CaiYong:      0.0059,
	ZhangChunHua: 0.0059,
	LuKang:       0.0059,
	YanYan:       0.0059,
	XunYou:       0.0059,
	ZouShi:       0.0059,
	DongBai:      0.0059,
	DongYun:      0.036,
	GuYong:       0.036,
	LiJue:        0.036,
	GuoSi:        0.036,
	XuSheng:      0.036,
	ZhuGeJin:     0.036,
	CaoHong:      0.036,
	CaoZhen:      0.036,
	YangXiu:      0.036,
	ZhangXiu:     0.036,
	FuShiRen:     0.0449,
	KanZe:        0.0449,
	ZhuGeZhan:    0.0449,
	ZhuRan:       0.0449,
	QuanCong:     0.0449,
	XiangChong:   0.0449,
	FeiShi:       0.0449,
	WuYi:         0.0449,
	CaoXiu:       0.0449,
	PanFeng:      0.0449,
	BianXi:       0.0449,
	HuaXin:       0.0449,
	SunQian:      0.0449,
}

//PK赛季名将2
var PKSaiJiMingJiang2Pool = map[General_Id]float64{
	SPZhuGeLiang: 0.0029,
	SPZhouYu:     0.0029,
	ManChong:     0.0029,
	SPZhuJun:     0.0029,
	WeiYan:       0.0029,
	JuShou:       0.0059,
	CaiYong:      0.0059,
	LuKang:       0.0059,
	GaoLan:       0.0059,
	WangShuang:   0.0059,
	ZouShi:       0.0059,
	ZhangRang:    0.0059,
	DongYun:      0.036,
	GuYong:       0.036,
	LiJue:        0.036,
	GuoSi:        0.036,
	XuSheng:      0.036,
	ZhuGeJin:     0.036,
	CaoHong:      0.036,
	CaoZhen:      0.036,
	YangXiu:      0.036,
	ZhangXiu:     0.036,
	FuShiRen:     0.049,
	KanZe:        0.049,
	ZhuGeZhan:    0.049,
	ZhuRan:       0.049,
	QuanCong:     0.049,
	XiangChong:   0.049,
	FeiShi:       0.049,
	WuYi:         0.049,
	CaoXiu:       0.049,
	PanFeng:      0.049,
	BianXi:       0.049,
	HuaXin:       0.049,
	SunQian:      0.049,
}

//PK赛季名将3
var PKSaiJiMingCheng3Pool = map[General_Id]float64{
	SPLvMeng:       0.0029,
	SPGuoJia:       0.0029,
	WangYuanJi:     0.0029,
	ShuGuoZhangBao: 0.0029,
	SPXunYu:        0.0029,
	ZhangChunHua:   0.0059,
	GaoLan:         0.0059,
	WangShuang:     0.0059,
	YanYan:         0.0059,
	DongBai:        0.0059,
	ZhangRang:      0.0059,
	GuanXing:       0.0059,
	DongYun:        0.036,
	GuYong:         0.036,
	LiJue:          0.036,
	GuoSi:          0.036,
	XuSheng:        0.036,
	ZhuGeJin:       0.036,
	CaoHong:        0.036,
	CaoZhen:        0.036,
	YangXiu:        0.036,
	ZhangXiu:       0.036,
	FuShiRen:       0.0449,
	KanZe:          0.0449,
	ZhuGeZhan:      0.0449,
	ZhuRan:         0.0449,
	QuanCong:       0.0449,
	XiangChong:     0.0449,
	FeiShi:         0.0449,
	WuYi:           0.0449,
	CaoXiu:         0.0449,
	PanFeng:        0.0449,
	BianXi:         0.0449,
	HuaXin:         0.0449,
	SunQian:        0.0449,
}

//PK赛季初始大卡池
var PKSaiJiChuShiDaKaChiPool = map[General_Id]float64{
	SunQuan:       0.0004,
	CaoCao:        0.0006,
	ZhuGeLiang:    0.0006,
	LiuBei:        0.0006,
	ZhouYu:        0.0006,
	ZhangFei:      0.0006,
	LvBu:          0.0006,
	MaChao:        0.0006,
	GuanYu:        0.0006,
	ZhangJiao:     0.0006,
	XuChu:         0.0006,
	ZhaoYun:       0.0006,
	DianWei:       0.0006,
	XunYu:         0.0012,
	ZhangZhao:     0.0012,
	CaoZhi:        0.0012,
	CaoPi:         0.0012,
	CaiWenJi:      0.0012,
	YuanShao:      0.0012,
	DengAi:        0.0012,
	ZhuRongFuRen:  0.0012,
	DongZhuo:      0.0012,
	XiaHouDun:     0.0012,
	HuaTuo:        0.0012,
	DaQiao:        0.0012,
	XiaoQiao:      0.0012,
	LiRu:          0.0012,
	SunCe:         0.0012,
	LvLingQi:      0.0012,
	MaTeng:        0.0012,
	XuShu:         0.0012,
	YuJin:         0.0012,
	LvMeng:        0.0012,
	WangPing:      0.0012,
	FaZheng:       0.0012,
	HuaXiong:      0.0012,
	HuangGai:      0.0012,
	GaoShun:       0.0012,
	ChengPu:       0.0012,
	CaoRen:        0.0012,
	ChenDao:       0.0012,
	TianFeng:      0.0012,
	ZhongHui:      0.0012,
	ChengYu:       0.0012,
	ZuoCi:         0.0012,
	HuangZhong:    0.0012,
	SunJian:       0.0012,
	GuoJia:        0.0012,
	XiaHouYuan:    0.0012,
	YuJi:          0.0012,
	XuHuang:       0.0012,
	DiaoChan:      0.0012,
	XuYou:         0.0012,
	XuSheng:       0.0061,
	CaoHong:       0.0061,
	LuZhi:         0.0061,
	GuanHai:       0.0061,
	HuangFuSong:   0.0061,
	ZhangManCheng: 0.0061,
	HanDang:       0.0061,
	DingFeng:      0.0061,
	ChenLin:       0.0061,
	ZhangBao:      0.0061,
	ZhangLiang:    0.0061,
	JiangQin:      0.0061,
	LiJue:         0.0061,
	LiDian:        0.0061,
	LvFan:         0.0061,
	MiZhu:         0.0061,
	LiuYe:         0.0061,
	ZhuJun:        0.0061,
	ZhangXiu:      0.0061,
	JiLing:        0.0061,
	ZhangYan:      0.0061,
	GuanPing:      0.0061,
	PanZhang:      0.0061,
	ZangBa:        0.0061,
	JianYong:      0.0061,
	LiYan:         0.0061,
	ZhangRen:      0.0061,
	WenPin:        0.0061,
	CaoZhen:       0.0061,
	GuYong:        0.0061,
	LiuFeng:       0.0061,
	ShaMoKe:       0.0061,
	MaLiang:       0.0061,
	CaoZhang:      0.0061,
	GuoHuai:       0.0061,
	LiaoHua:       0.0061,
	ZhongYao:      0.0061,
	MaSu:          0.0061,
	DongYun:       0.0061,
	ChenWu:        0.0061,
	ZhangLu:       0.0061,
	FengJi:        0.0061,
	KongRong:      0.0061,
	ShenPei:       0.0061,
	GuoTu:         0.0061,
	LiuBiao:       0.0061,
	GuoSi:         0.0061,
	HuCheEr:       0.0061,
	BuLianShi:     0.0061,
	HanSui:        0.0061,
	YangXiu:       0.0061,
	MaTie:         0.0061,
	FeiYi:         0.0061,
	DongXi:        0.0061,
	HuangQuan:     0.0061,
	ZhouCang:      0.0061,
	ZhuGeJin:      0.0061,
	ZhuHuan:       0.0061,
	WangLang:      0.0061,
	LiuYao:        0.0183,
	DingYuan:      0.0183,
	WangYun:       0.0183,
	TaoQian:       0.0183,
	MiFang:        0.0183,
	HeJin:         0.0183,
	CheZhou:       0.0183,
	PanFeng:       0.0183,
	BianXi:        0.0183,
	SunQian:       0.0183,
	SongXian:      0.0183,
	DongZhao:      0.0183,
	YuFan:         0.0183,
	LvQian:        0.0183,
	HuaXin:        0.0183,
	BuZhi:         0.0183,
	XiaHouEn:      0.0183,
	MaoJie:        0.0183,
	DengZhi:       0.0183,
	LiuBa:         0.0183,
	FuShiRen:      0.0183,
	KanZe:         0.0183,
	ZhuGeZhan:     0.0183,
	ZhuRan:        0.0183,
	QuanCong:      0.0183,
	XiangChong:    0.0183,
	FeiShi:        0.0183,
	WuYi:          0.0183,
	CaoXiu:        0.0183,
	SunHao:        0.0183,
	PanJun:        0.0183,
	LiuChan:       0.0183,
}

//PK赛季最终大卡池
var PKSaiJiZuiZhongDaKaChiPool = map[General_Id]float64{
	SunQuan:        0.0002,
	CaoCao:         0.0003,
	ZhuGeLiang:     0.0003,
	LiuBei:         0.0003,
	ZhouYu:         0.0003,
	ZhangFei:       0.0003,
	LvBu:           0.0003,
	MaChao:         0.0003,
	GuanYu:         0.0003,
	ZhangJiao:      0.0003,
	XuChu:          0.0003,
	TaiShiCi:       0.0003,
	SPXunYu:        0.0003,
	SPZhuGeLiang:   0.0003,
	SPZhuJun:       0.0003,
	GuanYinPing:    0.0003,
	SPZhouYu:       0.0003,
	LuXun:          0.0003,
	SPYuanShao:     0.0003,
	ZhouTai:        0.0003,
	ZhaoYun:        0.0003,
	ShuGuoZhangBao: 0.0003,
	DianWei:        0.0003,
	ManChong:       0.0003,
	WeiYan:         0.0003,
	SPGuoJia:       0.0003,
	WangYuanJi:     0.0003,
	JiangWei:       0.0003,
	SPLvMeng:       0.0003,
	XunYu:          0.0007,
	ZhangZhao:      0.0007,
	CaoZhi:         0.0007,
	CaoPi:          0.0007,
	CaiWenJi:       0.0007,
	YuanShao:       0.0007,
	DengAi:         0.0007,
	ZhuRongFuRen:   0.0007,
	DongZhuo:       0.0007,
	XiaHouDun:      0.0007,
	HuaTuo:         0.0007,
	DaQiao:         0.0007,
	XiaoQiao:       0.0007,
	LiRu:           0.0007,
	SunCe:          0.0007,
	LvLingQi:       0.0007,
	MaTeng:         0.0007,
	XuShu:          0.0007,
	YuJin:          0.0007,
	LvMeng:         0.0007,
	WangPing:       0.0007,
	FaZheng:        0.0007,
	HuaXiong:       0.0007,
	HuangGai:       0.0007,
	GaoShun:        0.0007,
	ChengPu:        0.0007,
	CaoRen:         0.0007,
	ChenDao:        0.0007,
	TianFeng:       0.0007,
	ZhongHui:       0.0007,
	ChengYu:        0.0007,
	ZuoCi:          0.0007,
	HuangZhong:     0.0007,
	SunJian:        0.0007,
	GuoJia:         0.0007,
	LeJin:          0.0007,
	PangDe:         0.0007,
	CaoChun:        0.0007,
	ZhangRang:      0.0007,
	XiaHouYuan:     0.0007,
	WangShuang:     0.0007,
	YanLiang:       0.0007,
	YuanShu:        0.0007,
	YuJi:           0.0007,
	WuTuGu:         0.0007,
	CaiYong:        0.0007,
	ZouShi:         0.0007,
	JuShou:         0.0007,
	GanNing:        0.0007,
	XuHuang:        0.0007,
	GuanXing:       0.0007,
	LuKang:         0.0007,
	GongSunZan:     0.0007,
	DongBai:        0.0007,
	ChenQun:        0.0007,
	YanYan:         0.0007,
	DiaoChan:       0.0007,
	XunYou:         0.0007,
	ZhangChunHua:   0.0007,
	SiMaHui:        0.0007,
	ZhenJi:         0.0007,
	MengHuo:        0.0007,
	HuangYueYing:   0.0007,
	GaoLan:         0.0007,
	ZhangHe:        0.0007,
	WenChou:        0.0007,
	XuYou:          0.0007,
	XuSheng:        0.0061,
	CaoHong:        0.0061,
	LuZhi:          0.0061,
	GuanHai:        0.0061,
	HuangFuSong:    0.0061,
	ZhangManCheng:  0.0061,
	HanDang:        0.0061,
	DingFeng:       0.0061,
	ChenLin:        0.0061,
	ZhangBao:       0.0061,
	ZhangLiang:     0.0061,
	JiangQin:       0.0061,
	LiJue:          0.0061,
	LiDian:         0.0061,
	LvFan:          0.0061,
	MiZhu:          0.0061,
	LiuYe:          0.0061,
	ZhuJun:         0.0061,
	ZhangXiu:       0.0061,
	JiLing:         0.0061,
	ZhangYan:       0.0061,
	GuanPing:       0.0061,
	JianYong:       0.0061,
	LiYan:          0.0061,
	ZhangRen:       0.0061,
	WenPin:         0.0061,
	CaoZhen:        0.0061,
	GuYong:         0.0061,
	LiuFeng:        0.0061,
	ShaMoKe:        0.0061,
	MaLiang:        0.0061,
	CaoZhang:       0.0061,
	GuoHuai:        0.0061,
	LiaoHua:        0.0061,
	ZhongYao:       0.0061,
	MaSu:           0.0061,
	DongYun:        0.0061,
	ChenWu:         0.0061,
	ZhangLu:        0.0061,
	LiuBiao:        0.0061,
	GuoSi:          0.0061,
	HuCheEr:        0.0061,
	BuLianShi:      0.0061,
	HanSui:         0.0061,
	YangXiu:        0.0061,
	MaTie:          0.0061,
	FeiYi:          0.0061,
	DongXi:         0.0061,
	HuangQuan:      0.0061,
	ZhouCang:       0.0061,
	ZhuGeJin:       0.0061,
	ZhuHuan:        0.0061,
	WangLang:       0.0061,
	LiuYao:         0.0183,
	DingYuan:       0.0183,
	WangYun:        0.0183,
	TaoQian:        0.0183,
	MiFang:         0.0183,
	HeJin:          0.0183,
	CheZhou:        0.0183,
	PanFeng:        0.0183,
	BianXi:         0.0183,
	SunQian:        0.0183,
	SongXian:       0.0183,
	DongZhao:       0.0183,
	YuFan:          0.0183,
	LvQian:         0.0183,
	HuaXin:         0.0183,
	BuZhi:          0.0183,
	XiaHouEn:       0.0183,
	MaoJie:         0.0183,
	DengZhi:        0.0183,
	LiuBa:          0.0183,
	FuShiRen:       0.0183,
	KanZe:          0.0183,
	ZhuGeZhan:      0.0183,
	ZhuRan:         0.0183,
	QuanCong:       0.0183,
	XiangChong:     0.0183,
	FeiShi:         0.0183,
	WuYi:           0.0183,
	CaoXiu:         0.0183,
	SunHao:         0.0183,
	PanJun:         0.0183,
	LiuChan:        0.0183,
}

//PK动如雷霆
var PKDongRuLeiTingPool = map[General_Id]float64{
	//五星 5.6%
	//孟获 0.35%
	MengHuo: 0.0035,
	//SP吕蒙 0.35%
	SPLvMeng: 0.0035,
	//左慈 0.35%
	ZuoCi: 0.0035,
	//姜维 0.35%
	JiangWei: 0.0035,
	//黄月英 0.7%
	HuangYueYing: 0.007,
	//高览 0.7%
	GaoLan: 0.007,
	//文丑 0.7%
	WenChou: 0.007,
	//张合 0.7%
	ZhangHe: 0.007,
	//黄忠 0.7%
	HuangZhong: 0.007,
	//孙坚 0.7%
	SunJian: 0.007,

	//四星 36%
	// 纪灵 4.0%
	JiLing: 0.04,
	// 曹真 4.0%
	CaoZhen: 0.04,
	// 张梁 4.0%
	ZhangLiang: 0.04,
	// 马谡 4.0%
	MaSu: 0.04,
	// 文聘 4.0%
	WenPin: 0.04,
	// 杨修 4.0%
	YangXiu: 0.04,
	// 李典 4.0%
	LiDian: 0.04,
	// 蒋钦 4.0%
	JiangQin: 0.04,
	// 韩遂 4.0%
	HanSui: 0.04,

	//三星 58.4%
	//邓芝 9.73%
	DengZhi: 0.0973,
	//华歆 9.73%
	HuaXin: 0.0973,
	//虞翻 9.73%
	YuFan: 0.0973,
	//朱然 9.73%
	ZhuRan: 0.0973,
	//吕虔 9.73%
	LvQian: 0.0973,
	//刘繇 9.73%
	LiuYao: 0.0973,
}

//PK其疾如风
var PKQiJiRuFengPool = map[General_Id]float64{
	LvBu:       0.0035,
	MaChao:     0.0035,
	TaiShiCi:   0.0035,
	LeJin:      0.0035,
	PangDe:     0.007,
	CaoChun:    0.007,
	MaTeng:     0.007,
	WangShuang: 0.007,
	ZhangRang:  0.007,
	YanLiang:   0.007,
	GuanPing:   0.04,
	ChenLin:    0.04,
	ZhangXiu:   0.04,
	MaTie:      0.04,
	HuCheEr:    0.04,
	ChenWu:     0.04,
	HanDang:    0.04,
	CaoHong:    0.04,
	DongYun:    0.04,
	WangYun:    0.1168,
	TaoQian:    0.1168,
	DingYuan:   0.1168,
	HeJin:      0.1168,
	PanFeng:    0.1168,
}

//PK其徐如林
var PKQiXuRuLinPool = map[General_Id]float64{
	SPZhuJun:     0.0035,
	SPXunYu:      0.0035,
	GuanYinPing:  0.0035,
	SPZhuGeLiang: 0.0035,
	YuanShu:      0.007,
	WuTuGu:       0.007,
	ZouShi:       0.007,
	LvMeng:       0.007,
	CaiYong:      0.007,
	WangPing:     0.007,
	ZangBa:       0.04,
	ZhangRen:     0.04,
	LiuFeng:      0.04,
	ZhuGeJin:     0.04,
	LiYan:        0.04,
	ZhongYao:     0.04,
	DongXi:       0.04,
	ZhuJun:       0.04,
	MiZhu:        0.04,
	SongXian:     0.1168,
	MaoJie:       0.1168,
	DongZhao:     0.1168,
	CheZhou:      0.1168,
	BianXi:       0.1168,
}

//PK难知如阴
var PKNanZhiRuYinPool = map[General_Id]float64{
	ZhangJiao:    0.0035,
	WeiYan:       0.0035,
	SPGuoJia:     0.0035,
	WangYuanJi:   0.0035,
	TianFeng:     0.007,
	ZhongHui:     0.007,
	ZhangChunHua: 0.007,
	XunYou:       0.007,
	SiMaHui:      0.007,
	ZhenJi:       0.007,
	GuoTu:        0.036,
	GuYong:       0.036,
	WangLang:     0.036,
	HuangQuan:    0.036,
	MaLiang:      0.036,
	JianYong:     0.036,
	LiuYe:        0.036,
	BuLianShi:    0.036,
	ZhangBao:     0.036,
	ShaMoKe:      0.036,
	QuanCong:     0.0976,
	SunHao:       0.0976,
	PanJun:       0.0976,
	ZhuGeZhan:    0.0976,
	LiuBa:        0.0976,
	FeiShi:       0.0976,
}

//PK不动如山
var PKBuDongRuShanPool = map[General_Id]float64{
	SPYuanShao:     0.0035,
	ZhouTai:        0.0035,
	ShuGuoZhangBao: 0.0035,
	ManChong:       0.0035,
	DongBai:        0.007,
	GaoShun:        0.007,
	ChenQun:        0.007,
	ChengPu:        0.007,
	ChenDao:        0.007,
	YanYan:         0.007,
	GuoHuai:        0.036,
	HuangFuSong:    0.036,
	ZhouCang:       0.036,
	ZhuHuan:        0.036,
	PanZhang:       0.036,
	FeiYi:          0.036,
	LiuBiao:        0.036,
	ZhangLu:        0.036,
	ShenPei:        0.036,
	HuangQuan:      0.036,
	MiFang:         0.1168,
	CaoXiu:         0.1168,
	WuYi:           0.1168,
	LiuChan:        0.1168,
	XiangChong:     0.1168,
}

//PK侵略如火
var PKQinLveRuHuoPool = map[General_Id]float64{
	SPZhouYu:      0.0035,
	GuanYu:        0.0035,
	LuXun:         0.0035,
	FaZheng:       0.0035,
	JuShou:        0.007,
	GanNing:       0.007,
	HuangGai:      0.007,
	GuanXing:      0.007,
	LuKang:        0.007,
	GongSunZan:    0.007,
	ZhangManCheng: 0.036,
	GuanHai:       0.036,
	ShaMoKe:       0.036,
	CaoZhang:      0.036,
	DingFeng:      0.036,
	LvFan:         0.036,
	LiaoHua:       0.036,
	LiJue:         0.036,
	XuSheng:       0.036,
	HanDang:       0.036,
	SunQian:       0.1168,
	FuShiRen:      0.1168,
	BuZhi:         0.1168,
	KanZe:         0.1168,
	XiaHouEn:      0.1168,
}

//PK霸业奖励1
var PKBaYeJiangLi1Pool = map[General_Id]float64{
	YuanShu:       0.01,
	ZhangChunHua:  0.01,
	JiangWei:      0.01,
	YanYan:        0.01,
	LuKang:        0.01,
	XunYou:        0.01,
	ZhouTai:       0.01,
	GuanYinPing:   0.01,
	SPYuanShao:    0.01,
	ZouShi:        0.01,
	DongBai:       0.01,
	CaiYong:       0.01,
	ZhangManCheng: 0.088,
	JianYong:      0.088,
	WenPin:        0.088,
	CaoHong:       0.088,
	ZhuJun:        0.088,
	ZhongYao:      0.088,
	DingFeng:      0.088,
	GuanPing:      0.088,
	ZhuGeJin:      0.088,
	JiangQin:      0.088,
}

//PK霸业奖励2
var PKBaYeJiangLi2Pool = map[General_Id]float64{
	SPZhuGeLiang:  0.01,
	CaiYong:       0.01,
	ManChong:      0.01,
	GaoLan:        0.01,
	SPZhouYu:      0.01,
	LuKang:        0.01,
	SPZhuJun:      0.01,
	JuShou:        0.01,
	WangShuang:    0.01,
	ZouShi:        0.01,
	WeiYan:        0.01,
	ZhangRang:     0.01,
	ZhangManCheng: 0.088,
	JianYong:      0.088,
	WenPin:        0.088,
	CaoHong:       0.088,
	ZhuJun:        0.088,
	ZhongYao:      0.088,
	DingFeng:      0.088,
	GuanPing:      0.088,
	ZhuGeJin:      0.088,
	JiangQin:      0.088,
}

//PK霸业奖励3
var PKBaYeJiangLi3Pool = map[General_Id]float64{
	SPXunYu:        0.01,
	YanYan:         0.01,
	SPLvMeng:       0.01,
	WangYuanJi:     0.01,
	GaoLan:         0.01,
	SPGuoJia:       0.01,
	GuanXing:       0.01,
	ShuGuoZhangBao: 0.01,
	ZhangChunHua:   0.01,
	WangShuang:     0.01,
	DongBai:        0.01,
	ZhangRang:      0.01,
	ZhangManCheng:  0.088,
	JianYong:       0.088,
	WenPin:         0.088,
	CaoHong:        0.088,
	ZhuJun:         0.088,
	ZhongYao:       0.088,
	DingFeng:       0.088,
	GuanPing:       0.088,
	ZhuGeJin:       0.088,
	JiangQin:       0.088,
}

//PK典藏卡包
var PKDianCangKaBaoPool = map[General_Id]float64{
	CaoCao:         0.013,
	ZhangLiao:      0.013,
	SiMaYi:         0.013,
	JiangWei:       0.013,
	SunShangXiang:  0.013,
	SPYuanShao:     0.013,
	WeiYan:         0.013,
	SPZhouYu:       0.013,
	SPZhuGeLiang:   0.013,
	SPZhuJun:       0.013,
	SPGuoJia:       0.013,
	SPXunYu:        0.013,
	GuanYu:         0.026,
	LvMeng:         0.026,
	DongZhuo:       0.026,
	XiaHouYuan:     0.026,
	LvBu:           0.026,
	MaChao:         0.026,
	HuangZhong:     0.026,
	ZhangFei:       0.026,
	XiaHouDun:      0.026,
	YanLiang:       0.026,
	WenChou:        0.026,
	HuangGai:       0.026,
	TaiShiCi:       0.026,
	PangDe:         0.026,
	MengHuo:        0.026,
	GuoJia:         0.026,
	SunJian:        0.026,
	LingTong:       0.026,
	DengAi:         0.026,
	DiaoChan:       0.026,
	HuangYueYing:   0.026,
	LvLingQi:       0.026,
	ChenDao:        0.026,
	MuLuDaWang:     0.026,
	WangPing:       0.026,
	ManChong:       0.026,
	GanNing:        0.026,
	GuanXing:       0.026,
	ShuGuoZhangBao: 0.026,
	TianFeng:       0.026,
	WangShuang:     0.026,
	FaZheng:        0.026,
}
