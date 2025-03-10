// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package index

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/game/mocks"
	"testing"
)

func TestIndex(t *testing.T) {

	tests := []struct {
		name          string
		requestMethod bool // true for GET, false otherwise
		expectError   bool
	}{
		{"GET Request", true, false},
		{"Non-GET Request", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResponse := &mocks.MockResponse{}
			mockRequest := &mocks.MockRequest{IsGet: tt.requestMethod}
			mockServer := mocks.NewMockServer()

			Index(mockResponse, mockRequest, mockServer)

			if tt.expectError {
				if !mockResponse.MethodNotAllowed {
					t.Errorf("Expected method not allowed error, but didn't get one")
				}
			} else {
				if mockResponse.SentStatusCode != 200 {
					t.Errorf("Expected status code 200, got %d", mockResponse.SentStatusCode)
				}
				if data, ok := mockResponse.SentData.(dtos.IndexDTO); !ok || data.Version != "0.0.1" {
					t.Errorf("Expected data version 0.0.1, got %v", mockResponse.SentData)
				}
			}
		})
	}

}
