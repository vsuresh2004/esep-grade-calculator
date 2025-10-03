package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Errorf("expected assignment, got %s", Assignment.String())
    }
    if Exam.String() != "exam" {
        t.Errorf("expected exam, got %s", Exam.String())
    }
    if Essay.String() != "essay" {
        t.Errorf("expected essay, got %s", Essay.String())
    }
}

func TestGetGradeC(t *testing.T) {
    g := NewGradeCalculator()
    g.AddGrade("assignment", 70, Assignment)
    g.AddGrade("exam", 70, Exam)
    g.AddGrade("essay", 70, Essay)
    if g.GetFinalGrade() != "C" {
        t.Errorf("expected C, got %s", g.GetFinalGrade())
    }
}

func TestGetGradeD(t *testing.T) {
    g := NewGradeCalculator()
    g.AddGrade("assignment", 60, Assignment)
    g.AddGrade("exam", 60, Exam)
    g.AddGrade("essay", 60, Essay)
    if g.GetFinalGrade() != "D" {
        t.Errorf("expected D, got %s", g.GetFinalGrade())
    }
}

func TestComputeAverageEmpty(t *testing.T) {
    grades := []Grade{}
    avg := computeAverage(grades)
    if avg != 0 {
        t.Errorf("expected 0 for empty slice, got %d", avg)
    }
}
