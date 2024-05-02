DROP TABLE IF EXISTS "service_account_permissions" cascade;
DROP TABLE IF EXISTS "service_accounts" cascade;

DROP TABLE IF EXISTS "user_permissions" cascade;
DROP TABLE IF EXISTS "users" cascade;
DROP TABLE IF EXISTS "emails" cascade;
DROP TABLE IF EXISTS "phone_numbers" cascade;
DROP TABLE IF EXISTS "user_tenent_totps" cascade;
DROP TABLE IF EXISTS "user_infos" cascade;

DROP TABLE IF EXISTS "permissions" cascade;
DROP TABLE IF EXISTS "tenents" cascade;
DROP TABLE IF EXISTS "applications" cascade;

DROP FUNCTION IF EXISTS "configs_notify" cascade;

DROP TABLE IF EXISTS "configs" cascade;

DROP FUNCTION IF EXISTS "trigger_updated_at" cascade;
