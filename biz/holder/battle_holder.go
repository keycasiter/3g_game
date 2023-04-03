package holder

import "sync"

// 当前执行武将 map<currentBattleId,generalId>
var BattleCurrentExecuteGeneralMap sync.Map

// 当前执行武将 map<currentBattleId,round>
var BattleCurrentExecuteRoundMap sync.Map
