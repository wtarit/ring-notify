/** Check for React Native WebView bridge (auto-injected by react-native-webview) */
export function isWebView(): boolean {
  return typeof window !== "undefined" && !!(window as any).ReactNativeWebView;
}
