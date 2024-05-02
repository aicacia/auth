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
 * @interface UpdateTenent
 */
export interface UpdateTenent {
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    algorithm?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    authorization_website?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    client_id?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    description: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    email_endpoint?: string;
    /**
     * 
     * @type {number}
     * @memberof UpdateTenent
     */
    expires_in_seconds?: number;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    phone_number_endpoint?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    private_key?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    public_key?: string;
    /**
     * 
     * @type {number}
     * @memberof UpdateTenent
     */
    refresh_expires_in_seconds?: number;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    registration_website?: string;
    /**
     * 
     * @type {number}
     * @memberof UpdateTenent
     */
    reset_expires_in_seconds?: number;
    /**
     * 
     * @type {string}
     * @memberof UpdateTenent
     */
    uri: string;
}

/**
 * Check if a given object implements the UpdateTenent interface.
 */
export function instanceOfUpdateTenent(value: object): boolean {
    if (!('description' in value)) return false;
    if (!('uri' in value)) return false;
    return true;
}

export function UpdateTenentFromJSON(json: any): UpdateTenent {
    return UpdateTenentFromJSONTyped(json, false);
}

export function UpdateTenentFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateTenent {
    if (json == null) {
        return json;
    }
    return {
        
        'algorithm': json['algorithm'] == null ? undefined : json['algorithm'],
        'authorization_website': json['authorization_website'] == null ? undefined : json['authorization_website'],
        'client_id': json['client_id'] == null ? undefined : json['client_id'],
        'description': json['description'],
        'email_endpoint': json['email_endpoint'] == null ? undefined : json['email_endpoint'],
        'expires_in_seconds': json['expires_in_seconds'] == null ? undefined : json['expires_in_seconds'],
        'phone_number_endpoint': json['phone_number_endpoint'] == null ? undefined : json['phone_number_endpoint'],
        'private_key': json['private_key'] == null ? undefined : json['private_key'],
        'public_key': json['public_key'] == null ? undefined : json['public_key'],
        'refresh_expires_in_seconds': json['refresh_expires_in_seconds'] == null ? undefined : json['refresh_expires_in_seconds'],
        'registration_website': json['registration_website'] == null ? undefined : json['registration_website'],
        'reset_expires_in_seconds': json['reset_expires_in_seconds'] == null ? undefined : json['reset_expires_in_seconds'],
        'uri': json['uri'],
    };
}

export function UpdateTenentToJSON(value?: UpdateTenent | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'algorithm': value['algorithm'],
        'authorization_website': value['authorization_website'],
        'client_id': value['client_id'],
        'description': value['description'],
        'email_endpoint': value['email_endpoint'],
        'expires_in_seconds': value['expires_in_seconds'],
        'phone_number_endpoint': value['phone_number_endpoint'],
        'private_key': value['private_key'],
        'public_key': value['public_key'],
        'refresh_expires_in_seconds': value['refresh_expires_in_seconds'],
        'registration_website': value['registration_website'],
        'reset_expires_in_seconds': value['reset_expires_in_seconds'],
        'uri': value['uri'],
    };
}

