import { useEffect } from "react";
import { Text, View } from "react-native";
import messaging from '@react-native-firebase/messaging'
import CallScreen from '@/lib/callscreen';
import RNCallKeep from 'react-native-callkeep';

export default function Index() {
  useEffect(() => {
    // const options = {
    //   ios: {
    //     appName: 'Ring Notify',
    //   },
    //   android: {
    //     alertTitle: 'Permissions required',
    //     alertDescription:
    //       'This application needs to access your phone accounts',
    //     cancelButton: 'Cancel',
    //     okButton: 'ok',
    //     imageName: 'phone_account_icon',
    //     // additionalPermissions: [PermissionsAndroid.PERMISSIONS.example],
    //     additionalPermissions: [],
    //     // Required to get audio in background when using Android 11
    //     foregroundService: {
    //       channelId: 'me.wtarit.ringnotify',
    //       channelName: 'Foreground service for my app',
    //       notificationTitle: 'My app is running on background',
    //       notificationIcon: 'Path to the resource icon of the notification',
    //     },
    //   },
    // };
    CallScreen.setupCallKeep();
  }, []);

  useEffect(() => {
    const unsubscribe = messaging().onMessage(remoteMessage => {
      if (remoteMessage.data != undefined) {
        console.log("foregound call")
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
      console.log('token', token);
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
