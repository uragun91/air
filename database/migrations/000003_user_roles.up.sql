BEGIN;

CREATE TABLE IF NOT EXISTS "roles" (
  "id"    UUID    NOT NULL DEFAULT (uuid_generate_v4()),
  "name"  VARCHAR NOT NULL,

  CONSTRAINT "roles_pkey" PRIMARY KEY (id),
  CONSTRAINT unique_idx_name UNIQUE ("name")
);


CREATE TABLE IF NOT EXISTS "roles_users" (
  "role_id"     UUID  NOT NULL,
  "user_id"     UUID  NOT NULL,
  "created_at"  TIMESTAMP(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "active"      BOOLEAN NOT NULL DEFAULT true,

  CONSTRAINT "roles_users_pkey" PRIMARY KEY (role_id, user_id),

  CONSTRAINT "fk_role_id" FOREIGN KEY (role_id)
    REFERENCES roles (id),

  CONSTRAINT "fk_user_id" FOREIGN KEY (user_id)
    REFERENCES users (id)
);

CREATE UNIQUE INDEX unique_idx_roles_users_active ON roles_users (role_id, user_id, active)
  WHERE active;

COMMIT;