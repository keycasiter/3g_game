CREATE TABLE `warbook`
(
    `id`    bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`  varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  DEFAULT NULL COMMENT '兵书名称',
    `level` int                                                           DEFAULT NULL COMMENT '兵书层级',
    `desc`  varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='兵书信息表'