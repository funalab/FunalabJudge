import { Td, Tr } from '@chakra-ui/react'
import React from 'react'

export interface SubmissionTableRowProps {
  submittedDate: string,
  problemId: number,
  userId: number,
  status: boolean,
}

const SubmissionTableRow: React.FC<SubmissionTableRowProps> = ({ submittedDate, problemId, userId, status }) => {
  return (
    <>
      <Tr>
        <Td>{submittedDate}</Td>
        <Td>{problemId}</Td>
        <Td>{userId}</Td>
        <Td>{status}</Td>
      </Tr>
    </>
  )
}
export default SubmissionTableRow
