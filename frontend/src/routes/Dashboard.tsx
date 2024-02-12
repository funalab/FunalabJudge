import {useState, useEffect } from 'react';
import {Button, Box, Text, VStack } from '@chakra-ui/react';
import { DefaultLayout } from "../components/DefaultLayout";
import axios from 'axios';

export const Dashboard = () => {
  const [data, setData] = useState(null);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    // バックエンドサーバーのエンドポイントURLを指定
    const apiUrl = 'http://localhost:3000/test'; 

    // HTTP GETリクエストの送信
    axios.get(apiUrl)
      .then(response => {
        console.log(response.data);
        // レスポンスを受け取り、stateにセットする
        setData(response.data.userID);
        console.log(response.data.userId);
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <DefaultLayout>
      <Box p={4}>
        <VStack spacing={4}>
          <Text>Data from Backend:</Text>
          {data ? (
            <Box>
              <Text>{JSON.stringify(data)}</Text>
            </Box>
          ) : (
            <Text>Loading...</Text>
          )}
        </VStack>
      </Box>
    </DefaultLayout>
  );
}
