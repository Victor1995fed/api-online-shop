ALTER TABLE product_tag 
  ADD CONSTRAINT "fk_product_tag-product"
      FOREIGN KEY(product_id) 
      REFERENCES product(id)
      ON UPDATE CASCADE ON DELETE CASCADE,
  ADD CONSTRAINT "fk_product_tag-tag"
      FOREIGN KEY(tag_id) 
      REFERENCES tag(id)
      ON UPDATE CASCADE ON DELETE CASCADE;


ALTER TABLE "order" 
  ADD CONSTRAINT "fk_order-order_status"
      FOREIGN KEY(status_id) 
      REFERENCES order_status(id)
      ON UPDATE CASCADE ON DELETE SET DEFAULT;

ALTER TABLE "order_customer" 
  ADD CONSTRAINT "fk_order_customer-customer"
      FOREIGN KEY(customer_id) 
      REFERENCES customer(id)
      ON UPDATE CASCADE ON DELETE CASCADE,
  ADD CONSTRAINT "fk_order_customer-order"
      FOREIGN KEY(order_id) 
      REFERENCES "order"(id)
      ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "order_data"
  ADD CONSTRAINT "fk_order_data-order"
      FOREIGN KEY(order_id)
      REFERENCES "order"(id)
      ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "order_product"   
  ADD CONSTRAINT "fk_order_product-product"
      FOREIGN KEY(product_id) 
      REFERENCES product(id)
      ON UPDATE CASCADE ON DELETE CASCADE,
  ADD CONSTRAINT "fk_order_product-order"
      FOREIGN KEY(order_id) 
      REFERENCES "order"(id)
      ON UPDATE CASCADE ON DELETE CASCADE;