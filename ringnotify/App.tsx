import React, {useEffect} from 'react';
import {SafeAreaView, StatusBar, useColorScheme} from 'react-native';
import messaging from '@react-native-firebase/messaging';

import {Colors} from 'react-native/Libraries/NewAppScreen';
import {
  NavigationContainer,
  DarkTheme as NavigationDarkTheme,
  DefaultTheme as NavigationDefaultTheme,
} from '@react-navigation/native';
import {createMaterialBottomTabNavigator} from '@react-navigation/material-bottom-tabs';
import {
  MD3DarkTheme,
  PaperProvider,
  adaptNavigationTheme,
} from 'react-native-paper';
import RNCallKeep from 'react-native-callkeep';

import Incomingvideocall from './src/utils/incoming-video-call';
import HomeScreen from './src/screens/HomeScreen';
import SettingsScreen from './src/screens/SettingsScreen';

const Tab = createMaterialBottomTabNavigator();

function App(): JSX.Element {
  useEffect(() => {
    const options = {
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
          notificationIcon: 'Path to the resource icon of the notification',
        },
      },
    };

    RNCallKeep.setup(options).then(accepted => {
      console.log('Callkeep Setup Done');
    });
  }, []);

  useEffect(() => {
    const unsubscribe = messaging().onMessage(remoteMessage => {
      if (remoteMessage.data != undefined) {
        Incomingvideocall.displayIncomingCall(remoteMessage.data.text);
      }
    });

    return () => {
      unsubscribe();
    };
  }, []);

  const isDarkMode = useColorScheme() === 'dark';

  const backgroundStyle = {
    backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
  };

  const {LightTheme, DarkTheme} = adaptNavigationTheme({
    reactNavigationLight: NavigationDefaultTheme,
    reactNavigationDark: NavigationDarkTheme,
  });

  const CombinedDarkTheme = {
    ...MD3DarkTheme,
    ...DarkTheme,
    colors: {
      ...MD3DarkTheme.colors,
      ...DarkTheme.colors,
    },
  };

  return (
    <NavigationContainer theme={CombinedDarkTheme}>
      <Tab.Navigator>
        <Tab.Screen name="Home" component={HomeScreen} />
        <Tab.Screen name="Settings" component={SettingsScreen} />
      </Tab.Navigator>
    </NavigationContainer>
  );
}

export default function Main() {
  return (
    <PaperProvider theme={MD3DarkTheme}>
      <App />
    </PaperProvider>
  );
}
