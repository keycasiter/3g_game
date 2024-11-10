package api

import (
	"encoding/json"
	"testing"

	"github.com/keycasiter/3g_game/biz/model/api"
)

func TestBattleExecute(t *testing.T) {
	//ctx := context.Background()
	apiReq := &api.BattleDoRequest{}
	json.Unmarshal([]byte(`{
    "FightingTeam": {
        "TeamType": 1,
        "ArmType": 2,
        "BattleGenerals": [
            {
                "BaseInfo": {
                    "Id": 4,
                    "IsSupportDynamics": true,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 2,
                        "Name": "士别三日",
                        "TacticsSource": 2,
                        "Type": 2,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 57,
                        "Name": "用武通神",
                        "TacticsSource": 1,
                        "Type": 3,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 3,
                        "Name": "军形",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 28,
                        "Name": "惜兵爱民",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 38,
                        "Name": "勇毅",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 33,
                        "Name": "守势",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 50
                },
                "IsMaster": true,
                "SoldierNum": 10000,
                "ArmsAbility": 1
            },
            {
                "BaseInfo": {
                    "Id": 5,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 10,
                        "Name": "魅惑",
                        "TacticsSource": 2,
                        "Type": 2,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 13,
                        "Name": "锋矢阵",
                        "TacticsSource": 2,
                        "Type": 5,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 4,
                        "Name": "九变",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 41,
                        "Name": "临敌不乱",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 45,
                        "Name": "掩虚",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 48,
                        "Name": "励军",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 2
            },
            {
                "BaseInfo": {
                    "Id": 20,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 14,
                        "Name": "刮骨疗毒",
                        "TacticsSource": 2,
                        "Type": 1,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 11,
                        "Name": "抚揖军民",
                        "TacticsSource": 2,
                        "Type": 3,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 4,
                        "Name": "九变",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 39,
                        "Name": "援其必攻",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 45,
                        "Name": "掩虚",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 48,
                        "Name": "励军",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 2
            }
        ]
    },
    "EnemyTeam": {
        "TeamType": 2,
        "ArmType": 3,
        "BattleGenerals": [
            {
                "BaseInfo": {
                    "Id": 53,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 16,
                        "Name": "夺魂挟魄",
                        "TacticsSource": 2,
                        "Type": 1,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 32,
                        "Name": "杯蛇鬼车",
                        "TacticsSource": 2,
                        "Type": 1,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 12,
                        "Name": "执锐",
                        "level": 3,
                        "category": "作战"
                    },
                    {
                        "Id": 13,
                        "Name": "文韬",
                        "level": 3,
                        "category": "作战"
                    },
                    {
                        "Id": 14,
                        "Name": "藏刀",
                        "level": 3,
                        "category": "作战"
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": true,
                "SoldierNum": 10000,
                "ArmsAbility": 1
            },
            {
                "BaseInfo": {
                    "Id": 41,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 19,
                        "Name": "太平道法",
                        "TacticsSource": 3,
                        "Type": 2,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 20,
                        "Name": "无当飞军",
                        "TacticsSource": 2,
                        "Type": 6,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 2,
                        "Name": "虚实",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 16,
                        "Name": "后发先至",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 23,
                        "Name": "将威",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 21,
                        "Name": "鬼谋",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 1
            },
            {
                "BaseInfo": {
                    "Id": 43,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false,
                    "ArmsAttr": {

                    }
                },
                "EquipTactics": [
                    {
                        "Id": 22,
                        "Name": "八门金锁阵",
                        "TacticsSource": 2,
                        "Type": 5,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 99,
                        "Name": "婴城自守",
                        "TacticsSource": 3,
                        "Type": 1,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 4,
                        "Name": "九变",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 39,
                        "Name": "援其必攻",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 47,
                        "Name": "驰援",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 45,
                        "Name": "掩虚",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {

                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 1
            }
        ]
    },
    "Uid": 1462451334
}`), &apiReq)

	//pretty.Logf("%v", util.ToJsonString(ctx, apiReq))
	//BattleDo(ctx, apiReq)
}
