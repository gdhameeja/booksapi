drop table if exists author;
create table if not exists author (
	id integer auto increment primary key not null,
	name varchar(20) not null
);

insert into author (id, name) values (1, 'J.K. Rowling');
insert into author (id, name) values (2, 'Chetan Bhagat');
insert into author (id, name) values (3, 'Gaurav Dhameeja');

drop table if exists book;
create table if not exists book (
	id integer auto increment primary key not null,
	name varchar(50) not null,
	author_id integer not null,
	status boolean default false,
	foreign key(author_id) references author(id)
);

insert into book (id, name, author_id, status) values (1, 'Harry Potter', (select id from author where name = 'Gaurav Dhameeja'), true);
insert into book (id, name, author_id, status) values (2, '3 idiots', (select id from author where name = 'J.K. Rowling'), true);
insert into book (id, name, author_id, status) values (3, 'No books written', (select id from author where name = 'Chetan Bhagat'), true);

drop table if exists user;
create table if not exists user (
	id integer auto increment primary key not null,
	username varchar(20) not null,
	role varchar(20),
	password varchar(10) not null,
	token text,
	check ( role like '%admin%' or role like '%regular%' )
);

insert into user (id, username, role, password, token) values (1, 'gauravdhameeja', 'admin', "password", "token1");
insert into user (id, username, role, password, token) values (2, 'User1', 'regular', "password", "token2");
