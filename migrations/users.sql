CREATE TABLE public.user (
  id varchar(40) primary key not null,
  name varchar(100) not null,
  email varchar(100) not null,
  password varchar(100) not null,
  is_active boolean default true,
  is_verified boolean default false,
  created_at timestamp with time zone not null default now()
); 

ALTER TABLE public.user ADD CONSTRAINT constraint_unique_email UNIQUE (email);