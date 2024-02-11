import { HStack, Spacer, Text, VStack } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'

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
          <CopyTestcase content={InputFile} />
        </HStack>
        <Spacer />
        <HStack>
          <Text>出力例{id}</Text>
          <CopyTestcase content={OutputFile} />
        </HStack>
      </VStack>
    </>
  )
}

export default Testcase
