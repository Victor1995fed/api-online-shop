create view product_tags_view as
select T.id tag_id, T.title tag_title, P.*
from "product" P
         LEFT JOIN  product_tag PT on P.id = pt.product_id
         LEFT JOIN tag T on PT.tag_id = T.id;