/*
 * Copyright 2023 The AeroSquirrel Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package models provides models for the application.
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Resource is a resource.
type Resource struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamId    primitive.ObjectID `json:"teamId" bson:"teamId"`
	Name      string             `json:"name" bson:"name"`
	Type      string             `json:"type" bson:"type"`
	Cost      float64            `json:"cost" bson:"cost"`
	Metadata  map[string]string  `json:"metadata" bson:"metadata"`
	Tags      []Tag              `json:"tags" bson:"tags"`
	Link      string             `json:"link" bson:"link"`
	Provider  string             `json:"provider" bson:"provider"`
	Region    string             `json:"region" bson:"region"`
	Account   string             `json:"account" bson:"account"`
	AccountId string             `json:"accountId" bson:"accountId"`
}

// Tag is a tag.
type Tag struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
