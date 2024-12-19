BEGIN;

ALTER TABLE users
DROP COLUMN reset_password_token,

ALTER TABLE users
DROP COLUMN verify_email_token,

ALTER TABLE users
DROP COLUMN is_verified;

COMMIT;