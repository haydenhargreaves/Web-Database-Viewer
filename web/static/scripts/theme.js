const themeStorageKey = "theme";

function applyTheme(theme) {
  const isDark = theme === "dark";
  document.documentElement.classList.toggle("dark", isDark);
  document.documentElement.style.colorScheme = isDark ? "dark" : "light";

  const toggle = document.getElementById("theme-toggle");
  if (toggle) {
    toggle.setAttribute("aria-pressed", String(isDark));
    toggle.setAttribute("aria-label", `Switch to ${isDark ? "light" : "dark"} mode`);
    toggle.setAttribute("title", `Switch to ${isDark ? "light" : "dark"} mode`);
  }
}

function getPreferredTheme() {
  const savedTheme = localStorage.getItem(themeStorageKey);
  if (savedTheme === "light" || savedTheme === "dark") {
    return savedTheme;
  }

  return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

function ToggleTheme() {
  const theme = document.documentElement.classList.contains("dark") ? "light" : "dark";
  localStorage.setItem(themeStorageKey, theme);
  applyTheme(theme);
}

applyTheme(getPreferredTheme());
