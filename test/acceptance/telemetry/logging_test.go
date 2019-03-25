/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package test

// Acceptance tests for logging. Sets up a small fake endpoint that logs are sent to.

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/creativesoftwarefdn/weaviate/client/actions"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)

// Ensure a request should be logged, get the most recently received log from the mock api, interpret it
func TestCreateActionLogging(t *testing.T) {
	t.Parallel()

	// send a request
	sendCreateActionRequest(t)

	// wait for the log to be posted
	time.Sleep(3 * time.Second)

	result := retrieveLogFromMockEndpoint(t)

	if result != nil {
		interpretedResult := interpretResult(t, result)
		if interpretedResult != nil {
			_, namePresent := interpretedResult["n"]
			_, typePresent := interpretedResult["t"]
			_, identifierPresent := interpretedResult["i"]
			_, amountPresent := interpretedResult["a"]
			_, whenPresent := interpretedResult["w"]

			assert.Equal(t, true, namePresent)
			assert.Equal(t, true, typePresent)
			assert.Equal(t, true, identifierPresent)
			assert.Equal(t, true, amountPresent)
			assert.Equal(t, true, whenPresent)

		}
	}
}

// sendCreateActionRequest is copied here to ensure at least one request should be logged when we check the mock api's most recently received request
func sendCreateActionRequest(t *testing.T) {
	// Set all action values to compare
	actionTestString := "Test string"
	actionTestInt := 1
	actionTestBoolean := true
	actionTestNumber := 1.337
	actionTestDate := "2017-10-06T08:15:30+01:00"

	params := actions.NewWeaviateActionsCreateParams().WithBody(actions.WeaviateActionsCreateBody{
		Action: &models.ActionCreate{
			AtContext: "http://example.org",
			AtClass:   "TestAction",
			Schema: map[string]interface{}{
				"testString":   actionTestString,
				"testInt":      actionTestInt,
				"testBoolean":  actionTestBoolean,
				"testNumber":   actionTestNumber,
				"testDateTime": actionTestDate,
			},
		},
	})

	resp, _, err := helper.Client(t).Actions.WeaviateActionsCreate(params, nil)

	// Ensure that the response is OK
	helper.AssertRequestOk(t, resp, err, func() {
		action := resp.Payload
		assert.Regexp(t, strfmt.UUIDPattern, action.ActionID)

		schema, ok := action.Schema.(map[string]interface{})
		if !ok {
			t.Fatal("The returned schema is not an JSON object")
		}

		// Check whether the returned information is the same as the data added
		assert.Equal(t, actionTestString, schema["testString"])
		assert.Equal(t, actionTestInt, int(schema["testInt"].(float64)))
		assert.Equal(t, actionTestBoolean, schema["testBoolean"])
		assert.Equal(t, actionTestNumber, schema["testNumber"])
		assert.Equal(t, actionTestDate, schema["testDateTime"])
	})
}

// retrieveLogFromMockEndpoint retrieves the most recently received log from the mock api
func retrieveLogFromMockEndpoint(t *testing.T) []byte {
	testURL, err := url.Parse("http://127.0.0.1:8087/mock/last")
	assert.Equal(t, nil, err)

	client := &http.Client{}
	resp, err := client.Get(testURL.String())
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return body
	}

	assert.Equal(t, nil, err)

	return nil
}

// interpretResult converts the received cbor-encoded log to a map[string]interface
func interpretResult(t *testing.T, resultBody []byte) map[string]interface{} {
	decoded := make(map[string]interface{}, 0)
	cborHandle := new(codec.CborHandle)
	encoder := codec.NewDecoderBytes(resultBody, cborHandle)
	err := encoder.Decode(decoded)

	assert.Equal(t, nil, err)

	if err == nil {
		return decoded
	}
	return nil
}