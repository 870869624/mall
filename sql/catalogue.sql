create table if not exists catalogue(
    id int(11) not null auto_increment primary key,
    product_name varchar(15) not null,
    picture varchar(256) not null,
)