CREATE TABLE public.products (
	id bigserial NOT NULL,
	category_id bigint NOT NULL,
	"name" varchar(255) NOT NULL,
	description text NOT NULL,
	photo varchar(255) NULL,
	price bigint NOT NULL DEFAULT 0,
	CONSTRAINT products_pk PRIMARY KEY (id),
	CONSTRAINT products_fk FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE INDEX products_category_id_idx ON public.products (category_id);
