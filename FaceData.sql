CREATE DATABASE FaceData;
USE FaceData;

DROP TABLE IF EXISTS department;
DROP TABLE IF EXISTS picture;


CREATE TABLE department(
	departmentNO INT(4) PRIMARY KEY AUTO_INCREMENT,
	departmentName VARCHAR(64) DEFAULT NULL
)AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE person(
	UID INT(8) PRIMARY KEY AUTO_INCREMENT,
    havePicture BOOL DEFAULT FALSE,
	picture MEDIUMBLOB DEFAULT NULL,
	departmentNO INT(2) DEFAULT NULL,
	studentNumber INT(16) UNIQUE ,
	itsName VARCHAR(64) DEFAULT NULL

)AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT person SET departmentNO = 1,itsName = 'Zhang San',studentNumber = 333;
INSERT person SET departmentNO = 2,itsName = 'Li Si',studentNumber = 444;

INSERT department SET departmentName = 'group 1';
INSERT department SET departmentName = 'group 2';
INSERT department SET departmentName = 'group 3';
INSERT department SET departmentName = 'group 4';