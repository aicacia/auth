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
			passwordConfirmation: 'Password confirmation'
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
			cannotDeleteAdmin: 'Can not delete Admin Application.'
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
		notification: {
			usernameChangedSuccess: 'Username changed',
			passwordResetSuccess: 'Your Password has been reset.'
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
				'This operation is permanent and will delete all permissions and data associated with application.',
			dangerZoneDeleteApplication: 'Delete Application',
			confirmTitle: 'Delete Application?',
			confirmMessage: 'Enter `{:uri}` to confirm delete.',
			confirm: 'Delete'
		}
	},
	users: {
		title: 'Users'
	},
	permissions: {
		title: 'Permissions'
	},
	tenents: {
		title: 'Tenents'
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
	}
} satisfies BaseTranslation;

export default en;
