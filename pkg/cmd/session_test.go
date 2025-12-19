// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/browserbase/stagehand-cli/internal/mocktest"
)

func TestSessionsAct(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "act",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--input", "Click the login button",
		"--frame-id", "frameId",
		"--options", "{model: openai/gpt-5-nano, timeout: 30000, variables: {username: john_doe}}",
		"--stream-response",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
		"--x-stream-response", "true",
	)
}

func TestSessionsEnd(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "end",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
		"--x-stream-response", "true",
	)
}

func TestSessionsExecute(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "execute",
		"--id", "c4dbf3a9-9a58-4b22-8a1c-9f20f9f9e123",
		"--agent-config", "{cua: true, model: openai/gpt-5-nano, systemPrompt: systemPrompt}",
		"--execute-options", "{instruction: 'Log in with username ''demo'' and password ''test123'', then navigate to settings', highlightCursor: true, maxSteps: 20}",
		"--frame-id", "frameId",
		"--stream-response",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
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
		"--options", "{model: openai/gpt-5-nano, selector: '#main-content', timeout: 30000}",
		"--schema", "{foo: bar}",
		"--stream-response",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
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
		"--stream-response",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
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
		"--options", "{model: openai/gpt-5-nano, selector: nav, timeout: 30000}",
		"--stream-response",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
		"--x-stream-response", "true",
	)
}

func TestSessionsStart(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sessions", "start",
		"--model-name", "gpt-4o",
		"--act-timeout-ms", "30000",
		"--browser", "{cdpUrl: ws://localhost:9222, launchOptions: {acceptDownloads: true, args: [string], cdpUrl: cdpUrl, chromiumSandbox: true, connectTimeoutMs: 0, deviceScaleFactor: 0, devtools: true, downloadsPath: downloadsPath, executablePath: executablePath, hasTouch: true, headless: true, ignoreDefaultArgs: true, ignoreHTTPSErrors: true, locale: locale, preserveUserDataDir: true, proxy: {server: server, bypass: bypass, password: password, username: username}, userDataDir: userDataDir, viewport: {height: 0, width: 0}}, type: local}",
		"--browserbase-session-create-params", "{browserSettings: {advancedStealth: true, blockAds: true, context: {id: id, persist: true}, extensionId: extensionId, fingerprint: {browsers: [chrome], devices: [desktop], httpVersion: '1', locales: [string], operatingSystems: [android], screen: {maxHeight: 0, maxWidth: 0, minHeight: 0, minWidth: 0}}, logSession: true, recordSession: true, solveCaptchas: true, viewport: {height: 0, width: 0}}, extensionId: extensionId, keepAlive: true, projectId: projectId, proxies: true, region: us-west-2, timeout: 0, userMetadata: {foo: bar}}",
		"--browserbase-session-id", "browserbaseSessionID",
		"--debug-dom",
		"--dom-settle-timeout-ms", "5000",
		"--experimental",
		"--self-heal",
		"--system-prompt", "systemPrompt",
		"--verbose", "1",
		"--wait-for-captcha-solves",
		"--x-language", "typescript",
		"--x-sdk-version", "3.0.6",
		"--x-sent-at", "2025-01-15T10:30:00Z",
		"--x-stream-response", "true",
	)
}
