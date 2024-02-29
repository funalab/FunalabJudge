import { useState } from 'react';
import { Button, Box, Stack } from '@chakra-ui/react';
import { Card, CardHeader, CardBody, CardFooter, Heading } from '@chakra-ui/react'
import { SimpleGrid } from '@chakra-ui/react'
import { useNavigate, useParams } from "react-router-dom"
import { Checkbox } from '@chakra-ui/react'
import { Problem } from '../types/DbTypes';

interface problemWithStatus {
	Problem: Problem,
	Status: boolean,
}

interface CardListProps {
  data: problemWithStatus[];
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
        {data.map((pws) => {
          // CheckBoxの状態に応じて表示するカードを変更する
          if (pws.Status && !checkedItems[0]) {
            return null;
          }
          if (new Date() < new Date(pws.Problem.OpenDate) && !checkedItems[1]) {
            return null;
          }
          return (
            <Card key={pws.Problem.Id} boxShadow={'dark-lg'}>
              <CardHeader>
                <Heading> {pws.Problem.Name}</Heading>
              </CardHeader>
              <CardBody>
                <Box>
                  Open: {new Date(pws.Problem.OpenDate).toLocaleString()}
                </Box>
                <Box>
                  Close: {new Date(pws.Problem.CloseDate).toLocaleString()}
                </Box>
                {pws.Status && <Box bg="red.500" w='100%' display='flex' justifyContent="center" alignItems="center" color="white">DONE!!</Box>}
              </CardBody>
              <CardFooter>
                {new Date() > new Date(pws.Problem.OpenDate) ? (
                  <Button colorScheme='teal' onClick={() => navigate(`/${userName}/problem/${pws.Problem.Id}`)}>
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

