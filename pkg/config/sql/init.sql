CREATE TABLE `user`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `user_name` varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User delete time',
    PRIMARY KEY (`id`),
    KEY `idx_user_name` (`user_name`) COMMENT 'UserName idx'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT 'User table';

CREATE TABLE `note`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `user_id` int(64) NOT NULL DEFAULT 0 COMMENT 'UserID',
    `title` varchar(128) NOT NULL DEFAULT '' COMMENT 'Title',
    `content` TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Note create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Note update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Note delete time',
    PRIMARY KEY (`id`),
    KEY `idx_user_id_title` (`user_id`, `title`) COMMENT 'UserID Title index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT 'Note table';