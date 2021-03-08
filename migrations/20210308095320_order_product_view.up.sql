create view order_product_view as
select P.*, O.id order_id, O.status_id, O.date_update as order_date_update, O.date_create order_date_create
from "order"  O   LEFT JOIN order_product OP ON O.id = OP.order_id LEFT JOIN product P ON OP.product_id = P.id;