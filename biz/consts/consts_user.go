package consts

/** 用户等级 **/
type UserLevel int

const (
	UserLevel_Common  UserLevel = 0 //普通用户
	UserLevel_Silver  UserLevel = 1 //白银用户
	UserLevel_Golden  UserLevel = 2 // 黄金用户
	UserLevel_Diamond UserLevel = 3 // 钻石用户
)
