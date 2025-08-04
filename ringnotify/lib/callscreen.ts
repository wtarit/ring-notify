import { Platform } from "react-native";
import RNCallKeep from "react-native-callkeep";

class IncomingCall {
  setupCallKeep = async () => {
    try {
      const result = await RNCallKeep.setup({
        ios: {
          appName: "Ring Notify",
        },
        android: {
          alertTitle: "Permissions required",
          alertDescription:
            "This application needs to access your phone accounts",
          cancelButton: "Cancel",
          okButton: "ok",
          imageName: "phone_account_icon",
          // additionalPermissions: [PermissionsAndroid.PERMISSIONS.example],
          additionalPermissions: [],
          // Required to get audio in background when using Android 11
          foregroundService: {
            channelId: "me.wtarit.ringnotify",
            channelName: "Foreground service for my app",
            notificationTitle: "My app is running on background",
            notificationIcon: "../assets/images/favicon.png",
          },
        },
      });
      console.log(result);
    } catch (error) {
      console.error("initializeCallKeep error:", (error as Error)?.message);
    }
  };

  //These method will display the incoming call
  displayIncomingCall = (callerName: string) => {
    Platform.OS === "android" && RNCallKeep.setAvailable(false);
    RNCallKeep.displayIncomingCall(
      "",
      "ringnotify",
      callerName,
      "generic",
      true,
      undefined
    );
  };
}

const CallScreen = new IncomingCall();
export default CallScreen;
