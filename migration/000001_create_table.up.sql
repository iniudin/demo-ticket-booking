CREATE TABLE concerts
(
    id        BYTEA PRIMARY KEY,
    name      VARCHAR(100) NOT NULL,
    location  VARCHAR(100) NOT NULL,
    date      TIMESTAMP    NOT NULL,
    remaining INTEGER      NOT NULL
);

CREATE TABLE orders
(
    id         BYTEA PRIMARY KEY,
    email      VARCHAR(100) NOT NULL,
    status     VARCHAR(50)  NOT NULL,
    date       TIMESTAMP    NOT NULL,
    concert_id BYTEA        NOT NULL REFERENCES concerts (id)
);

