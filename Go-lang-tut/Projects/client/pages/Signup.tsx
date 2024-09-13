import {
  StyleSheet,
  Text,
  TextInput,
  View,
  ProgressBarAndroid,
  Alert,
} from "react-native";
import React, { useEffect, useState } from "react";
import { SignupUser } from "../models/user";
import MyButton from "../components/Button";
import { Link } from "@react-navigation/native";
import useUserStore from "../store/auth";
import { register } from "../APIs/auth";

export default function Signup({ navigation }) {
  const signupUser = new SignupUser({});
  const [user, setUser] = useState(signupUser);
  const [fetchingResponse, setFetchingResponse] = useState(false);

  const register = useUserStore((state) => state.register);

  const isInputsValid = () => {
    return (
      user?.firstName !== "" &&
      user?.lastName !== "" &&
      user?.email !== "" &&
      new RegExp(/^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$/).test(user?.email) &&
      user?.password !== "" &&
      user?.password?.length >= 6 &&
      user?.confirmPassword !== "" &&
      user?.password === user?.confirmPassword
    );
  };

  const signUp = async () => {
    setFetchingResponse(true);
    try {
      await register(user);
      Alert.alert(
        "Success",
        "User has been registered successfully, please login to continue."
      );
    } catch (error) {
      Alert.alert("Error while registering user", "Please try again later.");
    }
    setFetchingResponse(false);
  };

  return (
    <View style={styles.container}>
      <View style={styles.innerContainer}>
        <Text style={styles.header}>Sign Up</Text>
        <TextInput
          style={styles.formControl}
          placeholder="First Name"
          value={signupUser.firstName}
          onChangeText={(text) => setUser({ ...user, firstName: text })}
        />
        <TextInput
          style={styles.formControl}
          placeholder="Last Name"
          value={user.lastName}
          onChangeText={(text) => setUser({ ...user, lastName: text })}
        />
        <TextInput
          style={styles.formControl}
          placeholder="Email"
          value={user.email}
          keyboardType="email-address"
          onChangeText={(text) => setUser({ ...user, email: text })}
        />
        <TextInput
          style={styles.formControl}
          placeholder="Password"
          secureTextEntry={true}
          value={user.password}
          onChangeText={(text) => setUser({ ...user, password: text })}
        />
        <TextInput
          style={styles.formControl}
          secureTextEntry={true}
          placeholder="Confirm Password"
          value={user.confirmPassword}
          onChangeText={(text) => setUser({ ...user, confirmPassword: text })}
        />
        <MyButton
          disabled={!isInputsValid() || fetchingResponse}
          title="Sign Up"
          customStyles={
            !isInputsValid() || fetchingResponse
              ? { backgroundColor: "#9BDDFF" }
              : null
          }
          onPress={signUp}
        />
        <Link
          style={{ marginTop: 10, color: "#007bff", marginBottom: 20 }}
          to={{ screen: "SignIn" }}
        >
          Already have an account? Sign In
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
