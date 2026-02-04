# Ring Notify - Frontend Documentation

## Project Structure

```
frontend/
├── app/
│   ├── components/          # Reusable UI components
│   │   ├── confirm-dialog.tsx
│   │   ├── copy-button.tsx
│   │   ├── empty-state.tsx
│   │   ├── navbar.tsx
│   │   └── protected-route.tsx
│   ├── context/
│   │   └── auth-context.tsx  # Supabase auth state provider
│   ├── lib/
│   │   ├── api.ts            # Backend API client
│   │   ├── supabase.ts       # Supabase client init
│   │   ├── theme.ts          # Light/dark theme management
│   │   ├── types.ts          # TypeScript interfaces
│   │   └── webview.ts        # WebView detection utility
│   ├── routes/
│   │   ├── home.tsx
│   │   ├── login.tsx
│   │   ├── signup.tsx
│   │   ├── forgot-password.tsx
│   │   ├── reset-password.tsx
│   │   ├── auth.callback.tsx
│   │   ├── dashboard.tsx
│   │   ├── dashboard.devices.tsx
│   │   └── dashboard.api-keys.tsx
│   ├── root.tsx              # Root layout with AuthProvider
│   ├── routes.ts             # Route definitions
│   └── app.css               # Tailwind + DaisyUI styles
├── vite.config.ts
├── react-router.config.ts    # ssr: false (SPA mode)
└── wrangler.jsonc            # Cloudflare Workers deployment
```

## Tech Stack

- React 19, React Router 7, TypeScript, Vite 7
- Tailwind CSS 4 + DaisyUI 5
- Supabase JS v2 (auth)
- Deployed on Cloudflare Workers (SPA mode, no SSR)

## Environment Variables

```env
VITE_SUPABASE_URL=<supabase project url>
VITE_SUPABASE_PUBLISHABLE_DEFAULT_KEY=<supabase anon key>
VITE_BACKEND_URL=<backend api base url>
```

## Routes

| Route | Auth | Description |
|---|---|---|
| `/` | No | Landing page |
| `/login` | No | Email/password + Google OAuth sign-in |
| `/signup` | No | Account registration |
| `/forgot-password` | No | Send password reset email |
| `/reset-password` | Yes* | Set new password after recovery link |
| `/auth/callback` | No | OAuth and password recovery redirect handler |
| `/dashboard/devices` | Yes | Device management |
| `/dashboard/api-keys` | Yes | API key management |

*Requires session from Supabase recovery link.

## Authentication Flow

Auth state is managed via `AuthProvider` in `root.tsx`, backed by Supabase.

**Sign in:** `supabase.auth.signInWithPassword()` or `signInWithOAuth({ provider: "google" })`

**Sign up:** `supabase.auth.signUp()`

**Password reset:**
1. `/forgot-password` calls `supabase.auth.resetPasswordForEmail(email, { redirectTo: origin + "/auth/callback" })`
2. User clicks email link → browser opens `/auth/callback`
3. `auth.callback.tsx` listens for `onAuthStateChange` — on `PASSWORD_RECOVERY` event, redirects to `/reset-password`
4. `/reset-password` calls `supabase.auth.updateUser({ password })`, then navigates to dashboard

**Protected routes:** `ProtectedRoute` component checks `session` from `useAuth()`. Redirects to `/login` if unauthenticated.

**Session token:** All backend API calls use the Supabase access token as `Authorization: Bearer <token>`.

## Backend API

Base URL from `VITE_BACKEND_URL`. All requests include Bearer token.

**Devices:**
```
GET    /devices           → list devices
POST   /devices           → register device
PATCH  /devices/{id}      → update device
DELETE /devices/{id}      → delete device
```

**API Keys:**
```
GET    /api-keys           → list keys
POST   /api-keys           → create key
DELETE /api-keys/{id}      → revoke key
```

## TypeScript Types (lib/types.ts)

```typescript
interface DeviceResponse {
  id: string;
  deviceName: string;
  deviceType: string;       // "android", "ios"
  registeredAt: string;
  lastActive: string;
  isActive: boolean;
}

interface RegisterDeviceRequest {
  fcmToken: string;
  deviceName: string;
  deviceType: string;
}

interface APIKeyResponse {
  id: string;
  key: string;              // only shown once at creation
  name: string;
  createdAt: string;
  expiresAt?: string;
  lastUsedAt?: string;
  isActive: boolean;
}

interface CreateAPIKeyRequest {
  name: string;
  expiresAt?: string;
}
```

## WebView Integration (Mobile Apps)

### Detection

`lib/webview.ts` checks for `window.ReactNativeWebView` to determine if the app is running inside a React Native WebView.

### Native Bridge (WebView ↔ Native)

When in WebView, the devices page shows a "Register This Device" button. On click:

```typescript
(window as any).ReactNativeWebView?.postMessage(
  JSON.stringify({ type: "REQUEST_FCM_TOKEN" })
);
```

**Expected native app response:** The native side should listen for this message, obtain the device FCM token, and post it back. The frontend then calls `POST /devices` with the token and device info.

When not in WebView, a QR code linking to GitHub releases is shown instead.

### Integration Checklist

- `window.ReactNativeWebView` is automatically set by `react-native-webview`
- Listen for `onMessage` events with `type: "REQUEST_FCM_TOKEN"`
- Post back FCM token and device info to the WebView
- Allow localStorage (theme persistence + Supabase session)
- Allow network access to both Supabase URL and backend URL
- Handle OAuth redirects — Google OAuth opens a browser; the redirect to `/auth/callback` must return to the WebView
- Consider injecting device metadata (model, OS) for the `deviceType` field
- Implement FCM token refresh to keep backend registration current

### Theme

Stored in `localStorage` key `ring-notify-theme`. Values: `"light"` or `"dark"`. Falls back to system preference. Applied via `data-theme` attribute on `<html>`.

## Development

```bash
pnpm install
cp .env.sample .env   # fill in values
pnpm dev              # http://localhost:5173
pnpm build            # outputs to build/client/
pnpm typecheck
```
