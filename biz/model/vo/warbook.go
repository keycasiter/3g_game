package vo

type QueryWarbookCondition struct {
	Id    int64  // 主键ID
	Name  string // 兵书名称
	Level int32  // 兵书层级
}

// 脱敏
func DesensitizeName(name string) string {
	if len(name) == 0 {
		return ""
	}
	if len(name) == 1 {
		return "*"
	}
	if len(name) > 1 {
		markNum := len(name) - 1
		markRes := ""
		markRes += "张"
		for i := 0; i < markNum; i++ {
			markRes += "*"
		}
		return markRes
	}
	return ""
}
