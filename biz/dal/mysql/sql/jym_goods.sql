CREATE TABLE `jym_goods`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `goods_url`    varchar(512) NOT NULL COMMENT '商品链接',
    `goods_detail` longtext COMMENT '商品信息',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='交易猫商品信息'