// Copyright 2023 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/apigee/registry/pkg/encoding"
)

var out = "apis/kubernetes.io"
var source = "import-kubernetes"

func main() {

	updated := time.Now().Format("2006-01-02")

	apiID := "kubernetes"

	api := &encoding.Api{
		Header: encoding.Header{
			ApiVersion: "apigeeregistry/v1",
			Kind:       "API",
			Metadata: encoding.Metadata{
				Name: apiID,
				Labels: map[string]string{
					"categories": "computing",
					"provider":   "kubernetes-io",
					"updated":    updated,
					"source":     source,
				},
			},
		},
		Data: encoding.ApiData{
			DisplayName: "Kubernetes",
			Description: "The Kubernetes API",
		},
	}

	// major, minor, latest patch
	versions := [][]int{
		{1, 5, 8},
		{1, 6, 13},
		{1, 7, 16},
		{1, 8, 15},
		{1, 9, 11},
		{1, 10, 13},
		{1, 11, 10},
		{1, 12, 10},
		{1, 13, 12},
		{1, 14, 10},
		{1, 15, 12},
		{1, 16, 15},
		{1, 17, 17},
		{1, 18, 20},
		{1, 19, 16},
		{1, 20, 15},
		{1, 21, 14},
		{1, 22, 17},
		{1, 23, 17},
		{1, 24, 12},
		{1, 25, 8},
		{1, 26, 3},
	}

	for _, version := range versions {
		for patch := 0; patch <= version[2]; patch++ {
			versionID := fmt.Sprintf("v%d.%d.%d", version[0], version[1], patch)
			sortableVersionID := fmt.Sprintf("v%d.%02d.%02d", version[0], version[1], patch)
			sourceURI := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/%s/api/openapi-spec/swagger.json", versionID)
			api.Data.ApiVersions = append(api.Data.ApiVersions,
				&encoding.ApiVersion{
					Header: encoding.Header{
						Metadata: encoding.Metadata{
							Name: sortableVersionID,
							Labels: map[string]string{
								"updated": updated,
								"source":  source,
							},
						},
					},
					Data: encoding.ApiVersionData{
						DisplayName: sortableVersionID,
						ApiSpecs: []*encoding.ApiSpec{
							{
								Header: encoding.Header{
									Metadata: encoding.Metadata{
										Name: "swagger",
										Labels: map[string]string{
											"updated": updated,
											"source":  source,
										},
									},
								},
								Data: encoding.ApiSpecData{
									FileName:  "swagger.json",
									MimeType:  "application/x.openapi+gzip;version=2.0",
									SourceURI: sourceURI,
								},
							},
						},
					},
				},
			)
		}
	}

	err := os.MkdirAll(filepath.Join(out, apiID), 0777)
	if err != nil {
		log.Fatalf("%s", err)
	}

	b, err := encoding.EncodeYAML(api)
	if err != nil {
		return
	}
	err = os.WriteFile(filepath.Join(out, apiID, "info.yaml"), b, 0666)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
