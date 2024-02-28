import { useState } from 'react';
import { Button, Box, Stack } from '@chakra-ui/react';
import { Card, CardHeader, CardBody, CardFooter, Heading, Text } from '@chakra-ui/react'
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
            <Card
              key={assignment.ProblemResp.Pid}
              boxShadow={'xl'}
              transition={"box-shadow 0.3s"}
              _hover={{ top: -3, boxShadow: 'dark-lg', }}
            >
              <CardHeader>
                <Heading> {assignment.ProblemResp.Name}</Heading>
              </CardHeader>
              <CardBody>
                <Text
                  fontWeight={"bold"}
                >
                  Open: {new Date(assignment.OpenDate).toLocaleString()}
                </Text>
                <Text
                  fontWeight={"bold"}
                >
                  Close: {new Date(assignment.CloseDate).toLocaleString()}
                </Text>
                {assignment.Status && (
                  <Text
                    bg="red.500"
                    w='100%'
                    h="50%"
                    display='flex'
                    justifyContent="center"
                    alignItems="center"
                    color="white"
                    rounded={'xl'}
                    mt="3"
                    fontWeight={"bold"}
                  >
                    DONE!!
                  </Text>
                )}
              </CardBody>
              <CardFooter>
                {new Date() > new Date(assignment.OpenDate) ? (
                  <Button colorScheme='teal' onClick={() => navigate(`/${userName}/problem/${assignment.ProblemResp.Pid}`)}>
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
        })
        }
      </SimpleGrid >
    );
  }

  return (
    <>
      <Box my={4}>
        <Heading>Assignments</Heading>
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
      <Stack mt={8}>
        <CardGrid data={data} />
      </Stack>
    </>
  );
}

