package main

import "testing"

func TestSortLines(t *testing.T) {
	// Проверка стандартной работы
	lines := []string{"Line 3", "Line 1", "Line 2"}
	sortedLines := sortLines(lines, 0, false, false, false)
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	if len(sortedLines) != len(expectedLines) || sortedLines[0] != expectedLines[0] || sortedLines[1] != expectedLines[1] || sortedLines[2] != expectedLines[2] {
		t.Errorf("Expected %v, got %v", expectedLines, sortedLines)
	}

	// Проверка сотировки по числовому значению
	lines = []string{"5", "1", "3"}
	sortedLines = sortLines(lines, 0, true, false, false)
	expectedLines = []string{"1", "3", "5"}
	if len(sortedLines) != len(expectedLines) || sortedLines[0] != expectedLines[0] || sortedLines[1] != expectedLines[1] || sortedLines[2] != expectedLines[2] {
		t.Errorf("Expected %v, got %v", expectedLines, sortedLines)
	}

	// Проверка сортировки в обратном порядке
	lines = []string{"Line 1", "Line 2", "Line 3"}
	sortedLines = sortLines(lines, 0, false, true, false)
	expectedLines = []string{"Line 3", "Line 2", "Line 1"}
	if len(sortedLines) != len(expectedLines) || sortedLines[0] != expectedLines[0] || sortedLines[1] != expectedLines[1] || sortedLines[2] != expectedLines[2] {
		t.Errorf("Expected %v, got %v", expectedLines, sortedLines)
	}

	// Проверка удаления дубликатов
	lines = append(lines, "Line 1")
	sortedLines = sortLines(lines, 0, false, false, true)
	expectedLines = []string{"Line 1", "Line 2", "Line 3"}
	if len(sortedLines) != len(expectedLines) || sortedLines[0] != expectedLines[0] || sortedLines[1] != expectedLines[1] || sortedLines[2] != expectedLines[2] {
		t.Errorf("Expected %v, got %v", expectedLines, sortedLines)
	}
}
