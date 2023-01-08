// idl/common.thrift
namespace go common

struct Meta {
  1: i64  status_code
  2: string status_msg
}

//武将资料
struct MetadataGeneral {
  1: string          Id
  2: string          Name
  3: string          Gender
  4: string          Group
  5: list<string>    GeneralTags
  6: string          AvatarUri
  7: AbilityAttr     AbilityAttr
  8: ArmsAttr        ArmsAttr
  9: Tactics         Tactics
  10: Biographies     Biographies
}

//能力属性
struct AbilityAttr {
  1: string ForceBase
  2: string ForceRate
  3: string IntelligenceBase
  4: string IntelligenceRate
  5: string CharmBase
  6: string CharmRate
  7: string CommandBase
  8: string CommandRate
  9: string PoliticsBase
  10: string PoliticsRate
  11: string SpeedBase
  12: string SpeedRate
}

//兵种属性
struct ArmsAttr {
  1: string Cavalry
  2: string Mauler
  3: string Archers
  4: string Spearman
  5: string Apparatus
}

//战法资料
struct Tactics {
  1: string SelfContained
  2: string Inherit
}

//列传资料
struct Biographies {
  1: string                  Desc
  2: list<Predestination> Predestinations
}

//缘分资料
struct Predestination {
  1: string          Name
  2: list<string> ReferGenerals
}
