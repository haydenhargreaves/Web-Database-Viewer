package templates

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const MANAGER string = `
    <div id="manager-modal" class="fixed inset-0 bg-gray-600 bg-opacity-50">
      <div class="flex items-center justify-center min-h-screen">
        <form hx-post="/v1/api/connections/delete" hx-swap="outerHTML" hx-target="#manager-modal" hx-trigger="submit" class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-2/3">
          <div class="flex justify-between items-start border-b border-gray-200 dark:border-gray-700 pb-2">
            <h2 class="text-xl font-bold">
              Manage Stored Connections
              <br>
              <span class="text-xs font-light">Connection data is stored in the browsers session and can be renamed or deleted here.</span>
            </h2>
            <button hx-get="/v1/web/manager/hide" hx-trigger="click" hx-swap="outerHTML" hx-target="#manager-modal" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          <div class="mt-4 w-full max-w-full overflow-x-auto">
            <table>
              <thead class="bg-gray-50 dark:bg-gray-700">
                <tr>
									<th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Delete</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
									<th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase jracking-wider">Driver</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">URL</th>
                </tr>
              </thead>
              <tbody class="w-full">
								%s
              </tbody>
            </table>
          </div>
          <div class="flex items-center space-x-4 mt-4 border-t border-gray-200 dark:border-gray-700 pt-4">
            <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded-md"> Save Changes </button>
          </div>
        </form>
      </div>
    </div>
`

const MANAGER_ENTRY string = `
	<tr class="overflow-x-auto">
		<td class="px-6 py-4 whitespace-nowrap flex items-center justify-center">
			<input type="checkbox" name="connections" value="%s" class="w-4 h-4 m-2">
		</td>
		<td class="px-6 py-4 whitespace-nowrap">
			<input type="text" name="%s" value="%s" class="p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-gray-100">
		</td>
		<td class="px-6 py-4 whitespace-nowrap">%s</td>
		<td class="px-6 py-4 whitespace-nowrap">%s</td>
	</tr>
`

const MANAGER_CLOSED string = `
	<div id="manager-modal" class="fixed inset-0 bg-gray-600 bg-opacity-50 hidden opacity-0" hx-get="/v1/web/connections" hx-trigger="load" hx-swap="outerHTML" hx-target="#connected-database" hx-params="none"></div>
`

func OpenManager(c *gin.Context) {
	session := sessions.Default(c)
	connections_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		fmt.Println("No connections found")
	}

	var connections map[string][2]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		fmt.Println(err)
	}

	var entries string
	for name, conn := range connections {
		entries += fmt.Sprintf(MANAGER_ENTRY, conn[0], conn[0], name, conn[1], conn[0])
	}

	c.String(200, fmt.Sprintf(MANAGER, entries))
}

func HideManager(c *gin.Context) {
	c.String(200, MANAGER_CLOSED)
}
