SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  数据条结构
-- ----------------------------


DROP TABLE IF EXISTS `iterms`;
CREATE TABLE `item` (
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `did`        int(10) unsigned NOT NULL DEFAULT '0',
    `light`      float(6,2)  DEFAULT '0',
    `mq2`        float(6,2)  DEFAULT '0',
    `mq135`      float(6,2)  DEFAULT '0',
    `temp`       float(6,2)  DEFAULT '0',
    `wet`        float(6,2)  DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户管理';

