import { Td, Tr, Button, Flex } from '@chakra-ui/react'
import React from 'react'
import { useNavigate, useParams } from "react-router-dom"
import StatusBlock from '../routes/StatusBlock'

export interface SubmissionWithStatusProps {
  Status: string
  Submission: SubmissionTableRowProps
}

export interface SubmissionTableRowProps {
  Id: number;
  UserId: number;
  ProblemId: number;
  SubmittedDate: string;
  Results: Result[];
  Status: string;
}

export interface Result {
  TestId: number;
  Status: string;
}

const SubmissionTableRow: React.FC<SubmissionTableRowProps> = ({ Id, SubmittedDate, ProblemId, UserId, Results, Status }) => {
  const navigate = useNavigate()
  const { userName } = useParams()
  return (
    <>
      <Tr>
        <Td>{new Date(SubmittedDate).toLocaleString()}</Td>
        <Td>{ProblemId}</Td>
        <Td>{userName}</Td>
        <Td>
          <Flex mr={2}>
            <StatusBlock status={Status} />
            <Button variant="link" onClick={() => navigate(`/${userName}/submission/${Id}`, { state: { status: Status } })}>
              詳細
            </Button>
          </Flex>
        </Td>
      </Tr>
    </>
  )
}
export default SubmissionTableRow
