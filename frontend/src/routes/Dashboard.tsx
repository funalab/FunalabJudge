import { useState, useEffect } from "react";
import { Box, Text, VStack } from "@chakra-ui/react";
import axios from "axios";
import DefaultLayout from "../components/DefaultLayout";

const DashboardPage = () => {
  const [data, setData] = useState(null);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    // バックエンドサーバーのエンドポイントURLを指定
    const apiUrl = "http://localhost:3000/";

    // HTTP GETリクエストの送信
    axios
      .get(apiUrl)
      .then((response) => {
        // レスポンスを受け取り、stateにセットする
        setData(response.data);
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
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
};

export default DashboardPage;
