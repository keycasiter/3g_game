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
var S2BaYeKaBaoPool = map[General_Id]float64{}

//S2赛季初始大卡池
var S2SaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//S2赛季最终大卡池
var S2SaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

//S3纵横天下
var S3ZongHengTianXiaPool = map[General_Id]float64{}

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
var S3DianCangKaBaoPool = map[General_Id]float64{}

//S3赛季初始大卡池
var S3SaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//S3赛季最终大卡池
var S3SaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

//PK赛季名将1
var PKSaiJiMingCheng1Pool = map[General_Id]float64{}

//PK赛季名将2
var PKSaiJiMingJiang2Pool = map[General_Id]float64{}

//PK赛季名将3
var PKSaiJiMingCheng3Pool = map[General_Id]float64{}

//PK赛季初始大卡池
var PKSaiJiChuShiDaKaChiPool = map[General_Id]float64{}

//PK赛季最终大卡池
var PKSaiJiZuiZhongDaKaChiPool = map[General_Id]float64{}

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
var PKQiJiRuFengPool = map[General_Id]float64{}

//PK其徐如林
var PKQiXuRuLinPool = map[General_Id]float64{}

//PK难知如阴
var PKNanZhiRuYinPool = map[General_Id]float64{}

//PK不动如山
var PKBuDongRuShanPool = map[General_Id]float64{}

//PK侵略如火
var PKQinLveRuHuoPool = map[General_Id]float64{}

//PK霸业奖励1
var PKBaYeJiangLi1Pool = map[General_Id]float64{}

//PK霸业奖励2
var PKBaYeJiangLi2Pool = map[General_Id]float64{}

//PK霸业奖励3
var PKBaYeJiangLi3Pool = map[General_Id]float64{}

//PK典藏卡包
var PKDianCangKaBaoPool = map[General_Id]float64{}
