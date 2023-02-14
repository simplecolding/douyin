CREATE TABLE `user_auth` (
                               `uid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                               `user_name` varchar(20) not null unique COMMENT '用户名',
                               `password` varchar(20) not null COMMENT '密码',
                               `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                               `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户鉴权表';



CREATE TABLE `video` (
                               `vid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                               `uid` bigint not null COMMENT '作者',
                               `play_url` varchar(255) not null COMMENT '播放地址',
                               `cover_url` varchar(255) not null COMMENT '封面地址',
                               `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                               `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`vid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE `comment` (
                         `cid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                         `vid` bigint not null COMMENT '视频id',
                         `uid` bigint not null COMMENT '用户id',
                         `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                         `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE `love` (
                           `lid` bigint unsigned AUTO_INCREMENT COMMENT '主键',
                           `vid` bigint not null COMMENT '视频id',
                           `uid` bigint not null COMMENT '用户id',
                           `status` tinyint(1) DEFAULT '1' COMMENT '是否删除 1:是  0:否',
                           `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`lid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='喜欢表';