-- DROP TABLE IF EXISTS `auth_tokens`;
CREATE TABLE `auth_tokens` (
  `token_id` char(32) NOT NULL,
  `user_id` int(4) NOT NULL,
  `refresh_hash` tinytext NOT NULL,
  `issued_at` datetime NOT NULL,
  `remote_addr` tinytext NOT NULL,
  `device_info` text NOT NULL CHECK (json_valid(device_info)),
  PRIMARY KEY (`token_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `auth_tokens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DELIMITER ;;

CREATE TRIGGER `drop_user_auth_tokens` AFTER UPDATE ON `users` FOR EACH ROW
IF NEW.password_hash != OLD.password_hash OR NEW.new_password_hash != OLD.new_password_hash THEN
  DELETE FROM `auth_tokens` WHERE user_id = NEW.user_id;
END IF;;

DELIMITER ;
