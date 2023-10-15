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
	UniqueId string
}

// 能力属性
type AbilityAttr struct {
	ForceBase        float64 `json:"force_base"`
	ForceRate        float64 `json:"force_rate"`
	IntelligenceBase float64 `json:"intelligence_base"`
	IntelligenceRate float64 `json:"intelligence_rate"`
	CharmBase        float64 `json:"charm_base"`
	CharmRate        float64 `json:"charm_rate"`
	CommandBase      float64 `json:"command_base"`
	CommandRate      float64 `json:"command_rate"`
	PoliticsBase     float64 `json:"politics_base"`
	PoliticsRate     float64 `json:"politics_rate"`
	SpeedBase        float64 `json:"speed_base"`
	SpeedRate        float64 `json:"speed_rate"`
}

// 能力属性
type AbilityAttrString struct {
	ForceBase        string `json:"force_base"`
	ForceRate        string `json:"force_rate"`
	IntelligenceBase string `json:"intelligence_base"`
	IntelligenceRate string `json:"intelligence_rate"`
	CharmBase        string `json:"charm_base"`
	CharmRate        string `json:"charm_rate"`
	CommandBase      string `json:"command_base"`
	CommandRate      string `json:"command_rate"`
	PoliticsBase     string `json:"politics_base"`
	PoliticsRate     string `json:"politics_rate"`
	SpeedBase        string `json:"speed_base"`
	SpeedRate        string `json:"speed_rate"`
}

// 兵种属性
type ArmsAttr struct {
	Cavalry   consts.ArmsAbility
	Mauler    consts.ArmsAbility
	Archers   consts.ArmsAbility
	Spearman  consts.ArmsAbility
	Apparatus consts.ArmsAbility
}

// 兵种属性
type ArmsAttrStr struct {
	Cavalry   string `json:"cavalry"`
	Mauler    string `json:"mauler"`
	Archers   string `json:"archers"`
	Spearman  string `json:"spearman"`
	Apparatus string `json:"apparatus"`
}

// 战法资料
type Tactics struct {
	Id            consts.TacticId      `json:"id"`
	Name          string               `json:"name"`
	TacticsSource consts.TacticsSource `json:"tactics_source"`
	Type          consts.TacticsType   `json:"type"`
}

// 列传资料
type Biographies struct {
	Desc           string            `json:"desc"`
	Predestination []*Predestination `json:"predestination"`
}

// 缘分资料
type Predestination struct {
	Name         string   `bson:"name"`
	ReferGeneral []string `bson:"refer_general"`
}
