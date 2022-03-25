create table if not exists users(
    id int(11) not null auto_increment primary key,
    name varchar(11) not null,
    gender int(1)not null,
    age int(3)not null,
    headpicture varchar(20),
    username varchar(15) not null,
    password int(15) not null
)
