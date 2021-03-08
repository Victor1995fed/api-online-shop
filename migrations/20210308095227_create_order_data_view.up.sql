create view order_data_view as
select OD.*, O.status_id, O.date_update as order_date_update, O.date_create order_date_create
from "order" O LEFT JOIN order_data OD ON O.id = OD.order_id;