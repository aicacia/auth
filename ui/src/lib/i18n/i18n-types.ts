// This file was auto-generated by 'typesafe-i18n'. Any manual changes will be overwritten.
/* eslint-disable */
import type { BaseTranslation as BaseTranslationType, LocalizedString, RequiredParams } from 'typesafe-i18n'

import type { uri } from './custom-types'

export type BaseTranslation = BaseTranslationType
export type BaseLocale = 'en'

export type Locales =
	| 'en'

export type Translation = RootTranslation

export type Translations = RootTranslation

type RootTranslation = {
	errors: {
		name: {
			/**
			 * A​p​p​l​i​c​a​t​i​o​n​ ​E​r​r​o​r
			 */
			internal: string
			/**
			 * R​e​q​u​e​s​t
			 */
			request: string
			/**
			 * U​s​e​r
			 */
			user: string
			/**
			 * E​m​a​i​l
			 */
			email: string
			/**
			 * U​s​e​r​n​a​m​e
			 */
			username: string
			/**
			 * P​a​s​s​w​o​r​d
			 */
			password: string
			/**
			 * P​a​s​s​w​o​r​d​ ​c​o​n​f​i​r​m​a​t​i​o​n
			 */
			passwordConfirmation: string
		}
		message: {
			/**
			 * i​f​ ​i​t​ ​p​r​e​s​i​s​t​s​ ​p​l​e​a​s​e​ ​c​o​n​t​a​c​t​ ​s​u​p​p​o​r​t
			 */
			application: string
			/**
			 * T​o​o​ ​m​a​n​y​ ​r​e​q​u​e​s​t​s
			 */
			tooManyRequests: string
			/**
			 * N​o​t​ ​f​o​u​n​d
			 */
			notFound: string
			/**
			 * I​n​v​a​l​i​d
			 */
			invalid: string
			/**
			 * R​e​q​u​i​r​e​d
			 */
			required: string
			/**
			 * N​o​ ​c​h​a​n​g​e
			 */
			noChange: string
			/**
			 * P​a​s​s​w​o​r​d​s​ ​d​o​ ​n​o​t​ ​m​a​t​c​h
			 */
			mismatch: string
			/**
			 * A​l​r​e​a​d​y​ ​u​s​e​d
			 */
			alreayUsed: string
			/**
			 * C​a​n​ ​n​o​t​ ​d​e​l​e​t​e​ ​A​d​m​i​n​ ​A​p​p​l​i​c​a​t​i​o​n​.
			 */
			cannotDeleteAdmin: string
		}
	}
	auth: {
		/**
		 * S​i​g​n​ ​i​n
		 */
		signIn: string
		/**
		 * S​i​g​n​ ​u​p
		 */
		signUp: string
		/**
		 * N​o​t​ ​a​ ​m​e​m​b​e​r​?
		 */
		notAMember: string
		/**
		 * A​l​r​e​a​d​y​ ​a​ ​m​e​m​b​e​r​?
		 */
		alreadyAMember: string
		/**
		 * B​a​c​k​ ​t​o​ ​s​i​g​n​ ​i​n
		 */
		backToSignIn: string
		/**
		 * E​m​a​i​l
		 */
		emailPlaceholder: string
		/**
		 * U​s​e​r​n​a​m​e
		 */
		usernamePlaceholder: string
		/**
		 * U​s​e​r​n​a​m​e​/​E​m​a​i​l
		 */
		usernameOrEmailPlaceholder: string
		/**
		 * P​a​s​s​w​o​r​d
		 */
		passwordPlaceholder: string
		/**
		 * P​a​s​s​w​o​r​d​ ​c​o​n​f​i​r​m​a​t​i​o​n
		 */
		passwordConfirmationPlaceholder: string
		/**
		 * R​e​s​e​t​ ​P​a​s​s​w​o​r​d​?
		 */
		resetPassword: string
		/**
		 * R​e​s​e​t
		 */
		reset: string
		/**
		 * S​e​n​d
		 */
		sendResetRequest: string
		/**
		 * P​l​e​a​s​e​ ​c​h​e​c​k​ ​y​o​u​r​ ​E​m​a​i​l
		 */
		checkYourEmail: string
		/**
		 * A​n​ ​e​m​a​i​l​ ​t​o​ ​r​e​s​e​t​ ​y​o​u​r​ ​p​a​s​s​w​o​r​d​ ​w​a​s​ ​s​e​n​t​ ​t​o​ ​<​b​>​{​e​m​a​i​l​}​<​/​b​>​.
		 * @param {string} email
		 */
		checkYourEmailMessage: RequiredParams<'email'>
	}
	home: {
		/**
		 * A​u​t​h
		 */
		title: string
	}
	header: {
		/**
		 * A​u​t​h
		 */
		title: string
		/**
		 * A​p​p​l​i​c​a​t​i​o​n​s
		 */
		applications: string
		/**
		 * P​r​o​f​i​l​e
		 */
		profile: string
		/**
		 * S​i​g​n​ ​o​u​t
		 */
		signOut: string
		/**
		 * S​i​g​n​ ​i​n
		 */
		signIn: string
	}
	exercises: {
		/**
		 * E​x​e​r​c​i​s​e​s
		 */
		title: string
	}
	dashboard: {
		/**
		 * D​a​s​h​b​o​a​r​d
		 */
		title: string
	}
	profile: {
		/**
		 * P​r​o​f​i​l​e
		 */
		title: string
		/**
		 * U​p​d​a​t​e​ ​U​s​e​r​n​a​m​e
		 */
		updateUsername: string
		/**
		 * U​p​d​a​t​e
		 */
		submitUpdateUsername: string
		/**
		 * U​p​d​a​t​e​ ​I​n​f​o
		 */
		updateUserInfo: string
		/**
		 * U​p​d​a​t​e​ ​E​m​a​i​l​s
		 */
		updateEmails: string
		/**
		 * U​p​d​a​t​e​ ​P​h​o​n​e​ ​N​u​m​b​e​r​s
		 */
		updatePhoneNumbers: string
		emails: {
			/**
			 * C​a​n​c​e​l
			 */
			cancel: string
			/**
			 * A​d​d
			 */
			add: string
			/**
			 * S​e​n​d​ ​C​o​n​f​i​r​m​a​t​i​o​n
			 */
			sendConfirmation: string
			/**
			 * S​e​t​ ​a​s​ ​P​r​i​m​a​r​y
			 */
			setAsPrimary: string
			/**
			 * D​e​l​e​t​e
			 */
			'delete': string
			/**
			 * C​h​e​c​k​ ​y​o​u​r​ ​E​m​a​i​l
			 */
			checkYourEmail: string
			/**
			 * C​o​n​f​i​r​m
			 */
			confirmCode: string
			/**
			 * D​e​l​e​t​e​ ​{​0​}​?
			 * @param {string} 0
			 */
			deleteEmail: RequiredParams<'0'>
		}
		phoneNumbers: {
			/**
			 * C​a​n​c​e​l
			 */
			cancel: string
			/**
			 * A​d​d
			 */
			add: string
			/**
			 * S​e​n​d​ ​C​o​n​f​i​r​m​a​t​i​o​n
			 */
			sendConfirmation: string
			/**
			 * S​e​t​ ​a​s​ ​P​r​i​m​a​r​y
			 */
			setAsPrimary: string
			/**
			 * D​e​l​e​t​e
			 */
			'delete': string
			/**
			 * C​h​e​c​k​ ​y​o​u​r​ ​P​h​o​n​e
			 */
			checkYourPhone: string
			/**
			 * C​o​n​f​i​r​m
			 */
			confirmCode: string
			/**
			 * D​e​l​e​t​e​ ​{​0​}​?
			 * @param {string} 0
			 */
			deletePhoneNumber: RequiredParams<'0'>
		}
		notification: {
			/**
			 * U​s​e​r​n​a​m​e​ ​c​h​a​n​g​e​d
			 */
			usernameChangedSuccess: string
			/**
			 * P​a​s​s​w​o​r​d​ ​h​a​s​ ​b​e​e​n​ ​r​e​s​e​t​.
			 */
			passwordResetSuccess: string
			/**
			 * U​s​e​r​ ​I​n​f​o​ ​u​p​d​a​t​e​d
			 */
			userInfoChangedSuccess: string
			/**
			 * S​e​n​t​ ​E​m​a​i​l​ ​C​o​n​f​i​r​m​a​t​i​o​n​ ​C​o​d​e
			 */
			sentEmailConfirmation: string
			/**
			 * S​e​n​t​ ​P​h​o​n​e​ ​N​u​m​b​e​r​ ​C​o​n​f​i​r​m​a​t​i​o​n​ ​C​o​d​e
			 */
			sentPhoneNumberConfirmation: string
			/**
			 * E​m​a​i​l​ ​C​o​n​f​i​r​m​e​d
			 */
			emailConfirmed: string
			/**
			 * P​h​o​n​e​ ​N​u​m​b​e​r​ ​C​o​n​f​i​r​m​e​d
			 */
			phoneNumberConfirmed: string
		}
	}
	applications: {
		/**
		 * A​p​p​l​i​c​a​t​i​o​n​s
		 */
		title: string
		/**
		 * F​i​l​t​e​r​.​.
		 */
		filter: string
	}
	application: {
		/**
		 * S​e​t​t​i​n​g​s
		 */
		title: string
		/**
		 * D​e​s​c​r​i​p​t​i​o​n
		 */
		description: string
		/**
		 * N​a​m​e
		 */
		name: string
		/**
		 * U​R​L​ ​S​a​f​e​ ​S​h​o​r​t​ ​N​a​m​e
		 */
		uri: string
		/**
		 * C​r​e​a​t​e
		 */
		create: string
		/**
		 * U​p​d​a​t​e
		 */
		update: string
		'delete': {
			/**
			 * D​a​n​g​e​r​ ​Z​o​n​e
			 */
			dangerZone: string
			/**
			 * T​h​i​s​ ​o​p​e​r​a​t​i​o​n​ ​i​s​ ​p​e​r​m​a​n​e​n​t​ ​a​n​d​ ​w​i​l​l​ ​d​e​l​e​t​e​ ​a​l​l​ ​p​e​r​m​i​s​s​i​o​n​s​ ​a​n​d​ ​d​a​t​a​ ​a​s​s​o​c​i​a​t​e​d​ ​w​i​t​h​ ​a​p​p​l​i​c​a​t​i​o​n​.
			 */
			dangerZoneMessage: string
			/**
			 * D​e​l​e​t​e​ ​A​p​p​l​i​c​a​t​i​o​n
			 */
			dangerZoneDeleteApplication: string
			/**
			 * D​e​l​e​t​e​ ​A​p​p​l​i​c​a​t​i​o​n​?
			 */
			confirmTitle: string
			/**
			 * E​n​t​e​r​ ​`​{​0​}​`​ ​t​o​ ​c​o​n​f​i​r​m​ ​d​e​l​e​t​e​.
			 * @param {uri} 0
			 */
			confirmMessage: RequiredParams<'0'>
			/**
			 * D​e​l​e​t​e
			 */
			confirm: string
		}
	}
	users: {
		/**
		 * U​s​e​r​s
		 */
		title: string
		/**
		 * I​d
		 */
		id: string
		/**
		 * U​s​e​r​n​a​m​e
		 */
		username: string
		/**
		 * E​m​a​i​l
		 */
		email: string
		/**
		 * P​h​o​n​e​ ​N​u​m​b​e​r
		 */
		phoneNumber: string
		edit: {
			/**
			 * E​d​i​t​ ​U​s​e​r
			 */
			title: string
			/**
			 * E​d​i​t
			 */
			button: string
		}
		'delete': {
			/**
			 * D​e​l​e​t​e
			 */
			button: string
			/**
			 * D​e​l​e​t​e​ ​U​s​e​r​?
			 */
			confirmTitle: string
			/**
			 * T​h​i​s​ ​o​p​e​r​a​t​i​o​n​ ​i​s​ ​p​e​r​m​a​n​e​n​t​ ​a​n​d​ ​w​i​l​l​ ​d​e​l​e​t​e​ ​e​v​e​n​y​t​h​i​n​g​ ​a​s​s​o​c​i​a​t​e​d​ ​w​i​t​h​ ​u​s​e​r​.
			 */
			confirmMessage: string
			/**
			 * D​e​l​e​t​e
			 */
			confirm: string
		}
	}
	permissions: {
		/**
		 * P​e​r​m​i​s​s​i​o​n​s
		 */
		title: string
	}
	tenents: {
		/**
		 * T​e​n​e​n​t​s
		 */
		title: string
	}
	templates: {
		/**
		 * T​e​m​p​l​a​t​e​s
		 */
		title: string
	}
	maintenance: {
		/**
		 * M​a​i​n​t​e​n​a​n​c​e
		 */
		title: string
		/**
		 * S​i​t​e​ ​d​o​w​n​ ​f​o​r​ ​M​a​i​n​t​e​n​a​n​c​e
		 */
		header: string
		/**
		 * S​o​r​r​y​,​ ​s​i​t​e​ ​i​s​ ​t​e​m​p​o​r​a​r​i​l​y​ ​d​o​w​n​ ​f​o​r​ ​m​a​i​n​t​e​n​a​n​c​e​,​ ​c​h​e​c​k​ ​b​a​c​k​ ​<​a​ ​h​r​e​f​=​"​{​l​i​n​k​}​"​>​h​e​r​e​<​/​a​>​ ​i​n​ ​a​ ​b​i​t​.
		 * @param {string} link
		 */
		body: RequiredParams<'link'>
	}
	health: {
		/**
		 * H​e​a​l​t​h
		 */
		title: string
		/**
		 * H​e​a​l​t​h​ ​C​h​e​c​k
		 */
		header: string
		/**
		 * H​e​a​l​t​h​y
		 */
		body: string
	}
	common: {
		/**
		 * U​p​d​a​t​e​d​ ​a​t
		 */
		updatedAt: string
		/**
		 * C​r​e​a​t​e​d​ ​a​t
		 */
		createdAt: string
	}
}

