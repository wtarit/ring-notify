import {View, StyleSheet} from 'react-native';
import React, {useEffect, useState} from 'react';
import messaging from '@react-native-firebase/messaging';
import {storage} from '../utils/storage';
import Config from 'react-native-config';
import {TextInput} from 'react-native-paper';
import Clipboard from '@react-native-clipboard/clipboard';
import ScreenWrapper from './ScreenWrapper';
import {useSafeAreaInsets} from 'react-native-safe-area-context';

function HomeScreen() {
  const [apiKey, setApiKey] = useState('');
  useEffect(() => {
    const apiKey = storage.getString('apiKey');
    async function getFCMtoken(): Promise<string> {
      const token = await messaging().getToken();
      return token;
    }
    async function getApiKey() {
      try {
        console.log(`${Config.API_URL}/newuser`);
        const response = await fetch(`${Config.API_URL}/newuser`, {
          method: 'POST',
          body: JSON.stringify({fcm_token: await getFCMtoken()}),
        });
        const json = await response.json();
        storage.set('apiKey', json.uuid);
        setApiKey(json.uuid);
      } catch (error) {
        console.log(error);
      }
    }
    if (apiKey == undefined) {
      getApiKey();
      console.log('called api');
    } else {
      setApiKey(apiKey);
    }
  }, []);

  const copyToClipboard = () => {
    Clipboard.setString(apiKey);
  };

  return (
    <ScreenWrapper style={styles.default}>
      <View>
        <TextInput
          label={'Your API Key:'}
          mode="outlined"
          value={apiKey}
          editable={false}
          right={
            <TextInput.Icon icon="content-copy" onPress={copyToClipboard} />
          }
        />
      </View>
    </ScreenWrapper>
  );
}

export default HomeScreen;

const styles = StyleSheet.create({
  default: {padding: 8},
});
