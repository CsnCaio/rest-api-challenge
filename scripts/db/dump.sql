-- CreateTable
create table if not exists users(
  "id" serial primary key not null,
  "name" text not null,
  "email" text not null
);

-- Seed
insert into users(name, email)
	values ('Gopher', 'hello@gopher.com'),
		   ('Halie', 'Lura33@yahoo.com'),
		   ('Laurine', 'laurine@gmail.com');
