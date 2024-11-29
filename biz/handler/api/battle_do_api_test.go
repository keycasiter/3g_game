package api

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/cache"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
)

func TestBattleExecute(t *testing.T) {

	conf.InitConfig()
	mysql.InitMysql()
	cache.InitCache()

	//ctx := context.Background()
	apiReq := []byte(`{
    "FightingTeam": {
        "TeamType": 1,
        "ArmType": 3,
        "BattleGenerals": [
            {
                "BaseInfo": {
                    "Id": 3,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 1,
                        "Mauler": 2,
                        "Archers": 1,
                        "Spearman": 4,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
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
                        "Id": 2,
                        "Name": "虚实",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 17,
                        "Name": "以治击乱",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 24,
                        "Name": "神机",
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
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {
                        "ForceBase": "",
                        "IntelligenceBase": "50",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": true,
                "SoldierNum": 10000,
                "ArmsAbility": 1,
                "RemainNum": 0
            },
            {
                "BaseInfo": {
                    "Id": 8,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 2,
                        "Mauler": 3,
                        "Archers": 1,
                        "Spearman": 3,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
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
                        "Id": 83,
                        "Name": "众动万计",
                        "TacticsSource": 3,
                        "Type": 2,
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
                        "Id": 35,
                        "Name": "铁甲",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 34,
                        "Name": "静心",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {
                        "ForceBase": "",
                        "IntelligenceBase": "",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 3,
                "RemainNum": 0
            },
            {
                "BaseInfo": {
                    "Id": 20,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 2,
                        "Mauler": 1,
                        "Archers": 2,
                        "Spearman": 3,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
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
                        "Id": 9,
                        "Name": "暂避其锋",
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
                        "ForceBase": "",
                        "IntelligenceBase": "",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 3,
                "RemainNum": 0
            }
        ],
        "BuildingTechAttrAddition": null,
        "BuildingTechGroupAddition": null,
        "SoliderNum": 0,
        "RemainNum": 0,
        "Name": "",
        "AvatarUrl": ""
    },
    "EnemyTeam": {
        "TeamType": 2,
        "ArmType": 4,
        "BattleGenerals": [
            {
                "BaseInfo": {
                    "Id": 42,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 1,
                        "Mauler": 3,
                        "Archers": 3,
                        "Spearman": 1,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
                },
                "EquipTactics": [
                    {
                        "Id": 98,
                        "Name": "裸衣血战",
                        "TacticsSource": 3,
                        "Type": 2,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 96,
                        "Name": "血刃争锋",
                        "TacticsSource": 3,
                        "Type": 2,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    }
                ],
                "WarBooks": [
                    {
                        "Id": 1,
                        "Name": "作战",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 8,
                        "Name": "胜而益强",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 15,
                        "Name": "胜战",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 10,
                        "Name": "武略",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {
                        "ForceBase": "",
                        "IntelligenceBase": "",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": true,
                "SoldierNum": 10000,
                "ArmsAbility": 1,
                "RemainNum": 0
            },
            {
                "BaseInfo": {
                    "Id": 50,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 1,
                        "Mauler": 3,
                        "Archers": 4,
                        "Spearman": 1,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
                },
                "EquipTactics": [
                    {
                        "Id": 71,
                        "Name": "青州兵",
                        "TacticsSource": 3,
                        "Type": 6,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 101,
                        "Name": "据水断桥",
                        "TacticsSource": 3,
                        "Type": 1,
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
                        "Id": 17,
                        "Name": "以治击乱",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 24,
                        "Name": "神机",
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
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {
                        "ForceBase": "",
                        "IntelligenceBase": "",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 0,
                "RemainNum": 0
            },
            {
                "BaseInfo": {
                    "Id": 54,
                    "Name": "",
                    "Gender": 0,
                    "Group": 0,
                    "GeneralTag": null,
                    "AvatarUrl": "",
                    "AbilityAttr": null,
                    "ArmsAttr": {
                        "Cavalry": 2,
                        "Mauler": 1,
                        "Archers": 3,
                        "Spearman": 1,
                        "Apparatus": 0
                    },
                    "GeneralBattleType": 0,
                    "SelfTactic": null,
                    "GeneralQuality": 0,
                    "IsSupportDynamics": false,
                    "IsSupportCollect": false
                },
                "EquipTactics": [
                    {
                        "Id": 31,
                        "Name": "破阵摧坚",
                        "TacticsSource": 2,
                        "Type": 1,
                        "Quality": 1,
                        "TriggerRate": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 30,
                        "Name": "所向披靡",
                        "TacticsSource": 2,
                        "Type": 1,
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
                        "Id": 17,
                        "Name": "以治击乱",
                        "Type": 0,
                        "Level": 0,
                        "Desc": ""
                    },
                    {
                        "Id": 24,
                        "Name": "神机",
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
                    }
                ],
                "SpecialTechs": [

                ],
                "Addition": {
                    "AbilityAttr": {
                        "ForceBase": "",
                        "IntelligenceBase": "",
                        "CommandBase": "",
                        "SpeedBase": ""
                    },
                    "GeneralLevel": 50,
                    "GeneralStarLevel": 0
                },
                "IsMaster": false,
                "SoldierNum": 10000,
                "ArmsAbility": 1,
                "RemainNum": 0
            }
        ],
        "BuildingTechAttrAddition": null,
        "BuildingTechGroupAddition": null,
        "SoliderNum": 0,
        "RemainNum": 0,
        "Name": "",
        "AvatarUrl": ""
    },
    "Uid": 1462451334
}`)

	c := &app.RequestContext{}
	c.Request.SetBodyRaw(apiReq)
	c.Request.SetHeader("Content-Type", "application/json")

	BattleDo(context.Background(), c)
}
