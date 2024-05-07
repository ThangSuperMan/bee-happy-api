# Bee happy API

## Table of contents

- Provisions services

## Provisions services

- Provision mysql server with docker

```bash
docker run --name mysql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=<mysql_your_custom_password> \
-d mysql:8.3.0
```

## Types of notification
- Someone follow me
- Someone liked my post
- Someone mentiond me on a post


- Followings notifications
```sql
CREATE TABLE IF NOT EXISTS notifications (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` enum('0', '1', '2') NOT NULL DEFAULT '0' COMMENT '0: followed-me, 1: liked-post, 2: mentioned-me',
  `sender_id` INT UNSIGNED NOT NULL,
  `receiver_id` INT UNSIGNED NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `options` VARCHAR(255) DEFAULT '',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  CONSTRAINT `fk_sender_id` FOREIGN KEY (sender_id) REFERENCES users(id),
  CONSTRAINT `receiver_id` FOREIGN KEY (receiver_id) REFERENCES users(id),
);
```

- Followings table

```sql
CREATE TABLE IF NOT EXISTS followings (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `following_user_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (id),
  CONSTRAINT `fk_user_id` FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT `fk_following_user_id` FOREIGN KEY (following_user_id) REFERENCES users(id),
  CONSTRAINT `unique_follow` UNIQUE (user_id, following_user_id)
);
```

MASTER_HOST='10.89.0.6',

```bash
CHANGE MASTER TO 
MASTER_PORT=3306,
MASTER_HOST='mysql8.3.0-master',
MASTER_USER='root',
MASTER_PASSWORD='ZQDpWn2zhj8SxD',
MASTER_LOG_FILE='mysql-bin.000002',
MASTER_LOG_POS=158,
MASTER_CONNECT_RETRY=60,
GET_MASTER_PUBLIC_KEY=1;
```
