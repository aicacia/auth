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
import type { Application } from './Application';
import {
    ApplicationFromJSON,
    ApplicationFromJSONTyped,
    ApplicationToJSON,
} from './Application';

/**
 * 
 * @export
 * @interface PaginationApplication
 */
export interface PaginationApplication {
    /**
     * 
     * @type {boolean}
     * @memberof PaginationApplication
     */
    hasMore: boolean;
    /**
     * 
     * @type {Array<Application>}
     * @memberof PaginationApplication
     */
    items: Array<Application>;
}

/**
 * Check if a given object implements the PaginationApplication interface.
 */
export function instanceOfPaginationApplication(value: object): boolean {
    if (!('hasMore' in value)) return false;
    if (!('items' in value)) return false;
    return true;
}

export function PaginationApplicationFromJSON(json: any): PaginationApplication {
    return PaginationApplicationFromJSONTyped(json, false);
}

export function PaginationApplicationFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaginationApplication {
    if (json == null) {
        return json;
    }
    return {
        
        'hasMore': json['has_more'],
        'items': ((json['items'] as Array<any>).map(ApplicationFromJSON)),
    };
}

export function PaginationApplicationToJSON(value?: PaginationApplication | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'has_more': value['hasMore'],
        'items': ((value['items'] as Array<any>).map(ApplicationToJSON)),
    };
}

