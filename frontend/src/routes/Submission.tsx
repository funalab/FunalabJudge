import React, { ChangeEvent, useEffect, useState } from 'react'
import { useLocation, useParams } from 'react-router-dom'
import DefaultLayout from '../components/DefaultLayout'
import { Box, Center, Flex, Heading, Select, Table, TableContainer, Tbody, Td, Textarea, Th, Thead, Tr } from '@chakra-ui/react'
import { SubmissionTableRowProps } from '../components/SubmissionTableRow'
import { Result } from "../components/SubmissionTableRow"
import { axiosClient } from '../providers/AxiosClientProvider'
import { getStatusColor } from '../api/GetStausColor'
import StatusBlock from './StatusBlock'

type SubmittedFile = {
  name: string
  content: string
}

const SubmissionPage: React.FC = () => {
  const { submissionId } = useParams()
  const location = useLocation();
  const [totalStatus, setTotalStatus] = useState<string>('')
  const [files, setFiles] = useState<SubmittedFile[]>([])
  const [selectedFileContent, setSelectedFileContent] = useState<string>('')
  const [score, setScore] = useState(0)
  const [submission, setSubmission] = useState<SubmissionTableRowProps>({
    Id: 0,
    UserId: 0,
    ProblemId: 0,
    SubmittedDate: "",
    Results: [] as Result[],
    Status: ""
  })

  const handleSelectFile = (ev: ChangeEvent<HTMLSelectElement>) => {
    setSelectedFileContent(ev.target.value);
  }

  useEffect(() => {
    axiosClient
      .get(`/getSubmission/${submissionId}`)
      .then(({ data }) => {
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
        setTotalStatus(location.state.status)
      })
      .catch(() => {
        console.log('error')
        alert("Failed to fetch data from database.")
      })

    axiosClient
      .get(`getSubmittedFiles/${submissionId}`)
      .then(({ data }) => {
        setFiles(data)
        setSelectedFileContent(data[0].content)
      })
      .catch(() => {
        console.log('error')
        alert('Failed to fetch submitted files from database.')
      })
  }, []);

  return (
    <DefaultLayout>
      <>
        {files && (
          <Box
            p={10}
            my={10}
            bg={"gray.50"}
            borderRadius={'2xl'}
            boxShadow={'xl'}
          >
            <p
              className='pb-5 font-bold text-2xl'
            >
              あなたが提出したファイル一覧
            </p>
            <Select
              value={selectedFileContent}
              onChange={handleSelectFile}
              mb={5}
              color={'blue.500'}
              fontStyle={'italic'}
              fontWeight={'bold'}
            >
              {files.length > 0 && (
                files.map((file) => (
                  <option
                    value={file.content}
                  >
                    {file.name}
                  </option>
                ))
              )}
            </Select>
            <Textarea
              value={selectedFileContent}
              height="40vh"
            />
          </Box>
        )}
        <Box
          p={10}
          my={10}
          bg={"gray.50"}
          borderRadius={'2xl'}
          boxShadow={'xl'}
        >
          <p className='pb-5 font-bold text-2xl'>ジャッジ結果</p>
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
                  <Td>{Date(submission.SubmittedDate)}</Td>
                  <Td>{submission.ProblemId}</Td>
                  <Td>{score} / {submission.Results.length}</Td>
                  <Td>
                    <StatusBlock status={totalStatus} />
                  </Td>
                </Tr>
              </Tbody>
            </Table>
          </TableContainer>


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
                  <Td>
                    <StatusBlock status={result.Status} />
                  </Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </Box>
      </>
    </DefaultLayout >
  )
}

export default SubmissionPage
