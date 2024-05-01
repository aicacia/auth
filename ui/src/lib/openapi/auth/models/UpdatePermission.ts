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
 * @interface UpdatePermission
 */
export interface UpdatePermission {
    /**
     * 
     * @type {string}
     * @memberof UpdatePermission
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdatePermission
     */
    uri?: string;
}

/**
 * Check if a given object implements the UpdatePermission interface.
 */
export function instanceOfUpdatePermission(value: object): boolean {
    return true;
}

export function UpdatePermissionFromJSON(json: any): UpdatePermission {
    return UpdatePermissionFromJSONTyped(json, false);
}

export function UpdatePermissionFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdatePermission {
    if (json == null) {
        return json;
    }
    return {
        
        'description': json['description'] == null ? undefined : json['description'],
        'uri': json['uri'] == null ? undefined : json['uri'],
    };
}

export function UpdatePermissionToJSON(value?: UpdatePermission | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'description': value['description'],
        'uri': value['uri'],
    };
}

