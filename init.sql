drop
    database if exists catalogue;
create
    database catalogue;
use
    catalogue;

drop table if exists catalogue;
create table catalogue
(
    product_id INT UNSIGNED NOT NULL AUTO_INCREMENT primary key,
    price      DOUBLE       NOT NULL DEFAULT 0,
    stock      INT UNSIGNED NOT NULL DEFAULT 0,
    updated_at TIMESTAMP    NOT NULL default current_timestamp
)