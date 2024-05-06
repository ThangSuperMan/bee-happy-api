ALTER TABLE `users`
ADD COLUMN `avatar_url` VARCHAR(255) DEFAULT NULL AFTER `date_of_birth`;
