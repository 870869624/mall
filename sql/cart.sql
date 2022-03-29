create table if not exists cart(
    id int(11) not null auto_increment primary key,
    user_id int(11) not null,
    catalogue_id int(11) not null,
    prodouct_id int(11) not null,
    priceINtotal int(11) not null,
    picture varchar(256) not null
)