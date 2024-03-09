CREATE TABLE user (
    id              INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    public_id       VARCHAR(20)  NOT NULL,
    fullname        VARCHAR(255) NOT NULL,
    email           VARCHAR(100) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    phone           VARCHAR(20)  NOT NULL,
    gender          VARCHAR(10)  NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP
) ENGINE=InnoDB;
