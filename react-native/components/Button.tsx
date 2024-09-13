import React from 'react'
import { Button, View } from 'react-native'

function CustomBTN({
    title = "Button",
    onPress = () => {},
    color = "blue",
    disabled = false,
    styles = {},
    ...props
}) {
  return (
    <View style={styles}>
      <Button
       title={title}
       onPress={onPress}
       color={color}
       disabled={disabled}
       {...props}
       />
      </View>
  )
}

export default CustomBTN
