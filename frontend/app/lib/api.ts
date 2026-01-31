import { supabase } from "./supabase";
import type {
  DeviceListResponse,
  DeviceResponse,
  UpdateDeviceRequest,
  RegisterDeviceRequest,
  APIKeyListResponse,
  APIKeyResponse,
  CreateAPIKeyRequest,
  SuccessResponse,
} from "./types";

const BASE_URL = import.meta.env.VITE_BACKEND_URL?.replace(/\/$/, "") || "http://localhost:1323";

async function getAuthHeaders(): Promise<Record<string, string>> {
  const { data } = await supabase.auth.getSession();
  const token = data.session?.access_token;
  if (!token) throw new Error("Not authenticated");
  return {
    Authorization: `Bearer ${token}`,
    "Content-Type": "application/json",
  };
}

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
  const headers = await getAuthHeaders();
  const res = await fetch(`${BASE_URL}${path}`, {
    ...options,
    headers: { ...headers, ...options.headers },
  });
  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: res.statusText }));
    throw new Error(body.error || `Request failed: ${res.status}`);
  }
  return res.json();
}

// Devices
export const listDevices = () => request<DeviceListResponse>("/devices");

export const registerDevice = (data: RegisterDeviceRequest) =>
  request<DeviceResponse>("/devices", {
    method: "POST",
    body: JSON.stringify(data),
  });

export const updateDevice = (id: string, data: UpdateDeviceRequest) =>
  request<DeviceResponse>(`/devices/${id}`, {
    method: "PATCH",
    body: JSON.stringify(data),
  });

export const deleteDevice = (id: string) =>
  request<SuccessResponse>(`/devices/${id}`, { method: "DELETE" });

// API Keys
export const listAPIKeys = () => request<APIKeyListResponse>("/api-keys");

export const createAPIKey = (data: CreateAPIKeyRequest) =>
  request<APIKeyResponse>("/api-keys", {
    method: "POST",
    body: JSON.stringify(data),
  });

export const revokeAPIKey = (id: string) =>
  request<SuccessResponse>(`/api-keys/${id}`, { method: "DELETE" });
