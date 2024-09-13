import "react-native-gesture-handler";
import { createDrawerNavigator } from "@react-navigation/drawer";
import { NavigationContainer } from "@react-navigation/native";
import React, { useState, useEffect } from "react";
import SignIn from "./pages/SignIn";
import Signup from "./pages/Signup";
import Conversation from "./pages/NewChat";
import SignOut from "./pages/SignOut";
import useUserStore from "./store/auth";

const drawerForHome = createDrawerNavigator();
const drawerForAuth = createDrawerNavigator();

const Drawer = createDrawerNavigator();

// const AuthStack = () => {
//   return (
//     <drawerForAuth.Navigator>
//       <drawerForAuth.Screen name="signin" component={SignIn} />
//       <drawerForAuth.Screen name="signup" component={Signup} />
//     </drawerForAuth.Navigator>
//   );
// };

// const AppStack = () => {
//   return (
//     <drawerForHome.Navigator>
//       <drawerForHome.Screen name="conversation" component={Conversation} />
//       <drawerForHome.Screen name="signout" component={SignOut} />
//     </drawerForHome.Navigator>
//   );
// };

export default function App() {
  const [user, setUser] = useState("");
  const getUserToken = useUserStore((state) => state.getUserToken);

  useEffect(() => {
    const checkUser = async () => {
      const token = await getUserToken();
      if (token) {
        setUser(token);
      }
    };
    checkUser();
  }, [user]);

  return (
    <NavigationContainer>
      {/* {user ? <AppStack /> : <AuthStack />} */}
      <Drawer.Navigator>
        {user ? (
          <>
            <Drawer.Screen name="conversation" component={Conversation} />
            <Drawer.Screen name="signout" component={SignOut} />
          </>
        ) : (
          <>
            <Drawer.Screen name="signin" component={SignIn} />
            <Drawer.Screen name="signup" component={Signup} />
          </>
        )}
      </Drawer.Navigator>
    </NavigationContainer>
  );
}
