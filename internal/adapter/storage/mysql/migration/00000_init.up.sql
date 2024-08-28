create table `users`
(
    `id`        bigint unsigned auto_increment primary key,
    `name`      varchar(20) not null,
    `cellphone` varchar(13) not null unique
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;