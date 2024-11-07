package api

import (
	"context"
	"testing"

	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/kr/pretty"
)

func TestBattleExecute(t *testing.T) {
	ctx := context.Background()
	apiReq := api.BattleDoRequest{
		FightingTeam: &api.BattleTeam{
			TeamType: enum.TeamType_Fighting,
			ArmType:  enum.ArmType_Spearman,
			BattleGenerals: []*api.BattleGeneral{
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                1,
						Name:              "刘备",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: []*api.Tactics{
						{
							Id: 111,
						},
						{
							Id: 222,
						},
					},
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    true,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                2,
						Name:              "关羽",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: nil,
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    false,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                3,
						Name:              "张飞",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: nil,
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    false,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
			},
		},
		EnemyTeam: &api.BattleTeam{
			TeamType: enum.TeamType_Enemy,
			ArmType:  enum.ArmType_Mauler,
			BattleGenerals: []*api.BattleGeneral{
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                1,
						Name:              "刘备",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: nil,
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    true,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                2,
						Name:              "关羽",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: nil,
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    false,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
				{
					BaseInfo: &api.MetadataGeneral{
						Id:                3,
						Name:              "张飞",
						Group:             enum.Group_ShuGuo,
						IsSupportDynamics: true,
						IsSupportCollect:  true,
					},
					EquipTactics: nil,
					WarBooks: []*api.WarBook{
						{
							Id:    111,
							Type:  int64(enum.WarbookType_Battle),
							Level: 1,
						},
						{
							Id:    222,
							Type:  int64(enum.WarbookType_Battle),
							Level: 2,
						},
						{
							Id:    333,
							Type:  int64(enum.WarbookType_Battle),
							Level: 3,
						},
					},
					SpecialTechs: []*api.SpecialTech{
						{
							Id:    111,
							Type:  enum.EquipType_Weapon,
							Level: enum.EquipLevel_S,
						},
					},
					Addition: &api.BattleGeneralAddition{
						AbilityAttr: &api.AbilityAttr{
							ForceBase:        "50",
							IntelligenceBase: "50",
							CommandBase:      "50",
							SpeedBase:        "50",
						},
					},
					IsMaster:    false,
					SoldierNum:  10000,
					ArmsAbility: enum.ArmsAbility_S,
				},
			},
		},
	}
	pretty.Logf("%v", util.ToJsonString(ctx, apiReq))
	//BattleDo(ctx, req)
}
