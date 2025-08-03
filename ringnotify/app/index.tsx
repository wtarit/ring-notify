import { useEffect, useState } from "react";
import {
  Text,
  View,
  TextInput,
  Button,
  StyleSheet,
  Alert,
  TouchableOpacity,
  StatusBar,
  Linking,
} from "react-native";
import {
  getMessaging,
  getToken,
  onMessage,
} from "@react-native-firebase/messaging";
import CallScreen from "@/lib/callscreen";
import { createUser, testCall, getStoredApiKey } from "@/lib/api";
import * as Clipboard from "expo-clipboard";

export default function Index() {
  const [apiKey, setApiKey] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false);
  const messaging = getMessaging();

  useEffect(() => {
    CallScreen.setupCallKeep();
  }, []);

  useEffect(() => {
    const unsubscribe = onMessage(messaging, (remoteMessage) => {
      if (remoteMessage.data && typeof remoteMessage.data.text === "string") {
        console.log("foreground call");
        CallScreen.displayIncomingCall(remoteMessage.data.text);
      }
    });

    return () => {
      unsubscribe();
    };
  }, []);

  useEffect(() => {
    async function initializeApp() {
      const storedApiKey = await getStoredApiKey();
      if (storedApiKey) {
        setApiKey(storedApiKey);
        return;
      }

      try {
        const token = await getToken(messaging);
        console.log("FCM Token:", token);
        const newApiKey = await createUser(token);
        setApiKey(newApiKey);
      } catch (error) {
        console.error("Error initializing app:", error);
        Alert.alert("Error", "Failed to initialize the app. Please try again.");
      }
    }

    initializeApp();
  }, []);

  const handleRegenerateApiKey = async () => {
    try {
      setIsLoading(true);
      const token = await getToken(messaging);
      const newApiKey = await createUser(token);
      setApiKey(newApiKey);
      Alert.alert("Success", "API Key regenerated successfully");
    } catch (error) {
      console.error("Error regenerating API key:", error);
      Alert.alert("Error", "Failed to regenerate API key");
    } finally {
      setIsLoading(false);
    }
  };

  const handleTestCall = async () => {
    try {
      setIsLoading(true);
      await testCall(apiKey);
      Alert.alert("Success", "Test call sent successfully");
    } catch (error) {
      console.error("Error testing call:", error);
      Alert.alert("Error", "Failed to send test call");
    } finally {
      setIsLoading(false);
    }
  };

  const copyToClipboard = async () => {
    await Clipboard.setStringAsync(apiKey);
    Alert.alert("Success", "API Key copied to clipboard");
  };

  const openDocumentation = () => {
    Linking.openURL("https://ringnotify.wtarit.me/");
  };

  return (
    <View style={styles.container}>
      <StatusBar barStyle="dark-content" backgroundColor="#fff" />
      <Text style={styles.title}>Ring Notify</Text>

      <View style={styles.apiKeyContainer}>
        <Text style={styles.label}>Your API Key:</Text>
        <View style={styles.inputContainer}>
          <TextInput
            style={styles.input}
            value={apiKey}
            editable={false}
            placeholder="No API Key available"
          />
          <TouchableOpacity style={styles.copyButton} onPress={copyToClipboard}>
            <Text style={styles.copyButtonText}>Copy</Text>
          </TouchableOpacity>
        </View>
      </View>

      <View style={styles.buttonContainer}>
        <Button
          title="Regenerate API Key"
          onPress={handleRegenerateApiKey}
          disabled={isLoading}
        />
        <View style={styles.spacer} />
        <Button
          title="Test Call"
          onPress={handleTestCall}
          disabled={isLoading}
        />
        <View style={styles.spacer} />
        <Button title="View Documentation" onPress={openDocumentation} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#fff",
    paddingTop: StatusBar.currentHeight || 20,
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 20,
    textAlign: "center",
    color: "#000",
  },
  apiKeyContainer: {
    marginBottom: 20,
  },
  label: {
    fontSize: 16,
    marginBottom: 8,
    color: "#000",
  },
  inputContainer: {
    flexDirection: "row",
    alignItems: "center",
  },
  input: {
    flex: 1,
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 4,
    padding: 10,
    marginRight: 10,
    color: "#000",
  },
  copyButton: {
    backgroundColor: "#007AFF",
    padding: 10,
    borderRadius: 4,
  },
  copyButtonText: {
    color: "#fff",
    fontWeight: "bold",
  },
  buttonContainer: {
    marginTop: 20,
  },
  spacer: {
    height: 10,
  },
});
