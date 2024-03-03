import { Td, Tr, Button, Flex } from '@chakra-ui/react'
import React from 'react'
import { useNavigate } from "react-router-dom"
import StatusBlock from './StatusBlock'

export interface SubmissionWithStatusProps {
  Status: string
  Submission: SubmissionTableRowProps
}

export interface SubmissionTableRowProps {
  Id: number;
  UserName: string;
  ProblemId: number;
  SubmittedDate: string;
  Results: Result[];
  Status: string;
}

export interface Result {
  TestId: number;
  Status: string;
}

const SubmissionTableRow: React.FC<SubmissionTableRowProps> = ({ Id, SubmittedDate, ProblemId, UserName, Status }) => {
  const navigate = useNavigate()
  return (
    <>
      <Tr>
        <Td>{new Date(SubmittedDate).toLocaleString()}</Td>
        <Td>{ProblemId}</Td>
        <Td>{UserName}</Td>
        <Td>
          <Flex mr={2}>
            <StatusBlock status={Status} />
            <Button variant="link" onClick={() => navigate(`/${UserName}/submission/${Id}`, { state: { status: Status } })}>
              詳細
            </Button>
          </Flex>
        </Td>
      </Tr>
    </>
  )
}
export default SubmissionTableRow
