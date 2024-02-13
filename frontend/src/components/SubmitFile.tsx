import { Input } from '@chakra-ui/react'
import React from 'react'

const SubmitFile = () => {
  const handleInputFile = (ev: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = ev.target.files![0]
    if (selectedFile) {
      /* selectedFile would be passed to compile phase. Temporary, console.log()*/
      console.log(selectedFile)
    }
  }
  return (
    <>
      <Input type="file" onChange={handleInputFile} />
    </>

  )
}

export default SubmitFile
