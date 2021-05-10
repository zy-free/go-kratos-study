CREATE TABLE `member`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `phone`      VARCHAR(30)     NOT NULL COMMENT '手机号',
    `name`       VARCHAR(30)     NOT NULL COMMENT '名字',
    `age`        TINYINT         NOT NULL DEFAULT 0 COMMENT '年龄',
    `address`    VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '地址',
    `order_num`  TINYINT(6)      NOT NULL DEFAULT '0' COMMENT '排序',
    `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at` DATETIME                 DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `ukx_phone` (`phone`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

CREATE TABLE `member_favorite`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '收藏夹id',
    `mid`        BIGINT UNSIGNED NOT NULL COMMENT '用户id',
    `name`       VARCHAR(30)     NOT NULL COMMENT '收藏夹名称',
    `hint_at`    DATETIME                 DEFAULT NULL COMMENT '提示时间(做null测试)',
    `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户收藏夹表';

