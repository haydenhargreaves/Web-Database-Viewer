package templates

import (
	"strings"
	"testing"

	"github.com/Azpect3120/Web-Database-Viewer/internal/model"
)

func TestRoutineTreeRendersEscapedLoadControls(t *testing.T) {
	routines := []model.Routine{{
		ID:        "42",
		Name:      `routine"name`,
		Signature: `public.routine"name(integer)`,
		Kind:      "function",
	}}

	output := RoutineTree(routines)
	for _, expected := range []string{
		`id="database-routine-tree"`,
		`data-routine-id="42"`,
		`data-routine-name="routine&#34;name"`,
		`data-routine-kind="function"`,
		`LoadRoutineDefinition(this);`,
		`public.routine&#34;name(integer)`,
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("RoutineTree() missing %q", expected)
		}
	}
}

func TestViewTreeRendersSortedEscapedLoadControls(t *testing.T) {
	output := ViewTree([]string{`z_view`, `a"view`})
	for _, expected := range []string{
		`id="database-view-tree"`,
		`data-view-name="a&#34;view"`,
		`LoadViewQuery(this.dataset.viewName);`,
		`a&#34;view`,
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("ViewTree() missing %q", expected)
		}
	}

	if strings.Index(output, `a&#34;view`) > strings.Index(output, `z_view`) {
		t.Error("ViewTree() did not sort views")
	}
}
