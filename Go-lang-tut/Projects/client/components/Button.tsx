import React from "react";
import { TouchableOpacity, Text, StyleSheet } from "react-native";

export default function MyButton({
  title,
  onPress,
  disabled,
  customStyles,
}: {
  title: string;
  onPress: any;
  disabled?: boolean;
  customStyles?: any;
}) {
  return (
    <TouchableOpacity
      disabled={disabled}
      style={[
        !customStyles
          ? styles.button
          : {
              ...styles.button,
              ...customStyles,
            },
      ]}
      onPress={onPress}
    >
      <Text style={styles.buttonText}>{title}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  button: {
    backgroundColor: "#007bff",
    borderRadius: 4,
    paddingVertical: 10,
    paddingHorizontal: 15,
    alignItems: "center",
    justifyContent: "center",
    marginBottom: 10,
    width: "80%",
  },
  buttonText: {
    color: "#fff",
    fontSize: 16,
    fontWeight: "bold",
  },
});
