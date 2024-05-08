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
 * @interface ModelRequestPasswordResetST
 */
export interface ModelRequestPasswordResetST {
    /**
     * 
     * @type {string}
     * @memberof ModelRequestPasswordResetST
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof ModelRequestPasswordResetST
     */
    phoneNumber: string;
}

/**
 * Check if a given object implements the ModelRequestPasswordResetST interface.
 */
export function instanceOfModelRequestPasswordResetST(value: object): boolean {
    if (!('email' in value)) return false;
    if (!('phoneNumber' in value)) return false;
    return true;
}

export function ModelRequestPasswordResetSTFromJSON(json: any): ModelRequestPasswordResetST {
    return ModelRequestPasswordResetSTFromJSONTyped(json, false);
}

export function ModelRequestPasswordResetSTFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelRequestPasswordResetST {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'],
        'phoneNumber': json['phoneNumber'],
    };
}

export function ModelRequestPasswordResetSTToJSON(value?: ModelRequestPasswordResetST | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'email': value['email'],
        'phoneNumber': value['phoneNumber'],
    };
}

