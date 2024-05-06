import { derived, get } from 'svelte/store';
import { localstorageWritable } from 'svelte-localstorage-writable';
import { isOnline } from './online';
import {
	authConfiguration,
	getAuthToken,
	registrationApi,
	setAuthToken,
	tokenApi,
	currentUserApi
} from '$lib/openapi';
import { type Token, type User } from '$lib/openapi/auth';
import EventEmitter from 'eventemitter3';
import { goto } from '$app/navigation';
import { base } from '$app/paths';

const tokenWritable = localstorageWritable<Token | null>('token', null);
const userWritable = localstorageWritable<User | null>('user', null);

export const currentUser = derived(userWritable, (user) => user);
export const signedIn = derived(userWritable, (user) => !!user);

export const userEmitter = new EventEmitter<{
	user(user: User): void;
	signOut(): void;
}>();

export function waitForUser() {
	const user = get(userWritable);
	if (getAuthToken() && user) {
		return Promise.resolve(user);
	} else {
		return new Promise<User>((resolve) => userEmitter.once('user', resolve));
	}
}

export function updateCurrentUser(user: User) {
	if (get(currentUser)?.id === user.id) {
		userWritable.update((currentUser) => (currentUser ? { ...currentUser, ...user } : null));
	}
}

export function isSignedIn() {
	return get(signedIn);
}

export function getCurrentUser() {
	return get(currentUser);
}

export async function signIn(usernameOrEmail: string, password: string) {
	const token = await tokenApi.createToken({
		grantType: 'password',
		username: usernameOrEmail,
		password
	});
	return signInWithToken(token);
}

export async function signUp(username: string, password: string, passwordConfirmation: string) {
	const token = await registrationApi.registerUser({
		username,
		password,
		passwordConfirmation: passwordConfirmation
	});
	return signInWithToken(token);
}

export async function signInWithToken(token: Token) {
	setAuthToken(token);
	const user = await currentUserApi.currentUser();
	userWritable.set(user);
	tokenWritable.set(token);
	userEmitter.emit('user', user);
	return user;
}

export function signOut() {
	userWritable.set(null);
	tokenWritable.set(null);
	setAuthToken(undefined);
	userEmitter.emit('signOut');
}

let initialCall = true;
export async function tryGetCurrentUser() {
	try {
		let user = get(userWritable);
		if (initialCall) {
			if (isOnline()) {
				const token = get(tokenWritable);
				if (token) {
					setAuthToken(token);
					user = await currentUserApi.currentUser();
					userWritable.set(user);
					userEmitter.emit('user', user);
				} else {
					signOut();
					user = null;
				}
			} else if (user) {
				userEmitter.emit('user', user);
			}
			initialCall = false;
		}
		return user;
	} catch (error) {
		console.error(error);
		signOut();
		return null;
	}
}

authConfiguration.middleware?.push({
	async post(context) {
		switch (context.response.status) {
			case 401:
				signOut();
				await goto(`${base}/signin`);
				break;
			case 503:
				await goto(`${base}/maintenance`);
		}
	}
});
