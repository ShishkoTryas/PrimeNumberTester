package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckPrimesHandler(t *testing.T) {
	router := gin.New()
	router.POST("/checkprimes", checkPrimes)

	testCases := []struct {
		name             string
		input            CheckPrimesRequest
		expectedStatus   int
		expectedResponse CheckPrimesResponse
	}{
		{
			name: "empty",
			input: CheckPrimesRequest{
				Numbers: []int{},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: CheckPrimesResponse{
				Results: []bool{},
			},
		},
		{
			name: "primes and no primes",
			input: CheckPrimesRequest{
				Numbers: []int{2, 3, 4, 5, 6},
			},
			expectedStatus: http.StatusOK,
			expectedResponse: CheckPrimesResponse{
				Results: []bool{true, true, false, true, false},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputBytes, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest(http.MethodPost, "/checkprimes", bytes.NewBuffer(inputBytes))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status code %d, but got %d", tc.expectedStatus, w.Code)
			}

			if tc.expectedStatus == http.StatusOK {
				var res CheckPrimesResponse
				if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
					t.Errorf("error decoding: %v", err)
				}

				if len(res.Results) != len(tc.expectedResponse.Results) {
					t.Errorf("expected %d results, but %d", len(tc.expectedResponse.Results), len(res.Results))
				}
			}
		})
	}
}
