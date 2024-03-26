import { Flex } from '@chakra-ui/react'
import React from 'react'
import CopyTestcase from './CopyTestcase'
import CopyTestcaseMulti from './CopyTestcaseMulti'
import { Testcase } from '../types/DbTypes'

const TestcaseBox: React.FC<Testcase> = (t: Testcase) => {
  return (
    <>
      <Flex mt={3} mb={6} overflowX="auto">
        <CopyTestcase
          text={`引数${t.TestcaseId}`}
          content={t.ArgsFileContent}
        />
        { t.InputFileList.length > 0 &&
          <CopyTestcaseMulti
            text={`使用ファイル${t.TestcaseId}`}
            files={t.InputFileList}
          />
        }
        <CopyTestcase
          text={`標準入力${t.TestcaseId}`}
          content={t.StdinFileContent}
        />
        <CopyTestcase
          text={`出力${t.TestcaseId}`}
          content={t.AnswerFileContent}
        />
      </Flex>
    </>
  )
}

export default TestcaseBox
