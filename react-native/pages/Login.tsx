import React, { useState } from "react";
import { Button, StyleSheet, Text, TextInput, View } from "react-native";
import AsyncStorage from "@react-native-async-storage/async-storage";

function Login({navigation}) {
  const [username, setUsername] = useState("");

  return (
    <View style={style.container}>
      <Text style={style.text}>Login</Text>
      <Text style={style.text}>Async Storage</Text>
      <TextInput
        style={style.input}
        placeholder="Username"
        onChangeText={(val) => setUsername(val)}
      />

      <View style={style.btn}>
        <Button
          title="Login"
          onPress={() => {
            AsyncStorage.setItem("username", username)
              .then(() => {
                console.log("Username was saved successfully");
                navigation.navigate("Home");
              })
              .catch(() => {
                console.log("There was an error saving the username");
              });
          }}
        />
      </View>
    </View>
  );
}
const style = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
  text: {
    fontSize: 20,
    color: "darkslateblue",
  },
  input: {
    borderWidth: 1,
    borderColor: "#777",
    padding: 8,
    margin: 10,
    width: 200,
  },
  btn: {
    backgroundColor: "coral",
    padding: 9,
    margin: 5,
  },
});
export default Login;
