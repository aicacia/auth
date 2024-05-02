import {
	Configuration,
	UserApi,
	CurrentUserApi,
	TokenApi,
	WellKnownApi,
	type ConfigurationParameters,
	ApplicationApi,
	RegistrationApi,
	type Token
} from './auth';
import { env } from '$env/dynamic/public';

let authToken: Token | undefined;

const defaultConfiguration: ConfigurationParameters = {
	middleware: [
		{
			pre: async (context) => ({ ...context, init: { ...context.init, mode: 'cors' } })
		}
	],
	apiKey(name: string) {
		switch (name) {
			case 'Tenent-Id':
				return env.PUBLIC_TENENT_ID as string;
			default:
				return `${authToken?.token_type} ${authToken?.access_token}`;
		}
	},
	credentials: 'same-origin'
};

const defaultAuthConfiguration: ConfigurationParameters = {
	...defaultConfiguration,
	get basePath() {
		return typeof __DEV_AUTH_API_URL__ !== 'undefined'
			? __DEV_AUTH_API_URL__
			: env.PUBLIC_AUTH_API_URL;
	}
};

export const authConfiguration = new Configuration(defaultAuthConfiguration);

export const currentUserApi = new CurrentUserApi(authConfiguration);
export const userApi = new UserApi(authConfiguration);
export const tokenApi = new TokenApi(authConfiguration);
export const registrationApi = new RegistrationApi(authConfiguration);
export const applicationApi = new ApplicationApi(authConfiguration);
export const wellKnownApi = new WellKnownApi(authConfiguration);

export function setAuthToken(newAuthToken?: Token) {
	authToken = newAuthToken;
}
export function getAuthToken() {
	return authToken;
}
