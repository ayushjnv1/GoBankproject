CREATE TABLE user (
id INT PRIMARY KEY AUTO_INCREMENT, 
name VARCHAR(50),
email VARCHAR(50) unique, 
password VARCHAR(500),
role VARCHAR(50),
created_at TIMESTAMP,
deleted_at TIMESTAMP);

CREATE TABLE customer (
    id VARCHAR(50) PRIMARY KEY,
    amount INT CHECK(amount>=0) DEFAULT 0,
    uid INT NOT NULL REFERENCES id(user),
    created_at TIME,
    deleted_at TIME
    );

CREATE TABLE transaction(
    id VARCHAR(52) PRIMARY KEY,
    amount INT NOT NULL,
    cid_credit VARCHAR(52) REFERENCES id(customer),
    cid_debit VARCHAR(52) REFERENCES id(customer),
    type VARCHAR(20),
    transaction_at varchar(100) not null
    );    
