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
 * @interface Email
 */
export interface Email {
    /**
     * 
     * @type {number}
     * @memberof Email
     */
    applicationId: number;
    /**
     * 
     * @type {boolean}
     * @memberof Email
     */
    confirmed: boolean;
    /**
     * 
     * @type {Date}
     * @memberof Email
     */
    createdAt: Date;
    /**
     * 
     * @type {string}
     * @memberof Email
     */
    email: string;
    /**
     * 
     * @type {number}
     * @memberof Email
     */
    id: number;
    /**
     * 
     * @type {Date}
     * @memberof Email
     */
    updatedAt: Date;
}

/**
 * Check if a given object implements the Email interface.
 */
export function instanceOfEmail(value: object): boolean {
    if (!('applicationId' in value)) return false;
    if (!('confirmed' in value)) return false;
    if (!('createdAt' in value)) return false;
    if (!('email' in value)) return false;
    if (!('id' in value)) return false;
    if (!('updatedAt' in value)) return false;
    return true;
}

export function EmailFromJSON(json: any): Email {
    return EmailFromJSONTyped(json, false);
}

export function EmailFromJSONTyped(json: any, ignoreDiscriminator: boolean): Email {
    if (json == null) {
        return json;
    }
    return {
        
        'applicationId': json['application_id'],
        'confirmed': json['confirmed'],
        'createdAt': (new Date(json['created_at'])),
        'email': json['email'],
        'id': json['id'],
        'updatedAt': (new Date(json['updated_at'])),
    };
}

export function EmailToJSON(value?: Email | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'application_id': value['applicationId'],
        'confirmed': value['confirmed'],
        'created_at': ((value['createdAt']).toISOString()),
        'email': value['email'],
        'id': value['id'],
        'updated_at': ((value['updatedAt']).toISOString()),
    };
}

