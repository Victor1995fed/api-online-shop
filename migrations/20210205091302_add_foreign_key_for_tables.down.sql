ALTER TABLE product_tag
 DROP CONSTRAINT "fk_product_tag-product", DROP CONSTRAINT "fk_product_tag-tag";

ALTER TABLE "order"
 DROP CONSTRAINT "fk_order-order_status";

ALTER TABLE "order_customer"
 DROP CONSTRAINT "fk_order_customer-customer", DROP CONSTRAINT "fk_order_customer-order";

ALTER TABLE "order_data"
 DROP CONSTRAINT "fk_order_data-order";

ALTER TABLE "order_product"
 DROP CONSTRAINT "fk_order_product-product", DROP CONSTRAINT "fk_order_product-order";