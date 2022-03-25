create table if not exists catalogue(
    id int(11) not null auto_increment primary key,
    name varchar(15) not null,
    number int(15) not null
)