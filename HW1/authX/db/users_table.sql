-- Use this scripts to create a user table according to our standard!

-- BECAREFULL: RERUNNING THIS CODE WILL REMOVE YOUR PREVIOUS DATABASE

-- Creating database for Code Runner Web Service!
DROP DATABASE IF EXISTS Code_Runner;
CREATE DATABASE IF NOT EXISTS Code_Runner;
USE Code_Runner;

-- Just in case :)
DROP TABLE IF EXISTS User_Info;


CREATE TABLE IF NOT EXISTS User_Info
(
    EMAIL VARCHAR(512) NOT NULL UNIQUE,
    TOKEN VARCHAR(512) NOT NULL UNIQUE,
    PWD VARCHAR(512) NOT NULL,
    PRIMARY KEY (EMAIL)
);

-- This user is for authx admin!
CREATE USER 'authx_local'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON Code_Runner.User_Info TO 'authx_local'@'localhost';
