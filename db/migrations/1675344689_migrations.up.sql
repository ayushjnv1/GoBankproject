CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT, 
    name VARCHAR(50),
    email VARCHAR(50) unique, 
    password VARCHAR(500),
    role VARCHAR(50),
    created_at TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE TABLE account (
    id VARCHAR(50) PRIMARY KEY,
    balance INT CHECK(balance>=0) DEFAULT 0,
    user_id INT NOT NULL ,
    created_at TIME,
    deleted_at TIME,
    constraint uid_fk FOREIGN KEY(uid) references user(id) 
    );

CREATE TABLE transaction(
    id VARCHAR(52) PRIMARY KEY,
    amount INT NOT NULL,
    aid_credit VARCHAR(52),
    aid_debit VARCHAR(52) ,
    type VARCHAR(20),
    transaction_at varchar(100) not null,
    constraint aid_debit_fk FOREIGN KEY(aid_credit) references account(id) ,
    constraint aid_credit_fk FOREIGN KEY(uid) references account(id) 
    ); 
