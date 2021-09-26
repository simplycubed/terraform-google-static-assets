package test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------
// When testing against dev-sandbox account, use these, instead
// ----------------------------------------------------------
//const ROOT_DOMAIN_NAME_FOR_TEST = "gcloud-dev.com"
//const MANAGED_ZONE_NAME_FOR_TEST = "gclouddev"

// const ROOT_DOMAIN_NAME_FOR_TEST = "simplycubed.net"
const MANAGED_ZONE_NAME_FOR_TEST = "terra-test"

const KEY_PROJECT = "project"
const KEY_DOMAIN_NAME = "domain-name"

const OUTPUT_LB_IP_ADDRESS = "load_balancer_ip_address"

const EXAMPLE_NAME_LB_SITE = "http-load-balancer-website"
const EXAMPLE_NAME_STATIC_SITE = "cloud-storage-static-website"

const TEST_WEBSITE_MAX_RETRIES = 10
const TEST_WEBSITE_SLEEP_TIME = 30 * time.Second

func testWebsite(t *testing.T, protocol string, url string, path string, expectedStatus int, expectedBody string) {
	finalUrl := fmt.Sprintf("%s://%s%s", protocol, url, path)
	// Resource propagation takes long, so we'll allow a lot of time to survive that
	err := HttpGetWithRetryE(t, finalUrl, expectedStatus, expectedBody, TEST_WEBSITE_MAX_RETRIES, TEST_WEBSITE_SLEEP_TIME)
	assert.NoError(t, err, "Failed to call URL %s", url)
}

// A lot of this is repetition from terratest http_helper, but to allow the custom TLS Config, we're
// implementing the methods here, instead.
// TODO: Look into possibility of incorporating the TLS flag into terratest

func HttpGetE(t *testing.T, url string) (int, string, error) {
	logger.Logf(t, "Making a GET call to URL %s", url)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	client := &http.Client{
		Transport: transCfg,
		// By default, Go does not impose a timeout, so an HTTP connection attempt can hang for a LONG time.
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return -1, "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, strings.TrimSpace(string(body)), nil
}

func HttpGetWithRetryE(t *testing.T, url string, expectedStatus int, expectedBody string, retries int, sleepBetweenRetries time.Duration) error {
	_, err := retry.DoWithRetryE(t, fmt.Sprintf("HTTP GET to URL %s", url), retries, sleepBetweenRetries, func() (string, error) {
		return "", HttpGetWithValidationE(t, url, expectedStatus, expectedBody)
	})

	return err
}

// HttpGetWithValidationE performs an HTTP GET on the given URL and verify that you get back the expected status code and body. If either
// doesn't match, return an error.
func HttpGetWithValidationE(t *testing.T, url string, expectedStatusCode int, expectedBody string) error {
	return HttpGetWithCustomValidationE(t, url, func(statusCode int, body string) bool {
		return statusCode == expectedStatusCode && body == expectedBody
	})
}

// HttpGetWithCustomValidationE performs an HTTP GET on the given URL and validate the returned status code and body using the given function.
func HttpGetWithCustomValidationE(t *testing.T, url string, validateResponse func(int, string) bool) error {
	statusCode, body, err := HttpGetE(t, url)

	if err != nil {
		return err
	}

	if !validateResponse(statusCode, body) {
		return ValidationFunctionFailed{Url: url, Status: statusCode, Body: body}
	}

	return nil
}

// ValidationFunctionFailed is an error that occurs if a validation function fails.
type ValidationFunctionFailed struct {
	Url    string
	Status int
	Body   string
}

func (err ValidationFunctionFailed) Error() string {
	return fmt.Sprintf("Validation failed for URL %s. Response status: %d. Response body:\n%s", err.Url, err.Status, err.Body)
}
