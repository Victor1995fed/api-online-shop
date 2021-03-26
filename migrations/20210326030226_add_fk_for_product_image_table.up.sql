ALTER TABLE product_image
  ADD CONSTRAINT "fk_product_image-product"
      FOREIGN KEY(product_id)
      REFERENCES product(id)
      ON UPDATE CASCADE ON DELETE CASCADE,
  ADD CONSTRAINT "fk_product_image-image"
      FOREIGN KEY(image_id)
      REFERENCES image(id)
      ON UPDATE CASCADE ON DELETE CASCADE;