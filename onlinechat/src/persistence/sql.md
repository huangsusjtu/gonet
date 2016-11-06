// mysqld -install
// net start mysql


mysql -h localhost -u root -p

// create user
mysql> insert into mysql.user(Host,User,Password) values("localhost","test",password("1234"));

update user set password=PASSWORD('12345678') where user="root"; 

// 
SHOW DATABASES;

//
CREATE DATABASE GolangOnlineChat;

USE databasename

CREATE TABLE IF NOT EXISTS test (id INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY, name VARCHAR(64), email VARCHAR(64), phone VARCHAR(32), description VARCHAR(1024), password VARCHAR(32));


CREATE TABLE IF NOT EXISTS user
(
name varchar(64),
email varchar(64),
phone varchar(16),
description varchar(16)
);