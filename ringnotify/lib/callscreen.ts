import { Platform } from "react-native";
import RNCallKeep from "react-native-callkeep";
import {randomUUID} from 'expo-crypto';

class IncomingCall {
  currentCallId: string;
  constructor() {
    this.currentCallId = "";
  }

  configure = (incomingcallAnswer, endIncomingCall) => {
    try {
      this.setupCallKeep();
      Platform.OS === "android" && RNCallKeep.setAvailable(true);
      RNCallKeep.addEventListener("answerCall", incomingcallAnswer);
      RNCallKeep.addEventListener("endCall", endIncomingCall);
    } catch (error) {
      console.error("initializeCallKeep error:", error?.message);
    }
  };

  //These emthod will setup the call keep.
  setupCallKeep = async () => {
    try {
      const result = await RNCallKeep.setup({
          ios: {
            appName: 'Ring Notify',
          },
          android: {
            alertTitle: 'Permissions required',
            alertDescription:
              'This application needs to access your phone accounts',
            cancelButton: 'Cancel',
            okButton: 'ok',
            imageName: 'phone_account_icon',
            // additionalPermissions: [PermissionsAndroid.PERMISSIONS.example],
            additionalPermissions: [],
            // Required to get audio in background when using Android 11
            foregroundService: {
              channelId: 'me.wtarit.ringnotify',
              channelName: 'Foreground service for my app',
              notificationTitle: 'My app is running on background',
              notificationIcon: '../assets/images/favicon.png',
            },
          },
        });
      console.log(result);
    } catch (error) {
      console.error("initializeCallKeep error:", error?.message);
    }
  };

  // Use startCall to ask the system to start a call - Initiate an outgoing call from this point
  startCall = ({ handle, localizedCallerName }) => {
    // Your normal start call action
    RNCallKeep.startCall(this.getCurrentCallId(), handle, localizedCallerName);
  };

  reportEndCallWithUUID = (callUUID, reason) => {
    RNCallKeep.reportEndCallWithUUID(callUUID, reason);
  };

  //These method will end the incoming call
  endIncomingcallAnswer = () => {
    RNCallKeep.endCall(this.currentCallId);
    this.currentCallId = "";
    this.removeEvents();
  };

  //These method will remove all the event listeners
  removeEvents = () => {
    RNCallKeep.removeEventListener("answerCall");
    RNCallKeep.removeEventListener("endCall");
  };

  //These method will display the incoming call
  displayIncomingCall = (callerName: string) => {
    Platform.OS === "android" && RNCallKeep.setAvailable(false);
    RNCallKeep.displayIncomingCall(
      this.getCurrentCallId(),
      "ringnotify",
      callerName,
      "generic",
      true,
      undefined
    );
  };

  //Bring the app to foreground
  backToForeground = () => {
    RNCallKeep.backToForeground();
  };

  //Return the ID of current Call
  getCurrentCallId = () => {
    if (!this.currentCallId) {
      this.currentCallId = randomUUID();
    }
    return this.currentCallId;
  };

  //These Method will end the call
  endAllCall = () => {
    RNCallKeep.endAllCalls();
    this.currentCallId = "";
    this.removeEvents();
  };

}

const CallScreen = new IncomingCall();
export default CallScreen;
