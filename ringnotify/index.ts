import "expo-router/entry";

import {
  getMessaging,
  setBackgroundMessageHandler,
} from "@react-native-firebase/messaging";
import { callHandler } from "./lib/fcm";
import { AppRegistry } from "react-native";

console.log("registering");
const messaging = getMessaging();
setBackgroundMessageHandler(messaging, callHandler);

AppRegistry.registerHeadlessTask(
  "RNCallKeepBackgroundMessage",
  () =>
    ({ name, callUUID, handle }) => {
      // Make your call here

      return Promise.resolve();
    }
);
