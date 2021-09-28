#IamAPerson

数据库的结构：
    这里facedata数据库名可自取

    CREATE DATABASE facedata;
    USE facedata;
    
    DROP TABLE IF EXISTS department;
    DROP TABLE IF EXISTS picture;
    
    
    CREATE TABLE department(
    departmentNO INT(4) PRIMARY KEY AUTO_INCREMENT,
    departmentName VARCHAR(64) DEFAULT NULL
    )AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8;
    
    CREATE TABLE picture(
    UID INT(8) PRIMARY KEY AUTO_INCREMENT,
    havePicture BOOL DEFAULT FALSE,
    picture MEDIUMBLOB DEFAULT NULL,
    departmentNO INT(2) DEFAULT NULL,
    itsName VARCHAR(64) DEFAULT NULL
    )AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8;