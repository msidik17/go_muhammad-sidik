-- Active: 1694339097796@@127.0.0.1@3306@alat_online_shop
USE alta_online_shop;

INSERT INTO operator (id_operator, Name)
VALUES ("1","operator1"),("2","operator2"),("3","operator3"),("4","operator4"),("5","operator5");

INSERT INTO type_product (id_tp, Name)
VALUES ("1","Baju Kemeja"),("2","Celana Jens"),("3","Hoodie");

INSERT INTO product_descriptions (id_descriptions, Description)
VALUES ("1","Description1"), ("2","Description2"), ("3","Description3");

INSERT INTO Product (id_product, code, name, status, id_tp, id_operator, id_descriptions)
VALUES ("1","10001", "Baju Kemeja2", "20", "1", "3", "1"),
("2","10002", "Baju Kemeja1", "20", "1", "3", "1"),
("3","10003", "Celana Jens1", "15", "2", "1", "2"),
("4","10004", "Celana Jens2", "15", "2", "1", "2"),
("5","10005", "Celana Jens3", "15", "2", "1", "2"),
("6","10006", "Hoodie1", "25", "3", "4", "3"),
("7","10007", "Hoodie2", "25", "3", "4", "3"),
("8","10008", "Hoodie3", "25", "3", "4", "3");

INSERT INTO payment_method (id_PM, Name, status)
VALUES ("1","Transfer BANK", "1"),("2","Gopay","2"),("3","Dana","3");

INSERT INTO users (id_users, name, DOB, Gender)
VALUES ("1","Sidik","2001-06-01","Male"),
("2","Fajar","2001-06-02","Male"),
("3","Haikal","2001-06-03","Male"),
("4","Laili","2001-06-04","Female"),
("5","Nada","2001-06-05","Female");

INSERT INTO transactions (id_transactions, Status, Total_QTY, Total_Price, ID_PM, ID_Users)
VALUES ("1","DONE", "3", "10,999", "1", "1"),
("2","DONE", "3", "10,999", "2", "1"),
("3","DONE", "3", "10,999", "3", "1"),
("4","DONE", "3", "10,999", "1", "2"),
("5","DONE", "3", "10,999", "2", "2"),
("6","DONE", "3", "10,999", "3", "2"),
("7","DONE", "3", "10,999", "1", "3"),
("8","DONE", "3", "10,999", "2", "3"),
("9","DONE", "3", "10,999", "3", "3"),
("10","DONE", "3", "10,999", "1", "4"),
("11","DONE", "3", "10,999", "2", "4"),
("12","DONE", "3", "10,999", "3", "4"),
("13","DONE", "3", "10,999", "1", "5"),
("14","DONE", "3", "10,999", "2", "5"),
("15","DONE", "3", "10,999", "3", "5");

INSERT INTO Transactions_details (ID_TD, status, QTY, ID_Transactions, ID_Product)
VALUES ("1","AVAILABLE", 20, 1, 1),("2","PRE-ORDER", 20, 1, 2),("3","AVAILABLE", 20, 1, 3),
("4","AVAILABLE", 20, 2, 1),("5","PRE-ORDER", 20, 2, 2),("6","AVAILABLE", 20, 2, 3),
("7","AVAILABLE", 20, 3, 1),("8","PRE-ORDER", 20, 3, 2),("9","AVAILABLE", 20, 3, 3),
("10","AVAILABLE", 20, 4, 1),("11","PRE-ORDER", 20, 4, 2),("12","AVAILABLE", 20, 4, 3),
("13","AVAILABLE", 20, 5, 1),("14","PRE-ORDER", 20, 5, 2),("15","AVAILABLE", 20, 5, 3),
("16","AVAILABLE", 20, 6, 1),("17","PRE-ORDER", 20, 6, 2),("18","AVAILABLE", 20, 6, 3),
("19","AVAILABLE", 20, 7, 1),("20","PRE-ORDER", 20, 7, 2),("21","AVAILABLE", 20, 7, 3),
("22","AVAILABLE", 20, 8, 1),("23","PRE-ORDER", 20, 8, 2),("24","AVAILABLE", 20, 8, 3),
("25","AVAILABLE", 20, 9, 1),("26","PRE-ORDER", 20, 9, 2),("27","AVAILABLE", 20, 9, 3),
("28","AVAILABLE", 20, 10, 1),("29","PRE-ORDER", 20, 10, 2),("30","AVAILABLE", 20, 10, 3),
("31","AVAILABLE", 20,11, 1),("32","PRE-ORDER", 20, 11, 2),("33","AVAILABLE", 20, 11, 3),
("34","AVAILABLE", 20, 12, 1),("35","PRE-ORDER", 20, 12, 2),("36","AVAILABLE", 20, 12, 3),
("37","AVAILABLE", 20, 13, 1),("38","PRE-ORDER", 20, 13, 2),("39","AVAILABLE", 20, 13, 3),
("40","AVAILABLE", 20, 14, 1),("41","PRE-ORDER", 20, 14, 2),("42","AVAILABLE", 20, 14, 3),
("43","AVAILABLE", 20, 15, 1),("44","PRE-ORDER", 20, 15, 2),("45","AVAILABLE", 20, 15, 3);

SELECT * FROM users WHERE gender = 'M';

SELECT * FROM product WHERE id_Product = 3;

SELECT * FROM users create_at WHERE create_at <= CURRENT_TIMESTAMP AND create_at >= (CURRENT_TIMESTAMP - INTERVAL 7 DAY) AND name LIKE '%a%';

SELECT COUNT(id_users) FROM users WHERE gender = 'F';

SElECT * FROM users ORDER BY name ASC;

SELECT * FROM product LIMIT 5;

UPDATE product
    SET name = 'product dummy'
    WHERE id_Product = 1;

SELECT * FROM product where id_Product = 1;

UPDATE transactions_details
SET QTY = 3
WHERE id_product = 1;

SELECT * FROM transactions_details WHERE id_product = 3;

SELECT * FROM product;

DELETE FROM Product WHERE id_product = 1;

SELECT * FROM product;




















select *FROM transactions;