/*
 * This file is used to generate the connection URL for the database connection
 * when the user changes inputs in the connection modal.
 *
 * This file is also used to show and hide the connection modal.
 */

// New connection modal
const modal = document.getElementById("connection-modal");

function ShowModal() {
  modal.classList.remove("hidden");
  modal.classList.remove("opacity-0");
}

function HideModal() {
  modal.classList.add("hidden");
  modal.classList.add("opacity-0");

  for (const key in input) {
    if (key != "driver") {
      input[key].value = "";
    }
  }
}

const input = {
  host: document.getElementById("db-host"),
  port: document.getElementById("db-port"),
  driver: document.getElementById("db-driver"),
  username: document.getElementById("db-username"),
  password: document.getElementById("db-password"),
  database: document.getElementById("db-database"),
  name: document.getElementById("db-conn-name"),
  connectionURL: document.getElementById("db-url")
}

function GenerateURL(data) {
  const driver = data.driver.value;
  if (driver == "postgres") {
    data.connectionURL.value = `postgres://${data.username.value || "<username>"}:${data.password.value || "<password>"}@${data.host.value || "<host>"}:${data.port.value || "<port>"}/${data.database.value || "<database>"}?sslmode=disable`;
  } else if (driver == "mysql" || driver == "mariadb") {
    data.connectionURL.value = `${data.username.value || "<username>"}:${data.password.value || "<password>"}@tcp(${data.host.value || "<host>"}:${data.port.value || "<port>"})/${data.database.value || "<database>"}`;
  } else if (driver == "sqlite3") {
    data.connectionURL.value = `${data.database.value || "<database>"}`;
  }
}

function ParseURL (data) {
  let regex;
  switch (data.driver.value) {
    case "postgres":
      regex = /^(?<protocol>[a-z]+):\/\/(?<username>[^:]+):(?<password>[^@]+)@(?<host>[^:]+):(?<port>[0-9]+)\/(?<database>[^\?]+)(\?(?<params>.*))?$/;
      break;
    case "mysql":
    case "mariadb":
      regex = /^(?<username>[^:]+):(?<password>[^@]+)@tcp\((?<host>[^:]+):(?<port>[0-9]+)\)\/(?<database>[^\?]+)(\?(?<params>.*))?$/;
      break;
    case "sqlite3":
      regex = /^(?<database>)/;
      break;
    default: 
      console.log("Parsing URL failed: Unsupported driver.")
      return;
  }

  const match = data.connectionURL.value.match(regex);
  if (match) {
    switch (data.driver.value) {
      case "postgres":
        var { protocol, username, password, host, port, database } = match.groups;
        data.host.value = host;
        data.port.value = port;
        data.password.value = password;
        data.username.value = username;
        data.driver.value = protocol;
        data.database.value = database;
        break;
      case "mysql":
      case "mariadb":
        var { username, password, host, port, database } = match.groups;
        data.host.value = host;
        data.port.value = port;
        data.password.value = password;
        data.username.value = username;
        data.database.value = database;
        break;
      case "sqlite":
        var { database } = match.groups;
        data.database.value = database;
        break;
      default:
        console.log("Parsing URL failed: Unsupported driver.")
        break;
    }

    
    data.connectionURL.classList.remove("text-red-500");
    document.getElementById("db-url-invalid").classList.add("hidden");

  } else {
    data.connectionURL.classList.add("text-red-500");
    document.getElementById("db-url-invalid").classList.remove("hidden");
  }
}

// Create the event listeners for the input fields
for (const key in input) {
  if (key == "connectionURL") {
    input[key].addEventListener("input", () => {
      ParseURL(input);
    })
  } else if (key == "name") {
    continue;
  } else {
    // If the input changed is the database name, update the connection name as well.
    // This will create a default connection name based on the database name.
    if (key == "database") {
      input[key].addEventListener("input", () => {
        GenerateURL(input);
        input.name.value = input[key].value;
      })
    } else {
      input[key].addEventListener("input", () => {
        GenerateURL(input);
      })
    }
  }
}

document.addEventListener("keydown", (event) => {
  if (event.key !== "Tab" || event.target.id !== "sql") {
    return;
  }

  event.preventDefault();

  const sql = event.target;
  const start = sql.selectionStart;
  const end = sql.selectionEnd;
  sql.value = `${sql.value.slice(0, start)}\t${sql.value.slice(end)}`;
  sql.selectionStart = start + 1;
  sql.selectionEnd = start + 1;
});
