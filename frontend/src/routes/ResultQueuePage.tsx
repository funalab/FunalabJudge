import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import axios from 'axios';
import SubmissionTableRow, { SubmissionWithStatusProps } from './SubmissionTableRow';

const ResultQueuePage: React.FC = () => {
  const { userId } = useParams()
  const [submissionsWithStatus, setSubmissionWithStatus] = useState<SubmissionWithStatusProps[]>([])

  useEffect(() => {
    /* fetch all submissions that submitted by user whose id is useId */
    axios
      .get(`/submissions/${userId}`)
      .then((response) => {
        const { data } = response;
        setSubmissionWithStatus(data)
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
              {submissionsWithStatus.map(submissionWithStatus => (
                <SubmissionTableRow
                  Id={submissionWithStatus.Submission.Id}
                  SubmittedDate={submissionWithStatus.Submission.SubmittedDate}
                  ProblemId={submissionWithStatus.Submission.ProblemId}
                  UserId={submissionWithStatus.Submission.UserId}
                  Results={submissionWithStatus.Submission.Results}
                  Status={submissionWithStatus.Status}
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
