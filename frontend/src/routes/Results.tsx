import React, { useEffect, useState } from 'react'
import { useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import SubmissionTableRow, { SubmissionTableRowProps } from '../components/SubmissionTableRow';
import { axiosClient } from '../providers/AxiosClientProvider';

const ResultsPage: React.FC = () => {
  const { userName } = useParams()
  const [submissions, setSubmissions] = useState<SubmissionTableRowProps[]>([])
  const [haveNotComplete, setHaveNotComplete] = useState<boolean>(false)

  useEffect(() => {
    axiosClient
      .get(`/getSubmissionList/${userName}`)
      .then((response) => {
        const { data } = response;
        setSubmissions(data.reverse())
        const complete = ["AC", "WA", "CE", "TLE", "RE"]
        data.map((submission: SubmissionTableRowProps) => {
          if (!complete.includes(submission.Status)) {
            setHaveNotComplete(true)
          }
        })
      })
      .catch((error) => {
        console.log(error)
        alert("Failed to fetch data from database")
      })
  }, [])

  /*未確定の奴があるなら1sずつリクエストを投げてレンダリングをする*/
  useEffect(() => {
    if (haveNotComplete) {
      const sendStatusRequest = () => {
        axiosClient.get(`/getSubmissionList/${userName}`)
          .then((response) => {
            const { data } = response;
            setSubmissions(data.reverse())
            const complete = ["AC", "WA", "CE", "TLE", "RE"]
            let completeFlag = true
            data.map((submission: SubmissionTableRowProps) => {
              if (!complete.includes(submission.Status)) {
                completeFlag = false
              }
            })
            if (completeFlag === true) {
              clearInterval(intervalId)
            }
          })
          .catch((error) => {
            console.log(error)
            alert("Failed to send status request")
          })
      }
      const intervalId = setInterval(sendStatusRequest, 10)
      return () => clearInterval(intervalId);
    }
  }, [haveNotComplete])

  return (
    <>
      <DefaultLayout>
        <Heading mt={5}>自分の提出</Heading>
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
              {/* This section is ongoing-judge submission row. */}

              {/* This section is existing submission list. */}
              {submissions?.map(submissions => (
                <SubmissionTableRow
                  Id={submissions.Id}
                  SubmittedDate={submissions.SubmittedDate}
                  ProblemId={submissions.ProblemId}
                  UserName={submissions.UserName}
                  Results={submissions.Results}
                  Status={submissions.Status}
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

export default ResultsPage
