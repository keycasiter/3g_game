CREATE TABLE `tactic`
(
    `id`      bigint                                                       NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`    varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '战法名称',
    `quality` int                                                          NOT NULL COMMENT '战法品质',
    `source`  int                                                          NOT NULL COMMENT '战法来源',
    `type`    int                                                          NOT NULL COMMENT '战法类型',
    `desc`    varchar(1028) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='战法信息表'