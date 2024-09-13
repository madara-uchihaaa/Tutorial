import { View, Text, StyleSheet, Alert } from "react-native";
import React from "react";
import useUserStore from "../store/auth";
import { TouchableOpacity } from "react-native-gesture-handler";

export default function SignOut({ navigation }) {
  const logout = useUserStore((state) => state.signOut);

  const logOutUser = async () => {
    await logout();
    navigation.navigate("signin");
    Alert.alert("Success", "User has been logged out successfully.");
  };

  return (
    <View style={styles.container}>
      <TouchableOpacity style={styles.addButton} onPress={logOutUser}>
        <Text style={styles.buttonText}>Log out</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    justifyContent: "center",
    alignItems: "center",
  },
  addButton: {
    backgroundColor: "red",
    padding: 15,
    borderRadius: 5,
    alignItems: "center",
  },
  buttonText: {
    color: "white",
    fontSize: 18, // Adjust the font size as needed
    fontWeight: "bold", // Adjust the font weight as needed
  },
  answer: {
    marginTop: 10,
  },
});
