export interface DeviceResponse {
  id: string;
  deviceName: string;
  deviceType: string;
  registeredAt: string;
  lastActive: string;
  isActive: boolean;
}

export interface DeviceListResponse {
  devices: DeviceResponse[];
}

export interface UpdateDeviceRequest {
  deviceName?: string;
  fcmToken?: string;
}

export interface RegisterDeviceRequest {
  fcmToken: string;
  deviceName: string;
  deviceType: string;
}

export interface APIKeyResponse {
  id: string;
  key: string;
  name: string;
  createdAt: string;
  expiresAt?: string;
  lastUsedAt?: string;
  isActive: boolean;
}

export interface APIKeyListResponse {
  apiKeys: APIKeyResponse[];
}

export interface CreateAPIKeyRequest {
  name: string;
  expiresAt?: string;
}

export interface ErrorResponse {
  error: string;
  details?: Record<string, string>;
}

export interface SuccessResponse {
  message: string;
}
