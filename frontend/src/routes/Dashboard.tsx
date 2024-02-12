import {useState, useEffect } from 'react';
import { Box, Text, VStack } from '@chakra-ui/react';
import { DefaultLayout } from "../components/DefaultLayout";
import { axiosClient } from '../providers/AxiosClientProvider';

export const Dashboard = () => {
  const [data, setData] = useState(null);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    axiosClient.get("/test")
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
