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
import type { ProtocolAuthenticatorTransport } from './ProtocolAuthenticatorTransport';
import {
    ProtocolAuthenticatorTransportFromJSON,
    ProtocolAuthenticatorTransportFromJSONTyped,
    ProtocolAuthenticatorTransportToJSON,
} from './ProtocolAuthenticatorTransport';
import type { ProtocolCredentialType } from './ProtocolCredentialType';
import {
    ProtocolCredentialTypeFromJSON,
    ProtocolCredentialTypeFromJSONTyped,
    ProtocolCredentialTypeToJSON,
} from './ProtocolCredentialType';

/**
 * 
 * @export
 * @interface ProtocolCredentialDescriptor
 */
export interface ProtocolCredentialDescriptor {
    /**
     * CredentialID The ID of a credential to allow/disallow.
     * @type {Array<number>}
     * @memberof ProtocolCredentialDescriptor
     */
    id?: Array<number>;
    /**
     * The authenticator transports that can be used.
     * @type {Array<ProtocolAuthenticatorTransport>}
     * @memberof ProtocolCredentialDescriptor
     */
    transports?: Array<ProtocolAuthenticatorTransport>;
    /**
     * The valid credential types.
     * @type {ProtocolCredentialType}
     * @memberof ProtocolCredentialDescriptor
     */
    type?: ProtocolCredentialType;
}

/**
 * Check if a given object implements the ProtocolCredentialDescriptor interface.
 */
export function instanceOfProtocolCredentialDescriptor(value: object): boolean {
    return true;
}

export function ProtocolCredentialDescriptorFromJSON(json: any): ProtocolCredentialDescriptor {
    return ProtocolCredentialDescriptorFromJSONTyped(json, false);
}

export function ProtocolCredentialDescriptorFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProtocolCredentialDescriptor {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'transports': json['transports'] == null ? undefined : ((json['transports'] as Array<any>).map(ProtocolAuthenticatorTransportFromJSON)),
        'type': json['type'] == null ? undefined : ProtocolCredentialTypeFromJSON(json['type']),
    };
}

export function ProtocolCredentialDescriptorToJSON(value?: ProtocolCredentialDescriptor | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'transports': value['transports'] == null ? undefined : ((value['transports'] as Array<any>).map(ProtocolAuthenticatorTransportToJSON)),
        'type': ProtocolCredentialTypeToJSON(value['type']),
    };
}
