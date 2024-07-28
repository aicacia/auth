package repository

import "time"

type PassKeysRowST struct {
	Id              []byte    `db:"id"`
	UserId          int32     `db:"user_id"`
	AplicationId    int32     `db:"application_id"`
	PublicKey       []byte    `db:"public_key"`
	AttestationType string    `db:"attestation_type"`
	Transports      []string  `db:"transports"`
	UserPresent     bool      `db:"user_present"`
	UserVerified    bool      `db:"user_verified"`
	BackupEligible  bool      `db:"backup_eligible"`
	BackupState     bool      `db:"backup_state"`
	AAGUID          []byte    `db:"aaguid"`
	SignCount       int32     `db:"sign_count"`
	CloneWarning    bool      `db:"clone_warning"`
	Attachment      string    `db:"attachment"`
	UpdatedAt       time.Time `db:"updated_at"`
	CreatedAt       time.Time `db:"created_at"`
}

func GetUserPassKeys(userId int32) ([]PassKeysRowST, error) {
	return All[PassKeysRowST](`SELECT pk.* 
		FROM passkeys pk 
		WHERE pk.user_id = $1;`, userId)
}

type UpsertPassKeyST struct {
	Id              []byte   `db:"id"`
	UserId          int32    `db:"user_id"`
	AplicationId    int32    `db:"application_id"`
	PublicKey       []byte   `db:"public_key"`
	AttestationType string   `db:"attestation_type"`
	Transports      []string `db:"transports"`
	UserPresent     bool     `db:"user_present"`
	UserVerified    bool     `db:"user_verified"`
	BackupEligible  bool     `db:"backup_eligible"`
	BackupState     bool     `db:"backup_state"`
	AAGUID          []byte   `db:"aaguid"`
	SignCount       int32    `db:"sign_count"`
	CloneWarning    bool     `db:"clone_warning"`
	Attachment      string   `db:"attachment"`
}

func UpsertUserPassKey(upsert UpsertPassKeyST) (PassKeysRowST, error) {
	return NamedGet[PassKeysRowST](`INSERT INTO passkeys (id, user_id, application_id, public_key, attestation_type, transports, user_present, user_verified, backup_eligible, backup_state, aaguid, sign_count, clone_warning, attachment)
			VALUES (:id, :user_id, :application_id, :public_key, :attestation_type, :transports, :user_present, :user_verified, :backup_eligible, :backup_state, :aaguid, :sign_count, :clone_warning, :attachment)
			ON CONFLICT (id) DO UPDATE SET
				user_id = :user_id,
				application_id = :application_id,
				public_key = :public_key,
				attestation_type = :attestation_type,
				transports = :transports,
				user_present = :user_present,
				user_verified = :user_verified,
				backup_eligible = :backup_eligible,
				backup_state = :backup_state,
				aaguid = :aaguid,
				sign_count = :sign_count,
				clone_warning = :clone_warning,
				attachment = :attachment,
				updated_at = NOW()
			RETURNING *;`, upsert)
}
