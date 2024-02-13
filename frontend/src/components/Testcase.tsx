import { Divider, SimpleGrid } from '@chakra-ui/react'
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
      <SimpleGrid mt={6} mb={6} columns={2} spacingX={'20px'}>
        <CopyTestcase
          text={`入力例${id}`}
          content={InputFileContent}
        />
        <CopyTestcase
          text={`出力例${id}`}
          content={OutputFileContent}
        />
      </SimpleGrid>
      <Divider />
    </>
  )
}

export default Testcase
