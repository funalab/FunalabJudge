import { Text, Stack, Divider, Flex, HStack, Input, InputRightAddon, Button, Textarea } from '@chakra-ui/react'
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

  return (
    <>
      <Divider />
      <HStack>
        <Text fontSize={30} fontWeight={'bold'}>Submit Form</Text>
      </HStack>
      <Stack>
        <Input type="file" onChange={handleInputFile} multiple />
        <Textarea placeholder="Your submitted files." value={filenames} />
        <Flex>
          <SubmitButton selectedFiles={selectedFiles} problemId={problemId} />
        </Flex>
      </Stack>
    </>
  )
}

export default SubmitForm
