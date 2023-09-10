-- Active: 1694339097796@@127.0.0.1@3306@alta_online_shop
CREATE TABLE Product (
    ID_Product INT(11) PRIMARY KEY,
    Code VARCHAR(50),
    Name VARCHAR(100),
    Status SMALLINT,
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ID_Operator INT,
    ID_TP INT,
    ID_Descriptions INT
);

CREATE TABLE Operator (
    ID_Operator INT(11) PRIMARY KEY,
    Name VARCHAR(100),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Type_Product (
    ID_TP INT(11) PRIMARY KEY,
    Name VARCHAR(100),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Product_Descriptions (
    ID_Descriptions INT(11) PRIMARY KEY,
    Description TEXT,
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Transactions (
    ID_Transactions INT(11) PRIMARY KEY,
    Status VARCHAR(255),
    Total_QTY INT,
    Total_Price NUMERIC(25,2),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ID_PM INT,
    ID_Users INT
);

CREATE TABLE Users (
    ID_Users INT(11) PRIMARY KEY,
    Name VARCHAR(255),
    DOB DATE,
    Gender CHAR(1),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Transactions_Details (
    ID_TD INT(11) PRIMARY KEY,
    Status VARCHAR(255),
    QTY INT,
    Price NUMERIC(25,2),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ID_Transactions INT,
    ID_Product INT(11)
);

CREATE TABLE Payment_Method (
    ID_PM INT(11) PRIMARY KEY,
    Name VARCHAR(255),
    Status SMALLINT,
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Kurir (
    ID_Kurir INT(11) PRIMARY KEY,
    Name VARCHAR(255),
    Create_at TIMESTAMP,
    Update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

