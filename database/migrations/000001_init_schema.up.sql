CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE IF NOT EXISTS "users" (
  "id"          UUID          NOT NULL DEFAULT (uuid_generate_v4()),
  "first_name"  VARCHAR       NOT NULL,
  "last_name"   VARCHAR       NOT NULL,
  "email"       VARCHAR       NOT NULL,
  "photo"       VARCHAR       NOT NULL,
  "verified"    BOOLEAN       NOT NULL,
  "password"    VARCHAR       NOT NULL,
  "role"        VARCHAR       NOT NULL,
  "created_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- Unique index for email
CREATE UNIQUE INDEX IF NOT EXISTS unique_idx_users_email ON "users"("email");

-- Merchants table
CREATE TABLE IF NOT EXISTS "merchants" (
  "id"          UUID          NOT NULL DEFAULT (uuid_generate_v4()),
  "name"        VARCHAR       NOT NULL,
  "created_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT "merchants_pkey" PRIMARY KEY ("id")
);

-- Table to link merchants with users
CREATE TABLE IF NOT EXISTS "merchants_users" (
  "user_id"     UUID NOT NULL,
  "merchant_id" UUID NOT NULL
);

-- Unique pair of ids
CREATE UNIQUE INDEX IF NOT EXISTS unique_idx_user_id_merchant_id ON merchants_users(user_id, merchant_id);
