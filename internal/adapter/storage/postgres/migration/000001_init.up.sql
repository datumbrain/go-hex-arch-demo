BEGIN;

CREATE TABLE "users" (
    "id"            SERIAL          PRIMARY KEY,
    "name"          VARCHAR(100)            NOT NULL,
    "email"         VARCHAR(100)    UNIQUE  NOT NULL,
    "password"      VARCHAR(255)            NOT NULL,
    "is_admin"      BOOLEAN                             DEFAULT FALSE,
    "created_at"    TIMESTAMP               NOT NULL    DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "updated_at"    TIMESTAMP               NOT NULL    DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "deleted_at"    TIMESTAMP
);

CREATE TABLE "products" (
    "id"            SERIAL PRIMARY KEY,
    "name"          VARCHAR(100)    NOT NULL,
    "price"         DECIMAL(10, 2)  NOT NULL,
    "quantity"      INT             NOT NULL,
    "created_at"    TIMESTAMP       NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "updated_at"    TIMESTAMP       NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "deleted_at"    TIMESTAMP
);

CREATE TABLE "orders" (
    "id"            SERIAL PRIMARY KEY,
    "user_id"       INT             NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
    "product_id"    INT             NOT NULL REFERENCES "products" (id) ON DELETE CASCADE,
    "purchase_date" TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    "quantity"      INT             NOT NULL,
    "total_price"   DECIMAL(10, 2)  NOT NULL,
    "created_at"    TIMESTAMP       NOT NULL    DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "updated_at"    TIMESTAMP       NOT NULL    DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    "deleted_at"    TIMESTAMP
);

COMMIT;
