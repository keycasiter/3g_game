api:
	hz update -idl idl/api.thrift
	# hz new -force -idl idl/api.thrift -module github.com/keycasiter/3g_game

dbmodel:
	# 武将信息表
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath general.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath special_tech.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath tactic.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath team.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath warbook.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath jym_goods.sql --dbFunc
	code_gen dbmodel --prefix github.com/keycasiter/3g_game --ddlPath user_battle_record.sql --dbFunc