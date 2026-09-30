const driverInput = document.getElementById("db-driver");
const hostInput = document.getElementById("db-host");
const portInput = document.getElementById("db-port");
const usernameInput = document.getElementById("db-username");
const databaseInput = document.getElementById("db-database");

function updateConnectionURL() {
  const host = hostInput.value.trim();
  const port = portInput.value.trim();
  const username = encodeURIComponent(usernameInput.value);
  const password = encodeURIComponent(passwordInput.value);
  const database = encodeURIComponent(databaseInput.value);
  const rawUsername = usernameInput.value;
  const rawPassword = passwordInput.value;
  const address = port ? `${host}:${port}` : host;

  switch (driverInput.value) {
    case "postgres":
      urlInput.value = `postgresql://${username}:${password}@${address}/${database}?sslmode=disable`;
      break;
    case "mysql":
    case "mariadb":
      urlInput.value = `${rawUsername}:${rawPassword}@tcp(${address})/${database}`;
      break;
    case "sqlite3":
      urlInput.value = databaseInput.value;
      break;
    case "sqlserver":
      urlInput.value = `sqlserver://${username}:${password}@${address}?database=${database}`;
      break;
    case "oracle":
      urlInput.value = `oracle://${username}:${password}@${address}/${database}`;
      break;
    case "db2":
      urlInput.value = `HOSTNAME=${host};PORT=${port};DATABASE=${databaseInput.value};UID=${rawUsername};PWD=${rawPassword}`;
      break;
  }
}

driverInput.addEventListener("change", updateConnectionURL);
for (const input of [hostInput, portInput, usernameInput, passwordInput, databaseInput]) {
  input.addEventListener("input", updateConnectionURL);
}
