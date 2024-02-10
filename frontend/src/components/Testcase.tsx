import { HStack, Spacer, Text, VStack } from '@chakra-ui/react'
import React from 'react'
import CopyButton from './CopyButton.tsx'
import InputOutputBox from './InputOutputBox'

export interface TestcaseProps {
  id: string,
  input: string,
  output: string
}

const Testcase: React.FC<TestcaseProps> = ({ id, input, output }) => {
  return (
    <>
      <VStack>
        <HStack>
          <Text>入力例{id}</Text>
          <CopyButton />
          <InputOutputBox content={input} />
        </HStack>
        <Spacer />
        <HStack>
          <Text>出力例{id}</Text>
          <CopyButton />
          <InputOutputBox content={output} />
        </HStack>
      </VStack>
    </>
  )
}

export default Testcase
