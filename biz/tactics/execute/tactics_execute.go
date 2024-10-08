package execute

import (
	"context"
	"github.com/keycasiter/3g_game/biz/tactics/interface"
)

// 战法执行器
// @ctx 执行环境
// @tactics 当前触发战法
// @tacticsParams 当前战法执行参数
func TacticsExecute(ctx context.Context, tactic _interface.Tactics) {
	//战法准备
	tactic.Prepare()
	//战法执行
	tactic.Execute()
}
