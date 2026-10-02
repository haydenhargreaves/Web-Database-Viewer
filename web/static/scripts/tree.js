/*
 * This file contains the functions that are used to toggle the visibility of the fields 
 * in the tree view of the tables.
 *
 * This file also contains the functions that are used to generate quick queries for the 
 * tables.
 *
 * This file also contains the functions that are used to toggle the visibility of the
 * enum values in the tree view of the tables.
 */
function ToggleFields(id) {
  const fields = document.getElementById(`fields-${id}`);
  const button_svg = document.getElementById(`icon-${id}`);
  if (fields.classList.contains("hidden")) {
    fields.classList.remove("hidden");
    button_svg.setAttribute("transform", "rotate(0)");
  } else {
    fields.classList.add("hidden");
    button_svg.setAttribute("transform", "rotate(-90)");
  }
}

function ToggleTreeSection(id) {
  const section = document.getElementById(`tree-section-${id}`);
  const button = document.getElementById(`tree-section-toggle-${id}`);
  const icon = document.getElementById(`tree-section-icon-${id}`);
  const isExpanded = section.classList.contains("hidden");

  section.classList.toggle("hidden", !isExpanded);
  button.setAttribute("aria-expanded", isExpanded.toString());
  icon.setAttribute("transform", isExpanded ? "rotate(0)" : "rotate(-90)");
}

function LoadTableQuery(table) {
  const sql = document.getElementById("sql")
  sql.value = `SELECT * FROM ${table};`;
  sql.dispatchEvent(new Event("input", { bubbles: true }));
}

function LoadTableQueryWithFields(table, fields) {
  const sql = document.getElementById("sql")
  sql.value = `SELECT ${fields} FROM ${table};`;
  sql.dispatchEvent(new Event("input", { bubbles: true }));
}

async function LoadRoutineDefinition(button) {
  const errorMessage = document.getElementById("routine-source-error");
  errorMessage.classList.add("hidden");
  errorMessage.textContent = "";

  const parameters = new URLSearchParams({
    id: button.dataset.routineId,
    name: button.dataset.routineName,
    kind: button.dataset.routineKind,
  });

  try {
    const response = await fetch(`/v1/web/connections/tree/routine/definition?${parameters}`);
    if (!response.ok) {
      throw new Error(await response.text());
    }

    const sql = document.getElementById("sql");
    sql.value = await response.text();
    sql.dispatchEvent(new Event("input", { bubbles: true }));
  } catch (error) {
    errorMessage.textContent = error.message || "Unable to load the routine definition.";
    errorMessage.classList.remove("hidden");
  }
}

function ToggleEnumValues(id) {
  const enum_values = document.getElementById(`enum-values-${id}`);
  const button_svg = document.getElementById(`icon-enum-${id}`);

  if (enum_values.classList.contains("hidden")) {
    enum_values.classList.remove("hidden");
    button_svg.setAttribute("transform", "rotate(0)");
  } else {
    enum_values.classList.add("hidden");
    button_svg.setAttribute("transform", "rotate(-90)");
  }
}

function FilterTables(query) {
  const normalizedQuery = query.trim().toLowerCase();
  const tables = document.querySelectorAll("#database-table-tree > li[data-table-name]");
  let matches = 0;

  tables.forEach((table) => {
    const matchesQuery = FuzzyMatches(table.dataset.tableName.toLowerCase(), normalizedQuery);
    table.classList.toggle("hidden", !matchesQuery);
    if (matchesQuery) {
      matches += 1;
    }
  });

  const emptyMessage = document.getElementById("table-search-empty");
  emptyMessage.classList.toggle("hidden", normalizedQuery === "" || matches > 0);
}

function FuzzyMatches(tableName, query) {
  let queryIndex = 0;

  for (const character of tableName) {
    if (character === query[queryIndex]) {
      queryIndex += 1;
    }
  }

  return queryIndex === query.length;
}

function ResetTableSearch() {
  const search = document.getElementById("table-search");
  if (!search) {
    return;
  }

  search.value = "";
  FilterTables("");
}

function FilterRoutines(query) {
  const normalizedQuery = query.trim().toLowerCase();
  const routines = document.querySelectorAll("#database-routine-tree > li[data-routine-name]");
  let matches = 0;

  routines.forEach((routine) => {
    const matchesQuery = FuzzyMatches(routine.dataset.routineName.toLowerCase(), normalizedQuery);
    routine.classList.toggle("hidden", !matchesQuery);
    if (matchesQuery) {
      matches += 1;
    }
  });

  const emptyMessage = document.getElementById("routine-search-empty");
  emptyMessage.classList.toggle("hidden", normalizedQuery === "" || matches > 0);
}

function ResetRoutineSearch() {
  const search = document.getElementById("routine-search");
  if (!search) {
    return;
  }

  search.value = "";
  FilterRoutines("");
}

document.addEventListener("htmx:afterSettle", (event) => {
  if (event.detail.xhr.responseText.includes('id="database-table-tree"')) {
    ResetTableSearch();
  }
  if (event.detail.xhr.responseText.includes('id="database-routine-tree"')) {
    ResetRoutineSearch();
  }
});
