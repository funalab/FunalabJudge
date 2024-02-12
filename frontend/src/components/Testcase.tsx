import { HStack, Spacer, Text, VStack } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'

export interface TestcaseProps {
  id: string,
  InputFilePath: string,
  OutputFilePath: string
}

const Testcase: React.FC<TestcaseProps> = ({ id, InputFilePath, OutputFilePath }) => {
  return (
    <>
      <VStack>
        <HStack>
          <Text>入力例{id}</Text>
          <CopyTestcase content={InputFilePath} />
        </HStack>
        <Spacer />
        <HStack>
          <Text>出力例{id}</Text>
          <CopyTestcase content={OutputFilePath} />
        </HStack>
      </VStack>
    </>
  )
}

export default Testcase
