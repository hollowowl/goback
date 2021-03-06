CREATE SCHEMA IF NOT EXISTS `data` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE IF NOT EXISTS `data`.`board`
(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `ip` VARCHAR(45) NULL,
    UNIQUE INDEX `name__unq` (`name` ASC)
);

CREATE TABLE IF NOT EXISTS `data`.`received_data`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `board_id` INT UNSIGNED NOT NULL,
    `data` LONGTEXT NOT NULL,
    INDEX `created__idx` (`created` ASC)
);

CREATE TABLE IF NOT EXISTS `data`.`decision`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `board_id` INT UNSIGNED NOT NULL,
    `received_data_id` BIGINT UNSIGNED NULL,
    `decision` LONGTEXT NOT NULL,
    INDEX `created__idx` (`created` ASC),
    INDEX `board_id__idx` (`created` ASC)
);
