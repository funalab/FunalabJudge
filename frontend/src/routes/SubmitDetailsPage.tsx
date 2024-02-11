import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useParams } from 'react-router-dom'
import DefaultLayout from '../components/DefaultLayout'
import { Code, Heading, Table, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react'
import { SubmissionTableRowProps } from './SubmissionTableRow'
import { Result } from "./SubmissionTableRow"

const SubmitDetailsPage: React.FC = () => {
  const { submitId } = useParams()
  const [submission, setSubmission] = useState<SubmissionTableRowProps>()
  const [score, setScore] = useState(0)
  useEffect(() => {
    axios
      .get(`/submission/${submitId}`)
      .then((response) => {
        const { data } = response
        setSubmission(data)
        console.log(data)
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
        <Heading>提出番号 {submitId}</Heading>
        <Code>import os</Code>
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
                <Td>submission!.SubmittedDate</Td>
                <Td>submission!.ProblemId</Td>
                <Td>{score}</Td>
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

            <Tr>
              <Td>result.testId</Td>
              <Td>result.status</Td>
            </Tr>


          </Tbody>
        </Table>
      </>
    </DefaultLayout>
  )
}

export default SubmitDetailsPage
