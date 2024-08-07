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
 * @interface ProtocolUserEntity
 */
export interface ProtocolUserEntity {
    /**
     * A human-palatable name for the user account, intended only for display.
     * For example, "Alex P. Müller" or "田中 倫". The Relying Party SHOULD let
     * the user choose this, and SHOULD NOT restrict the choice more than necessary.
     * @type {string}
     * @memberof ProtocolUserEntity
     */
    displayName?: string;
    /**
     * A serialized URL which resolves to an image associated with the entity. For example,
     * this could be a user’s avatar or a Relying Party's logo. This URL MUST be an a priori
     * authenticated URL. Authenticators MUST accept and store a 128-byte minimum length for
     * an icon member’s value. Authenticators MAY ignore an icon member’s value if its length
     * is greater than 128 bytes. The URL’s scheme MAY be "data" to avoid fetches of the URL,
     * at the cost of needing more storage.
     * 
     * Deprecated: this has been removed from the specification recommendations.
     * @type {string}
     * @memberof ProtocolUserEntity
     */
    icon?: string;
    /**
     * ID is the user handle of the user account entity. To ensure secure operation,
     * authentication and authorization decisions MUST be made on the basis of this id
     * member, not the displayName nor name members. See Section 6.1 of
     * [RFC8266](https://www.w3.org/TR/webauthn/#biblio-rfc8266).
     * @type {object}
     * @memberof ProtocolUserEntity
     */
    id?: object;
    /**
     * A human-palatable name for the entity. Its function depends on what the PublicKeyCredentialEntity represents:
     * 
     * When inherited by PublicKeyCredentialRpEntity it is a human-palatable identifier for the Relying Party,
     * intended only for display. For example, "ACME Corporation", "Wonderful Widgets, Inc." or "ОАО Примертех".
     * 
     * When inherited by PublicKeyCredentialUserEntity, it is a human-palatable identifier for a user account. It is
     * intended only for display, i.e., aiding the user in determining the difference between user accounts with similar
     * displayNames. For example, "alexm", "alex.p.mueller@example.com" or "+14255551234".
     * @type {string}
     * @memberof ProtocolUserEntity
     */
    name?: string;
}

/**
 * Check if a given object implements the ProtocolUserEntity interface.
 */
export function instanceOfProtocolUserEntity(value: object): boolean {
    return true;
}

export function ProtocolUserEntityFromJSON(json: any): ProtocolUserEntity {
    return ProtocolUserEntityFromJSONTyped(json, false);
}

export function ProtocolUserEntityFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProtocolUserEntity {
    if (json == null) {
        return json;
    }
    return {
        
        'displayName': json['displayName'] == null ? undefined : json['displayName'],
        'icon': json['icon'] == null ? undefined : json['icon'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function ProtocolUserEntityToJSON(value?: ProtocolUserEntity | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'displayName': value['displayName'],
        'icon': value['icon'],
        'id': value['id'],
        'name': value['name'],
    };
}

