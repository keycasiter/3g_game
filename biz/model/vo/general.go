package vo

type QueryGeneralCondition struct { // nolint:maligned
	Id                int64  // 主键ID
	Name              string // 姓名
	Gender            int8   // 性别
	Control           int32  // 统御
	Group             int8   // 阵营
	Quality           int8   // 品质
	Tags              []int  // 标签
	IsSupportDynamics int8   // 是否支持动态
	IsSupportCollect  int8   // 是否支持典藏
}
