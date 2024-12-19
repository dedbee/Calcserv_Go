package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dedbee/Calcserv_Go/internal/application"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		expression     string
		expectedStatus int
		expectedResult json.RawMessage
	}{
		{
			name:           "Valid expression",
			expression:     "2 + 2",
			expectedStatus: http.StatusOK,
			expectedResult: json.RawMessage(`{"result":4}`),
		},
		{
			name:           "Division by zero",
			expression:     "1 / 0",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: json.RawMessage(`{"error":"Division by zero"}`),
		},
		{
			name:           "Invalid expression",
			expression:     "2 +",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedResult: json.RawMessage(`{"error":"Expression is not valid"}`),
		},
		{
			name:           "Malformed JSON",
			expression:     "",
			expectedStatus: http.StatusInternalServerError,
			expectedResult: json.RawMessage(`{"error":"Internal server error"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			if tt.name != "Malformed JSON" {
				requestBody := application.Request{Expression: tt.expression}
				body, _ = json.Marshal(requestBody)
			}
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			application.CalcHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response json.RawMessage
			json.NewDecoder(w.Body).Decode(&response)

			if !reflect.DeepEqual(response, tt.expectedResult) {
				t.Errorf("expected response %s, got %s", tt.expectedResult, response)
			}
		})
	}
}
