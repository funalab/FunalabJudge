import React, { useEffect, useState } from 'react'
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import axios from 'axios';
import SubmissionTableRow, { SubmissionTableRowProps } from './SubmissionTableRow';

const ResultQueuePage: React.FC<number> = (userId: number) => {

  const [submissions, setSubmissions] = useState<SubmissionTableRowProps[]>([])

  useEffect(() => {
    (async () => {
      /* fetch all submissions that submitted by user whose id is useId */
      const { data } = await axios.get("/submissions/" + userId)
      setSubmissions(data)
    })()
  }, []);

  return (
    <>
      <DefaultLayout>
        <Heading>自分の提出</Heading>
        <Divider />
        <TableContainer>
          <Table variant='simple'>
            <TableCaption>Your Submission Queue</TableCaption>
            <Thead>
              <Tr>
                <Th>提出日時</Th>
                <Th>問題</Th>
                <Th>ユーザ</Th>
                <Th>結果</Th>
              </Tr>
            </Thead>
            <Tbody>
              {submissions.map(submission => (
                <SubmissionTableRow
                  submittedDate={submission.submittedDate}
                  problemId={submission.problemId}
                  userId={submission.userId}
                  status={submission.status}
                />
              ))}
            </Tbody>
            <Tfoot>
              {/* Nothing */}
            </Tfoot>
          </Table>
        </TableContainer>
      </DefaultLayout>
    </>
  )
}

export default ResultQueuePage
