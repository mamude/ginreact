/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/jsx-props-no-spreading */
import { TextField } from '@material-ui/core'
import { fieldToTextField, TextFieldProps } from 'formik-material-ui'
import React from 'react'
import InputMask from 'react-input-mask'
import { TextFieldCustomProps } from '../../common/interfaces/props'

const TextFieldPhoneMaskConfig = (props: TextFieldCustomProps) => {
  const { inputRef, ...others } = props
  return (
    <InputMask
      {...others}
      ref={(ref: any) => {
        inputRef(ref ? ref.inputElement : null)
      }}
    />
  )
}

const TextFieldPhoneMask = (props: TextFieldProps) => {
  return (
    <TextField
      {...fieldToTextField(props)}
      InputProps={{
        inputComponent: TextFieldPhoneMaskConfig as any,
        inputProps: {
          mask: '(99) 99999-9999',
        },
      }}
    />
  )
}

export default TextFieldPhoneMask
