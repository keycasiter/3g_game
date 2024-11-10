package consts

// 兵书类型
type WarBookType int

const (
	WarBookType_Fighting          WarBookType = 1 //作战
	WarBookType_TruthAndFalsehood WarBookType = 2 //虚实
	WarBookType_MilitaryForm      WarBookType = 3 //军形
	WarBookType_NineChanges       WarBookType = 4 //九变
)

// 兵书枚举
type WarBookDetailType int

const (
	//作战
	WarBookDetailType_1  WarBookDetailType = 1  //作战
	WarBookDetailType_2  WarBookDetailType = 2  //虚实
	WarBookDetailType_3  WarBookDetailType = 3  //军形
	WarBookDetailType_4  WarBookDetailType = 4  //九变
	WarBookDetailType_5  WarBookDetailType = 5  //奇正相生
	WarBookDetailType_6  WarBookDetailType = 6  //不勇则死
	WarBookDetailType_7  WarBookDetailType = 7  //一鼓作气
	WarBookDetailType_8  WarBookDetailType = 8  //胜而益强
	WarBookDetailType_9  WarBookDetailType = 9  //蛮勇非勇
	WarBookDetailType_10 WarBookDetailType = 10 //武略
	WarBookDetailType_11 WarBookDetailType = 11 //分险
	WarBookDetailType_12 WarBookDetailType = 12 //执锐
	WarBookDetailType_13 WarBookDetailType = 13 //文韬
	WarBookDetailType_14 WarBookDetailType = 14 //藏刀
	WarBookDetailType_15 WarBookDetailType = 15 //胜战
	//虚实
	WarBookDetailType_16 WarBookDetailType = 16 //后发先至
	WarBookDetailType_17 WarBookDetailType = 17 //以治击乱
	WarBookDetailType_18 WarBookDetailType = 18 //攻其不备
	WarBookDetailType_19 WarBookDetailType = 19 //大谋不谋
	WarBookDetailType_20 WarBookDetailType = 20 //顺应天时
	WarBookDetailType_21 WarBookDetailType = 21 //鬼谋
	WarBookDetailType_22 WarBookDetailType = 22 //妙算
	WarBookDetailType_23 WarBookDetailType = 23 //将威
	WarBookDetailType_24 WarBookDetailType = 24 //神机
	WarBookDetailType_25 WarBookDetailType = 25 //占卜
	WarBookDetailType_26 WarBookDetailType = 26 //合变
	//军形
	WarBookDetailType_27 WarBookDetailType = 27 //严阵以待
	WarBookDetailType_28 WarBookDetailType = 28 //惜兵爱民
	WarBookDetailType_29 WarBookDetailType = 29 //避其锐气
	WarBookDetailType_30 WarBookDetailType = 30 //无战而胜
	WarBookDetailType_31 WarBookDetailType = 31 //三里而还
	WarBookDetailType_32 WarBookDetailType = 32 //守而有道
	WarBookDetailType_33 WarBookDetailType = 33 //守势
	WarBookDetailType_34 WarBookDetailType = 34 //静心
	WarBookDetailType_35 WarBookDetailType = 35 //铁甲
	WarBookDetailType_36 WarBookDetailType = 36 //刚柔
	WarBookDetailType_37 WarBookDetailType = 37 //防备
	WarBookDetailType_38 WarBookDetailType = 38 //勇毅
	//九变
	WarBookDetailType_39 WarBookDetailType = 39 //援其必攻
	WarBookDetailType_40 WarBookDetailType = 40 //无功而励
	WarBookDetailType_41 WarBookDetailType = 41 //临敌不乱
	WarBookDetailType_42 WarBookDetailType = 42 //诱敌之策
	WarBookDetailType_43 WarBookDetailType = 43 //示敌以弱
	WarBookDetailType_44 WarBookDetailType = 44 //分而疾战
	WarBookDetailType_45 WarBookDetailType = 45 //掩虚
	WarBookDetailType_46 WarBookDetailType = 46 //速战
	WarBookDetailType_47 WarBookDetailType = 47 //驰援
	WarBookDetailType_48 WarBookDetailType = 48 //励军
	WarBookDetailType_49 WarBookDetailType = 49 //救主
	WarBookDetailType_50 WarBookDetailType = 50 //散仙
	WarBookDetailType_51 WarBookDetailType = 51 //百战
	WarBookDetailType_52 WarBookDetailType = 52 //疾战突围
	WarBookDetailType_53 WarBookDetailType = 53 //谋定后动
	//始计
	WarBookDetailType_57 WarBookDetailType = 57 //洞若观火
	WarBookDetailType_58 WarBookDetailType = 58 //神清气净
	WarBookDetailType_59 WarBookDetailType = 59 //应机立新
	WarBookDetailType_60 WarBookDetailType = 60 //久战
	WarBookDetailType_61 WarBookDetailType = 61 //远谋
	WarBookDetailType_62 WarBookDetailType = 62 //归心
	WarBookDetailType_63 WarBookDetailType = 63 //枕戈坐甲
	WarBookDetailType_64 WarBookDetailType = 64 //三军之众
	WarBookDetailType_65 WarBookDetailType = 65 //应机立断
	WarBookDetailType_66 WarBookDetailType = 66 //锤炼
	WarBookDetailType_67 WarBookDetailType = 67 //统军
	WarBookDetailType_71 WarBookDetailType = 71 //乐善好施
	WarBookDetailType_72 WarBookDetailType = 72 //锐利
	//用间
	WarBookDetailType_68 WarBookDetailType = 68 //兵行诡道
	WarBookDetailType_69 WarBookDetailType = 69 //以直报怨
	WarBookDetailType_70 WarBookDetailType = 70 //以退为进
	WarBookDetailType_73 WarBookDetailType = 73 //审时度势
	WarBookDetailType_74 WarBookDetailType = 74 //开阖
	WarBookDetailType_75 WarBookDetailType = 75 //持重
	WarBookDetailType_76 WarBookDetailType = 76 //仙姿
	WarBookDetailType_77 WarBookDetailType = 77 //善战
	WarBookDetailType_78 WarBookDetailType = 78 //分利
	WarBookDetailType_79 WarBookDetailType = 79 //精准
)

