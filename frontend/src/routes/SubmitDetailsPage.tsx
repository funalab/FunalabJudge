import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import DefaultLayout from '../components/DefaultLayout'
import { Heading, Table, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react'
import { SubmissionTableRowProps } from '../components/SubmissionTableRow'
import { Result } from "../components/SubmissionTableRow"
import { axiosClient } from '../providers/AxiosClientProvider'

const SubmitDetailsPage: React.FC = () => {
  const { submissionId } = useParams()
  const [submission, setSubmission] = useState<SubmissionTableRowProps>({
    Id: 0,
    UserId: 0,
    ProblemId: 0,
    SubmittedDate: "",
    Results: [] as Result[],
    Status: ""
  })

  const [score, setScore] = useState(0)
  useEffect(() => {
    axiosClient
      .get(`/submission/${submissionId}`)
      .then(({ data }) => {
        console.log(data)
        setSubmission(data)
        let newScore = 0;
        {
          data.Results.forEach((result: Result) => {
            if (result.Status == "AC") {
              newScore += 1
            }
          })
        }
        setScore(newScore)
      })
      .catch(() => {
        console.log('error')
        alert("Failed to fetch data from database")
      })
  }, []);


  return (
    <DefaultLayout>
      <>
        <Heading>提出番号 {submissionId}</Heading>
        <TableContainer>
          <Table variant='simple'>
            <Thead>
              <Tr>
                <Th>提出日時</Th>
                <Th>問題</Th>
                <Th>得点</Th>
                <Th>判定</Th>
              </Tr>
            </Thead>
            <Tbody>
              <Tr>
                <Td>{submission.SubmittedDate}</Td>
                <Td>{submission.ProblemId}</Td>
                <Td>{score} / {submission.Results.length}</Td>
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>

        <Heading>ジャッジ結果</Heading>
        <Table variant='simple'>
          <Thead>
            <Tr>
              <Th>ケース名</Th>
              <Th>結果</Th>
            </Tr>
          </Thead>
          <Tbody>
            {submission.Results.map((result) => (
              <Tr>
                <Td>{result.TestId}</Td>
                <Td>{result.Status}</Td>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </>
    </DefaultLayout>
  )
}

export default SubmitDetailsPage