export type TranslationFunctions = {
	errors: {
		name: {
			/**
			 * Application Error
			 */
			internal: () => LocalizedString
			/**
			 * Request
			 */
			request: () => LocalizedString
			/**
			 * User
			 */
			user: () => LocalizedString
			/**
			 * Email
			 */
			email: () => LocalizedString
			/**
			 * Username
			 */
			username: () => LocalizedString
			/**
			 * Password
			 */
			password: () => LocalizedString
			/**
			 * Password confirmation
			 */
			passwordConfirmation: () => LocalizedString
		}
		message: {
			/**
			 * if it presists please contact support
			 */
			application: () => LocalizedString
			/**
			 * Too many requests
			 */
			tooManyRequests: () => LocalizedString
			/**
			 * Not found
			 */
			notFound: () => LocalizedString
			/**
			 * Invalid
			 */
			invalid: () => LocalizedString
			/**
			 * Required
			 */
			required: () => LocalizedString
			/**
			 * No change
			 */
			noChange: () => LocalizedString
			/**
			 * Passwords do not match
			 */
			mismatch: () => LocalizedString
			/**
			 * Already used
			 */
			alreayUsed: () => LocalizedString
			/**
			 * Can not delete Admin Application.
			 */
			cannotDeleteAdmin: () => LocalizedString
		}
	}
	auth: {
		/**
		 * Sign in
		 */
		signIn: () => LocalizedString
		/**
		 * Sign up
		 */
		signUp: () => LocalizedString
		/**
		 * Not a member?
		 */
		notAMember: () => LocalizedString
		/**
		 * Already a member?
		 */
		alreadyAMember: () => LocalizedString
		/**
		 * Back to sign in
		 */
		backToSignIn: () => LocalizedString
		/**
		 * Email
		 */
		emailPlaceholder: () => LocalizedString
		/**
		 * Username
		 */
		usernamePlaceholder: () => LocalizedString
		/**
		 * Username/Email
		 */
		usernameOrEmailPlaceholder: () => LocalizedString
		/**
		 * Password
		 */
		passwordPlaceholder: () => LocalizedString
		/**
		 * Password confirmation
		 */
		passwordConfirmationPlaceholder: () => LocalizedString
		/**
		 * Reset Password?
		 */
		resetPassword: () => LocalizedString
		/**
		 * Reset
		 */
		reset: () => LocalizedString
		/**
		 * Send
		 */
		sendResetRequest: () => LocalizedString
		/**
		 * Please check your Email
		 */
		checkYourEmail: () => LocalizedString
		/**
		 * An email to reset your password was sent to <b>{email}</b>.
		 */
		checkYourEmailMessage: (arg: { email: string }) => LocalizedString
	}
	home: {
		/**
		 * Auth
		 */
		title: () => LocalizedString
	}
	header: {
		/**
		 * Auth
		 */
		title: () => LocalizedString
		/**
		 * Applications
		 */
		applications: () => LocalizedString
		/**
		 * Profile
		 */
		profile: () => LocalizedString
		/**
		 * Sign out
		 */
		signOut: () => LocalizedString
		/**
		 * Sign in
		 */
		signIn: () => LocalizedString
	}
	exercises: {
		/**
		 * Exercises
		 */
		title: () => LocalizedString
	}
	dashboard: {
		/**
		 * Dashboard
		 */
		title: () => LocalizedString
	}
	profile: {
		/**
		 * Profile
		 */
		title: () => LocalizedString
		/**
		 * Update Username
		 */
		updateUsername: () => LocalizedString
		/**
		 * Update
		 */
		submitUpdateUsername: () => LocalizedString
		/**
		 * Update Info
		 */
		updateUserInfo: () => LocalizedString
		/**
		 * Update Emails
		 */
		updateEmails: () => LocalizedString
		/**
		 * Update Phone Numbers
		 */
		updatePhoneNumbers: () => LocalizedString
		emails: {
			/**
			 * Cancel
			 */
			cancel: () => LocalizedString
			/**
			 * Add
			 */
			add: () => LocalizedString
			/**
			 * Send Confirmation
			 */
			sendConfirmation: () => LocalizedString
			/**
			 * Set as Primary
			 */
			setAsPrimary: () => LocalizedString
			/**
			 * Delete
			 */
			'delete': () => LocalizedString
			/**
			 * Check your Email
			 */
			checkYourEmail: () => LocalizedString
			/**
			 * Confirm
			 */
			confirmCode: () => LocalizedString
			/**
			 * Delete {0}?
			 */
			deleteEmail: (arg0: string) => LocalizedString
		}
		phoneNumbers: {
			/**
			 * Cancel
			 */
			cancel: () => LocalizedString
			/**
			 * Add
			 */
			add: () => LocalizedString
			/**
			 * Send Confirmation
			 */
			sendConfirmation: () => LocalizedString
			/**
			 * Set as Primary
			 */
			setAsPrimary: () => LocalizedString
			/**
			 * Delete
			 */
			'delete': () => LocalizedString
			/**
			 * Check your Phone
			 */
			checkYourPhone: () => LocalizedString
			/**
			 * Confirm
			 */
			confirmCode: () => LocalizedString
			/**
			 * Delete {0}?
			 */
			deletePhoneNumber: (arg0: string) => LocalizedString
		}
		notification: {
			/**
			 * Username changed
			 */
			usernameChangedSuccess: () => LocalizedString
			/**
			 * Password has been reset.
			 */
			passwordResetSuccess: () => LocalizedString
			/**
			 * User Info updated
			 */
			userInfoChangedSuccess: () => LocalizedString
			/**
			 * Sent Email Confirmation Code
			 */
			sentEmailConfirmation: () => LocalizedString
			/**
			 * Sent Phone Number Confirmation Code
			 */
			sentPhoneNumberConfirmation: () => LocalizedString
			/**
			 * Email Confirmed
			 */
			emailConfirmed: () => LocalizedString
			/**
			 * Phone Number Confirmed
			 */
			phoneNumberConfirmed: () => LocalizedString
		}
	}
	applications: {
		/**
		 * Applications
		 */
		title: () => LocalizedString
		/**
		 * Filter..
		 */
		filter: () => LocalizedString
	}
	application: {
		/**
		 * Settings
		 */
		title: () => LocalizedString
		/**
		 * Description
		 */
		description: () => LocalizedString
		/**
		 * Name
		 */
		name: () => LocalizedString
		/**
		 * URL Safe Short Name
		 */
		uri: () => LocalizedString
		/**
		 * Create
		 */
		create: () => LocalizedString
		/**
		 * Update
		 */
		update: () => LocalizedString
		'delete': {
			/**
			 * Danger Zone
			 */
			dangerZone: () => LocalizedString
			/**
			 * This operation is permanent and will delete all permissions and data associated with application.
			 */
			dangerZoneMessage: () => LocalizedString
			/**
			 * Delete Application
			 */
			dangerZoneDeleteApplication: () => LocalizedString
			/**
			 * Delete Application?
			 */
			confirmTitle: () => LocalizedString
			/**
			 * Enter `{0}` to confirm delete.
			 */
			confirmMessage: (arg0: uri) => LocalizedString
			/**
			 * Delete
			 */
			confirm: () => LocalizedString
		}
	}
	users: {
		/**
		 * Users
		 */
		title: () => LocalizedString
		/**
		 * Id
		 */
		id: () => LocalizedString
		/**
		 * Username
		 */
		username: () => LocalizedString
		/**
		 * Email
		 */
		email: () => LocalizedString
		/**
		 * Phone Number
		 */
		phoneNumber: () => LocalizedString
		edit: {
			/**
			 * Edit User
			 */
			title: () => LocalizedString
			/**
			 * Edit
			 */
			button: () => LocalizedString
		}
		'delete': {
			/**
			 * Delete
			 */
			button: () => LocalizedString
			/**
			 * Delete User?
			 */
			confirmTitle: () => LocalizedString
			/**
			 * This operation is permanent and will delete evenything associated with user.
			 */
			confirmMessage: () => LocalizedString
			/**
			 * Delete
			 */
			confirm: () => LocalizedString
		}
	}
	permissions: {
		/**
		 * Permissions
		 */
		title: () => LocalizedString
	}
	tenents: {
		/**
		 * Tenents
		 */
		title: () => LocalizedString
	}
	templates: {
		/**
		 * Templates
		 */
		title: () => LocalizedString
	}
	maintenance: {
		/**
		 * Maintenance
		 */
		title: () => LocalizedString
		/**
		 * Site down for Maintenance
		 */
		header: () => LocalizedString
		/**
		 * Sorry, site is temporarily down for maintenance, check back <a href="{link}">here</a> in a bit.
		 */
		body: (arg: { link: string }) => LocalizedString
	}
	health: {
		/**
		 * Health
		 */
		title: () => LocalizedString
		/**
		 * Health Check
		 */
		header: () => LocalizedString
		/**
		 * Healthy
		 */
		body: () => LocalizedString
	}
	common: {
		/**
		 * Updated at
		 */
		updatedAt: () => LocalizedString
		/**
		 * Created at
		 */
		createdAt: () => LocalizedString
	}
}

export type Formatters = {}
