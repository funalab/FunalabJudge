import { SimpleGrid } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'
import { Testcase } from '../types/DbTypes'

const TestcaseBox: React.FC<Testcase> = ({ TestcaseId, InputFileContent, OutputFileContent }) => {
  return (
    <>
      <SimpleGrid mt={6} mb={6} columns={2} spacingX={'20px'}>
        <CopyTestcase
          text={`入力例${TestcaseId}`}
          content={InputFileContent}
        />
        <CopyTestcase
          text={`出力例${TestcaseId}`}
          content={OutputFileContent}
        />
      </SimpleGrid>
    </>
  )
}

export default TestcaseBox
