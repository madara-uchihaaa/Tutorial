import React, { useEffect, useState } from "react";
import {
  StyleSheet,
  Text,
  View,
  ProgressBarAndroid,
  Alert,
} from "react-native";
import { TextInput } from "react-native-gesture-handler";
import MyButton from "../components/Button";
import { SigninUser } from "../models/user";
import { Link } from "@react-navigation/native";
import useUserStore from "../store/auth";

export default function SignIn({ navigation }) {
  const signinUser = new SigninUser({});
  const [user, setUser] = useState(signinUser);
  const [fetchingResponse, setFetchingResponse] = useState(false);

  const isInputValid = () => {
    return (
      user?.email !== "" &&
      new RegExp(/^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$/).test(user?.email) &&
      user?.password !== "" &&
      user?.password?.length >= 6
    );
  };

  const login = useUserStore((state) => state.signIn);
  const handleSubmit = () => {
    setFetchingResponse(true);
    try {
      login(user);
      navigation.navigate("conversation");
    } catch (error) {
      Alert.alert("Error while logging in", "Please try again later.");
    }
    setFetchingResponse(false);
  };

  return (
    <View style={styles.container}>
      <View style={styles.innerContainer}>
        <Text style={styles.header}>Sign In</Text>
        <TextInput
          style={styles.formControl}
          placeholder="Email"
          onChangeText={(text) => setUser({ ...user, email: text })}
        />
        <TextInput
          style={styles.formControl}
          placeholder="Password"
          onChangeText={(text) => setUser({ ...user, password: text })}
        />

        <MyButton
          disabled={!isInputValid() || fetchingResponse}
          title="Sign In"
          onPress={handleSubmit}
          customStyles={
            !isInputValid() || fetchingResponse
              ? { backgroundColor: "#9BDDFF" }
              : null
          }
        />
        <Link
          style={{ marginTop: 10, color: "#007bff", marginBottom: 20 }}
          to={{ screen: "SignUp" }}
        >
          Don't have an account? Sign Up
        </Link>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "rgba(210,146,255,1)",
    justifyContent: "center",
    alignItems: "center",
  },
  innerContainer: {
    width: "70%",
    backgroundColor: "#fff",
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 5,
  },
  header: {
    fontSize: 23,
    fontWeight: "bold",
    marginBottom: 10,
    marginTop: 20,
  },
  formControl: {
    height: 40,
    borderColor: "#ced4da",
    borderWidth: 1,
    borderRadius: 4,
    paddingHorizontal: 10,
    fontSize: 16,
    width: "80%",
    backgroundColor: "#fff",
    marginBottom: 15,
  },
});
