import { StatusBar } from "expo-status-bar";
import { useState } from "react";
import { StyleSheet, View,Text, Button } from "react-native";
import { NavigationContainer } from '@react-navigation/native';
import HomeScreen from "./pages/Home";
import {createDrawerNavigator} from "@react-navigation/drawer"
import Login from "./pages/Login";

const Stack = createDrawerNavigator();

export default function App() {
  return <>
    <NavigationContainer>
     <Stack.Navigator
     >
        <Stack.Screen name="Details" component={Login}/>
        <Stack.Screen name="Home" component={HomeScreen}/>
      </Stack.Navigator>
    </NavigationContainer>
  </>;
}

const styles = StyleSheet.create({
  container:{
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
  text:{
    fontSize: 20,
    color: "darkslateblue",
  },
});
