CREATE EXTENSION IF NOT EXISTS "pgcrypto";


CREATE FUNCTION "trigger_updated_at"()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = CURRENT_TIMESTAMP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE "configs" (
	"key" VARCHAR(255) NOT NULL PRIMARY KEY,
	"value" JSONB NOT NULL DEFAULT 'null',
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER "configs_updated_at_tgr" BEFORE UPDATE ON "configs" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "configs" ("key", "value") VALUES
	('host', '"0.0.0.0"'),
	('port', '3000'),
	('url', '"http://localhost:3000"'),
	('dashboard.enabled', 'true'),
	('openapi.enabled', 'true');


CREATE FUNCTION configs_notify() RETURNS trigger AS $$
DECLARE
  "key" VARCHAR(255);
  "value" JSONB;
BEGIN
	IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
		"key" = NEW."key";
	ELSE
		"key" = OLD."key";
	END IF;
	IF TG_OP != 'UPDATE' OR NEW."value" != OLD."value" THEN
		PERFORM pg_notify('configs_channel', json_build_object('table', TG_TABLE_NAME, 'key', "key", 'value', NEW."value", 'action_type', TG_OP)::text);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER "configs_notify_update" AFTER UPDATE ON "configs" FOR EACH ROW EXECUTE PROCEDURE configs_notify();
CREATE TRIGGER "configs_notify_insert" AFTER INSERT ON "configs" FOR EACH ROW EXECUTE PROCEDURE configs_notify();
CREATE TRIGGER "configs_notify_delete" AFTER DELETE ON "configs" FOR EACH ROW EXECUTE PROCEDURE configs_notify();


CREATE TABLE "applications" (
	"id" SERIAL PRIMARY KEY,
	"description" VARCHAR(255) NOT NULL,
	"uri" VARCHAR(255) NOT NULL,
	"website" VARCHAR(255),
	"is_admin" BOOLEAN NOT NULL DEFAULT false,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX "applications_uri_unique_idx" ON "applications" ("uri");
CREATE TRIGGER "applications_updated_at_tgr" BEFORE UPDATE ON "applications" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "applications" ("description", "uri", "is_admin", "website") VALUES
  ('Admin', 'admin', true, 'http://localhost:5173');


CREATE TABLE "tenents" (
	"id" SERIAL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"description" VARCHAR(255) NOT NULL,
	"uri" VARCHAR(255) NOT NULL,
	"authorization_website" VARCHAR(255) NOT NULL,
	"registration_website" VARCHAR(255),
	"email_endpoint" VARCHAR(255),
	"phone_number_endpoint" VARCHAR(255),
	"client_id" UUID NOT NULL DEFAULT gen_random_uuid(),
	"client_secret" VARCHAR(255) NOT NULL DEFAULT encode(gen_random_bytes(32), 'hex'),
	"algorithm" VARCHAR(255) NOT NULL DEFAULT 'HMAC',
	"public_key" TEXT,
	"private_key" TEXT NOT NULL DEFAULT encode(public.gen_random_bytes(255), 'base64'),
	"expires_in_seconds" INT8 NOT NULL DEFAULT 86400,
	"refresh_expires_in_seconds" INT8 NOT NULL DEFAULT 604800,
	"password_reset_expires_in_seconds" INT8 NOT NULL DEFAULT 86400,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "tenents_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "tenents_application_id_uri_unique_idx" ON "tenents" ("application_id", "uri");
CREATE UNIQUE INDEX "tenents_client_id_unique_idx" ON "tenents" ("client_id");
CREATE TRIGGER "tenents_updated_at_tgr" BEFORE UPDATE ON "tenents" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "tenents" ("client_id", "application_id", "description", "uri", "authorization_website") 
	VALUES
	('cbf7bbef-5132-4b2c-8622-06e28359c291', (SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Admin', 'admin', 'http://localhost:5173/signin');


CREATE TABLE "resources"(
	"id" SERIAL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"description" VARCHAR(255) NOT NULL,
	"uri" VARCHAR(255) NOT NULL,
	"actions" VARCHAR(255) ARRAY NOT NULL DEFAULT ARRAY[]::VARCHAR[],
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "resources_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "resources_application_id_uri_unique_idx" ON "resources" ("application_id", "uri");
CREATE TRIGGER "resources_updated_at_tgr" BEFORE UPDATE ON "resources" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "resources" ("application_id", "description", "uri", "actions")
  	VALUES
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Applications', 'applications', ARRAY['read', 'write']),
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Users', 'users', ARRAY['read', 'write']),
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Tenents', 'tenents', ARRAY['read', 'write']),
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Roles', 'roles', ARRAY['read', 'write']);


CREATE TABLE "roles"(
	"id" SERIAL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"description" VARCHAR(255) NOT NULL,
	"uri" VARCHAR(255) NOT NULL,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "roles_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "roles_application_id_uri_unique_idx" ON "roles" ("application_id", "uri");
CREATE TRIGGER "roles_updated_at_tgr" BEFORE UPDATE ON "roles" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "roles" ("application_id", "description", "uri")
  	VALUES
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'Admin', 'admin');


CREATE TABLE "role_resource_permissions"(
	"role_id" INT4 NOT NULL,
	"resource_id" INT4 NOT NULL,
	"actions" VARCHAR(255) ARRAY NOT NULL DEFAULT ARRAY[]::VARCHAR[],
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "role_resource_permissions_role_id_fk" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE CASCADE,
  	CONSTRAINT "role_resource_permissions_resource_id_fk" FOREIGN KEY("resource_id") REFERENCES "resources"("id") ON DELETE CASCADE,
	PRIMARY KEY("role_id", "resource_id")
);
CREATE TRIGGER "role_resource_permissions_updated_at_tgr" BEFORE UPDATE ON "role_resource_permissions" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "role_resource_permissions" ("role_id", "resource_id", "actions")
  	VALUES
	((SELECT id FROM "roles" WHERE uri='admin' LIMIT 1), (SELECT id FROM "resources" WHERE uri='applications' LIMIT 1), ARRAY['read', 'write']),
	((SELECT id FROM "roles" WHERE uri='admin' LIMIT 1), (SELECT id FROM "resources" WHERE uri='users' LIMIT 1), ARRAY['read', 'write']),
	((SELECT id FROM "roles" WHERE uri='admin' LIMIT 1), (SELECT id FROM "resources" WHERE uri='tenents' LIMIT 1), ARRAY['read', 'write']),
	((SELECT id FROM "roles" WHERE uri='admin' LIMIT 1), (SELECT id FROM "resources" WHERE uri='roles' LIMIT 1), ARRAY['read', 'write']);


CREATE TABLE "users"(
	"id" SERIAL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"username" VARCHAR(255) NOT NULL,
	"encrypted_password" VARCHAR(255) NOT NULL,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "users_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "users_username_unique_idx" ON "users" ("application_id", "username");
CREATE TRIGGER "users_updated_at_tgr" BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "users" ("application_id", "username", "encrypted_password")
	VALUES 
	((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), 'admin', '$argon2id$v=19$m=131072,t=10,p=16$CDXUDLS7JEAjO15qH557xg$39+nhOPQnsrghBc2yMRl7Qd4ifR+8AdQ4XdnYrTnhIw');


CREATE TABLE "user_infos"(
	"user_id" INT4 NOT NULL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"name" VARCHAR(255),
	"given_name" VARCHAR(255),
	"family_name" VARCHAR(255),
	"middle_name" VARCHAR(255),
	"nickname" VARCHAR(255),
	"profile" VARCHAR(255),
	"picture" VARCHAR(255),
	"website" VARCHAR(255),
	"gender" VARCHAR(255),
	"birthdate" TIMESTAMPTZ,
	"zoneinfo" VARCHAR(255),
	"locale" VARCHAR(255),
	"street_address" VARCHAR(255),
	"locality" VARCHAR(255),
	"region" VARCHAR(255),
	"postal_code" VARCHAR(255),
	"country" VARCHAR(255),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "user_infos_user_id_fk" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
  	CONSTRAINT "user_infos_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE TRIGGER "user_infos_updated_at_tgr" BEFORE UPDATE ON "user_infos" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

INSERT INTO "user_infos" ("application_id", "user_id")
	VALUES ((SELECT id FROM "applications" WHERE uri='admin' LIMIT 1), (SELECT id FROM "users" WHERE username='admin' LIMIT 1));


CREATE TABLE "emails"(
	"id" SERIAL PRIMARY KEY,
	"user_id" INT4 NOT NULL,
	"application_id" INT4 NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"confirmed" boolean NOT NULL DEFAULT false,
	"confirmation_token" VARCHAR(255),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "emails_user_id_fk" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
  	CONSTRAINT "emails_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "emails_email_unique_idx" ON "emails" ("application_id", "email");
CREATE TRIGGER "emails_updated_at_tgr" BEFORE UPDATE ON "emails" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

ALTER TABLE "users" ADD COLUMN "email_id" INT4;
ALTER TABLE "users" ADD CONSTRAINT "users_email_id_fk" FOREIGN KEY("email_id") REFERENCES "emails"("id") ON DELETE CASCADE;
CREATE UNIQUE INDEX "users_email_id_unique_idx" ON "users" ("email_id");


CREATE TABLE "phone_numbers"(
	"id" SERIAL PRIMARY KEY,
	"user_id" INT4 NOT NULL,
	"application_id" INT4 NOT NULL,
	"phone_number" VARCHAR(255) NOT NULL,
	"confirmed" boolean NOT NULL DEFAULT false,
	"confirmation_token" VARCHAR(255),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "phone_numbers_user_id_fk" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
  	CONSTRAINT "phone_numbers_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "phone_numbers_email_unique_idx" ON "phone_numbers" ("application_id", "phone_number");
CREATE TRIGGER "phone_numbers_updated_at_tgr" BEFORE UPDATE ON "phone_numbers" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();

ALTER TABLE "users" ADD COLUMN "phone_number_id" INT4;
ALTER TABLE "users" ADD CONSTRAINT "users_phone_number_id_fk" FOREIGN KEY("phone_number_id") REFERENCES "phone_numbers"("id") ON DELETE CASCADE;
CREATE UNIQUE INDEX "users_phone_number_id_unique_idx" ON "users" ("phone_number_id");


CREATE TABLE "user_tenent_totps"(
	"user_id" INT4 NOT NULL,
	"tenent_id" INT4 NOT NULL,
	"secret" VARCHAR(255) NOT NULL DEFAULT encode(gen_random_bytes(32), 'hex'),
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT "user_tenent_totps_user_id_fk" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
	CONSTRAINT "user_tenent_totps_tenent_id_fk" FOREIGN KEY("tenent_id") REFERENCES "tenents"("id") ON DELETE CASCADE,
	PRIMARY KEY("user_id", "tenent_id")
);


CREATE TABLE "user_roles"(
	"user_id" INT4 NOT NULL,
	"role_id" INT4 NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT "user_roles_user_id_fk" FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
	CONSTRAINT "user_roles_role_id_fk" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE CASCADE,
	PRIMARY KEY("user_id", "role_id")
);

INSERT INTO "user_roles" ("user_id", "role_id")
  	VALUES
    ((SELECT id FROM "users" WHERE username='admin' LIMIT 1), (SELECT id FROM "roles" WHERE uri='admin' LIMIT 1));


CREATE TABLE "service_accounts"(
	"id" SERIAL PRIMARY KEY,
	"application_id" INT4 NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"key" UUID DEFAULT gen_random_uuid(),
	"encrypted_secret" VARCHAR(255) NOT NULL,
	"updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	CONSTRAINT "service_accounts_application_id_fk" FOREIGN KEY("application_id") REFERENCES "applications"("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX "service_accounts_name_unique_idx" ON "service_accounts" ("application_id", "key");
CREATE TRIGGER "service_accounts_updated_at_tgr" BEFORE UPDATE ON "service_accounts" FOR EACH ROW EXECUTE PROCEDURE "trigger_updated_at"();


CREATE TABLE "service_account_roles"(
	"service_account_id" INT4 NOT NULL,
	"role_id" INT4 NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT "service_account_roles_service_account_id_fk" FOREIGN KEY("service_account_id") REFERENCES "service_accounts"("id") ON DELETE CASCADE,
	CONSTRAINT "service_account_roles_role_id_fk" FOREIGN KEY("role_id") REFERENCES "roles"("id") ON DELETE CASCADE,
	PRIMARY KEY("service_account_id", "role_id")
);