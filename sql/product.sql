create table if not exists prodouct(
    id int(11) not null auto_increment primary key,
    name varchar(15) not null,
    image varchar(20) not null,
    price int(11) not null,
    total int(2) not null,
    size varchar(4) not null
)