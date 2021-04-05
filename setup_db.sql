-- create database and user for cidenet database --
CREATE DATABASE IF NOT EXISTS cidenet_db;
CREATE user IF NOT EXISTS 'user_db'@'localhost' identified BY 'pswd_db';
GRANT usage ON *.* TO 'user_db'@'localhost';
GRANT all privileges ON cidenet_db.* TO 'user_db'@'localhost';
