import React, { useEffect, useState } from "react";
import {
  Text,
  View,
  StyleSheet,
  FlatList,
  TextInput,
  TouchableOpacity,
  SafeAreaView,
} from "react-native";
import { FontAwesome, AntDesign } from "@expo/vector-icons";
import { Chat } from "../models/chat";
import { createChat, getAllChats, getGPT } from "../APIs/chat";
import { v4 as uuidv4 } from "uuid";

export default function Conversation() {
  const [chats, setChats] = useState<Chat[]>([]);
  const [newQuestion, setNewQuestion] = useState("");
  const [newAnswer, setNewAnswer] = useState("");
  const [fetchingResponse, setFetchingResponse] = useState(false);

  useEffect(() => {
    const fetchAllChats = async () => {
      let chats: Chat[] = [];
      chats = await getAllChats();
      setChats(chats);
    };
    fetchAllChats();
  }, []);

  const handleAddQuestion = async () => {
    try {
      setFetchingResponse(true);
      const response = await getGPT(newQuestion);
      const answer = response.data;

      setChats((prevChats) => [
        ...prevChats,
        new Chat({
          question: newQuestion,
          answer,
          id: uuidv4(),
          image: "",
          userId: "",
        }),
      ]);

      setNewQuestion("");
      setNewAnswer("");
      await createChat(new Chat({ question: newQuestion, answer }));
    } catch (error) {
      console.error("Error while fetching response", error);
    }
    setFetchingResponse(false);
  };

  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.chatContainer}>
        {chats.length > 0 && (
          <FlatList
            style={styles.flatList}
            data={chats}
            keyExtractor={(item) => item.id.toString()}
            renderItem={({ item }) => (
              <View style={styles.item}>
                <View style={styles.view}>
                  <AntDesign name="user" size={24} color="black" />
                  <Text style={styles.text}>{item.question}</Text>
                </View>
                <View style={[styles.view, styles.answer]}>
                  <FontAwesome name="user-secret" size={24} color="black" />
                  <Text style={styles.text}>{item.answer}</Text>
                </View>
              </View>
            )}
          />
        )}
      </View>

      <View style={styles.inputContainer}>
        <TextInput
          style={styles.input}
          placeholder="Enter a new question"
          value={newQuestion}
          onChangeText={(text) => setNewQuestion(text)}
        />
        <TouchableOpacity
          style={[
            styles.addButton,
            (fetchingResponse || !newQuestion) && styles.btnDisabled,
          ]}
          onPress={handleAddQuestion}
          disabled={!newQuestion || fetchingResponse}
        >
          <Text style={styles.buttonText}>Add Question</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  chatContainer: {
    flex: 1,
  },
  flatList: {
    backgroundColor: "#fff",
    borderRadius: 10,
  },
  item: {
    marginHorizontal: 25,
    marginVertical: 10,
    borderRadius: 10,
    width: "100%",
  },
  view: {
    flexDirection: "row",
    width: "80%",
    margin: "auto",
    backgroundColor: "#F0F0F0",
    paddingVertical: 10,
    paddingHorizontal: 15,
    borderRadius: 10,
  },
  text: {
    marginLeft: 10,
    fontSize: 16,
    color: "#333",
    width: "90%",
  },
  inputContainer: {
    padding: 20,
    backgroundColor: "#fff",
  },
  input: {
    height: 40,
    borderColor: "#CCCCCC",
    borderWidth: 1,
    marginBottom: 10,
    paddingLeft: 15,
    borderRadius: 5,
    fontSize: 16,
    color: "#333",
  },
  addButton: {
    backgroundColor: "#4CAF50",
    padding: 15,
    borderRadius: 5,
    alignItems: "center",
  },
  buttonText: {
    color: "white",
    fontSize: 18,
    fontWeight: "bold",
  },
  answer: {
    marginTop: 10,
  },
  btnDisabled: {
    backgroundColor: "#ACE1AF",
  },
});
