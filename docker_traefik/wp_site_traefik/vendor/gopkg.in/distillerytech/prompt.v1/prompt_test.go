package prompt

import (
	"bytes"
	"fmt"
	"testing"
)

// TestAskString tests AskString() with input
func TestAskString(t *testing.T) {
	inBuf, outBuf := newTestBuffers()

	resp := "input"
	inBuf.WriteString(resp + "\n")

	recvResp := AskString("Test:")
	validateResp(resp, recvResp, t)

	expected := "Test: "
	validateBuf(outBuf, expected, t)
}

// TestAskString tests AskString() with empty input
func TestAskStringEmpty(t *testing.T) {
	inBuf, outBuf := newTestBuffers()

	resp := ""
	inBuf.WriteString(resp + "\n")

	recvResp := AskString("Test:")
	validateResp(resp, recvResp, t)

	expected := "Test: "
	validateBuf(outBuf, expected, t)
}

// TestAskStringRequired tests AskStringRequired()
func TestAskStringRequired(t *testing.T) {
	inBuf, outBuf := newTestBuffers()

	resp := "input"
	inBuf.WriteString("\n")
	inBuf.WriteString(resp + "\n")

	recvResp := AskStringRequired("Test:")
	validateResp(resp, recvResp, t)

	expected := "Test: Test: "
	validateBuf(outBuf, expected, t)
}

func TestAskStringLimit(t *testing.T) {
	inBuf, outBuf := newTestBuffers()

	resp := "valid"
	inBuf.WriteString("invalid\n")
	inBuf.WriteString(resp + "\n")

	recvResp := AskStringLimit("AskStringLimit:", "valid", "valid1")
	validateResp(resp, recvResp, t)

	expected := "AskStringLimit: AskStringLimit: "
	validateBuf(outBuf, expected, t)
}

func TestAskInteger(t *testing.T) {
	inBuf, outBuf := newTestBuffers()
	resp := 2

	inBuf.WriteString(fmt.Sprintf("%d", resp))

	recvResp := AskInteger("AskInteger:")
	validateResp(resp, recvResp, t)

	expected := "AskInteger: "
	validateBuf(outBuf, expected, t)
}

func TestAskIntegerInvalid(t *testing.T) {
	inBuf, outBuf := newTestBuffers()
	resp := 2

	inBuf.WriteString("invalid\n")
	inBuf.WriteString(fmt.Sprintf("%d", resp))

	recvResp := AskInteger("AskInteger:")
	validateResp(resp, recvResp, t)

	expected := "AskInteger: AskInteger: "
	validateBuf(outBuf, expected, t)
}

func TestAskIntegerDefault(t *testing.T) {
	inBuf, outBuf := newTestBuffers()

	inBuf.WriteString("invalid\n")

	resp := 2
	recvResp := AskIntegerDefault("AskIntegerDefault:", resp)

	validateResp(resp, recvResp, t)

	expected := "AskIntegerDefault: "
	validateBuf(outBuf, expected, t)
}

func TestConfirmInvalid(t *testing.T) {
	inBuf, outBuf := newTestBuffers()
	inBuf.WriteString("invalid\n")
	inBuf.WriteString("y")

	recvResp := Confirm("Confirm:")

	validateResp(true, recvResp, t)

	expected := "Confirm: Confirm: "
	validateBuf(outBuf, expected, t)
}

func TestConfirmValid(t *testing.T) {
	trues := []string{"y", "Y", "yes", "Yes"}
	falses := []string{"n", "N", "no", "No"}

	for _, input := range trues {
		testConfirm(input, true, t)
	}

	for _, input := range falses {
		testConfirm(input, false, t)
	}
}

func testConfirm(input string, expectedResp bool, t *testing.T) {
	inBuf, outBuf := newTestBuffers()
	inBuf.WriteString(input)
	recvResp := Confirm("Confirm:")
	validateResp(expectedResp, recvResp, t)
	expectedBuf := "Confirm: "
	validateBuf(outBuf, expectedBuf, t)
}

func newTestBuffers() (*bytes.Buffer, *bytes.Buffer) {
	return newTestInBuffer(), newTestOutBuffer()
}

func newTestInBuffer() *bytes.Buffer {
	var inBuf bytes.Buffer
	in = &inBuf

	return &inBuf
}

func newTestOutBuffer() *bytes.Buffer {
	var outBuf bytes.Buffer
	out = &outBuf

	return &outBuf
}

func validateResp(expectedResp, recvResp interface{}, t *testing.T) {
	if recvResp != expectedResp {
		t.Fatalf("Expected %v but received %v", expectedResp, recvResp)
	}
}

func validateBuf(outBuf *bytes.Buffer, expected string, t *testing.T) {
	if outBuf.String() != expected {
		t.Errorf("Output was %v but should have been %v", outBuf, expected)
	}
}
