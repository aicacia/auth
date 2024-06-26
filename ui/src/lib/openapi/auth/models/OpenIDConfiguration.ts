/* tslint:disable */
/* eslint-disable */
/**
 * Auth API
 * Auth API API
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: nathanfaucett@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface OpenIDConfiguration
 */
export interface OpenIDConfiguration {
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    authorizationEndpoint?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    claimsSupported: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    codeChallengeMethodsSupported: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    grantTypesSupported: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    idTokenSigningAlgValuesSupported: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    issuer: string;
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    jwksUri?: string;
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    registrationEndpoint?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    responseTypesSupported: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    scopesSupported: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    subjectTypesSupported: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    tokenEndpoint: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof OpenIDConfiguration
     */
    tokenEndpointAuthMethodsSupported: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof OpenIDConfiguration
     */
    userinfoEndpoint: string;
}

/**
 * Check if a given object implements the OpenIDConfiguration interface.
 */
export function instanceOfOpenIDConfiguration(value: object): boolean {
    if (!('claimsSupported' in value)) return false;
    if (!('codeChallengeMethodsSupported' in value)) return false;
    if (!('grantTypesSupported' in value)) return false;
    if (!('idTokenSigningAlgValuesSupported' in value)) return false;
    if (!('issuer' in value)) return false;
    if (!('responseTypesSupported' in value)) return false;
    if (!('scopesSupported' in value)) return false;
    if (!('subjectTypesSupported' in value)) return false;
    if (!('tokenEndpoint' in value)) return false;
    if (!('tokenEndpointAuthMethodsSupported' in value)) return false;
    if (!('userinfoEndpoint' in value)) return false;
    return true;
}

export function OpenIDConfigurationFromJSON(json: any): OpenIDConfiguration {
    return OpenIDConfigurationFromJSONTyped(json, false);
}

export function OpenIDConfigurationFromJSONTyped(json: any, ignoreDiscriminator: boolean): OpenIDConfiguration {
    if (json == null) {
        return json;
    }
    return {
        
        'authorizationEndpoint': json['authorization_endpoint'] == null ? undefined : json['authorization_endpoint'],
        'claimsSupported': json['claims_supported'],
        'codeChallengeMethodsSupported': json['code_challenge_methods_supported'],
        'grantTypesSupported': json['grant_types_supported'],
        'idTokenSigningAlgValuesSupported': json['id_token_signing_alg_values_supported'],
        'issuer': json['issuer'],
        'jwksUri': json['jwks_uri'] == null ? undefined : json['jwks_uri'],
        'registrationEndpoint': json['registration_endpoint'] == null ? undefined : json['registration_endpoint'],
        'responseTypesSupported': json['response_types_supported'],
        'scopesSupported': json['scopes_supported'],
        'subjectTypesSupported': json['subject_types_supported'],
        'tokenEndpoint': json['token_endpoint'],
        'tokenEndpointAuthMethodsSupported': json['token_endpoint_auth_methods_supported'],
        'userinfoEndpoint': json['userinfo_endpoint'],
    };
}

export function OpenIDConfigurationToJSON(value?: OpenIDConfiguration | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'authorization_endpoint': value['authorizationEndpoint'],
        'claims_supported': value['claimsSupported'],
        'code_challenge_methods_supported': value['codeChallengeMethodsSupported'],
        'grant_types_supported': value['grantTypesSupported'],
        'id_token_signing_alg_values_supported': value['idTokenSigningAlgValuesSupported'],
        'issuer': value['issuer'],
        'jwks_uri': value['jwksUri'],
        'registration_endpoint': value['registrationEndpoint'],
        'response_types_supported': value['responseTypesSupported'],
        'scopes_supported': value['scopesSupported'],
        'subject_types_supported': value['subjectTypesSupported'],
        'token_endpoint': value['tokenEndpoint'],
        'token_endpoint_auth_methods_supported': value['tokenEndpointAuthMethodsSupported'],
        'userinfo_endpoint': value['userinfoEndpoint'],
    };
}

