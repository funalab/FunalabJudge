import { HStack, Spacer, Text, VStack } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'

export interface TestcaseProps {
  id: string,
  InputFileContent: string,
  OutputFileContent: string
}

const Testcase: React.FC<TestcaseProps> = ({ id, InputFileContent, OutputFileContent }) => {
  return (
    <>
      <VStack>
        <HStack>
          <Text>入力例{id}</Text>
          <CopyTestcase content={InputFileContent} />
        </HStack>
        <Spacer />
        <HStack>
          <Text>出力例{id}</Text>
          <CopyTestcase content={OutputFileContent} />
        </HStack>
      </VStack>
    </>
  )
}

export default Testcase
