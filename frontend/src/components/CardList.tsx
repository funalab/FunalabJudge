import { useState } from 'react';
import { Button, Box, Stack } from '@chakra-ui/react';
import { Card, CardHeader, CardBody, CardFooter, Heading } from '@chakra-ui/react'
import { SimpleGrid } from '@chakra-ui/react'
import { useNavigate, useParams } from "react-router-dom"
import { Checkbox } from '@chakra-ui/react'

interface Problems {
  CloseDate: string,
  OpenDate: string,
  Status: boolean
  ProblemResp: ProblemResp,
}

interface ProblemResp {
  Pid: number,
  Name: string,
  ExTime: number,
  MemLim: number,
  Statement: string,
  Prbconst: string,
  InputFmt: string,
  OutputFmt: string,
  OpenDate: string,
  CloseDate: string,
  BorderScore: number,
}

interface CardListProps {
  data: Problems[];
}

export const CardList = ({ data }: CardListProps) => {

  const check_list = ["完了済みの課題も表示する", "未公開の課題も表示する"] //必要に応じて項目を増やすならここだけ変えればチェックボックスはそれに応じて増える
  const [checkedItems, setCheckedItems] = useState(new Array(check_list.length).fill(false))

  const navigate = useNavigate()
  const { userName } = useParams()

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
            <Card key={assignment.ProblemResp.Pid}>
              <CardHeader>
                <Heading> {assignment.ProblemResp.Name}</Heading>
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
                  <Button colorScheme='teal' onClick={() => navigate(`/${userName}/assignmentInfo/${assignment.ProblemResp.Pid}`)}>
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

