package consts

type General_Id int64

const (
	//未知武将
	UnknownGeneral General_Id = iota
	SPXunYu
	DianWei
	JiaXu
	SiMaYi
	CaoCao
	ZhangLiao
	SPLiuYe
	SPPangDe
	SPGuoJia
	XuChu
	ZhangHe
	HaoZhao
	CaoRen
	ChengYu
	XunYou
	XuHuang
	XiaHouDun
	PangDe
	ZhongHui
	ManChong
	WangShuang
	WangYuanJi
	CaoChun
	YuJin
	LeJin
	DengAi
	XiaHouYuan
	GuoJia
	ZhangChunHua
	LiuYe
	WenPin
	CaoZhang
	CaoHong
	CaoZhen
	WangLang
	ZangBa
	GuoHuai
	LiDian
	SPGuanYu
	LiuBei
	PangTong
	MaChao
	ZhuGeLiang
	GuanYu
	SPZhuGeLiang
	YiJi
	YanYan
	ShuGuoZhangBao
	GuanXing
	GuanYinPing
	MaYunLu
	ChenDao
	JiangWei
	WeiYan
	HuangZhong
	ZhaoYun
	ZhangFei
	WangPing
	XuShu
	ZhangJi
	HuangYueYing
	FaZheng
	HuangQuan
	ZhouCang
	LiuFeng
	GuanPing
	MaSu
	LiaoHua
	MaLiang
	ShaMoKe
	SPLvMeng
	SunShangXiang
	LuXun
	SPZhouYu
	MaZhong
	LingTong
	LuSu
	SunQuan
	GanNing
	ZhouTai
	LvMeng
	TaiShiCi
	SunJian
	LuKang
	ZhouYu
	HuangGai
	ChengPu
	SunCe
	ZhuGeKe
	ChenWu
	ZhuHuan
	HanDang
	DongXi
	JiangQin
	DingFeng
	PanZhang
	XuSheng
	SPZhuJun
	SPYuanShao
	YuanShu
	MengHuo
	YuJi
	DongZhuo
	LvBu
	XuYou
	DuoSiDaWang
	JuShou
	TianFeng
	LvLingQi
	ZhuRongFuRen
	WuTuGu
	GongSunZan
	YuanShao
	ZhangJiao
	ZhangRang
	GaoLan
	MuLuDaWang
	LiRu
	GaoShun
	MaTeng
	WenChou
	HuaXiong
	YanLiang
	HuaTuo
	ZuoCi
	DiaoChan
	ChenGong
	ZouShi
	DongBai
	CaiWenJi
	HuCheEr
	FengJi
	JiLing
	ZhangXiu
	ZhuJun
	KongRong
	GuoSi
	HanSui
	ZhangYan
	ZhangManCheng
	LiJue
	ShenPei
	HuangFuSong
	ZhangLiang
	ZhangRen
	GuoTu
	GuanHai
	ZhangBao
	SPZhangBao
	SPHuangFuSong
	SPZhangLiang
	YangXiu
	DengZhi
	HuaXin
	YuFan
	ZhuRan
	LvQian
	LiuYao
)

var General5LevMap = map[General_Id]bool{
	SPXunYu:        true,
	DianWei:        true,
	JiaXu:          true,
	SiMaYi:         true,
	CaoCao:         true,
	ZhangLiao:      true,
	SPLiuYe:        true,
	SPPangDe:       true,
	SPGuoJia:       true,
	XuChu:          true,
	ZhangHe:        true,
	HaoZhao:        true,
	CaoRen:         true,
	ChengYu:        true,
	XunYou:         true,
	XuHuang:        true,
	XiaHouDun:      true,
	PangDe:         true,
	ZhongHui:       true,
	ManChong:       true,
	WangShuang:     true,
	WangYuanJi:     true,
	CaoChun:        true,
	YuJin:          true,
	LeJin:          true,
	DengAi:         true,
	XiaHouYuan:     true,
	GuoJia:         true,
	ZhangChunHua:   true,
	SPGuanYu:       true,
	LiuBei:         true,
	PangTong:       true,
	MaChao:         true,
	ZhuGeLiang:     true,
	GuanYu:         true,
	SPZhuGeLiang:   true,
	YiJi:           true,
	YanYan:         true,
	ShuGuoZhangBao: true,
	GuanXing:       true,
	GuanYinPing:    true,
	MaYunLu:        true,
	ChenDao:        true,
	JiangWei:       true,
	WeiYan:         true,
	HuangZhong:     true,
	ZhaoYun:        true,
	ZhangFei:       true,
	WangPing:       true,
	XuShu:          true,
	ZhangJi:        true,
	HuangYueYing:   true,
	FaZheng:        true,
	SPLvMeng:       true,
	SunShangXiang:  true,
	LuXun:          true,
	SPZhouYu:       true,
	MaZhong:        true,
	LingTong:       true,
	LuSu:           true,
	SunQuan:        true,
	GanNing:        true,
	ZhouTai:        true,
	LvMeng:         true,
	TaiShiCi:       true,
	SunJian:        true,
	LuKang:         true,
	ZhouYu:         true,
	HuangGai:       true,
	ChengPu:        true,
	SunCe:          true,
	ZhuGeKe:        true,
	SPZhuJun:       true,
	SPYuanShao:     true,
	YuanShu:        true,
	MengHuo:        true,
	YuJi:           true,
	DongZhuo:       true,
	LvBu:           true,
	XuYou:          true,
	DuoSiDaWang:    true,
	JuShou:         true,
	TianFeng:       true,
	LvLingQi:       true,
	ZhuRongFuRen:   true,
	WuTuGu:         true,
	GongSunZan:     true,
	YuanShao:       true,
	ZhangJiao:      true,
	ZhangRang:      true,
	GaoLan:         true,
	MuLuDaWang:     true,
	LiRu:           true,
	GaoShun:        true,
	MaTeng:         true,
	WenChou:        true,
	HuaXiong:       true,
	YanLiang:       true,
	HuaTuo:         true,
	ZuoCi:          true,
	DiaoChan:       true,
	ChenGong:       true,
	ZouShi:         true,
	DongBai:        true,
	CaiWenJi:       true,
	SPZhangBao:     true,
	SPHuangFuSong:  true,
	SPZhangLiang:   true,
}
