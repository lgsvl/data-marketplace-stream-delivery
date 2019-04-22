//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

package resources

const (
	//Swaggers description
	Swaggers = `{
		"swagger": "2.0",
		"info": {
			"title": "api.proto",
			"version": "version not set"
		},
		"schemes": [
			"http",
			"https"
		],
		"consumes": [
			"application/json"
		],
		"produces": [
			"application/json"
		],
		"paths": {
			"/{DataStreamProtocol}/pullpushstream": {
				"post": {
					"operationId": "CreatePullPushDataStream",
					"responses": {
						"200": {
							"description": "A successful response.",
							"schema": {
								"$ref": "#/definitions/apiCreatePullPushDataStreamResponse"
							}
						}
					},
					"parameters": [
						{
							"name": "DataStreamProtocol",
							"in": "path",
							"required": true,
							"type": "string",
							"enum": [
								"KAFKA",
								"MQTT",
								"UNKNOWN"
							]
						},
						{
							"name": "body",
							"in": "body",
							"required": true,
							"schema": {
								"$ref": "#/definitions/apiCreatePullPushDataStreamRequest"
							}
						}
					],
					"tags": [
						"PullPushDataStream"
					]
				}
			}
		},
		"definitions": {
			"apiCreatePullPushDataStreamRequest": {
				"type": "object",
				"properties": {
					"DataContractTypeID": {
						"type": "string"
					},
					"DataStreamProtocol": {
						"$ref": "#/definitions/apiProtocol"
					},
					"DataStreamSourceURL": {
						"type": "string"
					},
					"DataStreamDestinationURL": {
						"type": "string"
					},
					"DataStreamTopic": {
						"type": "string"
					},
					"DataStreamParameters": {
						"type": "object",
						"additionalProperties": {
							"type": "string"
						},
						"title": "Duration StreamDuration = 5;"
					}
				}
			},
			"apiCreatePullPushDataStreamResponse": {
				"type": "object",
				"properties": {
					"DataStreamCreateStatus": {
						"type": "string"
					},
					"Error": {
						"$ref": "#/definitions/apiERROR"
					}
				}
			},
			"apiERROR": {
				"type": "object",
				"properties": {
					"Message": {
						"type": "string"
					}
				}
			},
			"apiProtocol": {
				"type": "string",
				"enum": [
					"KAFKA",
					"MQTT",
					"UNKNOWN"
				],
				"default": "KAFKA"
			}
		}
	}
	
	`
)
