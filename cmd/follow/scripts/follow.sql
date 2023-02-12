CREATE TABLE `follow_user` (
                               `id` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                               `user_id` int DEFAULT '' COMMENT '用户id',
                               `follow_id` int DEFAULT '' COMMENT '关注id',
                               `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                               `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户关注表';