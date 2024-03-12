
import { Button, Divider, Heading, Table, TableCaption, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import StatusBlock from './StatusBlock'
interface problemStatus {
  ProblemId: number,
  ProblemName: string,
  Status: string,
}

export interface B3StatusProps {
  UserName: string,
  ProblemsStatus: problemStatus[],
}


export const B3results = ({ data }: { data: B3StatusProps[] }) => {
  const navigate = useNavigate()
  return (
    <>
      <Heading mt={5}>B3の状態</Heading>
      <Divider />
      <TableContainer>
        <Table variant='simple'>
          <TableCaption>B3 results queue</TableCaption>
          <Thead>
            <Tr>
              <Th>Problem</Th>
              {data?.map(b3 => (
                <Th key={b3.UserName}>
                  <Button variant="link" onClick={() => navigate(`/${b3.UserName}/results`)}>
                    {b3.UserName}
                  </Button>
                </Th>
              ))}
            </Tr>
          </Thead>
          <Tbody>
            {data.length > 1 && data[0].ProblemsStatus?.map((problem, i) => (
              <Tr key={problem.ProblemId}>
                <Td>{problem.ProblemName}</Td>
                {data?.map(b3 => (
                  <Td><StatusBlock status={b3.ProblemsStatus[i].Status} onClick={() => navigate(`/${b3.UserName}/results`, { state: problem.ProblemName })} /></Td>
                ))}
              </Tr>
            ))}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  )
}

