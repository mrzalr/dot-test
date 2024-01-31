CREATE DATABASE `dot`;

USE `dot`;

CREATE TABLE IF NOT EXISTS `posts` (
  `id` uuid NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `like_count` bigint(20) DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=LATIN1_SWEDISH_CI;

CREATE TABLE IF NOT EXISTS `likes` (
  `user_id` longtext DEFAULT NULL,
  `post_id` uuid DEFAULT NULL,
  KEY `fk_likes_post` (`post_id`),
  CONSTRAINT `fk_likes_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=LATIN1_SWEDISH_CI;

DELIMITER //
CREATE PROCEDURE IF NOT EXISTS InsertRandomData()
BEGIN
    DECLARE counter INT DEFAULT 0;

    WHILE counter < 100 DO
        SET @uuid = UUID();
        SET @created_at = NOW() - INTERVAL FLOOR(RAND() * 365) DAY; -- Random date within the last year
        SET @updated_at = NOW() - INTERVAL FLOOR(RAND() * 365) DAY; -- Random date within the last year
        SET @title = CONCAT('Title ', counter); -- Unique title for each row
        SET @content = CONCAT('Content ', counter); -- Unique content for each row
        SET @like_count = FLOOR(RAND() * 401) + 100; -- Random value between 100 and 500
        SET @is_published = RAND() > 0.5; -- Random boolean value

        INSERT INTO posts (id, created_at, updated_at, title, content, like_count, is_published)
        VALUES (@uuid, @created_at, @updated_at, @title, @content, @like_count, @is_published);

        SET counter = counter + 1;
    END WHILE;
END //
DELIMITER ;

CALL InsertRandomData();