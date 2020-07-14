CREATE TABLE public.orders (
	id uuid NOT NULL,
	customer_name varchar(255) NOT NULL,
	customer_email varchar(255) NOT NULL,
	customer_address varchar(255) NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);

CREATE TABLE public.order_items (
	id bigint NOT NULL,
	order_id uuid NOT NULL,
	product_id bigint NULL DEFAULT NULL,
	"name" varchar(255) NOT NULL,
	description text NOT NULL,
	photo varchar(255) NOT NULL,
	price bigint NOT NULL,
	CONSTRAINT order_items_pk PRIMARY KEY (id),
	CONSTRAINT order_items_order_id_fk FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE RESTRICT ON UPDATE CASCADE,
	CONSTRAINT order_items_product_id_fk FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE INDEX order_items_order_id_idx ON public.order_items (order_id);

