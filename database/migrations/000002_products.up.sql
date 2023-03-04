BEGIN;

-- Products table
CREATE TABLE IF NOT EXISTS "products" (
  "id"          UUID          NOT NULL DEFAULT (uuid_generate_v4()),
  "title"       VARCHAR       NOT NULL,
  "price"       BIGINT        NOT NULL,
  "image"       VARCHAR,
  "created_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "products_pkey" PRIMARY KEY ("id")
);

-- Categories table
CREATE TABLE IF NOT EXISTS "categories" (
  "id"        UUID    NOT NULL DEFAULT (uuid_generate_v4()),
  "title"     VARCHAR NOT NULL,
  "parent_id" UUID    DEFAULT NULL,

  CONSTRAINT "categories_pkey" PRIMARY KEY ("id"),

  CONSTRAINT "fk_parent_id"
    FOREIGN KEY (parent_id)
      REFERENCES categories (id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

-- Categories-table junction table
CREATE TABLE IF NOT EXISTS "categories_products" (
  "product_id"  UUID  NOT NULL,
  "category_id" UUID  NOT NULL,

  CONSTRAINT "categories_products_pkey" PRIMARY KEY (product_id, category_id),

  CONSTRAINT "fk_product_id" FOREIGN KEY (product_id)
    REFERENCES products (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,

  CONSTRAINT "fk_category_id" FOREIGN KEY (category_id)
    REFERENCES categories (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

COMMIT;