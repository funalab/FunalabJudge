import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import axios from 'axios';
import SubmissionTableRow, { SubmissionTableRowProps } from './SubmissionTableRow';

const ResultQueuePage: React.FC = () => {
  const { userId } = useParams()
  const [submissions, setSubmissions] = useState<SubmissionTableRowProps[]>([])

  useEffect(() => {
    /* fetch all submissions that submitted by user whose id is useId */
    axios
      .get(`/submissions/${userId}`)
      .then((response) => {
        const { data } = response;
        setSubmissions(data)
      })
      .catch(() => {
        console.log('error')
        alert("Failed to fetch data from database")
      })
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
                  Id={submission.Id}
                  SubmittedDate={submission.SubmittedDate}
                  ProblemId={submission.ProblemId}
                  UserId={submission.UserId}
                  Results={submission.Results}
                  Status={submission.Status}
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
