import { Text, Stack, Divider, Flex, HStack, Input } from '@chakra-ui/react'
import { useState } from 'react'
import SubmitButton from './SubmitButton'

export interface SubmitFormProps {
  problemId: number
}

const SubmitForm: React.FC<SubmitFormProps> = ({ problemId }) => {
  const [selectedFiles, setSelectedFiles] = useState<File[]>([])

  const handleInputFile = (ev: React.ChangeEvent<HTMLInputElement>) => {
    const files: File[] = Array.from(ev.target.files!)
    handleSelectedFiles(files)
  }

  const handleSelectedFiles = (files: File[]) => {
    files.map((file) => {
      setSelectedFiles(prevFiles => [...prevFiles, file])
    })
  }
  return (
    <>
      <Divider />
      <HStack>
        <Text fontSize={30} fontWeight={'bold'}>Submit Form</Text>
      </HStack>
      <Stack>
        <Stack>
          <Input type="file" onChange={handleInputFile} multiple />
        </Stack>
        <Flex>
          <SubmitButton selectedFiles={selectedFiles} problemId={problemId} />
        </Flex>
      </Stack>
    </>
  )
}

export default SubmitForm
