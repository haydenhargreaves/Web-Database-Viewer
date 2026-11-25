package templates

import (
	"crypto/sha256"
	"fmt"
	"html"
)

// QueryHistoryID identifies the exact connection name without exposing it.
func QueryHistoryID(name string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(name)))
}

// List item templates
const LIST_OPEN string = `<select id="connected-database" name="connected-database" class="mt-1 block p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm md:text-base dark:bg-gray-700 dark:border-gray-600 dark:text-gray-100" title="Change the connection to query.">`
const LIST_ITEM string = `<option value='{"driver": "%s", "url": "%s"}'%s>%s</option>`
const LIST_CLOSE string = `</select>`

const TREE_VIEW_NAME string = `<span hx-swap-oob="outerHTML" id="database-name-tree">%s</span>`

func connectionItem(data [2]string, name string, selected, member bool) string {
	attributes := ""
	if selected {
		attributes += " selected"
	}
	if member {
		attributes += ` data-query-history-id="` + QueryHistoryID(name) + `"`
	}
	return fmt.Sprintf(LIST_ITEM, data[1], data[0], attributes, html.EscapeString(name))
}

// Generate a list of connections to display in the drop-down.
// Current connection will be toggled as selected
func ConnectionsList(connections map[string][2]string, current string) string {
	var html string = LIST_OPEN

	if len(connections) == 0 || connections == nil {
		html += fmt.Sprintf(LIST_ITEM, "", "", " selected", "No connections")
		return html + LIST_CLOSE
	}

	_, member := connections[current]
	html += connectionItem(connections[current], current, true, member)
	for _, name := range getSortedKeys(connections) {
		if name == current {
			continue
		} else {
			html += connectionItem(connections[name], name, false, true)
		}
	}

	treeName := fmt.Sprintf(TREE_VIEW_NAME, htmlEscape(current))

	html += LIST_CLOSE
	return html + treeName
}

func htmlEscape(s string) string { return html.EscapeString(s) }
