import 'expo-router/entry'

import messaging from '@react-native-firebase/messaging'
import { router } from 'expo-router'
import { callHandler } from './lib/fcm';
import { AppRegistry } from 'react-native';

console.log("registering");
messaging().setBackgroundMessageHandler(callHandler);

AppRegistry.registerHeadlessTask('RNCallKeepBackgroundMessage', () => ({ name, callUUID, handle }) => {
  // Make your call here
  
  return Promise.resolve();
});