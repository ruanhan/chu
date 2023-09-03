drop table if exists `users`;
CREATE TABLE `users` (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         username VARCHAR(50) NOT NULL UNIQUE ,
                         email VARCHAR(50) NOT NULL UNIQUE ,
                         password VARCHAR(32) NOT NULL,
                         salt VARCHAR(32) NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `users` ADD UNIQUE (`username`);
INSERT INTO `users` (`username`, `password`, `salt`, `email`) VALUES
                                                                  ('chenyi', MD5(CONCAT('123456', 'salt1')), 'salt1', 'chenyi@example.com'),
                                                                  ('lier', MD5(CONCAT('123456', 'salt2')), 'salt2', 'lier@example.com'),
                                                                  ('zhangsan', MD5(CONCAT('123456', 'salt3')), 'salt3', 'zhangsan@example.com'),
                                                                  ('lisi', MD5(CONCAT('123456', 'salt4')), 'salt4', 'lisi@example.com'),
                                                                  ('wangwu', MD5(CONCAT('123456', 'salt5')), 'salt5', 'wangwu@example.com');