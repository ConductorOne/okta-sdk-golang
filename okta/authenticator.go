/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"context"
	"fmt"
	"time"
)

type AuthenticatorResource resource

type Authenticator struct {
	Links       interface{}                     `json:"_links,omitempty"`
	Created     *time.Time                      `json:"created,omitempty"`
	Id          string                          `json:"id,omitempty"`
	Key         string                          `json:"key,omitempty"`
	LastUpdated *time.Time                      `json:"lastUpdated,omitempty"`
	Name        string                          `json:"name,omitempty"`
	Provider    *OnpremMfaAuthenticatorProvider `json:"provider,omitempty"`
	Settings    *AuthenticatorSettings          `json:"settings,omitempty"`
	Status      string                          `json:"status,omitempty"`
	Type        string                          `json:"type,omitempty"`
}

func (m *AuthenticatorResource) GetAuthenticator(ctx context.Context, authenticatorId string) (*Authenticator, *Response, error) {
	url := fmt.Sprintf("/api/v1/authenticators/%v", authenticatorId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authenticator *Authenticator

	resp, err := rq.Do(ctx, req, &authenticator)
	if err != nil {
		return nil, resp, err
	}

	return authenticator, resp, nil
}

// Updates an authenticator
func (m *AuthenticatorResource) UpdateAuthenticator(ctx context.Context, authenticatorId string, body Authenticator) (*Authenticator, *Response, error) {
	url := fmt.Sprintf("/api/v1/authenticators/%v", authenticatorId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authenticator *Authenticator

	resp, err := rq.Do(ctx, req, &authenticator)
	if err != nil {
		return nil, resp, err
	}

	return authenticator, resp, nil
}

func (m *AuthenticatorResource) ListAuthenticators(ctx context.Context) ([]*Authenticator, *Response, error) {
	url := fmt.Sprintf("/api/v1/authenticators")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authenticator []*Authenticator

	resp, err := rq.Do(ctx, req, &authenticator)
	if err != nil {
		return nil, resp, err
	}

	return authenticator, resp, nil
}

func (m *AuthenticatorResource) ActivateAuthenticator(ctx context.Context, authenticatorId string) (*Authenticator, *Response, error) {
	url := fmt.Sprintf("/api/v1/authenticators/%v/lifecycle/activate", authenticatorId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authenticator *Authenticator

	resp, err := rq.Do(ctx, req, &authenticator)
	if err != nil {
		return nil, resp, err
	}

	return authenticator, resp, nil
}

func (m *AuthenticatorResource) DeactivateAuthenticator(ctx context.Context, authenticatorId string) (*Authenticator, *Response, error) {
	url := fmt.Sprintf("/api/v1/authenticators/%v/lifecycle/deactivate", authenticatorId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authenticator *Authenticator

	resp, err := rq.Do(ctx, req, &authenticator)
	if err != nil {
		return nil, resp, err
	}

	return authenticator, resp, nil
}
