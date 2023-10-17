/**
 * Wasp API
 * REST API for the Wasp node
 *
 * OpenAPI spec version: 0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class EventJSON {
    /**
    * ID of the Contract that issued the event
    */
    'contractID': number;
    /**
    * payload
    */
    'payload': string;
    /**
    * timestamp
    */
    'timestamp': number;
    /**
    * topic
    */
    'topic': string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "contractID",
            "baseName": "contractID",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "payload",
            "baseName": "payload",
            "type": "string",
            "format": "string"
        },
        {
            "name": "timestamp",
            "baseName": "timestamp",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "topic",
            "baseName": "topic",
            "type": "string",
            "format": "string"
        }    ];

    static getAttributeTypeMap() {
        return EventJSON.attributeTypeMap;
    }

    public constructor() {
    }
}
