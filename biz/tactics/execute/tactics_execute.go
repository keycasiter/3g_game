package execute

import (
	"context"
	"github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法执行器
//@ctx 执行环境
//@tactics 当前触发战法
//@tacticsParams 当前战法执行参数
func TacticsExecute(ctx context.Context, tactics _interface.Tactics, tacticsParams model.TacticsParams) {
	//初始化当前回合战法参数
	tactics.Init(tacticsParams).Execute()
}
