-- 用户表
create table `user` (
                        `id` int auto_increment primary key,
                        `openid` varchar(50) not null unique,
                        `nickname` varchar(100),
                        `avatar` varchar(255),
                        `phone` varchar(11),
                        `gender` tinyint,
                        `role` tinyint default 1,
                        `status` tinyint default 1,
                        `created_at` bigint,
                        `updated_at` bigint
);

-- 简历表
create table `resume` (
                          `id` int auto_increment primary key,
                          `user_id` int not null unique,
                          `content` text,
                          `file_path` varchar(255),
                          `created_at` bigint,
                          `updated_at` bigint,
                          foreign key (`user_id`) references `user`(`id`)
);

-- 招聘者营业执照表
create table `business_license` (
                                    `id` int auto_increment primary key,
                                    `user_id` int not null unique,
                                    `content` text,
                                    `file_path` varchar(255),
                                    `created_at` bigint,
                                    `updated_at` bigint,
                                    foreign key (`user_id`) references `user`(`id`)
);

-- 实名审核表
create table `real_name_verification` (
                                          `id` int auto_increment primary key,
                                          `user_id` int not null unique,
                                          `real_name` varchar(50),
                                          `id_card_number` varchar(18),
                                          `id_card_front_img` varchar(255),
                                          `id_card_back_img` varchar(255),
                                          `verified_status` tinyint default 0,
                                          `verified_time` bigint,
                                          `created_at` bigint,
                                          `updated_at` bigint,
                                          foreign key (`user_id`) references `user`(`id`)
);

-- 管理员表
create table `admin` (
                         `id` int auto_increment primary key,
                         `email` varchar(100) not null unique,
                         `password` varchar(255) not null,
                         `phone` varchar(20) unique,
                         `avatar` varchar(255),
                         `nickname` varchar(32),
                         `signature` varchar(128),
                         `created_at` bigint,
                         `updated_at` bigint
);

-- 招聘订单表
create table `recruit_order` (
                                 `id` int auto_increment primary key,
                                 `recruiter_id` int not null,
                                 `title` varchar(100) not null,
                                 `description` text,
                                 `status` tinyint(20) default 1,
                                 `service_fee` bigint not null,
                                 `created_at` bigint,
                                 `updated_at` bigint,
                                 foreign key (`recruiter_id`) references `user`(`id`)
);

-- 求职申请表
create table `job_application` (
                                   `id` int auto_increment primary key,
                                   `order_id` int not null,
                                   `job_seeker_id` int not null,
                                   `status` tinyint(20) default 1,
                                   `created_at` bigint,
                                   `updated_at` bigint,
                                   foreign key (`order_id`) references `recruit_order`(`id`),
                                   foreign key (`job_seeker_id`) references `user`(`id`)
);

-- 评价表
create table `review` (
                          `id` int auto_increment primary key,
                          `reviewer_id` int not null,
                          `reviewed_user_id` int not null,
                          `order_id` int not null,
                          `content` text,
                          `rating` tinyint not null,
                          `created_at` bigint,
                          `updated_at` bigint,
                          foreign key (`reviewer_id`) references `user`(`id`),
                          foreign key (`reviewed_user_id`) references `user`(`id`),
                          foreign key (`order_id`) references `recruit_order`(`id`)
);

-- 金额操作日志表
create table `transaction_log` (
                                   `id` int auto_increment primary key,
                                   `user_id` int not null,
                                   `order_id` int not null,
                                   `amount` double not null,
                                   `type` tinyint(20),
                                   `description` varchar(255),
                                   `created_at` bigint,
                                   `updated_at` bigint,
                                   foreign key (`user_id`) references `user`(`id`),
                                   foreign key (`order_id`) references `recruit_order`(`id`)
);