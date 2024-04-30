CREATE TABLE IF NOT EXISTS `likes` (
  `user_id` INT UNSIGNED NOT NULL,
  `post_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (`user_id`, `post_id`),
  CONSTRAINT `fk_like_user_id` FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  CONSTRAINT `fk_like_post_id` FOREIGN KEY (`post_id`) REFERENCES posts(`id`) ON DELETE CASCADE
);
