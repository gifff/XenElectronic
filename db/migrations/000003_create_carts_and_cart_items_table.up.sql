CREATE TABLE public.carts (
	cart_id uuid NOT NULL,
	CONSTRAINT carts_pk PRIMARY KEY (cart_id)
);

CREATE TABLE public.cart_items (
	id bigserial NOT NULL,
	cart_id uuid NOT NULL,
	product_id bigint NOT NULL,
	CONSTRAINT cart_items_pk PRIMARY KEY (id),
	CONSTRAINT cart_items_cart_id_fk FOREIGN KEY (cart_id) REFERENCES public.carts(cart_id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT cart_items_product_id_fk FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE INDEX cart_items_cart_id_idx ON public.cart_items (cart_id,product_id);
