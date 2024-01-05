/**
 * @format
 */

import {AppRegistry} from 'react-native';
import Main from './App';
import {name as appName} from './app.json';
import messaging from '@react-native-firebase/messaging';

const firebaseListener = async (remoteMessage) => {
  Incomingvideocall.displayIncomingCall("test");
  
};

// Register background handler
messaging().setBackgroundMessageHandler(firebaseListener);

AppRegistry.registerComponent(appName, () => Main);

AppRegistry.registerHeadlessTask('RNCallKeepBackgroundMessage', () => ({ name, callUUID, handle }) => {
    // Make your call here
    
    return Promise.resolve();
  });
  