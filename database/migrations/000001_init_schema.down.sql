BEGIN;

DROP INDEX IF EXISTS unique_idx_users_email;
DROP INDEX IF EXISTS unique_idx_user_id_merchant_id;
DROP TABLE IF EXISTS merchants_users;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS merchants;

COMMIT;