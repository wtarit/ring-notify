import { FirebaseMessagingTypes } from "@react-native-firebase/messaging";
import CallScreen from "./callscreen";

export const callHandler = async (
  message: FirebaseMessagingTypes.RemoteMessage
) => {
  if (message.data?.text) {
    CallScreen.displayIncomingCall(message.data.text as string);
  }
};
