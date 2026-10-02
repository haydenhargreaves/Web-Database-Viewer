package templates

import "fmt"

const TABLE_TREE_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Tables: %s</p>
	</li>
`

const TABLE_FIELDS_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Table: %s</p>
	</li>
`

const ENUM_TREE_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Tables: %s</p>
	</li>
`

const ROUTINE_TREE_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Procedures and Functions: %s</p>
	</li>
`

// When an error occurs while generating the table tree,
// this function will return the HTML for the error message.
func TableTreeError(err error) string {
	var html string = TABLE_TREE_OPEN
	html += fmt.Sprintf(TABLE_TREE_ERROR_BODY, err.Error())
	return html + TABLE_TREE_CLOSE
}

// TableFieldsError returns an error suitable for insertion into one table's field list.
func TableFieldsError(err error) string {
	return fmt.Sprintf(TABLE_FIELDS_ERROR_BODY, err.Error())
}

// When an error occurs while generating the enum tree,
// this function will return the HTML for the error message.
func EnumTreeError(err error) string {
	var html string = ENUM_TREE_OPEN
	html += fmt.Sprintf(ENUM_TREE_ERROR_BODY, err.Error())
	return html + ENUM_TREE_CLOSE
}

// RoutineTreeError returns an error suitable for the routine tree.
func RoutineTreeError(err error) string {
	return ROUTINE_TREE_BODY_OPEN + fmt.Sprintf(ROUTINE_TREE_ERROR_BODY, err.Error()) + ROUTINE_TREE_CLOSE
}
