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
 * @interface Application
 */
export interface Application {
    /**
     * 
     * @type {Date}
     * @memberof Application
     */
    created_at: Date;
    /**
     * 
     * @type {string}
     * @memberof Application
     */
    description: string;
    /**
     * 
     * @type {number}
     * @memberof Application
     */
    id: number;
    /**
     * 
     * @type {Date}
     * @memberof Application
     */
    updated_at: Date;
    /**
     * 
     * @type {string}
     * @memberof Application
     */
    uri: string;
    /**
     * 
     * @type {string}
     * @memberof Application
     */
    website?: string;
}

/**
 * Check if a given object implements the Application interface.
 */
export function instanceOfApplication(value: object): boolean {
    if (!('created_at' in value)) return false;
    if (!('description' in value)) return false;
    if (!('id' in value)) return false;
    if (!('updated_at' in value)) return false;
    if (!('uri' in value)) return false;
    return true;
}

export function ApplicationFromJSON(json: any): Application {
    return ApplicationFromJSONTyped(json, false);
}

export function ApplicationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Application {
    if (json == null) {
        return json;
    }
    return {
        
        'created_at': (new Date(json['created_at'])),
        'description': json['description'],
        'id': json['id'],
        'updated_at': (new Date(json['updated_at'])),
        'uri': json['uri'],
        'website': json['website'] == null ? undefined : json['website'],
    };
}

export function ApplicationToJSON(value?: Application | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'created_at': ((value['created_at']).toISOString()),
        'description': value['description'],
        'id': value['id'],
        'updated_at': ((value['updated_at']).toISOString()),
        'uri': value['uri'],
        'website': value['website'],
    };
}

