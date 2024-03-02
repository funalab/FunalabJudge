import { SimpleGrid } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'
import { Testcase } from '../types/DbTypes'

const TestcaseBox: React.FC<Testcase> = ({ TestcaseId, ArgsFileContent, InputFileContent, OutputFileContent }) => {
  return (
    <>
      <SimpleGrid mt={3} mb={6} columns={3} spacingX={'20px'}>
        <CopyTestcase
          text={`引数例${TestcaseId}`}
          content={ArgsFileContent}
        />
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
