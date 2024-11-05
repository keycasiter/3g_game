CREATE TABLE `user_battle_record` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
      `uid` bigint NOT NULL COMMENT '用户id',
      `battle_record` longtext COMMENT '对战记录json',
      `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
      `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
      PRIMARY KEY (`id`),
      KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户对战记录'