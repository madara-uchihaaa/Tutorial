import AsyncStorage from '@react-native-async-storage/async-storage';
export const setTokenToLs = async (token: string) => {
    await AsyncStorage.setItem('userToken', token);
}

export const getTokenFromLS = async () => {
    const userToken = await AsyncStorage.getItem('userToken');
    return userToken ? userToken : null;
}

export const removeTokenFromLS = async () => {
    await AsyncStorage.removeItem('userToken')
}
