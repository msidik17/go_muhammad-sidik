-- Active: 1694339097796@@127.0.0.1@3306@alta_online_shop
-- SQL --
CREATE DATABASE alta_online_shop;

USE alta_online_shop; 

CREATE TABLE Product (
    id_product INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    create_at TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_operator INT,
    id_tp INT,
    id_descriptions INT
);

CREATE TABLE Operator (
    ID_Operator INT UNIQUE AUTO_INCREMENT  PRIMARY KEY,
    Name VARCHAR(100),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Type_Product (
    ID_TP INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Product_Descriptions (
    ID_Descriptions INT AUTO_INCREMENT PRIMARY KEY,
    Description TEXT,
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Transactions (
    ID_Transactions INT AUTO_INCREMENT PRIMARY KEY,
    Status VARCHAR(255),
    Total_QTY INT,
    Total_Price NUMERIC(25,2),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ID_PM INT,
    ID_Users INT
);

CREATE TABLE Users (
    ID_Users INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    DOB DATE,
    Gender CHAR(1),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Transactions_Details (
    ID_TD INT AUTO_INCREMENT PRIMARY KEY,
    Status VARCHAR(255),
    QTY INT,
    Price NUMERIC(25,2),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ID_Transactions INT,
    ID_Product INT(11)
);

CREATE TABLE Payment_Method (
    ID_PM INT UNIQUE AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Status SMALLINT,
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Kurir (
    ID_Kurir INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

