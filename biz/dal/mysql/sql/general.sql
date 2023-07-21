CREATE TABLE `general`
(
    `id`                  bigint                                                       NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`                varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
    `gender`              tinyint                                                      NOT NULL COMMENT '性别',
    `control`             int                                                          NOT NULL COMMENT '统御',
    `group`               tinyint                                                      NOT NULL COMMENT '阵营',
    `quality`             tinyint                                                      NOT NULL COMMENT '品质',
    `tag`                 varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '标签',
    `ability_attr`        varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '能力属性',
    `avatar_url`          varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '头像url',
    `arm_attr`            varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '兵种适性',
    `self_tactic_id`      int                                                           DEFAULT NULL COMMENT '自带战法ID',
    `is_support_dynamics` tinyint                                                       DEFAULT NULL COMMENT '是否支持动态',
    `is_support_collect`  tinyint                                                       DEFAULT NULL COMMENT '是否支持典藏',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='武将信息表'