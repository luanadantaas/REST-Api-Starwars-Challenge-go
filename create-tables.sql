create database if not exists planet;
use planet;
drop table if exists planet;
create table planet(
    ID      int auto_increment not null,
    name    varchar(128) not null,
    climate varchar(255) not null,
    terrain    varchar(128) not null,
    appearances int not null,
    primary key (`ID`)
);