func (w WarBookDetailType) String() string {
	switch w {
	case WarBookDetailType_1:
		return "作战"
	case WarBookDetailType_2:
		return "虚实"
	case WarBookDetailType_3:
		return "军形"
	case WarBookDetailType_4:
		return "九变"
	case WarBookDetailType_5:
		return "奇正相生"
	case WarBookDetailType_6:
		return "不勇则死"
	case WarBookDetailType_7:
		return "一鼓作气"
	case WarBookDetailType_8:
		return "胜而益强"
	case WarBookDetailType_9:
		return "蛮勇非勇"
	case WarBookDetailType_10:
		return "武略"
	case WarBookDetailType_11:
		return "分险"
	case WarBookDetailType_12:
		return "执锐"
	case WarBookDetailType_13:
		return "文韬"
	case WarBookDetailType_14:
		return "藏刀"
	case WarBookDetailType_15:
		return "胜战"
	case WarBookDetailType_16:
		return "后发先至"
	case WarBookDetailType_17:
		return "以治击乱"
	case WarBookDetailType_18:
		return "攻其不备"
	case WarBookDetailType_19:
		return "大谋不谋"
	case WarBookDetailType_20:
		return "顺应天时"
	case WarBookDetailType_21:
		return "鬼谋"
	case WarBookDetailType_22:
		return "妙算"
	case WarBookDetailType_23:
		return "将威"
	case WarBookDetailType_24:
		return "神机"
	case WarBookDetailType_25:
		return "占卜"
	case WarBookDetailType_26:
		return "合变"
	case WarBookDetailType_27:
		return "严阵以待"
	case WarBookDetailType_28:
		return "惜兵爱民"
	case WarBookDetailType_29:
		return "避其锐气"
	case WarBookDetailType_30:
		return "无战而胜"
	case WarBookDetailType_31:
		return "三里而还"
	case WarBookDetailType_32:
		return "守而有道"
	case WarBookDetailType_33:
		return "守势"
	case WarBookDetailType_34:
		return "静心"
	case WarBookDetailType_35:
		return "铁甲"
	case WarBookDetailType_36:
		return "刚柔"
	case WarBookDetailType_37:
		return "防备"
	case WarBookDetailType_38:
		return "勇毅"
	case WarBookDetailType_39:
		return "援其必攻"
	case WarBookDetailType_40:
		return "无功而励"
	case WarBookDetailType_41:
		return "临敌不乱"
	case WarBookDetailType_42:
		return "诱敌之策"
	case WarBookDetailType_43:
		return "示敌以弱"
	case WarBookDetailType_44:
		return "分而疾战"
	case WarBookDetailType_45:
		return "掩虚"
	case WarBookDetailType_46:
		return "速战"
	case WarBookDetailType_47:
		return "驰援"
	case WarBookDetailType_48:
		return "励军"
	case WarBookDetailType_49:
		return "救主"
	case WarBookDetailType_50:
		return "散仙"
	case WarBookDetailType_51:
		return "百战"
	case WarBookDetailType_52:
		return "疾战突围"
	case WarBookDetailType_53:
		return "谋定后动"
	case WarBookDetailType_57:
		return "洞若观火"
	case WarBookDetailType_58:
		return "神清气净"
	case WarBookDetailType_59:
		return "应机立新"
	case WarBookDetailType_60:
		return "久战"
	case WarBookDetailType_61:
		return "远谋"
	case WarBookDetailType_62:
		return "归心"
	case WarBookDetailType_63:
		return "枕戈坐甲"
	case WarBookDetailType_64:
		return "三军之众"
	case WarBookDetailType_65:
		return "应机立断"
	case WarBookDetailType_66:
		return "锤炼"
	case WarBookDetailType_67:
		return "统军"
	case WarBookDetailType_71:
		return "乐善好施"
	case WarBookDetailType_72:
		return "锐利"
	case WarBookDetailType_68:
		return "兵行诡道"
	case WarBookDetailType_69:
		return "以直报怨"
	case WarBookDetailType_70:
		return "以退为进"
	case WarBookDetailType_73:
		return "审时度势"
	case WarBookDetailType_74:
		return "开阖"
	case WarBookDetailType_75:
		return "持重"
	case WarBookDetailType_76:
		return "仙姿"
	case WarBookDetailType_77:
		return "善战"
	case WarBookDetailType_78:
		return "分利"
	case WarBookDetailType_79:
		return "精准"
	default:
		return "未知兵书"
	}
}
