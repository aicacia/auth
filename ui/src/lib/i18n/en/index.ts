import type { BaseTranslation } from '../i18n-types';

const en = {
	errors: {
		name: {
			internal: 'Application Error',
			request: 'Request',
			user: 'User',
			email: 'Email',
			username: 'Username',
			password: 'Password',
			passwordConfirmation: 'Password confirmation',
			phoneNumber: 'Phone Number',
			authorization: 'Authorization',
			mfa: 'Multi-factor Authorization'
		},
		message: {
			application: 'if it presists please contact support',
			tooManyRequests: 'Too many requests',
			notFound: 'Not found',
			invalid: 'Invalid',
			required: 'Required',
			noChange: 'No change',
			mismatch: 'Passwords do not match',
			alreayUsed: 'Already used',
			cannotDeleteAdmin: 'Can not delete Admin Application.',
			disabled: 'Disabled'
		}
	},
	auth: {
		signIn: 'Sign in',
		signUp: 'Sign up',
		notAMember: 'Not a member?',
		alreadyAMember: 'Already a member?',
		backToSignIn: 'Back to sign in',
		emailPlaceholder: 'Email',
		usernamePlaceholder: 'Username',
		usernameOrEmailPlaceholder: 'Username/Email',
		passwordPlaceholder: 'Password',
		passwordConfirmationPlaceholder: 'Password confirmation',
		resetPassword: 'Reset Password?',
		reset: 'Reset',
		sendResetRequest: 'Send',
		checkYourEmail: 'Please check your Email',
		checkYourEmailMessage: 'An email to reset your password was sent to <b>{email:string}</b>.'
	},
	mfa: {
		title: 'Multi-factor Authentication',
		totp: {
			codePlaceHolder: 'One-time Code'
		}
	},
	home: {
		title: 'Auth'
	},
	header: {
		title: 'Auth',
		applications: 'Applications',
		profile: 'Profile',
		signOut: 'Sign out',
		signIn: 'Sign in'
	},
	exercises: {
		title: 'Exercises'
	},
	dashboard: {
		title: 'Dashboard'
	},
	profile: {
		title: 'Profile',
		updateUsername: 'Update Username',
		submitUpdateUsername: 'Update',
		updateUserInfo: 'Info',
		updateEmails: 'Emails',
		updatePhoneNumbers: 'Phone Numbers',
		updateTOTPs: 'Time-based one-time passwords',
		emails: {
			cancel: 'Cancel',
			add: 'Add',
			sendConfirmation: 'Send Confirmation',
			setAsPrimary: 'Set as Primary',
			delete: 'Delete',
			checkYourEmail: 'Check your Email',
			confirmCode: 'Confirm',
			deleteEmail: 'Delete {0:string}?'
		},
		phoneNumbers: {
			cancel: 'Cancel',
			add: 'Add',
			sendConfirmation: 'Send Confirmation',
			setAsPrimary: 'Set as Primary',
			delete: 'Delete',
			checkYourPhone: 'Check your Phone',
			confirmCode: 'Confirm',
			deletePhoneNumber: 'Delete {0:string}?'
		},
		totps: {
			cancel: 'Cancel',
			add: 'Add',
			delete: 'Delete',
			deleteTOTP: 'Delete TOTP for {0:string}?'
		},
		mfa: {
			enabled: 'Enabled',
			disabled: 'Disabled'
		},
		notification: {
			usernameChangedSuccess: 'Username changed',
			passwordResetSuccess: 'Password has been reset.',
			userInfoChangedSuccess: 'User Info updated',
			sentEmailConfirmation: 'Sent Email Confirmation Code',
			sentPhoneNumberConfirmation: 'Sent Phone Number Confirmation Code',
			emailConfirmed: 'Email Confirmed',
			phoneNumberConfirmed: 'Phone Number Confirmed'
		}
	},
	applications: {
		title: 'Applications',
		filter: 'Filter..'
	},
	application: {
		title: 'Settings',
		description: 'Description',
		name: 'Name',
		uri: 'URL Safe Short Name',
		create: 'Create',
		update: 'Update',
		delete: {
			dangerZone: 'Danger Zone',
			dangerZoneMessage:
				'This operation is permanent and will delete all RBACs, users, tenents associated with application.',
			dangerZoneDeleteApplication: 'Delete Application',
			confirmTitle: 'Delete Application?',
			confirmMessage: 'Enter `{:uri}` to confirm delete.',
			confirm: 'Delete'
		}
	},
	users: {
		title: 'Users',
		id: 'Id',
		username: 'Username',
		email: 'Email',
		phoneNumber: 'Phone Number',
		newUser: {
			title: 'New User',
			button: 'Create'
		},
		edit: {
			title: 'Edit User',
			button: 'Edit'
		},
		delete: {
			button: 'Delete',
			confirmTitle: 'Delete User?',
			confirmMessage:
				'This operation is permanent and will delete everything associated with user {0:string}.',
			confirm: 'Delete'
		}
	},
	rbac: {
		title: 'RBAC'
	},
	tenents: {
		title: 'Tenents',
		id: 'Id',
		description: 'Description',
		uri: 'URI',
		descriptionPlaceholder: 'Description',
		uriPlaceholder: 'URI',
		newTenent: {
			title: 'New Tenent',
			button: 'Create'
		},
		edit: {
			title: 'Edit Tenent',
			button: 'Edit'
		},
		delete: {
			button: 'Delete',
			confirmTitle: 'Delete Tenent?',
			confirmMessage:
				'This operation is permanent and will delete everything associated with tenent {0:string}.',
			confirm: 'Delete'
		}
	},
	templates: {
		title: 'Templates'
	},
	maintenance: {
		title: 'Maintenance',
		header: 'Site down for Maintenance',
		body: 'Sorry, site is temporarily down for maintenance, check back <a href="{link:string}">here</a> in a bit.'
	},
	health: {
		title: 'Health',
		header: 'Health Check',
		body: 'Healthy'
	},
	common: {
		updatedAt: 'Updated at',
		createdAt: 'Created at'
	}
} satisfies BaseTranslation;

export default en;
