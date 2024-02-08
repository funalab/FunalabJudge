import { useState, useEffect } from 'react';
import { Button, Box, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import { Card, CardHeader, CardBody, CardFooter, Heading } from '@chakra-ui/react'
import { SimpleGrid } from '@chakra-ui/react'
import { useNavigate } from "react-router-dom"

export const CardList = () => {
  const [data, setData] = useState(null);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    // バックエンドサーバーのエンドポイントURLを指定
    const apiUrl = 'http://localhost:3000';

    // HTTP GETリクエストの送信
    axios.get(apiUrl)
      .then(response => {
        // レスポンスを受け取り、stateにセットする
        setData(response.data);
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }, []);
  interface Assignment {
    title: string;
    content: string;
    open: string;
    close: string;
    done: string;
    path: string;
  }
  interface CardGridProps {
    data: Record<string, Assignment>;
  }
  const navigate = useNavigate()

  const CardGrid: React.FC<CardGridProps> = ({ data }) => {
    return (
      <SimpleGrid columns={{ sm: 2, md: 3, lg: 4 }} spacing="20px">
        {Object.entries(data).map(([key, assignment]) => (
          <Card key={key}>
            <CardHeader>
              <Heading> {assignment.title}</Heading>
            </CardHeader>
            <CardBody>
              {assignment.content}
            </CardBody>
            <CardFooter>
              <Button variant="ghost" onClick={() => navigate(assignment.path)}>
                詳細
              </Button>
            </CardFooter>
          </Card>
        ))}
      </SimpleGrid>
    );
  }

  // 一時的に帰ってくるであろうデータを設定
  const data_tmp = ({
    q1: {
      title: 'C-x',
      content: 'this assignment is application of x function',
      open: '2023/03/01',
      close: '2023/03/05',
      done: 'False',
      path: '/message'
    },
    q2: {
      title: 'C-1',
      content: 'this is first assignment',
      open: '2023/03/01',
      close: '2023/03/07',
      done: 'True',
      path: '/account'
    },
    q3: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2023/03/05',
      close: '2023/03/11',
      done: 'False',
      path: '/schedule'
    },
    q4: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2023/03/05',
      close: '2023/03/11',
      done: 'False',
      path: '/schedule'
    },
    q5: {
      title: 'python',
      content: 'Drawing a graph using python',
      open: '2023/03/05',
      close: '2023/03/11',
      done: 'False',
      path: '/schedule'
    }
  })

  return (
    <div>
      <Box>
        <h1>There are some assignments</h1>
      </Box>
      <CardGrid data={data_tmp} />
    </div>
  );
}

