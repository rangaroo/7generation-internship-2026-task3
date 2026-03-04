package main

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	if !isLetter('a') || !isLetter('z') || !isLetter('A') || !isLetter('Z') {
		t.Error("should recognize letters")
	}
	if isLetter('0') || isLetter(' ') || isLetter('\n') {
		t.Error("should reject non-letters")
	}
}

func TestToLower(t *testing.T) {
	if toLower('A') != 'a' || toLower('Z') != 'z' {
		t.Error("should convert uppercase to lowercase")
	}
	if toLower('a') != 'a' {
		t.Error("should keep lowercase as is")
	}
}

func TestExtractWords(t *testing.T) {
	data := []byte("Hello World")
	words := extractWords(data)

	if len(words) != 2 {
		t.Errorf("expected 2 words, got %d", len(words))
	}
	if !bytesEqual(words[0], []byte("hello")) {
		t.Errorf("expected 'hello', got '%s'", words[0])
	}
	if !bytesEqual(words[1], []byte("world")) {
		t.Errorf("expected 'world', got '%s'", words[1])
	}
}

func TestExtractWordsWithNumbers(t *testing.T) {
	data := []byte("test123word")
	words := extractWords(data)

	if len(words) != 2 {
		t.Errorf("expected 2 words, got %d", len(words))
	}
}

func TestBytesEqual(t *testing.T) {
	if !bytesEqual([]byte("abc"), []byte("abc")) {
		t.Error("equal slices should match")
	}
	if bytesEqual([]byte("abc"), []byte("abd")) {
		t.Error("different slices should not match")
	}
	if bytesEqual([]byte("ab"), []byte("abc")) {
		t.Error("different lengths should not match")
	}
}

func TestCountWords(t *testing.T) {
	words := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("hello"),
	}
	counts := countWords(words)

	if len(counts) != 2 {
		t.Errorf("expected 2 unique words, got %d", len(counts))
	}

	for _, wc := range counts {
		if bytesEqual(wc.word, []byte("hello")) && wc.count != 2 {
			t.Errorf("expected hello count 2, got %d", wc.count)
		}
		if bytesEqual(wc.word, []byte("world")) && wc.count != 1 {
			t.Errorf("expected world count 1, got %d", wc.count)
		}
	}
}

func TestEmptyInput(t *testing.T) {
	words := extractWords([]byte{})
	if len(words) != 0 {
		t.Error("empty input should produce no words")
	}
}

func TestBinaryData(t *testing.T) {
	data := []byte{0x00, 0xFF, 'h', 'i', 0x00, 0x01}
	words := extractWords(data)

	if len(words) != 1 || !bytesEqual(words[0], []byte("hi")) {
		t.Error("should handle binary data safely")
	}
}
