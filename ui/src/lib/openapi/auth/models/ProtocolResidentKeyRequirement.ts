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
export const ProtocolResidentKeyRequirement = {
    ResidentKeyRequirementDiscouraged: 'discouraged',
    ResidentKeyRequirementPreferred: 'preferred',
    ResidentKeyRequirementRequired: 'required'
} as const;
export type ProtocolResidentKeyRequirement = typeof ProtocolResidentKeyRequirement[keyof typeof ProtocolResidentKeyRequirement];


export function ProtocolResidentKeyRequirementFromJSON(json: any): ProtocolResidentKeyRequirement {
    return ProtocolResidentKeyRequirementFromJSONTyped(json, false);
}

export function ProtocolResidentKeyRequirementFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProtocolResidentKeyRequirement {
    return json as ProtocolResidentKeyRequirement;
}

export function ProtocolResidentKeyRequirementToJSON(value?: ProtocolResidentKeyRequirement | null): any {
    return value as any;
}

