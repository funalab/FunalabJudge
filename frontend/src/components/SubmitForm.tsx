import { Text, Stack, Input, Textarea, VStack } from '@chakra-ui/react'
import { useState } from 'react'
import SubmitButton from './SubmitButton'

export interface SubmitFormProps {
  problemId: number
}

const SubmitForm: React.FC<SubmitFormProps> = ({ problemId }) => {
  const [selectedFiles, setSelectedFiles] = useState<File[]>([])
  const [filenames, setFilenames] = useState<string>('')

  const handleInputFile = (ev: React.ChangeEvent<HTMLInputElement>) => {
    const files: File[] = Array.from(ev.target.files!)
    const regexOk = handleRegex(files)
    if (!regexOk) {
      ev.target.value = ''
      return
    }
    handleSelectedFiles(files)
    handleFileNames(files)
  }

  const handleSelectedFiles = (files: File[]) => {
    files.map((file) => {
      setSelectedFiles(prevFiles => [...prevFiles, file])
    })
  }

  const handleFileNames = (files: File[]) => {
    let f = files.map(file => file.name).join(', ')
    setFilenames(f)
  }

  const handleRegex = (files: File[]) => {
    const regex = new RegExp('^Makefile$|\\.c$|\\.h$')
    const regexNotOkFile = files.find((selectedFile: File) => (regex.test(selectedFile.name) === false))
    if (regexNotOkFile) {
      alert("Cファイルおよびヘッダファイル、Makefileのみ提出してください。")
      return false
    }
    return true
  }

  return (
    <>
      <VStack mb={20}>
        <Text fontSize={30} fontWeight={'bold'}>Submit Form</Text>
        <Stack>
          <Input type="file" onChange={handleInputFile} multiple />
          <Textarea placeholder="Your submitted files." value={filenames} readOnly />
          <SubmitButton selectedFiles={selectedFiles} problemId={problemId} />
        </Stack>
      </VStack>
    </>
  )
}

export default SubmitForm
