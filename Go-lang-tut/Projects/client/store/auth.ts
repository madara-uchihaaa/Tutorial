import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { SigninUser, SignupUser, User } from '../models/user';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { persist, createJSONStorage } from 'zustand/middleware'
import { getTokenFromLS, removeTokenFromLS, setTokenToLs } from './storage';
import { login, register } from '../APIs/auth';

export interface UserStore {
    getUserToken: () => Promise<string | null>;
    signIn: (user: SigninUser) => Promise<void>;
    signOut: () => Promise<void>;
    register: (user: SignupUser) => Promise<void>
}


const useUserStore = create<UserStore>()(
    devtools(
        persist(
            (set) => ({
                signIn: async (signinUser: SigninUser) => {
                    try {
                        const response = await login(signinUser);
                        console.log(response.data, "response");
                        if (response.status === 200) {
                            setTokenToLs(response.data);
                        } else {
                            throw new Error('Something went wrong while logging in');
                        }
                    } catch (error) {
                        throw error;
                    }
                },
                signOut: async () => {
                    try {
                        await removeTokenFromLS();
                    } catch (error) {
                        throw error;
                    }
                },
                register: async (user: SignupUser) => {
                    try {
                        const response = await register(user);
                        if (response.status === 200) {
                        } else {
                            throw new Error('Something went wrong while registering');
                        }
                    } catch (error) {
                        throw error;
                    }
                },
                getUserToken: () => {
                    return getTokenFromLS();
                }
            }),
            { name: 'user-store', storage: createJSONStorage(() => AsyncStorage) },

        )
    )
);

export default useUserStore;
