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
	ForceBase        float64 `bson:"force_base"`
	ForceRate        float64 `bson:"force_rate"`
	IntelligenceBase float64 `bson:"intelligence_base"`
	IntelligenceRate float64 `bson:"intelligence_rate"`
	CharmBase        float64 `bson:"charm_base"`
	CharmRate        float64 `bson:"charm_rate"`
	CommandBase      float64 `bson:"command_base"`
	CommandRate      float64 `bson:"command_rate"`
	PoliticsBase     float64 `bson:"politics_base"`
	PoliticsRate     float64 `bson:"politics_rate"`
	SpeedBase        float64 `bson:"speed_base"`
	SpeedRate        float64 `bson:"speed_rate"`
}

// 兵种属性
type ArmsAttr struct {
	Cavalry   consts.ArmsAbility `bson:"cavalry"`
	Mauler    consts.ArmsAbility `bson:"mauler"`
	Archers   consts.ArmsAbility `bson:"archers"`
	Spearman  consts.ArmsAbility `bson:"spearman"`
	Apparatus consts.ArmsAbility `bson:"apparatus"`
}

// 战法资料
type Tactics struct {
	Id            consts.TacticId `bson:"id"`
	Name          string          `bson:"name"`
	SelfContained string          `bson:"self_contained"`
	Inherit       string          `bson:"inherit"`
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
