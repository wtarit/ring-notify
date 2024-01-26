import { Platform } from "react-native";
import RNCallKeep from "react-native-callkeep";
import uuid from "react-native-uuid";

class IncomingCall {
  constructor() {
    this.currentCallId = null;
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
  setupCallKeep = () => {
    try {
      RNCallKeep.setup({
        ios: {
          appName: "VideoSDK",
          supportsVideo: false,
          maximumCallGroups: "1",
          maximumCallsPerCallGroup: "1",
        },
        android: {
          alertTitle: "Permissions required",
          alertDescription:
            "This application needs to access your phone accounts",
          cancelButton: "Cancel",
          okButton: "Ok",
        },
      });
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
    this.currentCallId = null;
    this.removeEvents();
  };

  //These method will remove all the event listeners
  removeEvents = () => {
    RNCallKeep.removeEventListener("answerCall");
    RNCallKeep.removeEventListener("endCall");
  };

  //These method will display the incoming call
  displayIncomingCall = (callerName) => {
    Platform.OS === "android" && RNCallKeep.setAvailable(false);
    RNCallKeep.displayIncomingCall(
      this.getCurrentCallId(),
      callerName,
      callerName,
      "number",
      true,
      null
    );
  };

  //Bring the app to foreground
  backToForeground = () => {
    RNCallKeep.backToForeground();
  };

  //Return the ID of current Call
  getCurrentCallId = () => {
    if (!this.currentCallId) {
      this.currentCallId = uuid.v4();
    }
    return this.currentCallId;
  };

  //These Method will end the call
  endAllCall = () => {
    RNCallKeep.endAllCalls();
    this.currentCallId = null;
    this.removeEvents();
  };

}

export default Incomingvideocall = new IncomingCall();
