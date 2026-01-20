// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/browserbase/stagehand-cli/internal/mocktest"
	"github.com/browserbase/stagehand-cli/internal/requestflag"
)

func TestSessionsAct(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "act",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--input", "Click the login button",
		"--frame-id", "frameId",
		"--options", "{model: {modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}, timeout: 30000, variables: {username: john_doe}}",
		"--stream-response=false",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsAct)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "act",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--input", "Click the login button",
		"--frame-id", "frameId",
		"--options.model", "{modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}",
		"--options.timeout", "30000",
		"--options.variables", "{username: john_doe}",
		"--stream-response=false",
		"--x-stream-response", "true",
	)
}

func TestSessionsEnd(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "end",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"---force-body", "{}",
		"--x-stream-response", "true",
	)
}

func TestSessionsExecute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "execute",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--agent-config", "{cua: true, model: {modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}, provider: openai, systemPrompt: systemPrompt}",
		"--execute-options", "{instruction: 'Log in with username ''demo'' and password ''test123'', then navigate to settings', highlightCursor: true, maxSteps: 20}",
		"--frame-id", "frameId",
		"--stream-response=false",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsExecute)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "execute",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--agent-config.cua=true",
		"--agent-config.model", "{modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}",
		"--agent-config.provider", "openai",
		"--agent-config.systemPrompt", "systemPrompt",
		"--execute-options.instruction", "Log in with username 'demo' and password 'test123', then navigate to settings",
		"--execute-options.highlightCursor=true",
		"--execute-options.maxSteps", "20",
		"--frame-id", "frameId",
		"--stream-response=false",
		"--x-stream-response", "true",
	)
}

func TestSessionsExtract(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "extract",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--frame-id", "frameId",
		"--instruction", "Extract all product names and prices from the page",
		"--options", "{model: {modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}, selector: '#main-content', timeout: 30000}",
		"--schema", "{foo: bar}",
		"--stream-response=false",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsExtract)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "extract",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--frame-id", "frameId",
		"--instruction", "Extract all product names and prices from the page",
		"--options.model", "{modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}",
		"--options.selector", "#main-content",
		"--options.timeout", "30000",
		"--schema", "{foo: bar}",
		"--stream-response=false",
		"--x-stream-response", "true",
	)
}

func TestSessionsNavigate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "navigate",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--url", "https://example.com",
		"--frame-id", "frameId",
		"--options", "{referer: referer, timeout: 30000, waitUntil: networkidle}",
		"--stream-response=true",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsNavigate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "navigate",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--url", "https://example.com",
		"--frame-id", "frameId",
		"--options.referer", "referer",
		"--options.timeout", "30000",
		"--options.waitUntil", "networkidle",
		"--stream-response=true",
		"--x-stream-response", "true",
	)
}

func TestSessionsObserve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "observe",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--frame-id", "frameId",
		"--instruction", "Find all clickable navigation links",
		"--options", "{model: {modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}, selector: nav, timeout: 30000}",
		"--stream-response=false",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsObserve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "observe",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--frame-id", "frameId",
		"--instruction", "Find all clickable navigation links",
		"--options.model", "{modelName: openai/gpt-5-nano, apiKey: sk-some-openai-api-key, baseURL: https://api.openai.com/v1, provider: openai}",
		"--options.selector", "nav",
		"--options.timeout", "30000",
		"--stream-response=false",
		"--x-stream-response", "true",
	)
}

func TestSessionsStart(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "start",
		"--model-name", "openai/gpt-4o",
		"--act-timeout-ms", "0",
		"--browser", "{cdpUrl: ws://localhost:9222, launchOptions: {acceptDownloads: true, args: [string], cdpUrl: cdpUrl, chromiumSandbox: true, connectTimeoutMs: 0, deviceScaleFactor: 0, devtools: true, downloadsPath: downloadsPath, executablePath: executablePath, hasTouch: true, headless: true, ignoreDefaultArgs: true, ignoreHTTPSErrors: true, locale: locale, preserveUserDataDir: true, proxy: {server: server, bypass: bypass, password: password, username: username}, userDataDir: userDataDir, viewport: {height: 0, width: 0}}, type: local}",
		"--browserbase-session-create-params", "{browserSettings: {advancedStealth: true, blockAds: true, context: {id: id, persist: true}, extensionId: extensionId, fingerprint: {browsers: [chrome], devices: [desktop], httpVersion: '1', locales: [string], operatingSystems: [android], screen: {maxHeight: 0, maxWidth: 0, minHeight: 0, minWidth: 0}}, logSession: true, recordSession: true, solveCaptchas: true, viewport: {height: 0, width: 0}}, extensionId: extensionId, keepAlive: true, projectId: projectId, proxies: true, region: us-west-2, timeout: 0, userMetadata: {foo: bar}}",
		"--browserbase-session-id", "browserbaseSessionID",
		"--dom-settle-timeout-ms", "5000",
		"--experimental=true",
		"--self-heal=true",
		"--system-prompt", "systemPrompt",
		"--verbose", "1",
		"--wait-for-captcha-solves=true",
		"--x-stream-response", "true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sessionsStart)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "start",
		"--model-name", "openai/gpt-4o",
		"--act-timeout-ms", "0",
		"--browser.cdpUrl", "ws://localhost:9222",
		"--browser.launchOptions", "{acceptDownloads: true, args: [string], cdpUrl: cdpUrl, chromiumSandbox: true, connectTimeoutMs: 0, deviceScaleFactor: 0, devtools: true, downloadsPath: downloadsPath, executablePath: executablePath, hasTouch: true, headless: true, ignoreDefaultArgs: true, ignoreHTTPSErrors: true, locale: locale, preserveUserDataDir: true, proxy: {server: server, bypass: bypass, password: password, username: username}, userDataDir: userDataDir, viewport: {height: 0, width: 0}}",
		"--browser.type", "local",
		"--browserbase-session-create-params.browserSettings", "{advancedStealth: true, blockAds: true, context: {id: id, persist: true}, extensionId: extensionId, fingerprint: {browsers: [chrome], devices: [desktop], httpVersion: '1', locales: [string], operatingSystems: [android], screen: {maxHeight: 0, maxWidth: 0, minHeight: 0, minWidth: 0}}, logSession: true, recordSession: true, solveCaptchas: true, viewport: {height: 0, width: 0}}",
		"--browserbase-session-create-params.extensionId", "extensionId",
		"--browserbase-session-create-params.keepAlive=true",
		"--browserbase-session-create-params.projectId", "projectId",
		"--browserbase-session-create-params.proxies=true",
		"--browserbase-session-create-params.region", "us-west-2",
		"--browserbase-session-create-params.timeout", "0",
		"--browserbase-session-create-params.userMetadata", "{foo: bar}",
		"--browserbase-session-id", "browserbaseSessionID",
		"--dom-settle-timeout-ms", "5000",
		"--experimental=true",
		"--self-heal=true",
		"--system-prompt", "systemPrompt",
		"--verbose", "1",
		"--wait-for-captcha-solves=true",
		"--x-stream-response", "true",
	)
}
