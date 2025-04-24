import * as SecureStore from 'expo-secure-store';

const API_BASE_URL = 'https://api.ringnotify.wtarit.me';

export const createUser = async (fcmToken: string): Promise<string> => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ fcmToken }),
    });

    if (!response.ok) {
      throw new Error('Failed to create user');
    }

    const data = await response.json();
    await SecureStore.setItemAsync('apiKey', data.APIKey);
    return data.APIKey;
  } catch (error) {
    console.error('Error creating user:', error);
    throw error;
  }
};

export const testCall = async (apiKey: string): Promise<void> => {
  try {
    const response = await fetch(`${API_BASE_URL}/notify/call`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${apiKey}`,
      },
      body: JSON.stringify({ text: 'test' }),
    });

    if (!response.ok) {
      throw new Error('Failed to test call');
    }
  } catch (error) {
    console.error('Error testing call:', error);
    throw error;
  }
};

export const getStoredApiKey = async (): Promise<string | null> => {
  try {
    return await SecureStore.getItemAsync('apiKey');
  } catch (error) {
    console.error('Error getting stored API key:', error);
    return null;
  }
}; 