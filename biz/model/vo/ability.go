package vo

// 能力属性
type AbilityAttr struct {
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
