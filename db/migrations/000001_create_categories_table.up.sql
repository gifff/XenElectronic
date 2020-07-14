CREATE TABLE public.categories (
	id bigserial NOT NULL,
	"name" varchar(255) NOT NULL,
	CONSTRAINT categories_pk PRIMARY KEY (id)
);
