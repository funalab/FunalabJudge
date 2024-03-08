
import { Button, Divider, Heading, Table, TableCaption, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'

interface problemStatus {
  id: number,
  Status: String,
}

interface B3Status {
  user: string,
  problems: problemStatus[],
}

interface B3StatusProps {
  data: B3Status[]
}

export const B3results = ({ data }: B3StatusProps) => {
  const navigate = useNavigate()
  return (
    <>
      <Heading mt={5}>B3の提出</Heading>
      <Divider />
      <TableContainer>
        <Table variant='simple'>
          <TableCaption>B3 results queue</TableCaption>
          <Thead>
            <Tr>
              <Th>Problem Id</Th>
              {data?.map(b3 => (
                <Th>
                  <Button variant="link" onClick={() => navigate(`/${b3.user}/results`)}>
                    {b3.user}
                  </Button>
                </Th>
              ))}
            </Tr>
          </Thead>
          <Tbody>
            {data[0].problems.map((problem, i) => (
              <Tr key={problem.id}>
                <Td>{problem.id}</Td>
                {data?.map(b3 => (
                  <Td>{b3.problems[i].Status}</Td>
                ))}
              </Tr>
            ))}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  )
}

