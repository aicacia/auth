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


/**
 * 
 * @export
 */
export const ProtocolCredentialType = {
    PublicKeyCredentialType: 'public-key'
} as const;
export type ProtocolCredentialType = typeof ProtocolCredentialType[keyof typeof ProtocolCredentialType];


export function ProtocolCredentialTypeFromJSON(json: any): ProtocolCredentialType {
    return ProtocolCredentialTypeFromJSONTyped(json, false);
}

export function ProtocolCredentialTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProtocolCredentialType {
    return json as ProtocolCredentialType;
}

export function ProtocolCredentialTypeToJSON(value?: ProtocolCredentialType | null): any {
    return value as any;
}
