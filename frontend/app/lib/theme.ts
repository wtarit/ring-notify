export type Theme = "light" | "dark";

const STORAGE_KEY = "ring-notify-theme";

export function getStoredTheme(): Theme | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem(STORAGE_KEY) as Theme | null;
}

export function getSystemTheme(): Theme {
  if (typeof window === "undefined") return "light";
  return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

export function getTheme(): Theme {
  return getStoredTheme() || getSystemTheme();
}

export function setTheme(theme: Theme) {
  localStorage.setItem(STORAGE_KEY, theme);
  document.documentElement.setAttribute("data-theme", theme);
}

export function toggleTheme(): Theme {
  const next = getTheme() === "dark" ? "light" : "dark";
  setTheme(next);
  return next;
}

/** Inline script string to inject in <head> to prevent flash of wrong theme */
export const themeScript = `
(function() {
  var theme = localStorage.getItem("${STORAGE_KEY}");
  if (!theme) {
    theme = window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
  }
  document.documentElement.setAttribute("data-theme", theme);
})();
`;
