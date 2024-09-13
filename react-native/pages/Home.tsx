import AsyncStorage from "@react-native-async-storage/async-storage";
import { useEffect, useState } from "react";
import { Button, StyleSheet, Text, View } from "react-native";

function HomeScreen({ navigation }) {
  const [username, setUsername] = useState("");

  const getData = async () => {
    try {
      const value = await AsyncStorage.getItem("username");
      if (value !== null) {
        setUsername(value);
      }
    } catch (e) {
      console.log("There was an error retrieving the username");
    }
  };

  useEffect(() => {
    getData();
  }, [username]);

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Home Screen {username}</Text>
      <Button
        title="Go to Details"
        onPress={() => navigation.navigate("Details")}
      />
    </View>
  );
}

const styles = StyleSheet.create({
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
});

export default HomeScreen;
