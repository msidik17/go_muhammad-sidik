-- Active: 1694339097796@@127.0.0.1@3306@alta_online_shop
ALTER TABLE Product
ADD CONSTRAINT FK_Product_Operator
FOREIGN KEY (ID_Operator) REFERENCES Operator(ID_Operator);

ALTER TABLE Product
ADD CONSTRAINT FK_Product_Type_Product
FOREIGN KEY (ID_TP) REFERENCES Type_Product(ID_TP);

ALTER TABLE Product
ADD CONSTRAINT FK_Product_Product_Descriptions
FOREIGN KEY (ID_Discriptions) REFERENCES Product_Descriptions(ID_Discriptions);

ALTER TABLE Transactions
ADD CONSTRAINT FK_Transactions_Users
FOREIGN KEY (ID_Users) REFERENCES Users(ID_Users);

ALTER TABLE Transactions
ADD CONSTRAINT FK_Transactions_Payment_Method
FOREIGN KEY (ID_PM) REFERENCES Payment_Method(ID_PM);

ALTER TABLE Transactions_Details
ADD CONSTRAINT FK_Transactions_Details_Transactions
FOREIGN KEY (ID_Transactions) REFERENCES Transactions(ID_Transactions);

ALTER TABLE Transactions_Details
ADD CONSTRAINT FK_Transactions_Details_Product
FOREIGN KEY (ID_Product) REFERENCES Product(ID_Product);

ALTER TABLE Kurir RENAME TO Shipping;

DROP TABLE shipping;
