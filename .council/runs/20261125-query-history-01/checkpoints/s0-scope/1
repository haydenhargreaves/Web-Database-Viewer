package templates

import "fmt"

// List item templates
const LIST_OPEN string = `<select id="connected-database" name="connected-database" class="mt-1 block p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm md:text-base dark:bg-gray-700 dark:border-gray-600 dark:text-gray-100" title="Change the connection to query.">`
const LIST_ITEM string = `<option value='{"driver": "%s", "url": "%s"}'%s>%s</option>`
const LIST_CLOSE string = `</select>`

const TREE_VIEW_NAME string = `<span hx-swap-oob="outerHTML" id="database-name-tree">%s</span>`

// Generate a list of connections to display in the drop-down.
// Current connection will be toggled as selected
func ConnectionsList(connections map[string][2]string, current string) string {
	var html string = LIST_OPEN

	if len(connections) == 0 || connections == nil {
		html += fmt.Sprintf(LIST_ITEM, "", "", " selected", "No connections")
		return html + LIST_CLOSE
	}

	html += fmt.Sprintf(LIST_ITEM, connections[current][1], connections[current][0], " selected", current)
	for _, name := range getSortedKeys(connections) {
		if name == current {
			continue
		} else {
			html += fmt.Sprintf(LIST_ITEM, connections[name][1], connections[name][0], "", name)
		}
	}

	treeName := fmt.Sprintf(TREE_VIEW_NAME, current)

	html += LIST_CLOSE
	return html + treeName
}
