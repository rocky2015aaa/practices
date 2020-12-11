drop database if exists myBoard;
create database myBoard;
use myBoard;
-- TODO: check necessity of id for table
create table user (
	id int NOT NULL AUTO_INCREMENT,
	user_name varchar(30) NOT NULL,
	password varchar(50) NOT NULL,
	cookie_data varchar(255),
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at datetime ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY (user_name),
	UNIQUE KEY (password),
	UNIQUE KEY (cookie_data)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table category (
	id int NOT NULL AUTO_INCREMENT,
	name varchar(30) NOT NULL,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at datetime ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table sub_category (
	id int NOT NULL AUTO_INCREMENT,
	name varchar(30) NOT NULL,
	category_id int NOT NULL,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at datetime ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE KEY (name),
	KEY fk_category_id (category_id),
	CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES category(id)
	ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table post (
	id int NOT NULL AUTO_INCREMENT,
	title varchar(50) NOT NULL,
	body text NOT NULL,
	user_id int DEFAULT NULL, -- In the case of removing user, keep user's post
	sub_category_id int NOT NULL,
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at datetime ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	KEY fk_post_user_id (user_id),
	KEY fk_sub_category_id (sub_category_id),
	CONSTRAINT fk_post_user_id FOREIGN KEY (user_id) REFERENCES user(id) 
	ON DELETE SET NULL,
	CONSTRAINT fk_sub_category_id FOREIGN KEY (sub_category_id) REFERENCES sub_category(id) 
	ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table comment (
	id int NOT NULL AUTO_INCREMENT,
	body text NOT NULL,
	post_id int NOT NULL,
	user_id int DEFAULT NULL, -- In the case of removing user, keep user's comment
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at datetime ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	KEY fk_post_id (post_id),
	KEY fk_comment_user_id (user_id),
	CONSTRAINT fk_comment_user_id FOREIGN KEY (user_id) REFERENCES user(id) 
	ON DELETE SET NULL,
	CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES post(id) 
	ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
