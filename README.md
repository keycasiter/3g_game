# 3g_game
🎮 三国志战略版配将小程序

# swag
https://gitee.com/xvwen/hertz-swag
https://github.com/swaggo/swag/blob/master/README_zh-CN.md

# 📁 工程结构
&emsp;&emsp;没有过于复杂的架构设计，也不需要，简单的单体服务。

- biz 业务逻辑的核心代码
  - config 配置文件读取和存储
  - consts 静态常量、全局枚举等，如武将ID、战法效果、战法ID、兵书等
  - dal 数据持久层，目前用到了mongodb、mysql
  - handler 接口定义层，使用了herz作为http框架
  - logic 业务逻辑层
  - model 对象层
  - router 这里是router配置和分发的实现
  - tactics 这里是战法定义和具体实现模块，属于logic的一部分，独立出来维护
  - util 工具类，封装了大量的随机概率、伤害、恢复、战报、队伍、效果等能力
- conf 配置文件，多环境
- idl 接口定义
- 其他

# 🔰 对战生命周期定义

BattleAction_Unknow            BattleAction = iota //未知动作
BattleAction_BeginAction                           //开始行动
BattleAction_EndAction                             //结束行动
BattleAction_Attack                                //普通攻击开始
BattleAction_AttackEnd                             //普通攻击结束
BattleAction_ActiveTactic                          //发动主动战法开始
BattleAction_ActiveTacticEnd                       //发动主动战法结束
BattleAction_AssaultTactic                         //发动突击战法开始
BattleAction_AssaultTacticEnd                      //发动突击战法结束
BattleAction_CommandTactic                         //指挥战法攻击开始
BattleAction_CommandTacticEnd                      //指挥战法攻击结束
BattleAction_PassiveTactic                         //被动战法攻击开始
BattleAction_PassiveTacticEnd                      //被动战法攻击结束
BattleAction_TroopsTactic                          //阵法战法攻击开始
BattleAction_TroopsTacticEnd                       //阵法战法攻击结束
BattleAction_ArmTactic                             //兵种战法攻击开始
BattleAction_ArmTacticEnd                          //兵种战法攻击借宿
BattleAction_WeaponDamage                          //发动兵刃伤害开始
BattleAction_WeaponDamageEnd                       //发动兵刃伤害结束
BattleAction_StrategyDamage                        //发动谋略伤害开始
BattleAction_StrategyDamageEnd                     //发动谋略伤害结束
BattleAction_DebuffEffect                          //施加负面效果开始
BattleAction_DebuffEffectEnd                       //施加负面效果结束
BattleAction_BuffEffect                            //施加正面效果开始
BattleAction_BuffEffectEnd                         //施加正面效果结束

//遭受伤害
BattleAction_SufferDamage           //遭受伤害开始
BattleAction_SufferDamageEnd        //遭受伤害结束
BattleAction_SufferGeneralAttack    //被普通攻击开始
BattleAction_SufferGeneralAttackEnd //被普通攻击结束
BattleAction_SufferActiveTactic     //被主动战法攻击开始
BattleAction_SufferActiveTacticEnd  //被主动战法攻击结束
BattleAction_SufferAssaultTactic    //被突击战法攻击开始
BattleAction_SufferAssaultTacticEnd //被突击战法攻击结束
BattleAction_SufferCommandTactic    //被指挥战法攻击开始
BattleAction_SufferCommandTacticEnd //被指挥战法攻击结束
BattleAction_SufferArmTactic        //被兵种战法攻击开始
BattleAction_SufferArmTacticEnd     //被兵种战法攻击结束
BattleAction_SufferTroopsTactic     //被阵法战法攻击开始
BattleAction_SufferTroopsTacticEnd  //被阵法战法攻击结束
BattleAction_SufferPassiveTactic    //被被动战法攻击开始
BattleAction_SufferPassiveTacticEnd //被被动战法攻击结束
BattleAction_SufferDebuffEffect     //被施加负面效果开始
BattleAction_SufferDebuffEffectEnd  //被施加负面效果结束
BattleAction_SufferBuffEffect       //被施加正面效果开始
BattleAction_SufferBuffEffectEnd    //被施加正面效果结束
