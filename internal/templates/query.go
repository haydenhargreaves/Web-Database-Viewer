package templates

import (
	"github.com/gin-gonic/gin"
)

const MANUAL_QUERY string = `
	<div id="query-main" class="mb-4">
		<div class="flex items-center justify-between">
			<label for="sql" class="block text-sm font-medium text-gray-700 dark:text-gray-200">SQL Query</label>
			<div class="flex items-center space-x-6">
				<form class="flex items-center space-x-2" hx-get="/v1/web/query/auto" hx-swap="outerHTML" hx-target="#query-main" hx-trigger="input">
					<label for="auto-toggle" class="text-sm font-medium text-gray-700 dark:text-gray-200">Auto-Run</label>
					<input type="checkbox" name="toggle" class="toggle-checkbox">
				</form>
				<button hx-post="/v1/api/query" hx-trigger="click" hx-swap="outerHTML" hx-target="#query-results" hx-indicator="#spinner" hx-include="#sql" class="bg-blue-500 text-white py-2 px-3 rounded-md text-xs md:text-sm">Run Query</button>
			</div>
		</div>
		<textarea id="sql" name="sql" rows="12" class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm dark:bg-gray-800 dark:border-gray-600 dark:text-gray-100"></textarea>
		<p id="spinner" class="text-gray-700 dark:text-gray-300 text-sm px-1 py-2 htmx-indicator opacity-0"> Query running... </p>
		<p id="query-error" class="text-red-500 py-2 text-sm hidden"></p>
	</div>
`

const AUTO_QUERY string = `
	<div id="query-main" class="mb-4">
		<div class="flex items-center justify-between">
			<label for="sql" class="block text-sm font-medium text-gray-700 dark:text-gray-200">SQL Query</label>
			<div class="flex items-center space-x-6">
				<form class="flex items-center space-x-2" hx-get="/v1/web/query/auto" hx-swap="outerHTML" hx-target="#query-main" hx-trigger="input">
					<label for="auto-toggle" class="text-sm font-medium text-gray-700 dark:text-gray-200">Auto-Run</label>
					<input type="checkbox" name="toggle" class="toggle-checkbox" checked>
				</form>
				<button class="bg-blue-500 text-white py-2 px-3 rounded-md text-xs md:text-sm opacity-60 cursor-default" title="Disable Auto-Run to use manual queries." disabled>Run Query</button>
			</div>
		</div>
		<textarea id="sql" name="sql" rows="12" class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm dark:bg-gray-800 dark:border-gray-600 dark:text-gray-100" hx-post="/v1/api/query" hx-trigger="input delay:500ms" hx-swap="outerHTML" hx-target="#query-results" hx-indicator="#spinner"></textarea>
		<p id="spinner" class="text-gray-700 dark:text-gray-300 text-sm px-1 py-2 htmx-indicator opacity-0"> Query running... </p>
		<p id="query-error" class="text-red-500 py-2 text-sm hidden"></p>
	</div>
`

func ToggleQueryType(c *gin.Context) {
	toggled := c.Query("toggle")

	if toggled == "on" {
		c.String(200, AUTO_QUERY)
	} else {
		c.String(200, MANUAL_QUERY)
	}
}
