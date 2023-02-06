CREATE TABLE user (
id INT PRIMARY KEY AUTO_INCREMENT, 
name VARCHAR(50),
email VARCHAR(50) unique, 
password VARCHAR(500),
role VARCHAR(50),
created_at TIMESTAMP,
deleted_at TIMESTAMP);

CREATE TABLE account (
    id VARCHAR(50) PRIMARY KEY,
    balance INT CHECK(balance>=0) DEFAULT 0,
    uid INT NOT NULL REFERENCES id(user),
    created_at TIME,
    deleted_at TIME
    );

CREATE TABLE transaction(
    id VARCHAR(52) PRIMARY KEY,
    amount INT NOT NULL,
    aid_credit VARCHAR(52) REFERENCES id(account),
    aid_debit VARCHAR(52) REFERENCES id(account),
    type VARCHAR(20),
    transaction_at varchar(100) not null
    );    
