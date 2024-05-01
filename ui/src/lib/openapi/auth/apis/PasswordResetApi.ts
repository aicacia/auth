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


import * as runtime from '../runtime';
import type {
  Errors,
  ModelPasswordResetST,
  ModelRequestPasswordResetST,
  Token,
} from '../models/index';
import {
    ErrorsFromJSON,
    ErrorsToJSON,
    ModelPasswordResetSTFromJSON,
    ModelPasswordResetSTToJSON,
    ModelRequestPasswordResetSTFromJSON,
    ModelRequestPasswordResetSTToJSON,
    TokenFromJSON,
    TokenToJSON,
} from '../models/index';

export interface PasswordResetRequest {
    passwordReset: ModelPasswordResetST;
}

export interface RequestPasswordResetRequest {
    requestPasswordReset: ModelRequestPasswordResetST;
}

/**
 * PasswordResetApi - interface
 * 
 * @export
 * @interface PasswordResetApiInterface
 */
export interface PasswordResetApiInterface {
    /**
     * 
     * @summary Request Password Reset
     * @param {ModelPasswordResetST} passwordReset request password reset body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PasswordResetApiInterface
     */
    passwordResetRaw(requestParameters: PasswordResetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Token>>;

    /**
     * Request Password Reset
     */
    passwordReset(passwordReset: ModelPasswordResetST, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Token>;

    /**
     * 
     * @summary Request Password Reset
     * @param {ModelRequestPasswordResetST} requestPasswordReset request password reset body
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PasswordResetApiInterface
     */
    requestPasswordResetRaw(requestParameters: RequestPasswordResetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Request Password Reset
     */
    requestPasswordReset(requestPasswordReset: ModelRequestPasswordResetST, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;

}

/**
 * 
 */
export class PasswordResetApi extends runtime.BaseAPI implements PasswordResetApiInterface {

    /**
     * Request Password Reset
     */
    async passwordResetRaw(requestParameters: PasswordResetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Token>> {
        if (requestParameters['passwordReset'] == null) {
            throw new runtime.RequiredError(
                'passwordReset',
                'Required parameter "passwordReset" was null or undefined when calling passwordReset().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/password-reset`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelPasswordResetSTToJSON(requestParameters['passwordReset']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TokenFromJSON(jsonValue));
    }

    /**
     * Request Password Reset
     */
    async passwordReset(passwordReset: ModelPasswordResetST, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Token> {
        const response = await this.passwordResetRaw({ passwordReset: passwordReset }, initOverrides);
        return await response.value();
    }

    /**
     * Request Password Reset
     */
    async requestPasswordResetRaw(requestParameters: RequestPasswordResetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['requestPasswordReset'] == null) {
            throw new runtime.RequiredError(
                'requestPasswordReset',
                'Required parameter "requestPasswordReset" was null or undefined when calling requestPasswordReset().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/password-reset/request`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelRequestPasswordResetSTToJSON(requestParameters['requestPasswordReset']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Request Password Reset
     */
    async requestPasswordReset(requestPasswordReset: ModelRequestPasswordResetST, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.requestPasswordResetRaw({ requestPasswordReset: requestPasswordReset }, initOverrides);
    }

}
