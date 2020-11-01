import React from 'react'

export interface LayoutProps {
  children: React.ReactNode
}

export interface SubmitButtonProps {
  name: string
  loading: boolean
}

export interface TextFieldCustomProps {
  inputRef: (ref: HTMLInputElement | null) => void
  name: string
  label: string
  mask: string
  value: string
}
