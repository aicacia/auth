package service

import (
	"net/url"
	"time"

	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/aicacia/go-expiringmap"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	WebAuthnSessions = expiringmap.New[int32, *webauthn.SessionData]()
)

type WebAuthnUser struct {
	user     repository.UserRowST
	passkeys []repository.PassKeysRowST
}

func NewWebAuthnUser(user repository.UserRowST, passkeys []repository.PassKeysRowST) *WebAuthnUser {
	return &WebAuthnUser{
		user:     user,
		passkeys: passkeys,
	}
}

func (webAuthnUser *WebAuthnUser) WebAuthnID() []byte {
	return webAuthnUser.user.Key
}

func (webAuthnUser *WebAuthnUser) WebAuthnName() string {
	return webAuthnUser.user.Username
}

func (webAuthnUser *WebAuthnUser) WebAuthnDisplayName() string {
	return webAuthnUser.user.Username
}

func (webAuthnUser *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return util.Map(webAuthnUser.passkeys, WebAuthnCredentialsFromRow)
}

func (user *WebAuthnUser) WebAuthnIcon() string {
	return ""
}

func WebAuthnCredentialsFromRow(row repository.PassKeysRowST) webauthn.Credential {
	transports := make([]protocol.AuthenticatorTransport, 0, len(row.Transports))
	for _, transport := range row.Transports {
		transports = append(transports, protocol.AuthenticatorTransport(transport))
	}
	return webauthn.Credential{
		ID:              row.Id,
		PublicKey:       row.PublicKey,
		AttestationType: row.AttestationType,
		Transport:       transports,
		Flags: webauthn.CredentialFlags{
			UserPresent:    row.UserPresent,
			UserVerified:   row.UserVerified,
			BackupEligible: row.BackupEligible,
			BackupState:    row.BackupState,
		},
		Authenticator: webauthn.Authenticator{
			AAGUID:       row.AAGUID,
			SignCount:    uint32(row.SignCount),
			CloneWarning: row.CloneWarning,
			Attachment:   protocol.AuthenticatorAttachment(row.Attachment),
		},
	}
}

func WebAuthnCredentialsToUpsert(applicationId, userId int32, credential webauthn.Credential) repository.UpsertPassKeyST {
	transports := make([]string, 0, len(credential.Transport))
	for _, transport := range credential.Transport {
		transports = append(transports, string(transport))
	}
	return repository.UpsertPassKeyST{
		Id:              credential.ID,
		UserId:          userId,
		AplicationId:    applicationId,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
		Transports:      transports,
		UserPresent:     credential.Flags.UserPresent,
		UserVerified:    credential.Flags.UserVerified,
		BackupEligible:  credential.Flags.BackupEligible,
		BackupState:     credential.Flags.BackupState,
		AAGUID:          credential.Authenticator.AAGUID,
		SignCount:       int32(credential.Authenticator.SignCount),
		CloneWarning:    credential.Authenticator.CloneWarning,
		Attachment:      string(credential.Authenticator.Attachment),
	}
}

func WebAuthnFromTenent(tenent *repository.TenentRowST) (*webauthn.WebAuthn, error) {
	u, err := url.Parse(tenent.AuthorizationWebsite)
	if err != nil {
		return nil, err
	}
	origin := u.Scheme + "://" + u.Host
	return webauthn.New(&webauthn.Config{
		RPDisplayName:        tenent.Description,
		RPID:                 u.Hostname(),
		RPOrigins:            []string{origin},
		EncodeUserIDAsString: false,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: protocol.CrossPlatform,
			RequireResidentKey:      protocol.ResidentKeyNotRequired(),
			ResidentKey:             protocol.ResidentKeyRequirementDiscouraged,
			UserVerification:        protocol.VerificationRequired,
		},
		AttestationPreference: protocol.PreferNoAttestation,
		Timeouts: webauthn.TimeoutsConfig{
			Login: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Minute,
				TimeoutUVD: time.Minute,
			},
			Registration: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    time.Minute,
				TimeoutUVD: time.Minute,
			},
		},
	})
}
