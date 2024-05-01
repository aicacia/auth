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
 * @interface CreatePhoneNumber
 */
export interface CreatePhoneNumber {
    /**
     * 
     * @type {string}
     * @memberof CreatePhoneNumber
     */
    phone_number: string;
}

/**
 * Check if a given object implements the CreatePhoneNumber interface.
 */
export function instanceOfCreatePhoneNumber(value: object): boolean {
    if (!('phone_number' in value)) return false;
    return true;
}

export function CreatePhoneNumberFromJSON(json: any): CreatePhoneNumber {
    return CreatePhoneNumberFromJSONTyped(json, false);
}

export function CreatePhoneNumberFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreatePhoneNumber {
    if (json == null) {
        return json;
    }
    return {
        
        'phone_number': json['phone_number'],
    };
}

export function CreatePhoneNumberToJSON(value?: CreatePhoneNumber | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'phone_number': value['phone_number'],
    };
}

