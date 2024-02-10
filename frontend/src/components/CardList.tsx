import { useState } from 'react';
import { Button, Box, Stack } from '@chakra-ui/react';
import { Card, CardHeader, CardBody, CardFooter, Heading } from '@chakra-ui/react'
import { SimpleGrid } from '@chakra-ui/react'
import { useNavigate } from "react-router-dom"
import { Checkbox } from '@chakra-ui/react'

interface Problems {
  ProblemId: number,
  Name: string,
  ExecutionTime: number,
  MemoryLimit: number,
  Statement: string,
  ProblemConstraints: string,
  InputFormat: string,
  OutputFormat: string,
  OpenDate: string,
  CloseDate: string,
  BorderScore: number,
  Status: boolean
}

interface CardListProps {
  data: Problems[];
}

export const CardList = ({ data }: CardListProps) => {

  const check_list = ["完了済みの課題も表示する", "未公開の課題も表示する"] //必要に応じて項目を増やすならここだけ変えればチェックボックスはそれに応じて増える
  const [checkedItems, setCheckedItems] = useState(new Array(check_list.length).fill(false))

  const navigate = useNavigate()

  const CardGrid = ({ data }: CardListProps) => {
    return (
      // 渡された引数のkeyの数だけカードの一覧を表示する
      <SimpleGrid columns={{ sm: 2, md: 3, lg: 4 }} spacing="20px">
        {data.map((assignment) => {

          // CheckBoxの状態に応じて表示するカードを変更する
          if (assignment.Status && !checkedItems[0]) {
            return null;
          }
          if (new Date() < new Date(assignment.OpenDate) && !checkedItems[1]) {
            return null;
          }

          return (
            <Card key={assignment.ProblemId}>
              <CardHeader>
                <Heading> {assignment.Name}</Heading>
              </CardHeader>
              <CardBody>
                <Box>
                  Open: {new Date(assignment.OpenDate).toLocaleString()}
                </Box>
                <Box>
                  Close: {new Date(assignment.CloseDate).toLocaleString()}
                </Box>
                {assignment.Status && <Box bg="red.500" w='100%' display='flex' justifyContent="center" alignItems="center" color="white">DONE!!</Box>}
              </CardBody>
              <CardFooter>
                {new Date() > new Date(assignment.OpenDate) ? (
                  <Button colorScheme='teal' onClick={() => navigate(`/assignmentInfo/${assignment.ProblemId}`)}>
                    詳細
                  </Button>
                ) : (
                  <Button variant="ghost">
                    詳細
                  </Button>
                )}
              </CardFooter>
            </Card>
          );
        })}
      </SimpleGrid>
    );
  }

  // 一時的に返ってくるであろうデータを設定
  const data_tmp = ({
    q1: {
      title: 'C-x',
      content: 'this assignment is application of x function',
      open: '2024-04-10',
      close: '2024-04-25',
      status: false,
      path: '/message'
    },
    q2: {
      title: 'C-1',
      content: 'this is first assignment',
      open: '2024-01-10',
      close: '2024-01-30',
      status: true,
      path: '/account'
    },
    q3: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2024-02-10',
      close: '2024-02-26',
      status: false,
      path: '/schedule'
    },
    q4: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2024-02-10',
      close: '2024-04-10',
      status: false,
      path: '/schedule'
    },
    q5: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2024-01-01',
      close: '2024-04-30',
      status: false,
      path: '/schedule'
    }
  })
  return (
    <div>
      <Box>
        <h1>There are some assignments</h1>
      </Box>
      <Stack spacing={5} direction='row'>
        {check_list.map((check, index: number) => (
          <Checkbox
            key={check}
            isChecked={checkedItems[index]}
            onChange={(e) => setCheckedItems([...checkedItems.slice(0, index), e.target.checked, ...checkedItems.slice(index + 1)])}
          >
            {check}
          </Checkbox>
        ))}
      </Stack>
      <CardGrid data={data} />
    </div>
  );
}

