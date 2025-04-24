import { useEffect } from "react";
import { Text, View } from "react-native";
import messaging from "@react-native-firebase/messaging";
import CallScreen from "@/lib/callscreen";

export default function Index() {
  useEffect(() => {
    CallScreen.setupCallKeep();
  }, []);

  useEffect(() => {
    const unsubscribe = messaging().onMessage((remoteMessage) => {
      if (remoteMessage.data != undefined) {
        console.log("foregound call");
        CallScreen.displayIncomingCall(remoteMessage.data.text);
      }
    });

    return () => {
      unsubscribe();
    };
  }, []);

  useEffect(() => {
    async function getFCMtoken(): Promise<string> {
      const token = await messaging().getToken();
      console.log("token", token);
      return token;
    }
    getFCMtoken();
  });
  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text>Edit app/index.tsx to edit this screen.</Text>
    </View>
  );
}
