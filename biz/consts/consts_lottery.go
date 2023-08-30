package consts

//***抽卡卡池***

type GeneralLotteryPool int

const (
	//PK动如雷霆
	PK_DongRuLeiTing GeneralLotteryPool = 1
)

func GetGeneralPool(generalLotteryPool GeneralLotteryPool) map[General_Id]float64 {
	switch generalLotteryPool {
	//PK动如雷霆
	case PK_DongRuLeiTing:
		return PK_DongRuLeiTingPool
	}
	return nil
}

//PK赛季名将2

//PK动如雷霆
var PK_DongRuLeiTingPool = map[General_Id]float64{
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
