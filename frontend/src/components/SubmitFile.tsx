import { Input } from '@chakra-ui/react'
import React from 'react'

interface SubmitFileProps {
  handleSelectedFiles: (file: File) => void
}

const SubmitFile: React.FC<SubmitFileProps> = ({ handleSelectedFiles }) => {
  const handleInputFile = (ev: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = ev.target.files![0]
    if (selectedFile) {
      /* selectedFile would be passed to compile phase. Temporary, console.log()*/
      handleSelectedFiles(selectedFile)
    }
  }
  return (
    <>
      <Input type="file" onChange={handleInputFile} />
    </>

  )
}

export default SubmitFile
