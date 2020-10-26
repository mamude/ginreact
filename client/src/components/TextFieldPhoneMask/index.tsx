/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/jsx-props-no-spreading */
import React from 'react'
import InputMask from 'react-input-mask'
import { TextFieldCustomProps } from '../../common/interfaces/props'

const TextFieldPhoneMask = (props: TextFieldCustomProps) => {
  const { inputRef, ...others } = props
  return (
    <InputMask
      {...others}
      ref={(ref: any) => {
        inputRef(ref ? ref.inputElement : null)
      }}
      mask="(999) 99999-9999"
    />
  )
}

export default TextFieldPhoneMask
