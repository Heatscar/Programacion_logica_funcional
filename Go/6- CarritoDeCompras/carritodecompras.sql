CREATE DATABASE IF NOT EXISTS carritodecompras;

USE carritodecompras;

CREATE TABLE users(
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	lastName varchar(255) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE sales(
	id int(11) NOT NULL AUTO_INCREMENT,
	product int(11) NOT NULL,
    quantity int(255) NOT NULL,
    user_id int(11) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE product(
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	price int(11) NOT NULL,
    img_path varchar(255) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE warehouse(
	id int(11) NOT NULL AUTO_INCREMENT,
	stock varchar(255) NOT NULL,
	product_id int(11) NOT NULL,
	PRIMARY KEY(id)
);