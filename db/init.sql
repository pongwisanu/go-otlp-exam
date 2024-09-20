CREATE SCHEMA catalog;

CREATE TABLE catalog.product
(
    id          BIGINT PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(250) NOT NULL
);

INSERT INTO catalog.product(id, name, description)
VALUES  (1, 'first', 'just first product'),
        (2, 'second', 'another second product'),
        (3, 'third', 'finally third product');

CREATE SCHEMA pricing;

CREATE TABLE pricing.price
(
    id BIGINT PRIMARY KEY,
    value FLOAT NOT NULL
);

INSERT INTO pricing.price(id, value)
VALUES  (1, 0.49),
        (2, 1.27),
        (3, 2.21);