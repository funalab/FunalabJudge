import { HStack, Spacer, Text, VStack } from '@chakra-ui/react'
import React from 'react'
import CopyButton from './CopyButton.tsx'
import InputFileOutputBox from './InputOutputBox'

export interface TestcaseProps {
  id: string,
  InputFile: string,
  OutputFile: string
}

const Testcase: React.FC<TestcaseProps> = ({ id, InputFile, OutputFile }) => {
  return (
    <>
      <VStack>
        <HStack>
          <Text>入力例{id}</Text>
          <CopyButton />
          <InputFileOutputBox content={InputFile} />
        </HStack>
        <Spacer />
        <HStack>
          <Text>出力例{id}</Text>
          <CopyButton />
          <InputFileOutputBox content={OutputFile} />
        </HStack>
      </VStack>
    </>
  )
}

export default Testcase
