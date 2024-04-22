CREATE TABLE IF NOT EXISTS users (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(255) not null,
  `last_name` VARCHAR(255) not null,
  `email` VARCHAR(255) not null,
  `password` VARCHAR(255) not null,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  unique key (email)
);
