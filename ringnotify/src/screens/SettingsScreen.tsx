import {StyleSheet, TextComponent, View} from 'react-native';
import {Text} from 'react-native-paper';
import ScreenWrapper from './ScreenWrapper';

function SettingsScreen() {
  return (
    <ScreenWrapper>
      <View style={{flex: 1}}>
        <Text variant="titleSmall">Appearance</Text>
          <View style={styles.row}>
            <Text>Theme</Text>
            
          </View>
      </View>
    </ScreenWrapper>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingVertical: 8,
  },
  row: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingVertical: 8,
    paddingHorizontal: 16,
  },
});

export default SettingsScreen;
