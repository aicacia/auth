DROP TABLE IF EXISTS "service_account_roles" cascade;
DROP TABLE IF EXISTS "service_accounts" cascade;

DROP TABLE IF EXISTS "user_mfas" cascade;
DROP TYPE IF EXISTS MFA_TYPE cascade;

DROP TABLE IF EXISTS "user_roles" cascade;
DROP TABLE IF EXISTS "users" cascade;
DROP TABLE IF EXISTS "emails" cascade;
DROP TABLE IF EXISTS "phone_numbers" cascade;
DROP TABLE IF EXISTS "passkeys" cascade;
DROP TABLE IF EXISTS "totps" cascade;
DROP TABLE IF EXISTS "user_infos" cascade;

DROP TABLE IF EXISTS "role_resource_permissions" cascade;
DROP TABLE IF EXISTS "roles" cascade;
DROP TABLE IF EXISTS "permissions" cascade;
DROP TABLE IF EXISTS "resources" cascade;
DROP TABLE IF EXISTS "tenents" cascade;
DROP TABLE IF EXISTS "applications" cascade;

DROP FUNCTION IF EXISTS "configs_notify" cascade;

DROP TABLE IF EXISTS "configs" cascade;

DROP FUNCTION IF EXISTS "trigger_updated_at" cascade;
