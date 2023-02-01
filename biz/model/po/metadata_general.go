package po

// 武将资料
type MetadataGeneral struct {
	Id          string       `bson:"_id"`
	Name        string       `bson:"name"`
	Gender      string       `bson:"gender"`
	Group       string       `bson:"group"`
	GeneralTag  []string     `bson:"general_tag"`
	AvatarUri   string       `bson:"avatar_uri"`
	AbilityAttr *AbilityAttr `bson:"ability_attr"`
	ArmsAttr    *ArmsAttr    `bson:"arms_attr"`
	Tactics     *Tactics     `bson:"tactics"`
	Biographies *Biographies `bson:"biographies"`
}

//能力属性
type AbilityAttr struct {
	ForceBase        string `bson:"force_base"`
	ForceRate        string `bson:"force_rate"`
	IntelligenceBase string `bson:"intelligence_base"`
	IntelligenceRate string `bson:"intelligence_rate"`
	CharmBase        string `bson:"charm_base"`
	CharmRate        string `bson:"charm_rate"`
	CommandBase      string `bson:"command_base"`
	CommandRate      string `bson:"command_rate"`
	PoliticsBase     string `bson:"politics_base"`
	PoliticsRate     string `bson:"politics_rate"`
	SpeedBase        string `bson:"speed_base"`
	SpeedRate        string `bson:"speed_rate"`
}

//兵种属性
type ArmsAttr struct {
	Cavalry   string `bson:"cavalry"`
	Mauler    string `bson:"mauler"`
	Archers   string `bson:"archers"`
	Spearman  string `bson:"spearman"`
	Apparatus string `bson:"apparatus"`
}

//战法资料
type Tactics struct {
	Id            int64  `bson:"id"`
	Name          string `bson:"name"`
	SelfContained string `bson:"self_contained"`
	Inherit       string `bson:"inherit"`
}

//列传资料
type Biographies struct {
	Desc           string            `bson:"desc"`
	Predestination []*Predestination `bson:"predestination"`
}

//缘分资料
type Predestination struct {
	Name         string   `bson:"name"`
	ReferGeneral []string `bson:"refer_general"`
}
