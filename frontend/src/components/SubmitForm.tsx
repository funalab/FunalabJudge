import { Input } from '@chakra-ui/react'
import React from 'react'
import SubmitButton from './SubmitButton'

export interface SubmitFormProps {
}

const handleInputFile = (ev: React.ChangeEvent<HTMLInputElement>) => {
  const selectedFile = ev.target.files![0]
  if (selectedFile) {
    /* selectedFile would be passed to compile phase. Temporary, console.log()*/
    console.log(selectedFile)
  }
}

const SubmitForm: React.FC<SubmitFormProps> = () => {
  return (
    <>
      <Input placeholder="Your file" type="file" onChange={handleInputFile} />
      {/* <Input type="text" onChange={handleCommitHash} /> */}
      <SubmitButton />
    </>
  )
}

export default SubmitForm
