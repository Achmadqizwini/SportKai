CREATE TABLE club (
    id              INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    public_id       VARCHAR(40)  NOT NULL,
    name            VARCHAR(100) NOT NULL,
    address         VARCHAR(100) NOT NULL,
    city            VARCHAR(50)  NOT NULL,
    description     VARCHAR(150),
    joined_member   INT          DEFAULT 0,
    member_total    INT          DEFAULT 0,
    rules           VARCHAR(255) NOT NULL,
    requirements   VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP   
) ENGINE=InnoDB;


CREATE TABLE club_member (
    id          INT             NOT NULL PRIMARY KEY AUTO_INCREMENT,
    public_id   VARCHAR(40)     NOT NULL,
    user_id     INT             NOT NULL,
    club_id     INT             NOT NULL,
    status      VARCHAR(20)     NOT NULL, 
    joined_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    left_at     TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (club_id) REFERENCES club (id)
)