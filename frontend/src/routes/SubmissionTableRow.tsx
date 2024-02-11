import { Td, Tr, Button } from '@chakra-ui/react'
import React from 'react'
import { useNavigate } from "react-router-dom"
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
  return (
    <>
      <Tr>
        <Td>{new Date(SubmittedDate).toLocaleString()}</Td>
        <Td>{ProblemId}</Td>
        <Td>{UserId}</Td>
        <Td>{Status}   <Button variant="link" onClick={() => navigate(`/submission/${Id}`)}>詳細</Button> </Td>
      </Tr>
    </>
  )
}
export default SubmissionTableRow
