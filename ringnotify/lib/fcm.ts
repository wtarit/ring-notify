import {FirebaseMessagingTypes} from '@react-native-firebase/messaging'
import CallScreen from './callscreen'

export const callHandler = async (message: FirebaseMessagingTypes.RemoteMessage) => {
  console.log("message =", message);
  CallScreen.displayIncomingCall("test");
}