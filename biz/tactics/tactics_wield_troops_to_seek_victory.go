package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 挥兵谋胜
// 战斗中，每当我军抵御被消耗时，使我军单体恢复一定兵力（治疗率94%，受智力影响），
// 抵御持续时间结束但未消耗时，使我军武力最高武将对敌军随机武将发动一次兵刃攻击（伤害率94%）
// 战斗前3回合，每回合有60%概率（受智力影响），自身为主将时，基础概率提升至70%，
// 使我军群体（2～3人）获得1次抵御，持续1回合
// 指挥，100%
type WieldTroopsToSeekVictoryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WieldTroopsToSeekVictoryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Prepare() {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Id() consts.TacticId {
	return consts.WieldTroopsToSeekVictory
}

func (w WieldTroopsToSeekVictoryTactic) Name() string {
	return "挥兵谋胜"
}

func (w WieldTroopsToSeekVictoryTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Execute() {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
