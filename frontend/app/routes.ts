import { type RouteConfig, index, layout, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("login", "routes/login.tsx"),
  route("signup", "routes/signup.tsx"),
  route("forgot-password", "routes/forgot-password.tsx"),
  route("reset-password", "routes/reset-password.tsx"),
  route("auth/callback", "routes/auth.callback.tsx"),
  layout("routes/dashboard.tsx", [
    route("dashboard/devices", "routes/dashboard.devices.tsx"),
    route("dashboard/api-keys", "routes/dashboard.api-keys.tsx"),
  ]),
] satisfies RouteConfig;
