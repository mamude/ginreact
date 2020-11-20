/* eslint-disable react/jsx-props-no-spreading */
import React from 'react'
import { IconButton, IconButtonProps } from '@material-ui/core'
import AddIcon from '@material-ui/icons/Add'
import RemoveIcon from '@material-ui/icons/Remove'

const ButtonPlus = (props: IconButtonProps) => {
  return (
    <IconButton {...props} size="small">
      <AddIcon color="inherit" />
    </IconButton>
  )
}

const ButtonRemove = (props: IconButtonProps) => {
  return (
    <IconButton {...props} size="small">
      <RemoveIcon color="inherit" />
    </IconButton>
  )
}

export { ButtonPlus, ButtonRemove }
