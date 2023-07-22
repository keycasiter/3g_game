package po

import "github.com/keycasiter/3g_game/biz/consts"

// 武将资料
type MetadataGeneral struct {
	//mongodb 持久化字段
	Id          int64               `bson:"_id"`
	Name        string              `bson:"name"`
	Gender      consts.Gender       `bson:"gender"`
	Group       consts.Group        `bson:"group"`
	GeneralTag  []consts.GeneralTag `bson:"general_tag"`
	AvatarUri   string              `bson:"avatar_uri"`
	AbilityAttr *AbilityAttr        `bson:"ability_attr"`
	ArmsAttr    *ArmsAttr           `bson:"arms_attr"`
	Tactics     *Tactics            `bson:"tactics"`
	Biographies *Biographies        `bson:"biographies"`

	//业务字段
	GeneralBattleType consts.GeneralBattleType
	//唯一对战ID
	UniqueId int64
}

// 能力属性
type AbilityAttr struct {
	ForceBase        float64
	ForceRate        float64
	IntelligenceBase float64
	IntelligenceRate float64
	CharmBase        float64
	CharmRate        float64
	CommandBase      float64
	CommandRate      float64
	PoliticsBase     float64
	PoliticsRate     float64
	SpeedBase        float64
	SpeedRate        float64
}

// 兵种属性
type ArmsAttr struct {
	Cavalry   consts.ArmsAbility
	Mauler    consts.ArmsAbility
	Archers   consts.ArmsAbility
	Spearman  consts.ArmsAbility
	Apparatus consts.ArmsAbility
}

// 战法资料
type Tactics struct {
	Id            consts.TacticId      `bson:"id"`
	Name          string               `bson:"name"`
	TacticsSource consts.TacticsSource `bson:"tactics_source"`
	Type          consts.TacticsType   `bson:"type"`
}

// 列传资料
type Biographies struct {
	Desc           string            `bson:"desc"`
	Predestination []*Predestination `bson:"predestination"`
}

// 缘分资料
type Predestination struct {
	Name         string   `bson:"name"`
	ReferGeneral []string `bson:"refer_general"`
}
