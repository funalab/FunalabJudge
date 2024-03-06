import { SimpleGrid } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'
import { Testcase } from '../types/DbTypes'

const TestcaseBox: React.FC<Testcase> = (t: Testcase) => {
  return (
    <>
      <SimpleGrid mt={3} mb={6} columns={3} spacingX={'20px'}>
        <CopyTestcase
          text={`引数例${t.TestcaseId}`}
          content={t.ArgsFileContent}
        />
        <CopyTestcase
          text={`入力例${t.TestcaseId}`}
          content={t.InputFileContent}
        />
        <CopyTestcase
          text={`出力例${t.TestcaseId}`}
          content={t.OutputFileContent}
        />
      </SimpleGrid>
    </>
  )
}

export default TestcaseBox
