-- Active: 1694339097796@@127.0.0.1@3306@alta_online_shop

--1
SELECT * FROM transactions WHERE id_users = 1
UNION
SELECT * FROM transactions WHERE id_users = 2;

-- 2
SELECT SUM(total_price) FROM transactions WHERE id_users = 1;

-- 3
SELECT COUNT(ID_Transactions) FROM transactions_details WHERE id_product IN (SELECT id_product FROM product WHERE id_tp = 2);

-- 4
SELECT p.*, pt.name AS pt_name FROM product p INNER JOIN type_product pt ON p.id_tp = pt.id_tp;

-- 5
SELECT t.*, p.name AS p_name, u.name AS u_name FROM transactions t JOIN product p JOIN users u;

-- 6
DELIMITER //
CREATE TRIGGER after_delete_transactions_detail
AFTER DELETE ON transactions_details
FOR EACH ROW
BEGIN
    UPDATE transactions
    SET total_qyt = (
        SELECT SUM(QTY)
        FROM transactions_details
        WHERE id_transactions = OLD.ID_Transactions
    )
    WHERE id = OLD.ID_Transactions;
END;
//
DELIMITER ;

SHOW TRIGGERS;

-- 6 test case
SELECT COUNT(transaction_id) FROM transaction_details WHERE transaction_id = 1;
DELETE FROM transactions WHERE id = 1;
SELECT COUNT(transaction_id) FROM transaction_details WHERE transaction_id = 1;

-- 7
DELIMITER $$
CREATE TRIGGER update_transaction_total_qty
BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN

DECLARE v_transaction_id INT;
DECLARE new_total_qty INT;

SET v_transaction_id = OLD.transaction_id;
SET new_total_qty = (SELECT SUM(qty) FROM transaction_details WHERE transaction_id = v_transaction_id) - OLD.qty;

UPDATE transactions SET total_qty = new_total_qty WHERE id = v_transaction_id;

END$$
DELIMITER ;

SHOW TRIGGERS;

-- 7 test case
SELECT total_qty FROM transactions WHERE id = 2;
DELETE FROM transaction_details WHERE transaction_id = 2 AND product_id = 4;
SELECT total_qty FROM transactions WHERE id = 2;

-- 8
-- add data product baru yang belum ada di transactions
INSERT INTO products (
  id,
  product_type_id,
  operator_id,
  code,
  name,
  status
)
VALUES (9, 1, 3, 'RT3', 'Vacuum Cleaner', '200');

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details);

DELETE FROM products WHERE id = 9;